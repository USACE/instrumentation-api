package model

import (
	"context"
	"encoding/json"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
)

const listProjectsSQL = `
	SELECT
		id, federal_id, image, office_id, deleted, slug, name, creator, create_date,
		updater, update_date, instrument_count, instrument_group_count, timeseries
	FROM v_project
`

type District struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	Name             string     `json:"name" db:"name"`
	Initials         string     `json:"initials" db:"initials"`
	DivisionName     string     `json:"division_name" db:"division_name"`
	DivisionInitials string     `json:"division_initials" db:"division_initials"`
	OfficeID         *uuid.UUID `json:"office_id" db:"office_id"`
}

type Project struct {
	ID                   uuid.UUID          `json:"id"`
	FederalID            *string            `json:"federal_id" db:"federal_id"`
	OfficeID             *uuid.UUID         `json:"office_id" db:"office_id"`
	Image                *string            `json:"image" db:"image"`
	Deleted              bool               `json:"-"`
	Slug                 string             `json:"slug"`
	Name                 string             `json:"name"`
	Timeseries           dbSlice[uuid.UUID] `json:"timeseries" db:"timeseries"`
	InstrumentCount      int                `json:"instrument_count" db:"instrument_count"`
	InstrumentGroupCount int                `json:"instrument_group_count" db:"instrument_group_count"`
	AuditInfo
}

type ProjectCount struct {
	ProjectCount int `json:"project_count"`
}

type ProjectCollection struct {
	Projects []Project
}

