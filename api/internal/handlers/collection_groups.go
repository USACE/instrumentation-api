package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/util"

	"github.com/google/uuid"

	"github.com/USACE/instrumentation-api/api/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListCollectionGroups returns instrument groups
func ListCollectionGroups(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		cc, err := models.ListCollectionGroups(db, &pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &cc)
	}
}

// GetCollectionGroupDetails gets all data needed to render collection group form
func GetCollectionGroupDetails(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		cgID, err := uuid.Parse(c.Param("collection_group_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		d, err := models.GetCollectionGroupDetails(db, &pID, &cgID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &d)
	}
}

// CreateCollectionGroup creates a new collection group
func CreateCollectionGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cg models.CollectionGroup
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
		slugsTaken, err := models.ListCollectionGroupSlugs(db, &pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		slug, err := util.NextUniqueSlug(cg.Name, slugsTaken)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		cg.Slug = slug
		// Profile of user creating collection group
		p := c.Get("profile").(*models.Profile)
		cg.Creator, cg.CreateDate = p.ID, time.Now()
		// Create Collection Group
		cgNew, err := models.CreateCollectionGroup(db, &cg)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, []models.CollectionGroup{*cgNew})
	}
}

// UpdateCollectionGroup updates an existing collection group
func UpdateCollectionGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cg models.CollectionGroup
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
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`collection_group_id`"))
		}
		// Actor Information (creator, create_date, updater, update_date)
		p := c.Get("profile").(*models.Profile)
		t := time.Now()
		cg.Updater, cg.UpdateDate = &p.ID, &t
		// Update Collection Group
		cgUpdated, err := models.UpdateCollectionGroup(db, &cg)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, cgUpdated)
	}
}

// DeleteCollectionGroup deletes a collection group using the id of the collection group
func DeleteCollectionGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		cgID, err := uuid.Parse(c.Param("collection_group_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := models.DeleteCollectionGroup(db, &pID, &cgID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

// AddTimeseriesToCollectionGroup adds a timeseries from a collection group
func AddTimeseriesToCollectionGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		cgID, err := uuid.Parse(c.Param("collection_group_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		tsID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := models.AddTimeseriesToCollectionGroup(db, &cgID, &tsID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

// RemoveTimeseriesFromCollectionGroup removes a timeseries from a collection group
func RemoveTimeseriesFromCollectionGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		cgID, err := uuid.Parse(c.Param("collection_group_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		tsID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := models.RemoveTimeseriesFromCollectionGroup(db, &cgID, &tsID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
