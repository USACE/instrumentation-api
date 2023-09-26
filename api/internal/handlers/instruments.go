package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"
	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/paulmach/orb/geojson"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListInstruments returns instruments
func ListInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		nn, err := models.ListInstruments(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, nn)
	}
}

// GetInstrumentCount returns the total number of non deleted instruments in the system
func GetInstrumentCount(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := models.GetInstrumentCount(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"instrument_count": count})
	}
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		n, err := models.GetInstrument(db, &id)
		if err != nil {
			if err.Error() == messages.NotFound {
				return echo.NewHTTPError(http.StatusBadRequest, messages.NotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, n)
	}
}

// CreateInstruments accepts an array of instruments for bulk upload to the database
func CreateInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Sanatized instruments with ID, projectID, and slug assigned
		newInstrumentCollection := func(c echo.Context) (models.InstrumentCollection, error) {
			ic := models.InstrumentCollection{}
			if err := c.Bind(&ic); err != nil {
				return models.InstrumentCollection{}, err
			}

			// Get ProjectID of Instruments
			projectID, err := uuid.Parse(c.Param("project_id"))
			if err != nil {
				return models.InstrumentCollection{}, err
			}

			// slugs already taken in the database
			slugsTaken, err := models.ListInstrumentSlugs(db)
			if err != nil {
				return models.InstrumentCollection{}, err
			}

			// profile of user creating instruments
			p := c.Get("profile").(*models.Profile)

			// timestamp
			t := time.Now()

			for idx := range ic.Items {
				// Assign ProjectID
				ic.Items[idx].ProjectID = &projectID
				// Assign Slug
				s, err := util.NextUniqueSlug(ic.Items[idx].Name, slugsTaken)
				if err != nil {
					return models.InstrumentCollection{}, err
				}
				ic.Items[idx].Slug = s
				// Assign Creator
				ic.Items[idx].Creator = p.ID
				// Assign CreateDate
				ic.Items[idx].CreateDate = t
				// Add slug to array of slugs originally fetched from the database
				// to catch duplicate names/slugs from the same bulk upload
				slugsTaken = append(slugsTaken, s)
			}

			return ic, nil
		}

		// Instruments
		ic, err := newInstrumentCollection(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Validate POST
		if c.QueryParam("dry_run") == "true" {
			v, err := models.ValidateCreateInstruments(db, ic.Items)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusOK, v)
		}

		// Actually POST
		nn, err := models.CreateInstruments(db, ic.Items)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, nn)
	}
}

// UpdateInstrument modifies an existing instrument
func UpdateInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// instrument_id from url params
		iID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		// project_id from url params
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		// instrument from request payload
		i := models.Instrument{ID: iID, ProjectID: &pID}
		if err := c.Bind(&i); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// check project_id in url params matches project_id in request body
		if pID != *i.ProjectID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`project_id`"))
		}
		// check instrument_id in url params matches instrument_id in request body
		if iID != i.ID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`instrument_id`"))
		}

		// profile of user creating instruments
		p := c.Get("profile").(*models.Profile)

		// timestamp
		t := time.Now()
		i.Updater, i.UpdateDate = &p.ID, &t

		// update
		iUpdated, err := models.UpdateInstrument(db, &i)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// return updated instrument
		return c.JSON(http.StatusOK, iUpdated)
	}
}

// UpdateInstrumentGeometry updates only the geometry property of an instrument
func UpdateInstrumentGeometry(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		var geom geojson.Geometry
		if err := c.Bind(&geom); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// profile of user creating instruments
		p := c.Get("profile").(*models.Profile)
		instrument, err := models.UpdateInstrumentGeometry(db, &projectID, &instrumentID, &geom, p)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, instrument)
	}
}

// DeleteFlagInstrument changes deleted flag true for an instrument
func DeleteFlagInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		iID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		if err := models.DeleteFlagInstrument(db, &pID, &iID); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
		}

		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
