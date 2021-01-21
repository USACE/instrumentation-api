package models

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const listProjectsSQL = `SELECT id, image, office_id, deleted, slug, federal_id, name, creator, create_date,
     updater, update_date, instrument_count, instrument_group_count, timeseries
	 FROM v_project`

// Project is a project data structure
type Project struct {
	ID                   uuid.UUID   `json:"id"`
	OfficeID             *uuid.UUID  `json:"office_id" db:"office_id"`
	Image                *string     `json:"image" db:"image"`
	Deleted              bool        `json:"-"`
	FederalID            *string     `json:"federal_id" db:"federal_id"`
	Slug                 string      `json:"slug"`
	Name                 string      `json:"name"`
	Timeseries           []uuid.UUID `json:"timeseries" db:"timeseries"`
	InstrumentCount      int         `json:"instrument_count" db:"instrument_count"`
	InstrumentGroupCount int         `json:"instrument_group_count" db:"instrument_group_count"`
	AuditInfo
}

// ProjectCollection helps unpack unspecified JSON into an array of products
type ProjectCollection struct {
	Projects []Project
}

// UnmarshalJSON implements UnmarshalJSON interface
func (c *ProjectCollection) UnmarshalJSON(b []byte) error {

	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Projects); err != nil {
			return err
		}
	case "OBJECT":
		var p Project
		if err := json.Unmarshal(b, &p); err != nil {
			return err
		}
		c.Projects = []Project{p}
	default:
		c.Projects = make([]Project, 0)
	}
	return nil
}

// ProjectFactory converts database rows to Project objects
func ProjectFactory(rows *sqlx.Rows) ([]Project, error) {
	defer rows.Close()
	pp := make([]Project, 0) // Projects
	var p Project
	for rows.Next() {
		err := rows.Scan(
			&p.ID, &p.Image, &p.OfficeID, &p.Deleted, &p.Slug, &p.FederalID, &p.Name, &p.Creator, &p.CreateDate,
			&p.Updater, &p.UpdateDate, &p.InstrumentCount, &p.InstrumentGroupCount, pq.Array(&p.Timeseries),
		)
		if err != nil {
			return make([]Project, 0), err
		}
		pp = append(pp, p)
	}
	return pp, nil
}

