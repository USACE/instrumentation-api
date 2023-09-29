package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListProjectEvaluations lists evaluations for a single project optionally filtered by alert_config_id
func (h ApiHandler) ListProjectEvaluations(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var aa []model.Evaluation
	if qp := c.QueryParam("alert_config_id"); qp != "" {
		alertConfigID, err := uuid.Parse(qp)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		aa, err = h.EvaluationStore.ListProjectEvaluationsByAlertConfig(c.Request().Context(), projectID, alertConfigID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	} else {
		aa, err = h.EvaluationStore.ListProjectEvaluations(c.Request().Context(), projectID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, aa)
}

// ListInstrumentEvaluations lists evaluations for a single instrument
func (h ApiHandler) ListInstrumentEvaluations(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	aa, err := h.EvaluationStore.ListInstrumentEvaluations(c.Request().Context(), instrumentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, aa)
}

// GetEvaluation gets a single evaluation
func (h ApiHandler) GetEvaluation(c echo.Context) error {
	acID, err := uuid.Parse(c.Param("evaluation_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	a, err := h.EvaluationStore.GetEvaluation(c.Request().Context(), acID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, a)
}

// CreateEvaluation creates one evaluation
func (h ApiHandler) CreateEvaluation(c echo.Context) error {
	ac := model.Evaluation{}
	if err := c.Bind(&ac); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	profile := c.Get("profile").(*model.Profile)
	ac.ProjectID, ac.Creator, ac.CreateDate = projectID, profile.ID, time.Now()

	aa, err := h.EvaluationStore.CreateEvaluation(c.Request().Context(), ac)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, aa)
}

// UpdateEvaluation updates an existing evaluation
func (h ApiHandler) UpdateEvaluation(c echo.Context) error {
	var ac model.Evaluation
	if err := c.Bind(&ac); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	acID, err := uuid.Parse(c.Param("evaluation_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	p := c.Get("profile").(*model.Profile)
	t := time.Now()
	ac.Updater, ac.UpdateDate = &p.ID, &t
	aUpdated, err := h.EvaluationStore.UpdateEvaluation(c.Request().Context(), acID, ac)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, &aUpdated)
}

// DeleteEvaluation deletes an evaluation
func (h ApiHandler) DeleteEvaluation(c echo.Context) error {
	acID, err := uuid.Parse(c.Param("evaluation_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.EvaluationStore.DeleteEvaluation(c.Request().Context(), acID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
