package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/dbutils"
	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListPlotConfigurations returns plot groups
func ListPlotConfigurations(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		cc, err := models.ListPlotConfigurations(db, &pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &cc)
	}
}

// GetPlotConfiguration returns single instrument group
func GetPlotConfiguration(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Project ID
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		// Plot configuration ID
		cID, err := uuid.Parse(c.Param("plot_configuration_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		// Get the plot configuration
		g, err := models.GetPlotConfiguration(db, &pID, &cID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, g)
	}
}

// CreatePlotConfiguration add plot configuration for a project
func CreatePlotConfiguration(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var pc models.PlotConfiguration
		// Bind Information Provided in Request Body
		if err := c.Bind(&pc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// Default to 1 year if no date range provided
		if pc.DateRange == "" {
			pc.DateRange = "1 year"
		}
		if err := pc.ValidateDateRange(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// Project ID from Route Params
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// Check Project ID in payload vs. Project ID in Route Params
		if pID != pc.ProjectID {
			return echo.NewHTTPError(http.StatusBadRequest, "route parameter project_id does not match project_id in JSON payload")
		}
		// Generate Unique Slug
		slugsTaken, err := models.ListPlotConfigurationSlugs(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		slug, err := dbutils.NextUniqueSlug(pc.Name, slugsTaken)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		pc.Slug = slug

		// Profile of user creating collection group
		p := c.Get("profile").(*models.Profile)
		pc.Creator, pc.CreateDate = p.ID, time.Now()

		// Create Collection Group
		pcNew, err := models.CreatePlotConfiguration(db, &pc)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, pcNew)
	}
}

// UpdatePlotConfiguration updates a plot configuration for a project
func UpdatePlotConfiguration(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var pc models.PlotConfiguration
		// Bind Information Provided in Request Body
		if err := c.Bind(&pc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// Default to 1 year if no date range provided
		if pc.DateRange == "" {
			pc.DateRange = "1 year"
		}
		if err := pc.ValidateDateRange(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// Project ID from Route Params
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// Check Project ID in payload vs. Project ID in Route Params
		if pID != pc.ProjectID {
			return echo.NewHTTPError(http.StatusBadRequest, "route parameter project_id does not match project_id in JSON payload")
		}

		// Profile of user creating Plot Configuration
		p := c.Get("profile").(*models.Profile)
		tNow := time.Now()
		pc.Updater, pc.UpdateDate = &p.ID, &tNow

		// Update Plot Configuration
		pcUpdated, err := models.UpdatePlotConfiguration(db, &pc)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, pcUpdated)
	}
}

// DeletePlotConfiguration delete plot configuration for a project
func DeletePlotConfiguration(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		cID, err := uuid.Parse(c.Param("plot_configuration_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := models.DeletePlotConfiguration(db, &pID, &cID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
