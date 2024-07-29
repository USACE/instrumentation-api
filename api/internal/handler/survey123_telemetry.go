package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// CreateOrUpdateSurvey123Measurements godoc
//
//	@Summary creates or updates survey123 measurements
//	@Tags timeseries
//	@Produce json
//	@Param timeseries_collection_items body TODO true "timeseries collection items payload"
//	@Success 200 {array} map[string]uuid.UUID
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /telemetry/survey123/measurements [post]
//	@Security Bearer
func (h *TelemetryHandler) CreateOrUpdateSurvey123Measurements(c echo.Context) error {
	// TODO
	var sp model.Survey123Payload
	if err := c.Bind(&sp); err != nil {
		return httperr.MalformedBody(err)
	}

	if err := h.Survey123Service.CreateOrUpdateSurvey123Measurements(c.Request().Context(), sp); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.NoContent(http.StatusCreated)
}
