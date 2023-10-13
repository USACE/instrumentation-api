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

// ListCollectionGroups godoc
//
//	@Summary lists instrument groups
//	@Tags collection-groups
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups [get]
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

// GetCollectionGroupDetails godoc
//
//	@Summary gets all data needed to render collection group form
//	@Tags collection-groups
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param collection_group_id path string true "collection group uuid" Format(uuid)
//	@Success 200 {object} model.CollectionGroupDetails
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id} [get]
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
	return c.JSON(http.StatusOK, d)
}

// CreateCollectionGroup godoc
//
//	@Summary creates a new collection group
//	@Description lists alert configs for a single project optionally filtered by alert_type_id
//	@Tags collection-groups
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param collection_group body model.CollectionGroup true "collection group payload"
//	@Success 200 {array} model.CollectionGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups [post]
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

// UpdateCollectionGroup godoc
//
//	@Summary updates an existing collection group
//	@Tags collection-groups
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param collection_group_id path string true "collection group uuid"
//	@Param collection_group body model.CollectionGroup true "collection group payload"
//	@Success 200 {object} model.CollectionGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id} [put]
func (h *ApiHandler) UpdateCollectionGroup(c echo.Context) error {
	var cg model.CollectionGroup
	if err := c.Bind(&cg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cg.ProjectID = pID
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if cgID != cg.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`collection_group_id`"))
	}
	p := c.Get("profile").(model.Profile)
	t := time.Now()
	cg.Updater, cg.UpdateDate = &p.ID, &t
	cgUpdated, err := h.CollectionGroupService.UpdateCollectionGroup(c.Request().Context(), cg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, cgUpdated)
}

// DeleteCollectionGroup godoc
//
//	@Summary deletes a collection group using the id of the collection group
//	@Tags collection-groups
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param collection_group_id path string true "collection group uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id} [delete]
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

// AddTimeseriesToCollectionGroup godoc
//
//	@Summary adds a timeseries to a collection group
//	@Tags collection-groups
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param collection_group_id path string true "collection group uuid" Format(uuid)
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id}/timeseries/{timeseries_id} [post]
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

// RemoveTimeseriesFromCollectionGroup godoc
//
//	@Summary removes a timeseries from a collection group
//	@Tags collection-groups
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param collection_group_id path string true "collection group uuid" Format(uuid)
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id}/timeseries/{timeseries_id} [delete]
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
