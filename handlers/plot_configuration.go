package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/models"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListPlotConfigurations returns plot groups
func ListPlotConfigurations(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		cc, err := models.ListPlotConfigurations(db, &pID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
			// return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, &cc)
	}
}
