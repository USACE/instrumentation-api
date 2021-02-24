package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/models"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListPlotConfigurations returns plot groups
func ListPlotConfigurations(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		cc, err := models.ListPlotConfigurations(db, &pID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, &cc)
	}
}

// GetPlotConfiguration returns single instrument group
func GetPlotConfiguration(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Project ID
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		// Plot configuration ID
		cID, err := uuid.Parse(c.Param("plot_configuration_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		// Get the plot configuration
		g, err := models.GetPlotConfiguration(db, pID, cID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, g)
	}
}

// CreatePlotConfiguration add plot configuration for a project
func CreatePlotConfiguration(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Sanatized instruments with ID, PID, and slug assigned
		newPlotConfigCollection := func(c echo.Context) (models.PlotConfigurationCollection, error) {
			pc := models.PlotConfigurationCollection{}
			if err := c.Bind(&pc); err != nil {
				return models.PlotConfigurationCollection{}, err
			}

			// Get PID of Instruments
			projectID, err := uuid.Parse(c.Param("project_id"))
			if err != nil {
				return models.PlotConfigurationCollection{}, err
			}

			// slugs already taken in the database
			slugsTaken, err := models.ListInstrumentSlugs(db)
			if err != nil {
				return models.PlotConfigurationCollection{}, err
			}

			// profile of user creating instruments
			p := c.Get("profile").(*models.Profile)

			// timestamp
			t := time.Now()

			for idx := range pc.Items {
				// Assign projectID
				pc.Items[idx].ProjectID = &projectID
				// Assign Slug
				s, err := dbutils.NextUniqueSlug(pc.Items[idx].Name, slugsTaken)
				if err != nil {
					return models.PlotConfigurationCollection{}, err
				}
				pc.Items[idx].Slug = s
				// Assign Creator
				pc.Items[idx].Creator = p.ID
				// Assign CreateDate
				pc.Items[idx].CreateDate = t
				// Add slug to array of slugs originally fetched from the database
				// to catch duplicate names/slugs from the same bulk upload
				slugsTaken = append(slugsTaken, s)
			}

			return pc, nil
		}

		// Plot Configurations
		pc, err := newPlotConfigCollection(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// // Validate POST
		// if c.QueryParam("dry_run") == "true" {
		// 	v, err := models.ValidateCreateInstruments(db, pc.Items)
		// 	if err != nil {
		// 		return c.JSON(http.StatusBadRequest, err)
		// 	}
		// 	return c.JSON(http.StatusOK, v)
		// }

		// Actually POST
		nn, err := models.CreatePlotConfiguration(db, pc.Items)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, nn)
	}
}

// // UpdatePlotConfiguration update plot configuration for a project
// func UpdatePlotConfiguration() echo.HandlerFunc {
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	cc, err := models.UpdatePlotConfiguration(db, &pID)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}
// 	return c.JSON(http.StatusOK, &cc)
// }

// DeletePlotConfiguration delete plot configuration for a project
func DeletePlotConfiguration(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		cID, err := uuid.Parse(c.Param("plot_configuration_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := models.DeletePlotConfiguration(db, &pID, &cID); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
