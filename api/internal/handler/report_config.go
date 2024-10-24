package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
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
		return httperr.MalformedID(err)
	}
	rcs, err := h.ReportConfigService.ListProjectReportConfigs(c.Request().Context(), pID)
	if err != nil {
		return httperr.InternalServerError(err)
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
		return httperr.MalformedID(err)
	}
	var rc model.ReportConfig
	if err := c.Bind(&rc); err != nil {
		return httperr.MalformedBody(err)
	}
	rc.ProjectID = pID

	profile := c.Get("profile").(model.Profile)
	t := time.Now()
	rc.CreatorID, rc.CreateDate = profile.ID, t

	rcNew, err := h.ReportConfigService.CreateReportConfig(c.Request().Context(), rc)
	if err != nil {
		return httperr.InternalServerError(err)
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
		return httperr.MalformedID(err)
	}
	rcID, err := uuid.Parse(c.Param("report_config_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	var rc model.ReportConfig
	if err := c.Bind(&rc); err != nil {
		return httperr.MalformedBody(err)
	}
	rc.ID = rcID
	rc.ProjectID = pID

	profile := c.Get("profile").(model.Profile)
	t := time.Now()
	rc.UpdaterID, rc.UpdateDate = &profile.ID, &t

	if err := h.ReportConfigService.UpdateReportConfig(c.Request().Context(), rc); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": rcID})
}

// DeleteReportConfig godoc
//
//	@Summary updates a report config
//	@Tags report-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param report_config_id path string true "report config uuid" Format(uuid)
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
		return httperr.MalformedID(err)
	}
	if err := h.ReportConfigService.DeleteReportConfig(c.Request().Context(), rcID); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": rcID})
}

// GetReportConfigWithPlotConfigs godoc
//
//	@Summary Lists all plot configs for a report config
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
		return httperr.MalformedID(err)
	}
	rcs, err := h.ReportConfigService.GetReportConfigWithPlotConfigs(c.Request().Context(), rcID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, rcs)
}

// CreateReportDownloadJob godoc
//
//	@Summary starts a job to create a pdf report
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
		return httperr.MalformedID(err)
	}
	isLandscape := strings.ToLower(c.QueryParam("is_landscape")) == "true"
	p := c.Get("profile").(model.Profile)

	j, err := h.ReportConfigService.CreateReportDownloadJob(c.Request().Context(), rcID, p.ID, isLandscape)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusCreated, j)
}

// GetReportDownloadJob godoc
//
//	@Summary gets a job that creates a pdf report
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
		return httperr.MalformedID(err)
	}
	p := c.Get("profile").(model.Profile)

	j, err := h.ReportConfigService.GetReportDownloadJob(c.Request().Context(), jobID, p.ID)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, j)
}

// UpdateReportDownloadJob godoc
//
//	@Summary updates a job that creates a pdf report
//	@Tags report-config
//	@Param job_id path string true "download job uuid" Format(uuid)
//	@Param report_download_job body model.ReportDownloadJob true "report download job payload"
//	@Param key query string true "api key"
//	@Accept application/json
//	@Produce application/json
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /report_jobs/{job_id} [put]
func (h *ApiHandler) UpdateReportDownloadJob(c echo.Context) error {
	jobID, err := uuid.Parse(c.Param("job_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	var j model.ReportDownloadJob
	if err := c.Bind(&j); err != nil {
		return httperr.MalformedBody(err)
	}
	j.ID = jobID
	j.ProgressUpdateDate = time.Now()

	if err := h.ReportConfigService.UpdateReportDownloadJob(c.Request().Context(), j); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": j.ID})
}

// DownloadReport godoc
//
//	@Summary downloads a report for a given report job id
//	@Tags report-config
//	@Produce application/pdf
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param report_config_id path string true "report config uuid" Format(uuid)
//	@Param job_id path string true "job uuid" Format(uuid)
//	@Success 200 {file} file
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/report_configs/{report_config_id}/jobs/{job_id}/downloads [get]
//	@Security Bearer
func (h *ApiHandler) DownloadReport(c echo.Context) error {
	jobID, err := uuid.Parse(c.Param("job_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	p := c.Get("profile").(model.Profile)

	j, err := h.ReportConfigService.GetReportDownloadJob(c.Request().Context(), jobID, p.ID)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	if j.Status != "SUCCESS" {
		return httperr.Message(http.StatusBadRequest, fmt.Sprintf("job status %s", j.Status))
	}
	if j.FileExpiry != nil && time.Now().After(*j.FileExpiry) {
		return httperr.Message(http.StatusBadRequest, fmt.Sprintf("file no longer exists, expired at %s", *j.FileExpiry))
	}
	if j.FileKey == nil {
		return httperr.InternalServerError(errors.New("file key returned nil"))
	}

	r, err := h.BlobService.NewReaderContext(c.Request().Context(), *j.FileKey, "")
	if err != nil {
		return httperr.InternalServerError(err)
	}
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment")
	c.Response().Header().Set("Cache-Control", "public, max-age=31536000")

	return c.Stream(http.StatusOK, "application/pdf", r)
}
