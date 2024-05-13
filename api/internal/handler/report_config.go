package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

// ListProjectReportConfigs godoc
//
//	@Summary lists all report configs for a project
//	@Tags report-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Accept application/json
//	@Success 200 {object} model.ReportConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/report_configs [get]
//	@Security Bearer
func (h *ApiHandler) ListProjectReportConfigs(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	rcs, err := h.ReportConfigService.ListProjectReportConfigs(c.Request().Context(), pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rcs)
}

// CreateReportConfig godoc
//
//	@Summary creates a report config
//	@Tags report-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param report_config body model.ReportConfig true "report config payload"
//	@Param key query string false "api key"
//	@Accept application/json
//	@Success 201 {object} model.ReportConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/report_configs [post]
//	@Security Bearer
func (h *ApiHandler) CreateReportConfig(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	var rc model.ReportConfig
	if err := c.Bind(&rc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	rc.ProjectID = pID

	profile := c.Get("profile").(model.Profile)
	t := time.Now()
	rc.CreatorID, rc.CreateDate = profile.ID, t

	rcNew, err := h.ReportConfigService.CreateReportConfig(c.Request().Context(), rc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, rcNew)
}

// UpdateReportConfig godoc
//
//	@Summary updates a report config
//	@Tags report-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param report_config_id path string true "report config uuid" Format(uuid)
//	@Param report_config body model.ReportConfig true "report config payload"
//	@Param key query string false "api key"
//	@Accept application/json
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/report_configs/{report_config_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateReportConfig(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	rcID, err := uuid.Parse(c.Param("report_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	var rc model.ReportConfig
	if err := c.Bind(&rc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	rc.ID = rcID
	rc.ProjectID = pID

	profile := c.Get("profile").(model.Profile)
	t := time.Now()
	rc.UpdaterID, rc.UpdateDate = &profile.ID, &t

	if err := h.ReportConfigService.UpdateReportConfig(c.Request().Context(), rc); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": rcID})
}

// DeleteReportConfig godoc
//
//	@Summary updates a report config
//	@Tags report-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/report_configs/{report_config_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteReportConfig(c echo.Context) error {
	rcID, err := uuid.Parse(c.Param("report_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	if err := h.ReportConfigService.DeleteReportConfig(c.Request().Context(), rcID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": rcID})
}

// GetReportConfigWithPlotConfigs godoc
//
//	@Sumary Lists all plot configs for a report config
//	@Tags report-config
//	@Produce json
//	@Param report_config_id path string true "report config uuid" Format(uuid)
//	@Param key query string true "api key"
//	@Success 200 {object} model.ReportConfigWithPlotConfigs
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /report_configs/{report_config_id}/plot_configs [get]
func (h *ApiHandler) GetReportConfigWithPlotConfigs(c echo.Context) error {
	rcID, err := uuid.Parse(c.Param("report_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	rcs, err := h.ReportConfigService.GetReportConfigWithPlotConfigs(c.Request().Context(), rcID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rcs)
}

// CreateReportDownloadJob godoc
//
//	@Sumary starts a job to create a pdf report
//	@Tags report-config
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param report_config_id path string true "report config uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Produce application/json
//	@Success 201 {object} model.ReportDownloadJob
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/report_configs/{report_config_id}/jobs [post]
//	@Security Bearer
func (h *ApiHandler) CreateReportDownloadJob(c echo.Context) error {
	rcID, err := uuid.Parse(c.Param("report_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	p := c.Get("profile").(model.Profile)

	j, err := h.ReportConfigService.CreateReportDownloadJob(c.Request().Context(), rcID, p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, j)
}

// GetReportDownloadJob godoc
//
//	@Sumary gets a job that creates a pdf report
//	@Tags report-config
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param report_config_id path string true "report config uuid" Format(uuid)
//	@Param job_id path string true "download job uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Produce application/json
//	@Success 200 {object} model.ReportDownloadJob
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/report_configs/{report_config_id}/jobs/{job_id} [get]
//	@Security Bearer
func (h *ApiHandler) GetReportDownloadJob(c echo.Context) error {
	jobID, err := uuid.Parse(c.Param("job_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	p := c.Get("profile").(model.Profile)

	j, err := h.ReportConfigService.GetReportDownloadJob(c.Request().Context(), jobID, p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, j)
}

// UpdateReportDownloadJob godoc
//
//	@Sumary updates a job that creates a pdf report
//	@Tags report-config
//	@Param job_id path string true "download job uuid" Format(uuid)
//	@Param report_download_job body model.ReportDownloadJob true "report download job payload"
//	@Param key query string true "api key"
//	@Produce application/json
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/report_configs/{report_config_id}/jobs/{job_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateReportDownloadJob(c echo.Context) error {
	jobID, err := uuid.Parse(c.Param("job_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	var j model.ReportDownloadJob
	if err := c.Bind(&j); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	j.ID = jobID
	j.ProgressUpdateDate = time.Now()

	if err := h.ReportConfigService.UpdateReportDownloadJob(c.Request().Context(), j); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": j.ID})
}
