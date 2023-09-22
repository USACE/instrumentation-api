package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"
	"github.com/USACE/instrumentation-api/api/internal/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListDataLoggers
func ListDataLoggers(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID := c.QueryParam("project_id")
		if pID != "" {

			pID, err := uuid.Parse(pID)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
			}

			dls, err := models.ListProjectDataLoggers(db, &pID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return echo.NewHTTPError(http.StatusOK, dls)
		}

		dls, err := models.ListAllDataLoggers(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, dls)
	}
}

// CreateDataLogger
func CreateDataLogger(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		n := models.DataLogger{}
		if err := c.Bind(&n); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		p := c.Get("profile").(*models.Profile)
		n.Creator = p.ID

		if n.Name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "valid `name` field required")
		}
		// Generate unique slug
		slug, err := utils.CreateUniqueSlug(db, `SELECT slug FROM datalogger`, n.Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}
		n.Slug = slug

		model, err := models.GetDataLoggerModel(db, &n.ModelID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("data logger model id %s not found", n.ModelID))
		}

		// check if datalogger with model and sn already exists and is not deleted
		exists, err := models.DataLoggerActive(db, model, n.SN)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if exists {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				"active data logger model with this model and serial number already exist",
			)
		}

		dl, err := models.CreateDataLogger(db, &n)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}

		return c.JSON(http.StatusCreated, dl)
	}
}

// CycleDataLoggerKey
func CycleDataLoggerKey(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		u := models.DataLogger{ID: dlID}

		if err := models.VerifyDataLoggerExists(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		profile := c.Get("profile").(*models.Profile)
		t := time.Now()
		u.Updater, u.UpdateDate = &profile.ID, &t

		dl, err := models.CycleDataLoggerKey(db, &u)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}

		return c.JSON(http.StatusOK, dl)
	}
}

// GetDataLogger
func GetDataLogger(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		dl, err := models.GetDataLogger(db, &dlID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
		}

		return c.JSON(http.StatusOK, dl)
	}
}

// UpdateDataLogger
func UpdateDataLogger(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		u := models.DataLogger{ID: dlID}
		if err := c.Bind(&u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if dlID != u.ID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`id`"))
		}

		if err := models.VerifyDataLoggerExists(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		profile := c.Get("profile").(*models.Profile)
		t := time.Now()
		u.Updater, u.UpdateDate = &profile.ID, &t

		dlUpdated, err := models.UpdateDataLogger(db, &u)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, dlUpdated)
	}
}

// DeleteDataLogger
func DeleteDataLogger(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		if err := models.VerifyDataLoggerExists(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		d := models.DataLogger{ID: dlID}
		profile := c.Get("profile").(*models.Profile)
		t := time.Now()
		d.Updater, d.UpdateDate = &profile.ID, &t

		if err := models.DeleteDataLogger(db, &d); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"id": dlID})
	}
}

func GetDataLoggerPreview(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return err
		}

		// Get preview from db
		preview, err := models.GetDataLoggerPreview(db, &dlID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, preview)
	}
}
