package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func ListTimeseriesMeasurementsByTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tsID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		f := models.MeasurementsFilter{TimeseriesID: tsID}

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f)
		return streamMeasurementsHandler(c)
	}
}

func ListTimeseriesMeasurementsByInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		iID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		f := models.MeasurementsFilter{InstrumentID: iID}

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f)
		return streamMeasurementsHandler(c)
	}
}

func ListTimeseriesMeasurementsByInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		igID, err := uuid.Parse(c.Param("instrument_group_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		f := models.MeasurementsFilter{InstrumentID: igID}

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f)
		return streamMeasurementsHandler(c)
	}
}

func ListTimeseriesMeasurementsExplorer(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var iIDs []uuid.UUID

		// Instrument IDs from POST
		if err := (&echo.DefaultBinder{}).BindBody(c, &iIDs); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		f := models.MeasurementsFilter{}
		if err := f.InstrumentIDs.Set(&iIDs); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f)
		return streamMeasurementsHandler(c)
	}
}

func StreamTimeseriesMeasurements(db *sqlx.DB, f *models.MeasurementsFilter) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tw timeseries.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err := tw.SetWindow(a, b); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		f.After = tw.After
		f.Before = tw.Before

		rows, err := models.QueryTimeseriesMeasurementsRows(db, f)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		defer func() {
			if err := rows.Close(); err != nil {
				log.Fatal(err.Error())
			}
		}()

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		enc := json.NewEncoder(c.Response())

		// LOCF (Last Observation Carried Forward)
		remember := make(map[uuid.UUID]map[string]interface{})

		for rows.Next() {
			var mfs models.MeasurementsFromStream

			if err = rows.StructScan(&mfs); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			var env map[string]interface{}
			if err = mfs.MeasurementsJSON.AssignTo(&env); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			// Simply stream stored timeseries and continue
			if !mfs.IsComputed {
				val, exists := env["value"]
				if !exists {
					log.Warnf("bad measurements_json %v for row with date %s and timeseries_id %v", mfs.MeasurementsJSON, mfs.Time, mfs.TimeseriesID)
					continue
				}
				m := map[string]interface{}{"timeseries_id": mfs.TimeseriesID, "time": mfs.Time, "value": val}

				if masked, exists := env["masked"]; exists {
					m["masked"] = masked
				}
				if validated, exists := env["validated"]; exists {
					m["validated"] = validated
				}
				if annotation, exists := env["annotation"]; exists {
					m["annotation"] = annotation
				}

				if err := enc.Encode(m); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				c.Response().Flush()
				continue
			}

			// Carry forward any values that don't exist for this timestamp with the last known observation
			if _, exists := remember[mfs.TimeseriesID]; !exists {
				remember[mfs.TimeseriesID] = make(map[string]interface{})
			}
			for k, v := range remember[mfs.TimeseriesID] {
				if _, exists := env[k]; !exists {
					env[k] = v
				}
			}
			// Add/Update the most recent values
			for k, v := range env {
				remember[mfs.TimeseriesID][k] = v
			}

			expression, err := govaluate.NewEvaluableExpression(mfs.Formula)
			if err != nil {
				log.Warn(err.Error())
				return err
			}

			val, err := expression.Evaluate(env)
			if err != nil {
				// Any evaluation errors are passed back to client
				if err := enc.Encode(map[string]interface{}{"timeseries_id": mfs.TimeseriesID, "time": mfs.Time, "error": err.Error()}); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				c.Response().Flush()
				continue
			}

			val64, ok := val.(float64)
			if !ok {
				log.Warnf("unable to convert %v interface{} to float64", val)
				continue
			}

			m := map[string]interface{}{"timeseries_id": mfs.TimeseriesID, "time": mfs.Time, "value": val64}

			if err := enc.Encode(m); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Response().Flush()
		}

		if err := rows.Err(); err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}
