package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListTimeseriesCwmsForProject(ctx context.Context, projectID uuid.UUID) ([]model.TimeseriesCwms, error)
// ListTimeseriesCwmsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.TimeseriesCwms, error)
// ListTimeseriesCwmsMeasurements(ctx context.Context, timeseriesID uuid.UUID) (model.MeasurementCollection, error)
// CreateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) (model.TimeseriesCwms, error)
// UpdateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) error

// ListTimeseriesCwmsForProject godoc
//
//	@Summary lists cwms timeseries for a project uuid
//	@Tags timeseries-cwms
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.TimeseriesCwms
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries/cwms/{timeseries_id} [get]
func (h *ApiHandler) ListTimeseriesCwmsForProject(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	tss, err := h.TimeseriesCwmsService.ListTimeseriesCwmsForProject(c.Request().Context(), projectID)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, tss)
}
