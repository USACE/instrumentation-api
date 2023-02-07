package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListDataLoggers
func ListDataLoggers(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Check user has datalogger role permissions

		// Check for project_id in url query
		pID := c.QueryParam("project_id")
		if pID != "" {

			pID, err := uuid.Parse(pID)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// TODO: Check if user has permissions to project

			dls, err := models.ListProjectDataLoggers(db, &pID)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
			}

			return echo.NewHTTPError(http.StatusOK, dls)
		}

		dls, err := models.ListAllDataLoggers(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
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

		// TODO: Check user has datalogger role permissions

		if n.Name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "valid `name` field required")
		}
		// Generate unique slug
		slug, err := dbutils.CreateUniqueSlug(db, `SELECT slug FROM datalogger`, n.Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}
		n.Slug = slug

		// check if datalogger with sn already exists and is not deleted
		err = models.VerifyUniqueSN(db, n.SN)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
		// TODO: Check user has datalogger role permissions

		dl, err := models.GetDataLogger(db, &dlID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
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

		// check id in url params matches id in request body
		if dlID != u.ID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`id`"))
		}

		// TODO: Check user has datalogger role permissions

		profile := c.Get("profile").(*models.Profile)
		t := time.Now()
		u.Updater, u.UpdateDate = &profile.ID, &t

		dlUpdated, err := models.UpdateDataLogger(db, &u)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
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

		d := models.DataLogger{ID: dlID}
		profile := c.Get("profile").(*models.Profile)
		t := time.Now()
		d.Updater, d.UpdateDate = &profile.ID, &t

		// TODO: Check user has datalogger role permissions
		if err := models.DeleteDataLogger(db, &d); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}

		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

func GetEquivalencyTable(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		t, err := models.GetEquivalencyTable(db, &dlID)
		if err != nil {
			return c.JSON(http.StatusNotFound, t)
		}

		return c.JSON(http.StatusOK, t)
	}
}

func CreateOrUpdateEquivalencyTable(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		t := models.EquivalencyTable{DataLoggerID: dlID}
		if err := c.Bind(&t); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// check id in url params matches id in request body
		if dlID != t.DataLoggerID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`datalogger_id`"))
		}

		for _, r := range t.Rows {
			_, err := models.CreateOrUpdateEquivalencyTableRow(db, &r)
			if err != nil {
				log.Printf("error updating table row %v: %s", &r, err.Error())
			}
		}

		eqt, err := models.GetEquivalencyTable(db, &dlID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}

		return c.JSON(http.StatusOK, eqt)
	}
}

func DeleteEquivalencyTable(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		if err := models.DeleteEquivalencyTable(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}

		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

func DeleteEquivalencyTableRow(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		field := c.QueryParam("field_name")
		if field == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing field_name query parameter")
		}

		if err := models.DeleteEquivalencyTableRow(db, &dlID, field); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}

		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

func GetDataLoggerPreview(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sn := c.Param("sn")
		if sn == "" {
			return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
		}

		// Get preview from db
		preview, err := models.GetDataLoggerPreview(db, sn)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}

		return c.JSON(http.StatusOK, preview)
	}
}
