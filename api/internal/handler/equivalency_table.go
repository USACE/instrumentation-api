package handler

import (
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
//	@Success 200 {array} model.EquivalencyTable
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/equivalency_table [get]
//	@Security Bearer
func (h *ApiHandler) GetEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	if err := h.DataloggerService.VerifyDataloggerExists(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	t, err := h.EquivalencyTableService.GetEquivalencyTable(c.Request().Context(), dlID)
	if err != nil {
		return c.JSON(http.StatusNotFound, t)
	}

	return c.JSON(http.StatusOK, t)
}

// CreateEquivalencyTable godoc
//
//	@Summary creates an equivalency table for a datalogger
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param equivalency_table body model.EquivalencyTable true "equivalency table payload"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
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

	if dlID != t.DataloggerID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`datalogger_id`"))
	}

	if err := h.DataloggerService.VerifyDataloggerExists(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := h.EquivalencyTableService.CreateEquivalencyTable(c.Request().Context(), t); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"datalogger_id": dlID})
}

// UpdateEquivalencyTable godoc
//
//	@Summary updates an equivalency table for a datalogger
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param equivalency_table body model.EquivalencyTable true "equivalency table payload"
//	@Success 200 {object} model.EquivalencyTable
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/equivalency_table [put]
//	@Security Bearer
func (h *ApiHandler) UpdateEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	t := model.EquivalencyTable{DataloggerID: dlID}
	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if dlID != t.DataloggerID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`datalogger_id`"))
	}

	ctx := c.Request().Context()

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	eqtUpdated, err := h.EquivalencyTableService.UpdateEquivalencyTable(ctx, dlID, t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, eqtUpdated)
}

// DeleteEquivalencyTable godoc
//
//	@Summary deletes an equivalency table
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/equivalency_table [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteEquivalencyTable(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	if err := h.DataloggerService.VerifyDataloggerExists(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := h.EquivalencyTableService.DeleteEquivalencyTable(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"datalogger_id": dlID})
}

// DeleteEquivalencyTableRow godoc
//
//	@Summary deletes an equivalency table row
//	@Tags equivalency-table
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param id query string true "equivalency table row uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/equivalency_table/row [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteEquivalencyTableRow(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	rID, err := uuid.Parse(c.QueryParam("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	if err := h.DataloggerService.VerifyDataloggerExists(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := h.EquivalencyTableService.DeleteEquivalencyTableRow(c.Request().Context(), dlID, rID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"row_id": rID})
}
