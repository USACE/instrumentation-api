package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProjectStore interface {
}

type projectStore struct {
	db *model.Database
	q  *model.Queries
}

func NewProjectStore(db *model.Database, q *model.Queries) *projectStore {
	return &projectStore{db, q}
}

// ProjectSearch returns search result for projects
func (s projectStore) ProjectSearch(ctx context.Context, searchInput string, limit int) ([]model.SearchResult, error) {
	return s.q.ProjectSearch(ctx, searchInput, limit)
}

func (s projectStore) ListDistricts(ctx context.Context) ([]model.District, error) {
	return s.q.ListDistricts(ctx)
}

// ListProjectSlugs returns a list of used slugs for projects
func (s projectStore) ListProjectSlugs(ctx context.Context) ([]string, error) {
	return s.q.ListProjectSlugs(ctx)
}

// ListProjects returns a slice of projects
func (s projectStore) ListProjects(ctx context.Context) ([]model.Project, error) {
	return s.q.ListProjects(ctx)
}

// ListProjects returns a slice of projects
func (s projectStore) ListProjectsByFederalID(ctx context.Context, federalID string) ([]model.Project, error) {
	return s.q.ListProjectsByFederalID(ctx, federalID)
}

func (s projectStore) ListProjectsForProfile(ctx context.Context, profileID uuid.UUID) ([]model.Project, error) {
	return s.q.ListProjectsForProfile(ctx, profileID)
}

// ListProjectInstruments returns a slice of instruments for a project
func (s projectStore) ListProjectInstruments(ctx context.Context, projectID uuid.UUID) ([]model.Instrument, error) {
	return s.q.ListProjectInstruments(ctx, projectID)
}

// ListProjectInstrumentNames returns a slice of instrument names for a project
func (s projectStore) ListProjectInstrumentNames(ctx context.Context, projectID uuid.UUID) ([]string, error) {
	return s.q.ListProjectInstrumentNames(ctx, projectID)
}

// ListProjectInstrumentGroups returns a list of instrument groups for a project
func (s projectStore) ListProjectInstrumentGroups(ctx context.Context, projectID uuid.UUID) ([]model.InstrumentGroup, error) {
	return s.q.ListProjectInstrumentGroups(ctx, projectID)
}

// GetProjectCount returns the number of projects in the database that are not deleted
func (s projectStore) GetProjectCount(ctx context.Context) (int, error) {
	return s.q.GetProjectCount(ctx)
}

// GetProject returns a pointer to a project
func (s projectStore) GetProject(ctx context.Context, projectID uuid.UUID) (model.Project, error) {
	return s.q.GetProject(ctx, projectID)
}

const createProject = `
	INSERT INTO project (federal_id, slug, name, creator, create_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, slug
`

func (s projectStore) CreateProject(ctx context.Context, p model.Project) (model.IDAndSlug, error) {
	return s.q.CreateProject(ctx, p)
}

// CreateProjectBulk creates one or more projects from an array of projects
func (s projectStore) CreateProjectBulk(ctx context.Context, projects []model.Project) ([]model.IDAndSlug, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	pp := make([]model.IDAndSlug, len(projects))
	for idx, p := range projects {
		aa, err := qtx.CreateProject(ctx, p)
		if err != nil {
			return nil, err
		}
		pp[idx] = aa
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return pp, nil
}

// UpdateProject updates a project
func (s projectStore) UpdateProject(ctx context.Context, p model.Project) (model.Project, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Project{}, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	if err := qtx.UpdateProject(ctx, p); err != nil {
		return model.Project{}, err
	}

	updated, err := qtx.GetProject(ctx, p.ID)
	if err != nil {
		return model.Project{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.Project{}, err
	}

	return updated, nil
}

// DeleteFlagProject sets deleted to true for a project
func (s projectStore) DeleteFlagProject(ctx context.Context, projectID uuid.UUID) error {
	return s.q.DeleteFlagProject(ctx, projectID)
}

// CreateProjectTimeseries promotes a timeseries to the project level
func (s projectStore) CreateProjectTimeseries(ctx context.Context, projectID, timeseriesID uuid.UUID) error {
	return s.q.CreateProjectTimeseries(ctx, projectID, timeseriesID)
}

// DeleteProjectTimeseries removes a timeseries from the project level; Does not delete underlying timeseries
func (s projectStore) DeleteProjectTimeseries(ctx context.Context, projectID, timeseriesID uuid.UUID) error {
	return s.q.DeleteProjectTimeseries(ctx, projectID, timeseriesID)
}
