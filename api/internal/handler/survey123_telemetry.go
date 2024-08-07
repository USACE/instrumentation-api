package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *TelemetryHandler) CreateOrUpdateSurvey123Measurements(c echo.Context) error {
	survey123ID, err := uuid.Parse(c.Param("survey123_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	var raw map[string]json.RawMessage
	if err := c.Bind(&raw); err != nil {
		return httperr.MalformedBody(err)
	}

	previewRaw, err := json.Marshal(raw)
	if err != nil {
		return httperr.MalformedBody(err)
	}

	t := time.Now()
	pv := model.Survey123Preview{Survey123ID: survey123ID, Preview: previewRaw, UpdateDate: &t}

	ctx := c.Request().Context()
	if err := h.Survey123Service.CreateOrUpdateSurvey123Preview(ctx, pv); err != nil {
		return httperr.InternalServerError(err)
	}

	eq, err := h.Survey123Service.ListSurvey123EquivalencyTableRows(ctx, survey123ID)
	if err != nil {
		return httperr.ServerErrorOrNotFound(err)
	}

	var et string
	if err := json.Unmarshal(raw["eventType"], &et); err != nil {
		return httperr.MalformedBody(err)
	}

	var se []model.Survey123Edits
	if err := json.Unmarshal(raw["applyEdits"], &se); err != nil {
		return httperr.MalformedBody(err)
	}

	sp := model.Survey123Payload{EventType: et, Edits: se}

	if err := h.Survey123Service.CreateOrUpdateSurvey123Measurements(ctx, survey123ID, sp, eq); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.NoContent(http.StatusCreated)
}
