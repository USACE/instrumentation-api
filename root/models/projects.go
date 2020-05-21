package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// Project is a project data structure
type Project struct {
	ID         uuid.UUID `json:"id"`
	Deleted    bool      `json:"-"`
	FederalID  *string   `json:"federal_id" db:"federal_id"`
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	Creator    int       `json:"creator"`
	CreateDate time.Time `json:"create_date" db:"create_date"`
	Updater    int       `json:"updater"`
	UpdateDate time.Time `json:"update_date" db:"update_date"`
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
	if err := db.Select(&pp, "SELECT * FROM project WHERE NOT deleted"); err != nil {
		log.Printf(err.Error())
		return make([]Project, 0), err
	}
	return pp, nil
}

// GetProject returns a pointer to a project
func GetProject(db *sqlx.DB, id uuid.UUID) (*Project, error) {
	var p Project
	if err := db.Get(&p, "SELECT * FROM project WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &p, nil
}

// GetProjectByFederalID returns a pointer to a project, looked-up by FederalID
func GetProjectByFederalID(db *sqlx.DB, federalID string) (*Project, error) {
	var p Project
	if err := db.Get(&p, "SELECT * FROM project WHERE federal_id = $1", federalID); err != nil {
		return nil, err
	}
	return &p, nil
}

// CreateProjectBulk creates one or more projects from an array of projects
func CreateProjectBulk(db *sqlx.DB, projects []Project) error {

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

		if _, err = stmt.Exec(i.ID, i.FederalID, i.Slug, i.Name, i.Creator, i.CreateDate, i.Updater, i.UpdateDate); err != nil {
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

	var pUpdated Project
	if err := db.QueryRowx(
		"UPDATE project SET federal_id=$2, name=$3, updater=$4, update_date=$5 WHERE id=$1 RETURNING *",
		p.ID, p.FederalID, p.Name, p.Updater, p.UpdateDate,
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
