package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetEquivalencyTable(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		if err := models.VerifyDataLoggerExists(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		t, err := models.GetEquivalencyTable(db, &dlID)
		if err != nil {
			return c.JSON(http.StatusNotFound, t)
		}

		return c.JSON(http.StatusOK, t)
	}
}

func CreateEquivalencyTable(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		t := models.EquivalencyTable{DataLoggerID: dlID}
		if err := c.Bind(&t); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if dlID != t.DataLoggerID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`datalogger_id`"))
		}

		if err := models.VerifyDataLoggerExists(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		if err := models.CreateEquivalencyTable(db, &t); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		eqt, err := models.GetEquivalencyTable(db, &dlID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, eqt)
	}
}

func UpdateEquivalencyTable(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dlID, err := uuid.Parse(c.Param("datalogger_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		t := models.EquivalencyTable{DataLoggerID: dlID}
		if err := c.Bind(&t); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if dlID != t.DataLoggerID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`datalogger_id`"))
		}

		if err := models.VerifyDataLoggerExists(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		if err := models.UpdateEquivalencyTable(db, &t); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		eqt, err := models.GetEquivalencyTable(db, &dlID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

		if err := models.VerifyDataLoggerExists(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		if err := models.DeleteEquivalencyTable(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

		rID, err := uuid.Parse(c.QueryParam("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		if err := models.VerifyDataLoggerExists(db, &dlID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		if err := models.DeleteEquivalencyTableRow(db, &dlID, &rID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
