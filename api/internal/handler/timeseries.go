package handler

import (
	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListTimeseries godoc
//
//	@Summary lists all timeseries
//	@Tags timeseries
//	@Produce json
//	@Success 200 {array} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries [get]
func (h *ApiHandler) ListTimeseries(c echo.Context) error {
	tt, err := h.TimeseriesService.ListTimeseries(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, tt)
}

// GetTimeseries godoc
//
//	@Summary gets a single timeseries by id
//	@Tags timeseries
//	@Produce json
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Success 200 {object} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries/{timeseries_id} [get]
//	@Router /instruments/{instrument_id}/timeseries/{timeseries_id} [get]
func (h *ApiHandler) GetTimeseries(c echo.Context) error {
	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	t, err := h.TimeseriesService.GetTimeseries(c.Request().Context(), tsID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, t)
}

// ListInstrumentTimeseries godoc
//
//	@Summary lists timeseries for an instrument
//	@Tags timeseries
//	@Produce json
//	@Param project_id path string false "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/timeseries [get]
func (h *ApiHandler) ListInstrumentTimeseries(c echo.Context) error {
	nID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	tt, err := h.TimeseriesService.ListInstrumentTimeseries(c.Request().Context(), nID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tt)
}

// ListInstrumentGroupTimeseries godoc
//
//	@Summary lists timeseries for instruments in an instrument group
//	@Tags timeseries
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Success 200 {array} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id}/timeseries [get]
func (h *ApiHandler) ListInstrumentGroupTimeseries(c echo.Context) error {
	gID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	tt, err := h.TimeseriesService.ListInstrumentGroupTimeseries(c.Request().Context(), gID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tt)
}

// ListProjectTimeseries godoc
//
//	@Summary lists all timeseries for a single project
//	@Tags timeseries
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/timeseries [get]
func (h *ApiHandler) ListProjectTimeseries(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	tt, err := h.TimeseriesService.ListProjectTimeseries(c.Request().Context(), pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tt)
}

// CreateTimeseries godoc
//
//	@Summary creates one or more timeseries
//	@Tags timeseries
//	@Produce json
//	@Param timeseries_collection_items body model.TimeseriesCollectionItems true "timeseries collection items payload"
//	@Success 200 {array} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries [post]
func (h *ApiHandler) CreateTimeseries(c echo.Context) error {
	var tc model.TimeseriesCollectionItems
	if err := c.Bind(&tc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	slugsTaken, err := h.TimeseriesService.ListTimeseriesSlugs(ctx)
	if err != nil {
		return err
	}
	for idx := range tc.Items {
		tc.Items[idx].ID = uuid.Must(uuid.NewRandom())
		s, err := util.NextUniqueSlug(tc.Items[idx].Name, slugsTaken)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		tc.Items[idx].Slug = s
		slugsTaken = append(slugsTaken, s)
	}

	tt, err := h.TimeseriesService.CreateTimeseriesBatch(ctx, tc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, tt)
}

// UpdateTimeseries godoc
//
//	@Summary updates a single timeseries by id
//	@Tags timeseries
//	@Produce json
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Param timeseries body model.Timeseries true "timeseries payload"
//	@Success 200 {object} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries/{timeseries_id} [put]
func (h *ApiHandler) UpdateTimeseries(c echo.Context) error {
	id, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	t := model.Timeseries{}
	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if id != t.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`id`"))
	}
	if _, err := h.TimeseriesService.UpdateTimeseries(c.Request().Context(), t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, t)
}

// DeleteTimeseries godoc
//
//	@Summary deletes a single timeseries by id
//	@Tags timeseries
//	@Produce json
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries/{timeseries_id} [delete]
func (h *ApiHandler) DeleteTimeseries(c echo.Context) error {
	id, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.TimeseriesService.DeleteTimeseries(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
