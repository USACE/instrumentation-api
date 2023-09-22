package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// PlotConfigurationSettings describes options for displaying the plot consistently.
// Specifically, whether to ignore data entries in a timeseries that have been masked,
// or whether to display user comments.
type PlotConfigurationSettings struct {
	ShowMasked       bool   `json:"show_masked" db:"show_masked"`
	ShowNonValidated bool   `json:"show_nonvalidated" db:"show_nonvalidated"`
	ShowComments     bool   `json:"show_comments" db:"show_comments"`
	AutoRange        bool   `json:"auto_range" db:"auto_range"`
	DateRange        string `json:"date_range" db:"date_range"`
	Threshold        int    `json:"threshold" db:"threshold"`
}

// PlotConfiguration holds information for entity PlotConfiguration
type PlotConfiguration struct {
	ID           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	ProjectID    uuid.UUID   `json:"project_id" db:"project_id"`
	TimeseriesID []uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
	AuditInfo
	PlotConfigurationSettings
}

func (pc *PlotConfiguration) ValidateDateRange() error {
	dr := strings.ToLower(pc.DateRange)
	// Check for standard settings
	if dr == "lifetime" {
		return nil
	}
	if dr == "5 years" {
		return nil
	}
	if dr == "1 year" {
		return nil
	}

	// Check for custom date range
	cdr := strings.Split(dr, " - ")
	if len(cdr) == 2 {
		for _, v := range cdr {
			if _, err := time.Parse("01/02/2006", v); err != nil {
				return fmt.Errorf("custom date values must be in format \"MM/DD/YYYY - MM/DD/YYYY\"")
			}
		}
		return nil
	}

	// No match found
	return fmt.Errorf("invalid date range provided")
}

// ListPlotConfigurationsSQL is the base SQL statement for above functions
var ListPlotConfigurationsSQL = `
	SELECT
		id,
		slug,
		name,
		project_id,
		timeseries_id,
		creator,
		create_date,
		updater,
		update_date,
		show_masked,
		show_nonvalidated,
		show_comments,
		auto_range,
		date_range,
		threshold
	FROM v_plot_configuration
`

// PlotConfigFactory converts database rows to PlotConfiguration objects
func PlotConfigFactory(rows *sqlx.Rows) ([]PlotConfiguration, error) {
	defer rows.Close()
	pp := make([]PlotConfiguration, 0) // PlotConfigurations
	var p PlotConfiguration
	for rows.Next() {
		err := rows.Scan(
			&p.ID,
			&p.Slug,
			&p.Name,
			&p.ProjectID,
			pq.Array(&p.TimeseriesID),
			&p.Creator,
			&p.CreateDate,
			&p.Updater,
			&p.UpdateDate,
			&p.ShowMasked,
			&p.ShowNonValidated,
			&p.ShowComments,
			&p.AutoRange,
			&p.DateRange,
			&p.Threshold,
		)
		if err != nil {
			return make([]PlotConfiguration, 0), err
		}
		pp = append(pp, p)
	}
	return pp, nil
}

// ListPlotConfigurationSlugs lists used instrument group slugs in the database
func ListPlotConfigurationSlugs(db *sqlx.DB) ([]string, error) {

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
		return make([]PlotConfiguration, 0), err
	}
	return PlotConfigFactory(rows)
}

// GetPlotConfiguration returns a single plot configuration
func GetPlotConfiguration(db *sqlx.DB, projectID *uuid.UUID, plotconfigID *uuid.UUID) (*PlotConfiguration, error) {

	rows, err := db.Queryx(ListPlotConfigurationsSQL+" WHERE project_id = $1 AND id = $2", projectID, plotconfigID)
	if err != nil {
		return nil, err
	}
	pp, err := PlotConfigFactory(rows)
	if err != nil {
		return nil, err
	}
	if len(pp) == 0 {
		return nil, nil
	}

	return &pp[0], nil
}

