package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreatePlotConfigBullseyePlot godoc
//
//	@Summary adds a bullseye plot configuration to a project
//	@Tags plot-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param plot_config body model.PlotConfigBullseyePlot true "plot config payload"
//	@Param key query string false "api key"
//	@Success 200 {object} model.PlotConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/plot_configs/bullseye_plots [post]
//	@Security Bearer
func (h *ApiHandler) CreatePlotConfigBullseyePlot(c echo.Context) error {
	var pc model.PlotConfigBullseyePlot
	if err := c.Bind(&pc); err != nil {
		return httperr.MalformedBody(err)
	}
	// Default to 1 year if no date range provided
	if pc.DateRange == "" {
		pc.DateRange = "1 year"
	}
	if _, err := pc.DateRangeTimeWindow(); err != nil {
		return httperr.MalformedDate(err)
	}
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	pc.ProjectID = pID

	p := c.Get("profile").(model.Profile)
	pc.CreatorID, pc.CreateDate = p.ID, time.Now()

	pcNew, err := h.PlotConfigService.CreatePlotConfigBullseyePlot(c.Request().Context(), pc)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusCreated, pcNew)
}

// UpdatePlotConfigBullseyePlot godoc
//
//	@Summary updates a bullseye plot configuration in a project
//	@Tags plot-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param plot_configuration_id path string true "plot config uuid" Format(uuid)
//	@Param plot_config body model.PlotConfigBullseyePlot true "plot config payload"
//	@Param key query string false "api key"
//	@Success 200 {object} model.PlotConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/plot_configs/bullseye_plots/{plot_configuration_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdatePlotConfigBullseyePlot(c echo.Context) error {
	var pc model.PlotConfigBullseyePlot
	if err := c.Bind(&pc); err != nil {
		return httperr.MalformedBody(err)
	}
	// Default to 1 year if no date range provided
	if pc.DateRange == "" {
		pc.DateRange = "1 year"
	}
	if _, err := pc.DateRangeTimeWindow(); err != nil {
		return httperr.MalformedDate(err)
	}
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	pc.ProjectID = pID

	pcID, err := uuid.Parse(c.Param("plot_configuration_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	pc.ID = pcID

	p := c.Get("profile").(model.Profile)
	tNow := time.Now()
	pc.UpdaterID, pc.UpdateDate = &p.ID, &tNow

	pcUpdated, err := h.PlotConfigService.UpdatePlotConfigBullseyePlot(c.Request().Context(), pc)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, pcUpdated)
}

// ListPlotConfigMeasurementsBullseyePlot godoc
//
//	@Summary lists plot config measurements for a bullseye plot
//	@Tags plot-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param plot_configuration_id path string true "plot config uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} []model.PlotConfigMeasurementBullseyePlot
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/plot_configs/bullseye_plots/{plot_configuration_id}/measurements [get]
//	@Security Bearer
func (h *ApiHandler) ListPlotConfigMeasurementsBullseyePlot(c echo.Context) error {
	_, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	pcID, err := uuid.Parse(c.Param("plot_configuration_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	var tw model.TimeWindow
	a, b := c.QueryParam("after"), c.QueryParam("before")
	if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
		return httperr.MalformedDate(err)
	}
	mm, err := h.PlotConfigService.ListPlotConfigMeasurementsBullseyePlot(c.Request().Context(), pcID, tw)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, mm)
}
