package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/passwords"
	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// TODO: Finish implementation
func CreateOrUpdateDataLoggerMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sn := c.Param("sn")
		if sn == "" {
			return c.JSON(http.StatusBadRequest, messages.BadRequest)
		}

		// Check header for api key
		ak, exists := c.Request().Header["X-Api-Key"]
		if !exists || len(ak) != 1 {
			// Missing API key header
			return c.JSON(http.StatusUnauthorized, messages.Unauthorized)
		}

		// Get data logger hash
		hash, err := models.GetDataLoggerHashBySN(db, sn)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, messages.Unauthorized)
		}

		// Check that API Key exists in database
		if match, err := passwords.ComparePasswordAndHash(ak[0], hash); err != nil || !match {
			return c.JSON(http.StatusUnauthorized, messages.Unauthorized)
		}

		// Datalogger Authenticated
		// Update DataLogger Preview
		body := make(map[string]interface{})

		err = json.NewDecoder(c.Request().Body).Decode(&body)
		if err != nil {
			return err
		}

		raw, err := json.Marshal(body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, messages.BadRequest)
		}

		prv := models.DataLoggerPreview{SN: sn}
		prv.Payload.Set(raw)

		err = models.UpdateDataLoggerPreviewBySN(db, &prv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, messages.InternalServerError)
		}

		// Check that data logger exists
		_, err = models.GetDataLoggerBySN(db, sn)
		if err != nil {
			return c.JSON(http.StatusBadRequest, messages.BadRequest)
		}

		// if dl.Model == "CR6" {
		cr6Handler := getCR6Handler(db, sn)
		return cr6Handler(c)
		// }
	}
}

// getCR6Handler handles parsing and uploading of Campbell Scientific CR6 measurement payloads
func getCR6Handler(db *sqlx.DB, sn string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Upload DataLogger Measurements
		var pl models.DataLoggerPayload
		if err := c.Bind(&pl); err != nil {
			return c.JSON(http.StatusBadRequest, messages.BadRequest)
		}

		// Check sn from route param matches sn in request body
		if sn != pl.Head.Environment.SerialNo {
			return c.JSON(http.StatusBadRequest, messages.MatchRouteParam("`sn`"))
		}

		fields := pl.Head.Fields
		eqt, err := models.GetEquivalencyTableBySN(db, sn)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, messages.InternalServerError)
		}

		eqtFields := make(map[string]models.EquivalencyTableRow)
		for _, r := range eqt.Rows {
			eqtFields[r.FieldName] = models.EquivalencyTableRow{
				TimeseriesID: r.TimeseriesID,
				InstrumentID: r.InstrumentID,
			}
		}

		mcs := make([]timeseries.MeasurementCollection, len(fields))

		for i, f := range fields {
			// Map field to timeseries id
			row, exists := eqtFields[f.Name]
			if !exists {
				// TODO: Update validation status
				continue
			}

			// collect measurements
			items := make([]timeseries.Measurement, len(pl.Data))
			for j, d := range pl.Data {
				t, err := time.Parse(time.RFC3339, d.Time)
				if err != nil {
					// TODO: Handle error parsing time
					continue
				}
				items[j] = timeseries.Measurement{TimeseriesID: *row.TimeseriesID, Time: t, Value: d.Vals[i]}
			}

			mcs[i] = timeseries.MeasurementCollection{TimeseriesID: *row.TimeseriesID, Items: items}
		}

		ret, err := models.CreateOrUpdateTimeseriesMeasurements(db, mcs)
		if err != nil {
			return c.JSON(http.StatusBadRequest, messages.BadRequest)
		}

		return c.JSON(http.StatusCreated, &ret)
	}
}
