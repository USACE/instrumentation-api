package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/util"

	"github.com/google/uuid"

	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/labstack/echo/v4"
)

// ListCollectionGroups returns instrument groups
func (h *ApiHandler) ListCollectionGroups(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cc, err := h.CollectionGroupService.ListCollectionGroups(c.Request().Context(), pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &cc)
}

// GetCollectionGroupDetails gets all data needed to render collection group form
func (h *ApiHandler) GetCollectionGroupDetails(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	d, err := h.CollectionGroupService.GetCollectionGroupDetails(c.Request().Context(), pID, cgID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &d)
}

// CreateCollectionGroup creates a new collection group
func (h *ApiHandler) CreateCollectionGroup(c echo.Context) error {
	var cg model.CollectionGroup
	// Bind Information Provided
	if err := c.Bind(&cg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Project ID from Route Params
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cg.ProjectID = pID
	// Generate Unique Slug
	slugsTaken, err := h.CollectionGroupService.ListCollectionGroupSlugs(c.Request().Context(), pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	slug, err := util.NextUniqueSlug(cg.Name, slugsTaken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cg.Slug = slug
	// Profile of user creating collection group
	p := c.Get("profile").(model.Profile)
	cg.Creator, cg.CreateDate = p.ID, time.Now()
	// Create Collection Group
	cgNew, err := h.CollectionGroupService.CreateCollectionGroup(c.Request().Context(), cg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, []model.CollectionGroup{cgNew})
}

// UpdateCollectionGroup updates an existing collection group
func (h *ApiHandler) UpdateCollectionGroup(c echo.Context) error {
	var cg model.CollectionGroup
	// Bind Information Provided
	if err := c.Bind(&cg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Project ID from Route Params
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cg.ProjectID = pID
	// Collection Group ID
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Check ID in Route Params vs. ID in Payload
	if cgID != cg.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`collection_group_id`"))
	}
	// Actor Information (creator, create_date, updater, update_date)
	p := c.Get("profile").(model.Profile)
	t := time.Now()
	cg.Updater, cg.UpdateDate = &p.ID, &t
	// Update Collection Group
	cgUpdated, err := h.CollectionGroupService.UpdateCollectionGroup(c.Request().Context(), cg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, cgUpdated)
}

// DeleteCollectionGroup deletes a collection group using the id of the collection group
func (h *ApiHandler) DeleteCollectionGroup(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.CollectionGroupService.DeleteCollectionGroup(c.Request().Context(), pID, cgID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// AddTimeseriesToCollectionGroup adds a timeseries from a collection group
func (h *ApiHandler) AddTimeseriesToCollectionGroup(c echo.Context) error {
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.CollectionGroupService.AddTimeseriesToCollectionGroup(c.Request().Context(), cgID, tsID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// RemoveTimeseriesFromCollectionGroup removes a timeseries from a collection group
func (h *ApiHandler) RemoveTimeseriesFromCollectionGroup(c echo.Context) error {
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.CollectionGroupService.RemoveTimeseriesFromCollectionGroup(c.Request().Context(), cgID, tsID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
