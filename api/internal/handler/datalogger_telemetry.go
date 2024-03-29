package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

const preparse = "preparse"

// CreateOrUpdateDataloggerMeasurements creates or updates measurements for a timeseries using an equivalency table
func (h *TelemetryHandler) CreateOrUpdateDataloggerMeasurements(c echo.Context) error {
	modelName := c.Param("model")
	if modelName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing route param `model`")
	}
	sn := c.Param("sn")
	if sn == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing route param `sn`")
	}

	ctx := c.Request().Context()

	// Make sure datalogger is active
	dl, err := h.DataloggerTelemetryService.GetDataloggerByModelSN(ctx, modelName, sn)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	body := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return err
	}
	rawJSON, err := json.Marshal(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var prv model.DataloggerTablePreview
	if err := prv.Preview.Set(rawJSON); err != nil {
		return err
	}
	prv.UpdateDate = time.Now()

	if _, err := h.DataloggerTelemetryService.UpdateDataloggerTablePreview(ctx, dl.ID, preparse, prv); err != nil {
		return err
	}

	if modelName == "CR6" || modelName == "CR1000X" {
		cr6Handler := getCR6Handler(h, dl, rawJSON)
		return cr6Handler(c)
	}

	return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
}

// getCR6Handler handles parsing and uploading of Campbell Scientific CR6 measurement payloads
// File format must adhere to "CSIJSON" schema, which can be referenced in the CRBASIC documentation:
//
// CSIJSON Output Format: https://help.campbellsci.com/crbasic/cr350/#parameters/mqtt_outputformat.htm?Highlight=CSIJSON
//
// HTTPPost: https://help.campbellsci.com/crbasic/cr350/#Instructions/httppost.htm?Highlight=httppost
func getCR6Handler(h *TelemetryHandler, dl model.Datalogger, rawJSON []byte) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Errors are cellected and sent to datalogger preview for debugging since datalogger clients cannot parse responses
		em := make([]string, 0)
		ctx := c.Request().Context()
		tn := "preparse"

		// Since these HTTP responses will be returned to datalogger clients, this operates on a "best effort" basis
		// to collect logs to be previewed in the core web application. The error code returned to the client datalogger
		// will sill be relavent to the arm of control flow that raised it.
		defer func() {
			if err := h.DataloggerTelemetryService.UpdateDataloggerTableError(ctx, dl.ID, &tn, &model.DataloggerError{Errors: em}); err != nil {
				log.Printf(err.Error())
			}
		}()

		// Upload Datalogger Measurements
		var pl model.DataloggerPayload
		if err := json.Unmarshal(rawJSON, &pl); err != nil {
			em = append(em, fmt.Sprintf("%d: %s", http.StatusBadRequest, err.Error()))
			return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
		}

		// Check sn from route param matches sn in request body
		if dl.SN != pl.Head.Environment.SerialNo {
			em = append(em, fmt.Sprintf("%d: %s", http.StatusBadRequest, fmt.Sprint(message.MatchRouteParam("`sn`"), dl.SN)))
			return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
		}
		// Check sn from route param matches model in request body
		if *dl.Model != pl.Head.Environment.Model {
			em = append(em, fmt.Sprintf("%d: %s", http.StatusBadRequest, fmt.Sprint(message.MatchRouteParam("`model`"), *dl.Model)))
			return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
		}

		// reroute deferred errors and previews to respective table
		tn = pl.Head.Environment.TableName

		var prv model.DataloggerTablePreview
		if err := prv.Preview.Set(rawJSON); err != nil {
			return err
		}
		prv.UpdateDate = time.Now()

		tableID, err := h.DataloggerTelemetryService.UpdateDataloggerTablePreview(ctx, dl.ID, tn, prv)
		if err != nil {
			em = append(em, fmt.Sprintf("%d: %s", http.StatusInternalServerError, err.Error()))
			return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
		}

		eqt, err := h.EquivalencyTableService.GetEquivalencyTable(ctx, tableID)
		if err != nil {
			em = append(em, fmt.Sprintf("%d: %s", http.StatusInternalServerError, err.Error()))
			return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
		}

		eqtFields := make(map[string]model.EquivalencyTableRow)
		for _, r := range eqt.Rows {
			eqtFields[r.FieldName] = model.EquivalencyTableRow{
				TimeseriesID: r.TimeseriesID,
				InstrumentID: r.InstrumentID,
			}
		}

		fields := pl.Head.Fields
		mcs := make([]model.MeasurementCollection, len(fields))

		// Error if there is no field name in equivalency table to map the field name in the raw payload to
		// delete the keys that were used, check for any dangling afterwards
		for i, f := range fields {
			// Map field to timeseries id
			row, exists := eqtFields[f.Name]
			if !exists {
				em = append(em, fmt.Sprintf("field '%s' from datalogger does not exist in equivalency table", f.Name))
				continue
			}
			if row.InstrumentID == nil {
				em = append(em, fmt.Sprintf("field '%s' not mapped to instrument in equivalency table", f.Name))
				delete(eqtFields, f.Name)
				continue
			}
			if row.TimeseriesID == nil {
				em = append(em, fmt.Sprintf("field '%s' not mapped to timeseries in equivalency table", f.Name))
				delete(eqtFields, f.Name)
				continue
			}

			// collect measurements
			items := make([]model.Measurement, len(pl.Data))
			for j, d := range pl.Data {
				// To avoid complications of daylight savings and related issues,
				// all incoming datalogger timestamps are expected to be in UTC
				t, err := time.Parse("2006-01-02T15:04:05", d.Time)
				if err != nil {
					em = append(em, fmt.Sprintf("unable to parse timestamp for field '%s': %s", f.Name, err.Error()))
					delete(eqtFields, f.Name)
					continue
				}
				v := float64(d.Vals[i])
				if math.IsNaN(v) || math.IsInf(v, 1) {
					em = append(em, fmt.Sprintf("unable to upload '%s' at %s: %.f", f.Name, t, v))
					// don't upload nan or inf
					delete(eqtFields, f.Name)
					continue
				}
				items[j] = model.Measurement{TimeseriesID: *row.TimeseriesID, Time: t, Value: v}
			}

			mcs[i] = model.MeasurementCollection{TimeseriesID: *row.TimeseriesID, Items: items}

			delete(eqtFields, f.Name)
		}

		// This map should be empty if all fields are mapped, otherwise the error is added
		for eqtName := range eqtFields {
			em = append(em, fmt.Sprintf("field '%s' in equivalency table does not match any fields from datalogger", eqtName))
		}

		if _, err = h.MeasurementService.CreateOrUpdateTimeseriesMeasurements(ctx, mcs); err != nil {
			em = append(em, fmt.Sprintf("%d: %s", http.StatusInternalServerError, err.Error()))
			return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"model": *dl.Model, "sn": dl.SN})
	}
}
