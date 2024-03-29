package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListProjectEvaluations godoc
//
//	@Summary lists evaluations for a single project optionally filtered by alert_config_id
//	@Tags evaluation
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.Evaluation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/evaluations [get]
func (h *ApiHandler) ListProjectEvaluations(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var ee []model.Evaluation
	if qp := c.QueryParam("alert_config_id"); qp != "" {
		alertConfigID, err := uuid.Parse(qp)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		ee, err = h.EvaluationService.ListProjectEvaluationsByAlertConfig(c.Request().Context(), projectID, alertConfigID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	} else {
		ee, err = h.EvaluationService.ListProjectEvaluations(c.Request().Context(), projectID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, ee)
}

// ListInstrumentEvaluations godoc
//
//	@Summary lists evaluations for a single instrument
//	@Tags evaluation
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.Evaluation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/evaluations [get]
func (h *ApiHandler) ListInstrumentEvaluations(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ee, err := h.EvaluationService.ListInstrumentEvaluations(c.Request().Context(), instrumentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ee)
}

// GetEvaluation godoc
//
//	@Summary gets a single evaluation by id
//	@Tags evaluation
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param evaluation_id path string true "evaluation uuid" Format(uuid)
//	@Success 200 {object} model.Evaluation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/evaluations/{evaluation_id} [get]
func (h *ApiHandler) GetEvaluation(c echo.Context) error {
	acID, err := uuid.Parse(c.Param("evaluation_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ev, err := h.EvaluationService.GetEvaluation(c.Request().Context(), acID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, message.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ev)
}

// CreateEvaluation godoc
//
//	@Summary creates one evaluation
//	@Tags evaluation
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param evaluation body model.Evaluation true "evaluation payload"
//	@Success 200 {object} model.Evaluation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/evaluations [post]
//	@Security Bearer
func (h *ApiHandler) CreateEvaluation(c echo.Context) error {
	ev := model.Evaluation{}
	if err := c.Bind(&ev); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	profile := c.Get("profile").(model.Profile)
	ev.ProjectID, ev.CreatorID, ev.CreateDate = projectID, profile.ID, time.Now()

	evNew, err := h.EvaluationService.CreateEvaluation(c.Request().Context(), ev)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, evNew)
}

// UpdateEvaluation godoc
//
//	@Summary updates an existing evaluation
//	@Tags evaluation
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param evaluation_id path string true "evaluation uuid" Format(uuid)
//	@Param evaluation body model.Evaluation true "evaluation payload"
//	@Success 200 {object} model.Evaluation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/evaluations/{evaluation_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateEvaluation(c echo.Context) error {
	var ev model.Evaluation
	if err := c.Bind(&ev); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	evID, err := uuid.Parse(c.Param("evaluation_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	p := c.Get("profile").(model.Profile)
	t := time.Now()
	ev.UpdaterID, ev.UpdateDate = &p.ID, &t
	evUpdated, err := h.EvaluationService.UpdateEvaluation(c.Request().Context(), evID, ev)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, evUpdated)
}

// DeleteEvaluation godoc
//
//	@Summary deletes an evaluation
//	@Tags evaluation
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param evaluation_id path string true "evaluation uuid" Format(uuid)
//	@Success 200 {array} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/evaluations/{evaluation_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteEvaluation(c echo.Context) error {
	acID, err := uuid.Parse(c.Param("evaluation_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.EvaluationService.DeleteEvaluation(c.Request().Context(), acID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
