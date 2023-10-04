package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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

	if err := h.DataloggerService.VerifyDataloggerExists(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := h.EquivalencyTableService.UpdateEquivalencyTable(c.Request().Context(), &t); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	eqt, err := h.EquivalencyTableService.GetEquivalencyTable(c.Request().Context(), dlID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, eqt)
}

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
