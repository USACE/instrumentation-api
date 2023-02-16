package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// CreateOrUpdateDataLoggerMeasurements creates or updates measurements for a timeseires
// that a datalogger is mapped to using the DataLoggerEquivalencyTable
//
// DataLoggerKeyAuth middleware is applied to the group where the corresponding route
// to this handler is configured
func CreateOrUpdateDataLoggerMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		model := c.Param("model")
		if model == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing route param `model`")
		}
		sn := c.Param("sn")
		if sn == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing route param `sn`")
		}

		// Make sure data logger is active
		dl, err := models.GetDataLoggerByModelSN(db, model, sn)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		body := make(map[string]interface{})
		err = json.NewDecoder(c.Request().Body).Decode(&body)
		if err != nil {
			return err
		}
		raw, err := json.Marshal(body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		prv := models.DataLoggerPreview{DataLoggerID: dl.ID}
		prv.Preview.Set(raw)
		prv.UpdateDate = time.Now()

		err = models.UpdateDataLoggerPreview(db, &prv)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if model == "CR6" || model == "CR1000X" {
			cr6Handler := getCR6Handler(db, dl, &raw)
			return cr6Handler(c)
		}

		return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
	}
}

// getCR6Handler handles parsing and uploading of Campbell Scientific CR6 measurement payloads
// File format must adhere to "CSIJSON" schema, which can be referenced in the CRBASIC documentation:
//
// CSIJSON Output Format: https://help.campbellsci.com/crbasic/cr350/#parameters/mqtt_outputformat.htm?Highlight=CSIJSON
//
// HTTPPost: https://help.campbellsci.com/crbasic/cr350/#Instructions/httppost.htm?Highlight=httppost
func getCR6Handler(db *sqlx.DB, dl *models.DataLogger, rawJSON *[]byte) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Upload DataLogger Measurements
		var pl models.DataLoggerPayload
		if err := json.Unmarshal(*rawJSON, &pl); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Check sn from route param matches sn in request body
		if dl.SN != pl.Head.Environment.SerialNo {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprint(messages.MatchRouteParam("`sn`"), dl.SN))
		}
		// Check sn from route param matches model in request body
		if *dl.Model != pl.Head.Environment.Model {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprint(messages.MatchRouteParam("`model`"), *dl.Model))
		}

		fields := pl.Head.Fields
		eqt, err := models.GetEquivalencyTable(db, &dl.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

		_, err = models.CreateOrUpdateTimeseriesMeasurements(db, mcs)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{"model": *dl.Model, "sn": dl.SN})
	}
}
