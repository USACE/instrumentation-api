package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"

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
		return httperr.MalformedID(err)
	}
	cc, err := h.CollectionGroupService.ListCollectionGroups(c.Request().Context(), pID)
	if err != nil {
		return httperr.InternalServerError(err)
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
		return httperr.MalformedID(err)
	}
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	d, err := h.CollectionGroupService.GetCollectionGroupDetails(c.Request().Context(), pID, cgID)
	if err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {array} model.CollectionGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups [post]
//	@Security Bearer
func (h *ApiHandler) CreateCollectionGroup(c echo.Context) error {
	var cg model.CollectionGroup
	// Bind Information Provided
	if err := c.Bind(&cg); err != nil {
		return httperr.MalformedBody(err)
	}
	// Project ID from Route Params
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	cg.ProjectID = pID
	p := c.Get("profile").(model.Profile)
	cg.CreatorID, cg.CreateDate = p.ID, time.Now()

	cgNew, err := h.CollectionGroupService.CreateCollectionGroup(c.Request().Context(), cg)
	if err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} model.CollectionGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateCollectionGroup(c echo.Context) error {
	var cg model.CollectionGroup
	if err := c.Bind(&cg); err != nil {
		return httperr.MalformedBody(err)
	}

	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	cg.ProjectID = pID

	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	cg.ID = cgID

	p := c.Get("profile").(model.Profile)
	t := time.Now()
	cg.UpdaterID, cg.UpdateDate = &p.ID, &t
	cgUpdated, err := h.CollectionGroupService.UpdateCollectionGroup(c.Request().Context(), cg)
	if err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteCollectionGroup(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	if err := h.CollectionGroupService.DeleteCollectionGroup(c.Request().Context(), pID, cgID); err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id}/timeseries/{timeseries_id} [post]
//	@Security Bearer
func (h *ApiHandler) AddTimeseriesToCollectionGroup(c echo.Context) error {
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	if err := h.CollectionGroupService.AddTimeseriesToCollectionGroup(c.Request().Context(), cgID, tsID); err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/collection_groups/{collection_group_id}/timeseries/{timeseries_id} [delete]
//	@Security Bearer
func (h *ApiHandler) RemoveTimeseriesFromCollectionGroup(c echo.Context) error {
	cgID, err := uuid.Parse(c.Param("collection_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	if err := h.CollectionGroupService.RemoveTimeseriesFromCollectionGroup(c.Request().Context(), cgID, tsID); err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
