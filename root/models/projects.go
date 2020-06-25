package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// Project is a project data structure
type Project struct {
	ID                   uuid.UUID `json:"id"`
	Deleted              bool      `json:"-"`
	FederalID            *string   `json:"federal_id" db:"federal_id"`
	Slug                 string    `json:"slug"`
	Name                 string    `json:"name"`
	InstrumentCount      int       `json:"instrument_count" db:"instrument_count"`
	InstrumentGroupCount int       `json:"instrument_group_count" db:"instrument_group_count"`
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
	pp := make([]Project, 0)
	if err := db.Select(&pp, listProjectsSQL()+" WHERE NOT p.deleted"); err != nil {
		return make([]Project, 0), err
	}
	return pp, nil
}

// ListProjectInstruments returns a slice of instruments for a project
func ListProjectInstruments(db *sqlx.DB, id uuid.UUID) ([]Instrument, error) {

	rows, err := db.Queryx(
		listInstrumentsSQL()+" WHERE NOT I.deleted AND I.project_id = $1",
		id,
	)
	if err != nil {
		return make([]Instrument, 0), err
	}
	return InstrumentsFactory(db, rows)
}

// ListProjectInstrumentGroups returns a list of instrument groups for a project
func ListProjectInstrumentGroups(db *sqlx.DB, id uuid.UUID) ([]InstrumentGroup, error) {
	gg := make([]InstrumentGroup, 0)
	if err := db.Select(
		&gg,
		listInstrumentGroupsSQL()+" WHERE NOT deleted AND project_id = $1",
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
	var p Project
	if err := db.Get(&p, listProjectsSQL()+" WHERE p.id = $1", id); err != nil {
		return nil, err
	}
	return &p, nil
}

// GetProjectByFederalID returns a pointer to a project, looked-up by FederalID
func GetProjectByFederalID(db *sqlx.DB, federalID string) (*Project, error) {
	var p Project
	if err := db.Get(&p, listProjectsSQL()+" WHERE federal_id = $1", federalID); err != nil {
		return nil, err
	}
	return &p, nil
}

// CreateProjectBulk creates one or more projects from an array of projects
func CreateProjectBulk(db *sqlx.DB, a *Action, projects []Project) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("project", "id", "federal_id", "slug", "name", "creator", "create_date", "updater", "update_date"))
	if err != nil {
		return err
	}

	for _, i := range projects {

		if err != nil {
			return err
		}

		if _, err = stmt.Exec(i.ID, i.FederalID, i.Slug, i.Name, a.Actor, a.Time, a.Actor, a.Time); err != nil {
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
func UpdateProject(db *sqlx.DB, a *Action, p *Project) (*Project, error) {

	var pUpdated Project
	if err := db.QueryRowx(
		"UPDATE project SET federal_id=$2, name=$3, updater=$4, update_date=$5 WHERE id=$1 RETURNING *",
		p.ID, p.FederalID, p.Name, a.Actor, a.Time,
	).StructScan(&pUpdated); err != nil {
		return nil, err
	}

	return &pUpdated, nil
}

// DeleteFlagProject sets deleted to true for a project
func DeleteFlagProject(db *sqlx.DB, id uuid.UUID) error {
	if _, err := db.Exec("UPDATE project SET deleted=true WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}

// ListProjectsSQL is the standard SQL for listing all projects
func listProjectsSQL() string {
	return `SELECT p.*,
	               COALESCE(i.count, 0) AS instrument_count,
	               COALESCE(g.count, 0) AS instrument_group_count
            FROM   project p
            LEFT JOIN (
                SELECT project_id,
               	       COUNT(instrument) as count
				FROM   instrument
				WHERE NOT instrument.deleted
                GROUP BY project_id
            ) i ON i.project_id = p.id
            LEFT JOIN (
                SELECT project_id,
            		   COUNT(instrument_group) as count
				FROM   instrument_group
				WHERE NOT instrument_group.deleted
                GROUP BY project_id
            ) g ON g.project_id = p.id
	`
}
