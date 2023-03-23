package handlers

import (
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
		f := models.MeasurementsFilter{TimeseriesID: &tsID}

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
		f := models.MeasurementsFilter{InstrumentID: &iID}

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
		f := models.MeasurementsFilter{InstrumentGroupID: &igID}

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

		f := models.MeasurementsFilter{InstrumentIDs: iIDs}

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f)
		return streamMeasurementsHandler(c)
	}
}

// StreamTimeseriesMeasurements emits newline delimited json objects. The buffer flushes to the client
// every 1000 records, plus any remaining records in the buffer when complete
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

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)

		enc := json.NewEncoder(c.Response())

		// LOCF (Last Observation Carried Forward)
		remember := make(map[uuid.UUID]map[string]interface{})
		rowsInChunk := 0

		for rows.Next() {
			// Buffer is chunked to send for every 1000 records
			if rowsInChunk > 0 && rowsInChunk%1000 == 0 {
				c.Response().Flush()
				rowsInChunk = 0
			}
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

				val64, ok := val.(float64)
				if !ok {
					log.Warnf("unable to convert %v interface{} to float64", val)
					continue
				}

				m := models.MeasurementsStreamResponse{InstrumentID: mfs.InstrumentID, TimeseriesID: mfs.TimeseriesID, Time: mfs.Time, Value: val64}

				if masked, exists := env["masked"]; exists {
					maskedBool, ok := masked.(bool)
					if !ok {
						log.Warnf("unable to convert %v interface{} to bool", masked)
					}
					m.Masked = maskedBool
				}
				if validated, exists := env["validated"]; exists {
					validatedBool, ok := validated.(bool)
					if !ok {
						log.Warnf("unable to convert %v interface{} to bool", validated)
					}
					m.Validated = validatedBool
				}
				if annotation, exists := env["annotation"]; exists {
					annotationStr, ok := annotation.(string)
					if !ok {
						log.Warnf("unable to convert %v interface{} to string", annotation)
					}
					m.Annotation = annotationStr
				}

				if err := enc.Encode(m); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				rowsInChunk++
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
				m := models.MeasurementsStreamResponse{InstrumentID: mfs.InstrumentID, TimeseriesID: mfs.TimeseriesID, Time: mfs.Time, Error: err.Error()}

				// Any evaluation errors are passed back to client
				if err := enc.Encode(m); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				rowsInChunk++
				continue
			}

			val64, ok := val.(float64)
			if !ok {
				log.Warnf("unable to convert %v interface{} to float64", val)
				continue
			}

			m := models.MeasurementsStreamResponse{InstrumentID: mfs.InstrumentID, TimeseriesID: mfs.TimeseriesID, Time: mfs.Time, Value: val64}

			if err := enc.Encode(m); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			rowsInChunk++
		}

		// Send any remianing records
		if rowsInChunk > 0 {
			c.Response().Flush()
		} else {
			echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
		}

		return nil
	}
}
