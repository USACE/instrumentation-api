package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

// ListPlotConfigurations returns plot groups
func (h ApiHandler) ListPlotConfigurations(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cc, err := model.ListPlotConfigurations(db, &pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &cc)
}

// GetPlotConfiguration returns single instrument group
func (h ApiHandler) GetPlotConfiguration(c echo.Context) error {
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
	g, err := model.GetPlotConfiguration(db, &pID, &cID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, g)
}

// CreatePlotConfiguration add plot configuration for a project
func (h ApiHandler) CreatePlotConfiguration(c echo.Context) error {

	var pc model.PlotConfiguration
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
	slugsTaken, err := model.ListPlotConfigurationSlugs(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	slug, err := util.NextUniqueSlug(pc.Name, slugsTaken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	pc.Slug = slug

	// Profile of user creating collection group
	p := c.Get("profile").(*model.Profile)
	pc.Creator, pc.CreateDate = p.ID, time.Now()

	// Create Collection Group
	pcNew, err := model.CreatePlotConfiguration(db, &pc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, pcNew)
}

// UpdatePlotConfiguration updates a plot configuration for a project
func (h ApiHandler) UpdatePlotConfiguration(c echo.Context) error {

	var pc model.PlotConfiguration
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
	p := c.Get("profile").(*model.Profile)
	tNow := time.Now()
	pc.Updater, pc.UpdateDate = &p.ID, &tNow

	// Update Plot Configuration
	pcUpdated, err := model.UpdatePlotConfiguration(db, &pc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pcUpdated)
}

// DeletePlotConfiguration delete plot configuration for a project
func (h ApiHandler) DeletePlotConfiguration(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cID, err := uuid.Parse(c.Param("plot_configuration_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := model.DeletePlotConfiguration(db, &pID, &cID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
