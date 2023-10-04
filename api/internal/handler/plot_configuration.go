package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

// ListPlotConfigs returns plot groups
func (h *ApiHandler) ListPlotConfigs(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cc, err := h.PlotConfigService.ListPlotConfigs(c.Request().Context(), pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &cc)
}

// GetPlotConfig returns single instrument group
func (h *ApiHandler) GetPlotConfig(c echo.Context) error {
	cID, err := uuid.Parse(c.Param("plot_configuration_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	g, err := h.PlotConfigService.GetPlotConfig(c.Request().Context(), cID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, g)
}

// CreatePlotConfig add plot configuration for a project
func (h *ApiHandler) CreatePlotConfig(c echo.Context) error {
	var pc model.PlotConfig
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
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if pID != pc.ProjectID {
		return echo.NewHTTPError(http.StatusBadRequest, "route parameter project_id does not match project_id in JSON payload")
	}
	slugsTaken, err := h.PlotConfigService.ListPlotConfigSlugs(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	slug, err := util.NextUniqueSlug(pc.Name, slugsTaken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	pc.Slug = slug

	p := c.Get("profile").(model.Profile)
	pc.Creator, pc.CreateDate = p.ID, time.Now()

	pcNew, err := h.PlotConfigService.CreatePlotConfig(c.Request().Context(), pc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, pcNew)
}

// UpdatePlotConfig updates a plot configuration for a project
func (h *ApiHandler) UpdatePlotConfig(c echo.Context) error {
	var pc model.PlotConfig
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
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if pID != pc.ProjectID {
		return echo.NewHTTPError(http.StatusBadRequest, "route parameter project_id does not match project_id in JSON payload")
	}

	p := c.Get("profile").(model.Profile)
	tNow := time.Now()
	pc.Updater, pc.UpdateDate = &p.ID, &tNow

	if err := h.PlotConfigService.UpdatePlotConfig(c.Request().Context(), pc); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pc)
}

// DeletePlotConfig delete plot configuration for a project
func (h *ApiHandler) DeletePlotConfig(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cID, err := uuid.Parse(c.Param("plot_configuration_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.PlotConfigService.DeletePlotConfig(c.Request().Context(), pID, cID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
