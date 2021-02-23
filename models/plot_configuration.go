package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// PlotConfiguration holds information for entity Plot_group
type PlotConfiguration struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	ProjectID  uuid.UUID   `json:"project_id" db:"project_id"`
	Timeseries []uuid.UUID `json:"timeseries" db:"timeseries"`
	AuditInfo
}

// PlotConfigurationCollection is a collection of Plot items
type PlotConfigurationCollection struct {
	Items []PlotConfiguration
}

// UnmarshalJSON implements UnmarshalJSON interface
// Allows unpacking object or array of objects into array of objects
func (c *PlotConfigurationCollection) UnmarshalJSON(b []byte) error {

	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var g PlotConfiguration
		if err := json.Unmarshal(b, &g); err != nil {
			return err
		}
		c.Items = []PlotConfiguration{g}
	default:
		c.Items = make([]PlotConfiguration, 0)
	}
	return nil
}

// PlotConfigFactory converts database rows to PlotConfiguration objects
func PlotConfigFactory(rows *sqlx.Rows) ([]PlotConfiguration, error) {
	defer rows.Close()
	pp := make([]PlotConfiguration, 0) // PlotConfigurations
	var p PlotConfiguration
	for rows.Next() {
		err := rows.Scan(
			&p.ID, &p.Name, &p.ProjectID, pq.Array(&p.Timeseries), &p.Creator,
			&p.CreateDate, &p.Updater, &p.UpdateDate,
		)
		if err != nil {
			return make([]PlotConfiguration, 0), err
		}
		pp = append(pp, p)
	}
	return pp, nil
}

// ListPlotConfigurations returns a list of Plot groups
func ListPlotConfigurations(db *sqlx.DB, projectID *uuid.UUID) ([]PlotConfiguration, error) {
	var ListPlotConfigurationsSQL = `SELECT id,
									name,
									project_id,
									t.timeseries,
									creator,
									create_date,
									updater,
									update_date
									FROM plot_configuration
									LEFT JOIN
									(SELECT array_agg(timeseries_id) as timeseries, plot_configuration_id
									FROM plot_configuration_timeseries
									GROUP BY plot_configuration_id) as t
									ON id = t.plot_configuration_id`

	rows, err := db.Queryx(ListPlotConfigurationsSQL+" WHERE project_id = $1", projectID)
	if err != nil {
		return nil, err
	}
	return PlotConfigFactory(rows)
}

// GetPlotConfiguration returns a single plot configuration
func GetPlotConfiguration(db *sqlx.DB, ID uuid.UUID) (*PlotConfiguration, error) {
	var GetPlotConfigurationSQL = `SELECT id,
									deleted,
									slug,
									name,
									description,
									project_id,
									creator,
									create_date,
									updater,
									update_date
									FROM plot_configuration`

	var g PlotConfiguration
	if err := db.QueryRowx(
		GetPlotConfigurationSQL+" WHERE id = $1",
		ID,
	).StructScan(&g); err != nil {
		return nil, err
	}
	return &g, nil
}
