package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/passwords"
	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// TODO: Finish implementation
func CreateOrUpdateDataLoggerMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the API key from the header, make sure its hash is in the database
		// This could also be done in the middleware but might not be as condusive
		// to change if we move to another API not using echo

		// Bind request payload
		var dlp models.DataLoggerPayload
		if err := c.Bind(&dlp); err != nil {
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}

		sn := dlp.Head.Environment.SerialNo

		// Check that data logger exists
		_, err := models.GetDataLoggerBySerialNumber(db, sn)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}

		// Check header for api key
		ak, exists := c.Request().Header["X-Api-Key"]

		if !exists || len(ak) != 1 {
			// Missing API key header
			return c.JSON(http.StatusUnauthorized, models.DefaultMessageUnauthorized)
		}

		// Get data logger hash
		hash, err := models.GetDataLoggerHash(db, sn)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.DefaultMessageUnauthorized)
		}

		// Check that API Key exists in database
		if match, err := passwords.ComparePasswordAndHash(ak[0], hash); err != nil || !match {
			return c.JSON(http.StatusUnauthorized, models.DefaultMessageUnauthorized)
		}

		fields := dlp.Head.Fields
		eq, err := models.GetEquivalencyTable(db, sn)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.DefaultMessageInternalServerError)
		}

		mcs := make([]timeseries.MeasurementCollection, len(fields))

		for i, f := range fields {
			// Map field to timeseries id
			row, exists := eq.FieldMap[f.Name]
			if !exists {
				// TODO: Update validation status
				continue
			}

			// collect measurements
			items := make([]timeseries.Measurement, len(dlp.Data))
			for j, d := range dlp.Data {
				t, err := time.Parse(time.RFC3339, d.Time)
				if err != nil {
					// TODO: Hanlde error parsing time
					continue
				}
				items[j] = timeseries.Measurement{TimeseriesID: row.TimeseriesID, Time: t, Value: d.Vals[i]}
			}

			mcs[i] = timeseries.MeasurementCollection{TimeseriesID: row.TimeseriesID, Items: items}
		}

		ret, err := models.CreateOrUpdateTimeseriesMeasurements(db, mcs)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}

		return c.JSON(http.StatusCreated, &ret)
	}
}

func UpdateDataLoggerPreview(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sn := c.Param("sn")
		if sn == "" {
			return c.String(http.StatusBadRequest, models.DefaultMessageBadRequest.Message)
		}

		body := make(map[string]interface{})

		err := json.NewDecoder(c.Request().Body).Decode(&body)
		if err != nil {
			return err
		}

		pl, err := json.Marshal(body)
		if err != nil {
			return c.String(http.StatusBadRequest, models.DefaultMessageBadRequest.Message)
		}

		dlp := models.DataLoggerPreview{SN: sn}
		dlp.Payload.Set(pl)

		err = models.UpdateDataLoggerPreview(db, &dlp)
		if err != nil {
			return c.String(http.StatusInternalServerError, models.DefaultMessageInternalServerError.Message)
		}

		return c.JSON(http.StatusAccepted, sn)
	}
}

func GetDataLoggerPreview(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sn := c.Param("sn")
		if sn == "" {
			return c.String(http.StatusBadRequest, models.DefaultMessageBadRequest.Message)
		}

		// Get preview from db
		preview, err := models.GetDataLoggerPreview(db, sn)
		if err != nil {
			return c.String(http.StatusInternalServerError, models.DefaultMessageInternalServerError.Message)
		}

		return c.JSON(http.StatusOK, preview)
	}
}
