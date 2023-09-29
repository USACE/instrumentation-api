package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInstrumentStatus lists all Status for an instrument
func (h ApiHandler) ListInstrumentStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	ss, err := model.ListInstrumentStatus(db, &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, ss)
}

// GetInstrumentStatus returns a single Status
func (h ApiHandler) GetInstrumentStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("status_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	s, err := model.GetInstrumentStatus(db, &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, s)
}

// CreateOrUpdateInstrumentStatus creates a Status for an instrument
func (h ApiHandler) CreateOrUpdateInstrumentStatus(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	var sc model.InstrumentStatusCollection
	if err := c.Bind(&sc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Assign Fresh UUID to each Status
	for idx := range sc.Items {
		id, err := uuid.NewRandom()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		sc.Items[idx].ID = id
	}

	if err := model.CreateOrUpdateInstrumentStatus(db, &instrumentID, sc.Items); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, make(map[string]interface{}))
}

// DeleteInstrumentStatus deletes a Status for an instrument
func (h ApiHandler) DeleteInstrumentStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("status_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	if err := model.DeleteInstrumentStatus(db, &id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
