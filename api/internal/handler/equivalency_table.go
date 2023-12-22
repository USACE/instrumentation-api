package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetEquivalencyTable godoc
//
//	@Summary gets an equivalency table for a datalogger
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param datalogger_table_id path string true "datalogger table uuid" Format(uuid)
//	@Success 200 {array} model.EquivalencyTable
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table [get]
//	@Security Bearer
func (h *ApiHandler) GetEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	dataloggerTableID, err := uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	t, err := h.EquivalencyTableService.GetEquivalencyTable(ctx, dataloggerTableID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, message.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, t)
}

// CreateEquivalencyTable godoc
//
//	@Summary creates an equivalency table for a datalogger and auto create data logger table if not exists
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param datalogger_table_id path string true "datalogger table uuid" Format(uuid)
//	@Param equivalency_table body model.EquivalencyTable true "equivalency table payload"
//	@Success 200 {object} model.EquivalencyTable
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table [post]
//	@Router /datalogger/{datalogger_id}/equivalency_table [post]
//	@Security Bearer
func (h *ApiHandler) CreateEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	t := model.EquivalencyTable{DataloggerID: dlID}
	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var dataloggerTableID uuid.UUID
	tableIDParam := c.Param("datalogger_table_id")

	ctx := c.Request().Context()

	if tableIDParam != "" {
		dataloggerTableID, err = uuid.Parse(tableIDParam)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
		}
	} else {
		if t.DataloggerTableName == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "payload must contain datalogger_table_name field")
		}
		dataloggerTableID, err = h.DataloggerService.GetOrCreateDataloggerTable(ctx, dlID, t.DataloggerTableName)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	t.DataloggerID = dlID
	t.DataloggerTableID = dataloggerTableID

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := h.EquivalencyTableService.GetIsValidDataloggerTable(ctx, dataloggerTableID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid datalogger table %s %s", t.DataloggerID, t.DataloggerTableName))
	}

	eqt, err := h.EquivalencyTableService.CreateEquivalencyTable(ctx, t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, eqt)
}

// UpdateEquivalencyTable godoc
//
//	@Summary updates an equivalency table for a datalogger
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param datalogger_table_id path string true "datalogger table uuid" Format(uuid)
//	@Param equivalency_table body model.EquivalencyTable true "equivalency table payload"
//	@Success 200 {object} model.EquivalencyTable
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table [put]
//	@Security Bearer
func (h *ApiHandler) UpdateEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	dataloggerTableID, err := uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	t := model.EquivalencyTable{DataloggerID: dlID, DataloggerTableID: dataloggerTableID}
	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if dlID != t.DataloggerID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`datalogger_id`"))
	}
	if dataloggerTableID != t.DataloggerTableID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`datalogger_table_id`"))
	}

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	eqtUpdated, err := h.EquivalencyTableService.UpdateEquivalencyTable(ctx, t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, eqtUpdated)
}

// DeleteEquivalencyTable godoc
//
//	@Summary deletes an equivalency table and corresponding datalogger table
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param datalogger_table_id path string true "datalogger table uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	dataloggerTableID, err := uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := h.DataloggerService.DeleteDataloggerTable(ctx, dataloggerTableID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"datalogger_id": dlID, "datalogger_table_id": dataloggerTableID})
}

// DeleteEquivalencyTableRow godoc
//
//	@Summary deletes an equivalency table row
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param datalogger_table_id path string true "datalogger table uuid" Format(uuid)
//	@Param row_id path string true "equivalency table row uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table/row/{row_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteEquivalencyTableRow(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	_, err = uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	rowID, err := uuid.Parse(c.Param("row_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := h.EquivalencyTableService.DeleteEquivalencyTableRow(ctx, rowID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"row_id": rowID})
}
