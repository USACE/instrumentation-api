package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Knetic/govaluate"
	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Request types enum
const (
	byTimeseries = iota
	byInstrument
	byInstrumentGroup
	explorer
)

func ListTimeseriesMeasurementsByTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tsID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		f := models.MeasurementsFilter{TimeseriesID: &tsID}

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f, byTimeseries)
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

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f, byInstrument)
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

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f, byInstrumentGroup)
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

		streamMeasurementsHandler := StreamTimeseriesMeasurements(db, &f, explorer)
		return streamMeasurementsHandler(c)
	}
}

// StreamTimeseriesMeasurements emits newline delimited json objects. The buffer flushes to the client
// every 1000 records, plus any remaining records in the buffer when complete
func StreamTimeseriesMeasurements(db *sqlx.DB, f *models.MeasurementsFilter, requestType int) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tw timeseries.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err := tw.SetWindow(a, b); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		f.After = tw.After
		f.Before = tw.Before

		// temporal_resolution defaults to 1 (raw - no downsamping)
		// examples:
		// 		3600 (seconds) will keep one datapoint per hour
		//		1800 will keep one per 30 minutes
		//		900 will keep one per  15 minutes
		// 		60 will keep one per minute
		//		...
		// 		1 (default) will not resample and returns raw data
		trs := c.QueryParam("temporal_resolution")
		if trs == "" {
			f.TemporalResolution = 1
		} else {
			tr, err := strconv.Atoi(trs)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			f.TemporalResolution = tr
		}

		rows, err := models.QueryTimeseriesMeasurementsRows(db, f)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		defer func() {
			if err := rows.Close(); err != nil {
				log.Fatal(err.Error())
			}
		}()

		stream := c.Request().Header.Get("Accept") == "application/x-ndjson"

		var enc *json.Encoder
		var mrc models.MeasurementsResponseCollection

		if stream {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
			c.Response().WriteHeader(http.StatusOK)
			enc = json.NewEncoder(c.Response())
		} else {
			mrc = make(models.MeasurementsResponseCollection, 0)
		}

		// LOCF (Last Observation Carried Forward)
		remember := make(map[uuid.UUID]map[string]interface{})
		rowsInChunk := 0

		for rows.Next() {
			// Buffer is chunked to send for every 1000 records
			if stream && rowsInChunk > 0 && rowsInChunk%1000 == 0 {
				c.Response().Flush()
				rowsInChunk = 0
			}
			var mfr models.MeasurementsFromRow

			if err = rows.StructScan(&mfr); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			var env map[string]interface{}
			if err = mfr.MeasurementsJSON.AssignTo(&env); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			// Simply stream stored timeseries and continue
			if !mfr.IsComputed {
				val, exists := env["value"]
				if !exists {
					log.Warnf("bad measurements_json %v for row with date %s and timeseries_id %v", mfr.MeasurementsJSON, mfr.Time, mfr.TimeseriesID)
					continue
				}

				val64, ok := val.(float64)
				if !ok {
					log.Warnf("unable to convert %v interface{} to float64", val)
					continue
				}

				mmt := models.Measurement{Time: mfr.Time, Value: val64}
				tsn := models.TimeseriesNote{}

				if masked, exists := env["masked"]; exists {
					maskedBool, ok := masked.(bool)
					if !ok {
						log.Warnf("unable to convert %v interface{} to bool", masked)
					}
					tsn.Masked = maskedBool
				}
				if validated, exists := env["validated"]; exists {
					validatedBool, ok := validated.(bool)
					if !ok {
						log.Warnf("unable to convert %v interface{} to bool", validated)
					}
					tsn.Validated = validatedBool
				}
				if annotation, exists := env["annotation"]; exists {
					annotationStr, ok := annotation.(string)
					if !ok {
						log.Warnf("unable to convert %v interface{} to string", annotation)
					}
					tsn.Annotation = annotationStr
				}

				mr := models.MeasurementsResponse{
					InstrumentID:   mfr.InstrumentID,
					TimeseriesID:   mfr.TimeseriesID,
					Measurement:    mmt,
					TimeseriesNote: tsn,
				}
				if stream {
					if err := enc.Encode(mr); err != nil {
						return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
					}
				} else {
					mrc = append(mrc, mr)
				}
				rowsInChunk++
				continue
			}

			// Carry forward any values that don't exist for this timestamp with the last known observation
			if _, exists := remember[mfr.TimeseriesID]; !exists {
				remember[mfr.TimeseriesID] = make(map[string]interface{})
			}
			for k, v := range remember[mfr.TimeseriesID] {
				if _, exists := env[k]; !exists {
					env[k] = v
				}
			}
			// Add/Update the most recent values
			for k, v := range env {
				remember[mfr.TimeseriesID][k] = v
			}

			expr, err := govaluate.NewEvaluableExpression(mfr.Formula)
			if err != nil {
				log.Warn(err.Error())
				return err
			}

			val, err := expr.Evaluate(env)
			if err != nil {
				// Any evaluation errors are passed back to client
				// TODO: Apply once UI appropriately filters errors, as to not incorrectly plot 0 values

				// mmt := models.Measurement{Time: mfr.Time, Error: err.Error()}
				// mr := models.MeasurementsResponse{InstrumentID: mfr.InstrumentID, TimeseriesID: mfr.TimeseriesID, Measurement: mmt}
				// if stream {
				// 	if err := enc.Encode(mr); err != nil {
				// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				// 	}
				// } else {
				// 	mrc = append(mrc, mr)
				// }
				// rowsInChunk++
				continue
			}

			val64, ok := val.(float64)
			if !ok {
				log.Warnf("unable to convert %v interface{} to float64", val)
				continue
			}

			mmt := models.Measurement{Time: mfr.Time, Value: val64}
			mr := models.MeasurementsResponse{InstrumentID: mfr.InstrumentID, TimeseriesID: mfr.TimeseriesID, Measurement: mmt}

			if stream {
				if err := enc.Encode(mr); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			} else {
				mrc = append(mrc, mr)
			}
			rowsInChunk++
		}

		// Send any remianing records
		if stream {
			if rowsInChunk > 0 {
				c.Response().Flush()
			}
			return nil
		}

		var resBody interface{}

		if requestType == byTimeseries {
			resBody, err = mrc.CollectSingleTimeseries()
			if err != nil {
				if err.Error() == "no rows" {
					return c.JSON(
						http.StatusOK,
						timeseries.MeasurementCollection{
							TimeseriesID: *f.TimeseriesID,
							Items:        make([]timeseries.Measurement, 0),
						},
					)
				}
				return err
			}
		} else {
			resBody, err = mrc.GroupByInstrument()
			if err != nil {
				if err.Error() == "no rows" {
					return c.JSON(http.StatusOK, make([]map[string]interface{}, 0))
				}
				return err
			}
		}

		return c.JSON(http.StatusOK, resBody)
	}
}