// ListProjectSlugs returns a list of used slugs for projects
func ListProjectSlugs(db *sqlx.DB) ([]string, error) {
	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM project"); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

// ListProjects returns a slice of projects
func ListProjects(db *sqlx.DB) ([]Project, error) {
	rows, err := db.Queryx(listProjectsSQL + " WHERE NOT deleted")
	if err != nil {
		return make([]Project, 0), err
	}
	return ProjectFactory(rows)
}

// ListProjectInstruments returns a slice of instruments for a project
func ListProjectInstruments(db *sqlx.DB, id uuid.UUID) ([]Instrument, error) {

	rows, err := db.Queryx(listInstrumentsSQL+" WHERE project_id = $1 AND NOT deleted", id)
	if err != nil {
		return make([]Instrument, 0), err
	}
	return InstrumentsFactory(rows)
}

// ListProjectInstrumentNames returns a slice of instrument names for a project
func ListProjectInstrumentNames(db *sqlx.DB, id *uuid.UUID) ([]string, error) {
	var names []string
	if err := db.Select(
		&names,
		"SELECT name FROM instrument WHERE project_id = $1",
		id,
	); err != nil {
		return make([]string, 0), err
	}
	return names, nil
}

// ListProjectInstrumentGroups returns a list of instrument groups for a project
func ListProjectInstrumentGroups(db *sqlx.DB, id uuid.UUID) ([]InstrumentGroup, error) {
	gg := make([]InstrumentGroup, 0)
	if err := db.Select(
		&gg,
		listInstrumentGroupsSQL+" WHERE project_id = $1 AND NOT deleted",
		id,
	); err != nil {
		return make([]InstrumentGroup, 0), err
	}
	return gg, nil
}

// GetProjectCount returns the number of projects in the database that are not deleted
func GetProjectCount(db *sqlx.DB) (int, error) {
	var count int
	if err := db.Get(&count, "SELECT COUNT(id) FROM project WHERE NOT deleted"); err != nil {
		return 0, err
	}
	return count, nil
}

// GetProject returns a pointer to a project
func GetProject(db *sqlx.DB, id uuid.UUID) (*Project, error) {
	rows, err := db.Queryx(listProjectsSQL+" WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	pp, err := ProjectFactory(rows)
	if err != nil {
		return nil, err
	}
	return &pp[0], nil
}

// CreateProjectBulk creates one or more projects from an array of projects
func CreateProjectBulk(db *sqlx.DB, projects []Project) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("project", "federal_id", "slug", "name", "creator", "create_date"))
	if err != nil {
		return err
	}

	for _, i := range projects {

		if err != nil {
			return err
		}

		if _, err = stmt.Exec(i.FederalID, i.Slug, i.Name, i.Creator, i.CreateDate); err != nil {
			return err
		}
	}

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	if err := stmt.Close(); err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// UpdateProject updates a project
func UpdateProject(db *sqlx.DB, p *Project) (*Project, error) {

	_, err := db.Exec(
		"UPDATE project SET federal_id=$2, name=$3, updater=$4, update_date=$5, office_id=$6 WHERE id=$1 RETURNING id",
		p.ID, p.FederalID, p.Name, p.Updater, p.UpdateDate, p.OfficeID,
	)
	if err != nil {
		return nil, err
	}
	return GetProject(db, p.ID)
}

// DeleteFlagProject sets deleted to true for a project
func DeleteFlagProject(db *sqlx.DB, id uuid.UUID) error {
	if _, err := db.Exec("UPDATE project SET deleted=true WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}

// projectInstrumentNamesMap returns a map of key: project_id , value: map[string]bool ;  string is name of instrument Upper
func projectInstrumentNamesMap(db *sqlx.DB, projectIDs []uuid.UUID) (map[uuid.UUID]map[string]bool, error) {
	sql := `SELECT project_id, name
			FROM instrument
			WHERE project_id IN (?)
			ORDER BY project_id
			`
	query, args, err := sqlx.In(sql, projectIDs)
	if err != nil {
		return nil, err
	}
	var nn []struct {
		ProjectID      uuid.UUID `db:"project_id"`
		InstrumentName string    `db:"name"`
	}
	if err := db.Select(&nn, db.Rebind(query), args...); err != nil {
		return nil, err
	}

	// Make Map
	m := make(map[uuid.UUID]map[string]bool)
	var _pID uuid.UUID
	for _, n := range nn {
		if n.ProjectID != _pID {
			// Starting on a new project of instrument names
			m[n.ProjectID] = make(map[string]bool)
			_pID = n.ProjectID // Increment ProjectID
		}

		m[n.ProjectID][strings.ToUpper(n.InstrumentName)] = true
	}
	return m, nil

}

// CreateProjectTimeseries promotes a timeseries to the project level
func CreateProjectTimeseries(db *sqlx.DB, projectID *uuid.UUID, timeseriesID *uuid.UUID) error {

	// if the timeseries_id is already promoted to the project level, do nothing (i.e. RESTful 200)
	if _, err := db.Exec(
		`INSERT INTO project_timeseries (project_id, timeseries_id) VALUES ($1, $2)
		 ON CONFLICT ON CONSTRAINT project_unique_timeseries DO NOTHING`,
		projectID, timeseriesID,
	); err != nil {
		return err
	}
	return nil
}

// DeleteProjectTimeseries removes a timeseries from the project level; Does not delete underlying timeseries
func DeleteProjectTimeseries(db *sqlx.DB, projectID *uuid.UUID, timeseriesID *uuid.UUID) error {

	// if the timeseries_id is already promoted to the project level, do nothing (i.e. RESTful 200)
	if _, err := db.Exec(
		`DELETE FROM project_timeseries WHERE project_id = $1 AND timeseries_id = $2`,
		projectID, timeseriesID,
	); err != nil {
		return err
	}
	return nil
}
