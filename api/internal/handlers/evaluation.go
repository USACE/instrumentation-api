package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListProjectEvaluations lists evaluations for a single project optionally filtered by alert_config_id
func ListProjectEvaluations(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		var aa []models.Evaluation
		if qp := c.QueryParam("alert_config_id"); qp != "" {
			alertConfigID, err := uuid.Parse(qp)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			aa, err = models.ListProjectEvaluationsByAlertConfig(db, &projectID, &alertConfigID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		} else {
			aa, err = models.ListProjectEvaluations(db, &projectID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusOK, aa)
	}
}

// ListInstrumentEvaluations lists evaluations for a single instrument
func ListInstrumentEvaluations(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		aa, err := models.ListInstrumentEvaluations(db, &instrumentID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, aa)
	}
}

// GetEvaluation gets a single evaluation
func GetEvaluation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		acID, err := uuid.Parse(c.Param("evaluation_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		a, err := models.GetEvaluation(db, &acID)
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, a)
	}
}

// CreateEvaluation creates one evaluation
func CreateEvaluation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := models.Evaluation{}
		if err := c.Bind(&ac); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		profile := c.Get("profile").(*models.Profile)
		ac.ProjectID, ac.Creator, ac.CreateDate = projectID, profile.ID, time.Now()

		aa, err := models.CreateEvaluation(db, &ac)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, aa)
	}
}

// UpdateEvaluation updates an existing evaluation
func UpdateEvaluation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ac models.Evaluation
		if err := c.Bind(&ac); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		acID, err := uuid.Parse(c.Param("evaluation_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		p := c.Get("profile").(*models.Profile)
		t := time.Now()
		ac.Updater, ac.UpdateDate = &p.ID, &t
		aUpdated, err := models.UpdateEvaluation(db, &acID, &ac)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, &aUpdated)
	}
}

// DeleteEvaluation deletes an evaluation
func DeleteEvaluation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		acID, err := uuid.Parse(c.Param("evaluation_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := models.DeleteEvaluation(db, &acID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
