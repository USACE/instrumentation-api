package handler

import (
	"fmt"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
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
//	@Param key query string false "api key"
//	@Success 200 {array} model.EquivalencyTable
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table [get]
//	@Security Bearer
func (h *ApiHandler) GetEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	dataloggerTableID, err := uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return httperr.ServerErrorOrNotFound(err)
	}

	t, err := h.EquivalencyTableService.GetEquivalencyTable(ctx, dataloggerTableID)
	if err != nil {
		return httperr.ServerErrorOrNotFound(err)
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
//	@Param key query string false "api key"
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
		return httperr.MalformedID(err)
	}

	t := model.EquivalencyTable{DataloggerID: dlID}
	if err := c.Bind(&t); err != nil {
		return httperr.MalformedBody(err)
	}

	var dataloggerTableID uuid.UUID
	tableIDParam := c.Param("datalogger_table_id")

	ctx := c.Request().Context()

	if tableIDParam != "" {
		dataloggerTableID, err = uuid.Parse(tableIDParam)
		if err != nil {
			return httperr.MalformedID(err)
		}
	} else {
		if t.DataloggerTableName == "" {
			return httperr.Message(http.StatusBadRequest, "payload must contain datalogger_table_name field")
		}
		dataloggerTableID, err = h.DataloggerService.GetOrCreateDataloggerTable(ctx, dlID, t.DataloggerTableName)
		if err != nil {
			httperr.InternalServerError(err)
		}
	}

	t.DataloggerID = dlID
	t.DataloggerTableID = dataloggerTableID

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return httperr.ServerErrorOrNotFound(err)
	}

	if err := h.EquivalencyTableService.GetIsValidDataloggerTable(ctx, dataloggerTableID); err != nil {
		return httperr.Message(http.StatusBadRequest, fmt.Sprintf("invalid datalogger table %s %s", t.DataloggerID, t.DataloggerTableName))
	}

	eqt, err := h.EquivalencyTableService.CreateOrUpdateEquivalencyTable(ctx, t)
	if err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} model.EquivalencyTable
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table [put]
//	@Security Bearer
func (h *ApiHandler) UpdateEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	dataloggerTableID, err := uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	t := model.EquivalencyTable{DataloggerID: dlID, DataloggerTableID: dataloggerTableID}
	if err := c.Bind(&t); err != nil {
		return httperr.MalformedBody(err)
	}

	t.DataloggerID = dlID
	t.DataloggerTableID = dataloggerTableID

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return httperr.ServerErrorOrNotFound(err)
	}

	eqtUpdated, err := h.EquivalencyTableService.UpdateEquivalencyTable(ctx, t)
	if err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	dataloggerTableID, err := uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return httperr.ServerErrorOrNotFound(err)
	}

	if err := h.DataloggerService.DeleteDataloggerTable(ctx, dataloggerTableID); err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/equivalency_table/row/{row_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteEquivalencyTableRow(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	_, err = uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	rowID, err := uuid.Parse(c.Param("row_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return httperr.ServerErrorOrNotFound(err)
	}

	if err := h.EquivalencyTableService.DeleteEquivalencyTableRow(ctx, rowID); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"row_id": rowID})
}