// CreatePlotConfiguration add plot configuration for a project
func CreatePlotConfiguration(db *sqlx.DB, pc *PlotConfiguration) (*PlotConfiguration, error) {

	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	// Create Batch Plot
	stmt1, err := txn.Preparex(
		`INSERT INTO plot_configuration (slug, name, project_id, creator, create_date) VALUES ($1, $2, $3, $4, $5)
		 RETURNING id`,
	)
	if err != nil {
		return nil, err
	}
	// Insert any timeseries_id in payload, not in table
	stmt2, err := txn.Preparex(
		`INSERT INTO plot_configuration_timeseries (plot_configuration_id, timeseries_id) VALUES ($1, $2)`,
	)
	if err != nil {
		return nil, err
	}
	stmt3, err := txn.Preparex(
		`INSERT INTO plot_configuration_settings
			(id, show_masked, show_nonvalidated, show_comments, auto_range, date_range, threshold) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
	)
	if err != nil {
		return nil, err
	}

	// ID of newly created plot configuration
	var pcID uuid.UUID
	if err := stmt1.Get(&pcID, pc.Slug, pc.Name, pc.ProjectID, pc.Creator, pc.CreateDate); err != nil {
		return nil, err
	}
	// Create associated plot_configuration_timeseries records
	for _, tsid := range pc.TimeseriesID {
		if _, err := stmt2.Exec(&pcID, &tsid); err != nil {
			return nil, err
		}
	}
	// Create settings.
	if _, err := stmt3.Exec(&pcID, pc.ShowMasked, pc.ShowNonValidated, pc.ShowComments, pc.AutoRange, pc.DateRange, pc.Threshold); err != nil {
		return nil, err
	}
	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return GetPlotConfiguration(db, &pc.ProjectID, &pcID)
}

// UpdatePlotConfiguration update plot configuration for a project
func UpdatePlotConfiguration(db *sqlx.DB, pc *PlotConfiguration) (*PlotConfiguration, error) {

	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	// Prepared Statement; Update Existing Plot Configuration
	stmt1, err := txn.Preparex(`UPDATE plot_configuration SET name = $3, updater = $4, update_date = $5 WHERE project_id = $1 AND id = $2`)
	if err != nil {
		return nil, err
	}

	// Prepared Statement; Delete plot_configuration_timeseries in table that are not in updated plot config
	// Note: "IN" queries w/ sqlx require use of sqlx.In and Query Re-Binding
	query, args, err := sqlx.In(
		`DELETE FROM plot_configuration_timeseries WHERE plot_configuration_id = ? AND timeseries_id NOT IN (?)`,
		pc.ID, pc.TimeseriesID,
	)
	if err != nil {
		return nil, err
	}
	stmt2, err := txn.Preparex(txn.Rebind(query))
	if err != nil {
		return nil, err
	}

	// Prepared Statement; Insert plot_configuration_timeseries from updated plot config
	// DO NOTHING if record already exists for given timeseries_id
	stmt3, err := txn.Preparex(
		`INSERT INTO plot_configuration_timeseries (plot_configuration_id, timeseries_id) VALUES ($1, $2)
		 ON CONFLICT ON CONSTRAINT plot_configuration_unique_timeseries DO NOTHING`,
	)
	if err != nil {
		return nil, err
	}

	// Prepared Statement; Delete plot_configuration_timeseries in table that are not in updated plot config
	// Note: "IN" queries w/ sqlx require use of sqlx.In and Query Re-Binding
	stmt4, err := txn.Preparex(
		`DELETE FROM plot_configuration_settings WHERE id = $1`,
	)
	if err != nil {
		return nil, err
	}

	// Prepared Statement; Update exiting plot configuration settings
	stmt5, err := txn.Preparex(`
		INSERT INTO plot_configuration_settings
		(id, show_masked, show_nonvalidated, show_comments, auto_range, date_range, threshold)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return nil, err
	}

	if _, err := stmt1.Exec(pc.ProjectID, pc.ID, pc.Name, pc.Updater, pc.UpdateDate); err != nil {
		return nil, err
	}

	if _, err := stmt2.Exec(args...); err != nil {
		return nil, err
	}
	if _, err := stmt4.Exec(pc.ID); err != nil {
		return nil, err
	}
	if _, err := stmt5.Exec(pc.ID, pc.ShowMasked, pc.ShowNonValidated, pc.ShowComments, pc.AutoRange, pc.DateRange, pc.Threshold); err != nil {
		return nil, err
	}

	for _, tsid := range pc.TimeseriesID {
		if _, err := stmt3.Exec(pc.ID, tsid); err != nil {
			return nil, err
		}
	}

	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := stmt3.Close(); err != nil {
		return nil, err
	}
	if err := stmt4.Close(); err != nil {
		return nil, err
	}
	if err := stmt5.Close(); err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return GetPlotConfiguration(db, &pc.ProjectID, &pc.ID)

}

// DeletePlotConfiguration delete plot configuration for a project
func DeletePlotConfiguration(db *sqlx.DB, projectID *uuid.UUID, plotConfigID *uuid.UUID) error {
	if _, err := db.Exec(`DELETE from plot_configuration WHERE project_id = $1 AND id = $2`, projectID, plotConfigID); err != nil {
		return err
	}
	return nil
}