func (c *ProjectCollection) UnmarshalJSON(b []byte) error {
	switch util.JSONType(b) {
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

const projectSearch = listProjectsSQL + `
	WHERE NOT deleted AND name ILIKE '%' || $1 || '%' LIMIT $2 ORDER BY name
`

// SearchProjects returns search result for projects
func (q *Queries) SearchProjects(ctx context.Context, searchInput string, limit int) ([]SearchResult, error) {
	ss := make([]SearchResult, 0)
	if err := q.db.SelectContext(ctx, &ss, projectSearch, searchInput, limit); err != nil {
		return nil, err
	}
	rr := make([]SearchResult, len(ss))
	for idx, p := range ss {
		rr[idx] = SearchResult{ID: p.ID, Type: "project", Item: p}
	}
	return rr, nil
}

const listDistricts = `
	SELECT * FROM v_district
`

func (q *Queries) ListDistricts(ctx context.Context) ([]District, error) {
	dd := make([]District, 0)
	if err := q.db.SelectContext(ctx, &dd, listDistricts); err != nil {
		return nil, err
	}
	return dd, nil
}

const listProjectSlugs = `
	SELECT slug FROM project
`

// ListProjectSlugs returns a list of used slugs for projects
func (q *Queries) ListProjectSlugs(ctx context.Context) ([]string, error) {
	ss := make([]string, 0)
	if err := q.db.SelectContext(ctx, &ss, listProjectSlugs); err != nil {
		return nil, err
	}
	return ss, nil
}

const listProjects = listProjectsSQL + `
	WHERE NOT deleted ORDER BY name
`

// ListProjects returns a slice of projects
func (q *Queries) ListProjects(ctx context.Context) ([]Project, error) {
	pp := make([]Project, 0)
	if err := q.db.SelectContext(ctx, &pp, listProjects); err != nil {
		return nil, err
	}
	return pp, nil
}

const listProjectsByFederalID = listProjectsSQL + `
	WHERE federal_id IS NOT NULL AND federal_id = $1 AND NOT deleted ORDER BY name
`

// ListProjects returns a slice of projects
func (q *Queries) ListProjectsByFederalID(ctx context.Context, federalID string) ([]Project, error) {
	pp := make([]Project, 0)
	if err := q.db.SelectContext(ctx, &pp, listProjectsByFederalID, federalID); err != nil {
		return nil, err
	}
	return pp, nil
}

const listProjectsForProfile = `
	SELECT DISTINCT
		p.id, p.federal_id, p.image, p.office_id, p.deleted, p.slug, p.name, p.creator, p.create_date,
		p.updater, p.update_date, p.instrument_count, p.instrument_group_count, p.timeseries
	FROM profile_project_roles ppr
	INNER JOIN v_project p on p.id = ppr.project_id
	WHERE ppr.profile_id = $1 AND NOT p.deleted
	ORDER BY p.name
`

func (q *Queries) ListProjectsForProfile(ctx context.Context, profileID uuid.UUID) ([]Project, error) {
	pp := make([]Project, 0)
	if err := q.db.SelectContext(ctx, &pp, listProjectsForProfile, profileID); err != nil {
		return nil, err
	}
	return pp, nil
}

const listProjectInstruments = listInstrumentsSQL + `
	WHERE project_id = $1 AND NOT deleted
`

// ListProjectInstruments returns a slice of instruments for a project
func (q *Queries) ListProjectInstruments(ctx context.Context, projectID uuid.UUID) ([]Instrument, error) {
	ii := make([]Instrument, 0)
	if err := q.db.SelectContext(ctx, &ii, listProjectInstruments, projectID); err != nil {
		return nil, err
	}
	return ii, nil
}

const listProjectInstrumentNames = `
	SELECT name FROM instrument WHERE project_id = $1
`

// ListProjectInstrumentNames returns a slice of instrument names for a project
func (q *Queries) ListProjectInstrumentNames(ctx context.Context, projectID uuid.UUID) ([]string, error) {
	names := make([]string, 0)
	if err := q.db.SelectContext(ctx, &names, listProjectInstrumentNames, projectID); err != nil {
		return nil, err
	}
	return names, nil
}

const listProjectInstrumentGroups = listInstrumentGroupsSQL + `
	WHERE project_id = $1 AND NOT deleted
`

// ListProjectInstrumentGroups returns a list of instrument groups for a project
func (q *Queries) ListProjectInstrumentGroups(ctx context.Context, projectID uuid.UUID) ([]InstrumentGroup, error) {
	gg := make([]InstrumentGroup, 0)
	if err := q.db.SelectContext(ctx, &gg, listProjectInstrumentGroups, projectID); err != nil {
		return nil, err
	}
	return gg, nil
}

const getProjectCount = `
	SELECT COUNT(id) FROM project WHERE NOT deleted
`

// GetProjectCount returns the number of projects in the database that are not deleted
func (q *Queries) GetProjectCount(ctx context.Context) (ProjectCount, error) {
	var pc ProjectCount
	if err := q.db.GetContext(ctx, &pc.ProjectCount, getProjectCount); err != nil {
		return pc, err
	}
	return pc, nil
}

const getProject = listProjectsSQL + `
	WHERE id = $1
`

// GetProject returns a pointer to a project
func (q *Queries) GetProject(ctx context.Context, id uuid.UUID) (Project, error) {
	var p Project
	err := q.db.GetContext(ctx, &p, getProject, id)
	return p, err
}

const createProject = `
	INSERT INTO project (federal_id, slug, name, creator, create_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, slug
`

func (q *Queries) CreateProject(ctx context.Context, p Project) (IDAndSlug, error) {
	var aa IDAndSlug
	err := q.db.GetContext(ctx, &aa, createProject, p.FederalID, p.Slug, p.Name, p.Creator, p.CreateDate)
	return aa, err
}

const updateProject = `
	UPDATE project SET name=$2, updater=$3, update_date=$4, office_id=$5, federal_id=$6 WHERE id=$1 RETURNING id
`

// UpdateProject updates a project
func (q *Queries) UpdateProject(ctx context.Context, p Project) error {
	_, err := q.db.ExecContext(ctx, updateProject, p.ID, p.Name, p.Updater, p.UpdateDate, p.OfficeID, p.FederalID)
	return err
}

const deleteFlagProject = `
	UPDATE project SET deleted=true WHERE id = $1
`

// DeleteFlagProject sets deleted to true for a project
func (q *Queries) DeleteFlagProject(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFlagProject, id)
	return err
}

const createProjectTimeseries = `
	INSERT INTO project_timeseries (project_id, timeseries_id) VALUES ($1, $2)
	ON CONFLICT ON CONSTRAINT project_unique_timeseries DO NOTHING
`

// CreateProjectTimeseries promotes a timeseries to the project level
func (q *Queries) CreateProjectTimeseries(ctx context.Context, projectID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createProjectTimeseries, projectID, timeseriesID)
	return err
}

const deleteProjectTimeseries = `
	DELETE FROM project_timeseries WHERE project_id = $1 AND timeseries_id = $2
`

// DeleteProjectTimeseries removes a timeseries from the project level; Does not delete underlying timeseries
func (q *Queries) DeleteProjectTimeseries(ctx context.Context, projectID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProjectTimeseries, projectID, timeseriesID)
	return err
}