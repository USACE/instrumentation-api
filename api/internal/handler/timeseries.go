package handler

import (
	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListTimeseries returns an array of timeseries
func (h *ApiHandler) ListTimeseries(c echo.Context) error {
	tt, err := h.TimeseriesService.ListTimeseries(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, tt)
}

// GetTimeseries returns a single timeseries
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

// ListInstrumentTimeseries lists timeseries for an instrument
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

// ListInstrumentGroupTimeseries lists timeseries for instruments in an instrument group
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

// ListProjectTimeseries lists all timeseries for a single project
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

// CreateTimeseries accepts a timeseries object or array of timeseries objects
// Can handle objects with or without TimeseriesMeasurements
func (h *ApiHandler) CreateTimeseries(c echo.Context) error {
	var tc model.TimeseriesCollectionItems
	if err := c.Bind(&tc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// slugs already taken in the database
	slugsTaken, err := h.TimeseriesService.ListTimeseriesSlugs(c.Request().Context())
	if err != nil {
		return err
	}
	for idx := range tc.Items {
		// Assign UUID
		tc.Items[idx].ID = uuid.Must(uuid.NewRandom())
		// Assign Slug
		s, err := util.NextUniqueSlug(tc.Items[idx].Name, slugsTaken)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		tc.Items[idx].Slug = s
		// Add slug to array of slugs originally fetched from the database
		// to catch duplicate names/slugs from the same bulk upload
		slugsTaken = append(slugsTaken, s)
	}

	tt, err := h.TimeseriesService.CreateTimeseriesBatch(c.Request().Context(), tc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, tt)
}

// UpdateTimeseries updates a single timeseries
func (h *ApiHandler) UpdateTimeseries(c echo.Context) error {

	// id from url params
	id, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	// id from request
	t := model.Timeseries{}
	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// check :id in url params matches id in request body
	if id != t.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`id`"))
	}
	// update
	if _, err := h.TimeseriesService.UpdateTimeseries(c.Request().Context(), t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// return updated instrument
	return c.JSON(http.StatusOK, t)
}

// DeleteTimeseries deletes a single timeseries
func (h *ApiHandler) DeleteTimeseries(c echo.Context) error {
	// id from url params
	id, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.TimeseriesService.DeleteTimeseries(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
