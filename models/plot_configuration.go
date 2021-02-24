package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// PlotConfiguration holds information for entity PlotConfiguration
type PlotConfiguration struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	Slug       string      `json:"slug"`
	ProjectID  *uuid.UUID  `json:"project_id" db:"project_id"`
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

// ListPlotConfigurationsSlugs lists used instrument group slugs in the database
func ListPlotConfigurationsSlugs(db *sqlx.DB) ([]string, error) {

	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM plot_configuration"); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

// ListPlotConfigurations returns a list of Plot groups
func ListPlotConfigurations(db *sqlx.DB, projectID *uuid.UUID) ([]PlotConfiguration, error) {
	rows, err := db.Queryx(ListPlotConfigurationsSQL+" WHERE project_id = $1", projectID)
	if err != nil {
		return nil, err
	}
	return PlotConfigFactory(rows)
}

// GetPlotConfiguration returns a single plot configuration
func GetPlotConfiguration(db *sqlx.DB, PID uuid.UUID, CID uuid.UUID) (*PlotConfiguration, error) {
	var p PlotConfiguration

	row := db.QueryRow(ListPlotConfigurationsSQL+" WHERE project_id = $1 AND id = $2", PID, CID)

	if err := row.Scan(&p.ID, &p.Name, &p.ProjectID, pq.Array(&p.Timeseries), &p.Creator,
		&p.CreateDate, &p.Updater, &p.UpdateDate); err != nil {
		return nil, err
	}

	return &p, nil
}

// CreatePlotConfiguration add plot configuration for a project
func CreatePlotConfiguration(db *sqlx.DB, configurations []PlotConfiguration) ([]IDAndSlug, error) {
	// Begin the transation
	txn, err := db.Beginx()
	if err != nil {
		return make([]IDAndSlug, 0), err
	}

	// Prepare the SQL statement
	stmt, err := txn.Preparex(`INSERT INTO plot_configuration
		(slug, name, project_id, creator, create_date, updater, update_date) VALUES
		($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, slug`)

	// Process the array of configurations
	cc := make([]IDAndSlug, len(configurations))
	for idx, c := range configurations {
		if err := stmt.Get(
			&cc[idx],
			c.Slug, c.Name, c.ProjectID, c.Creator, c.CreateDate, c.Updater,
			c.UpdateDate,
		); err != nil {
			return make([]IDAndSlug, 0), err
		}
	}
	// close the statement
	if err = stmt.Close(); err != nil {
		return make([]IDAndSlug, 0), err
	}
	// Commit the transaction
	if err = txn.Commit(); err != nil {
		return make([]IDAndSlug, 0), err
	}
	// return the array of id's and slugs
	return cc, nil
}

// UpdatePlotConfiguration update plot configuration for a project

// DeletePlotConfiguration delete plot configuration for a project
func DeletePlotConfiguration(db *sqlx.DB, projectID *uuid.UUID, plotConfigID *uuid.UUID) error {
	if _, err := db.Exec(
		`WITH config_id AS (
			DELETE from plot_configuration where id = $2
			and project_id = $1 returning id
		)
		delete from plot_configuration_timeseries where plot_configuration_id in (select id from config_id)`,
		projectID, plotConfigID,
	); err != nil {
		return err
	}
	return nil
}

// ListPlotConfigurationsSQL is the base SQL statement for above functions
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
