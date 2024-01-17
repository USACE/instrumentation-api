package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

// ListPlotConfigs godoc
//
//	@Summary lists plot configs
//	@Tags plot-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.PlotConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/plot_configurations [get]
func (h *ApiHandler) ListPlotConfigs(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cc, err := h.PlotConfigService.ListPlotConfigs(c.Request().Context(), pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cc)
}

// GetPlotConfig godoc
//
//	@Sumary gets a single plot configuration by id
//	@Tags plot-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param plot_configuration_id path string true "plot config uuid" Format(uuid)
//	@Success 200 {object} model.PlotConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/plot_configurations/{plot_configuration_id} [get]
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

// CreatePlotConfig godoc
//
//	@Sumary adds a plot configuration to a project
//	@Tags plot-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param plot_config body model.PlotConfig true "plot config payload"
//	@Success 200 {object} model.PlotConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/plot_configurations [post]
//	@Security Bearer
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
	p := c.Get("profile").(model.Profile)
	pc.Creator, pc.CreateDate = p.ID, time.Now()

	pcNew, err := h.PlotConfigService.CreatePlotConfig(c.Request().Context(), pc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, pcNew)
}

// UpdatePlotConfig godoc
//
//	@Sumary updates a plot configuration in a project
//	@Tags plot-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param plot_configuration_id path string true "plot config uuid" Format(uuid)
//	@Param plot_config body model.PlotConfig true "plot config payload"
//	@Success 200 {object} model.PlotConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/plot_configurations/{plot_configuration_id} [put]
//	@Security Bearer
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

	pcUpdated, err := h.PlotConfigService.UpdatePlotConfig(c.Request().Context(), pc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pcUpdated)
}

// DeletePlotConfig godoc
//
//	@Sumary deletes a plot configuration in a project
//	@Tags plot-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param plot_configuration_id path string true "plot config uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/plot_configurations/{plot_configuration_id} [delete]
//	@Security Bearer
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
