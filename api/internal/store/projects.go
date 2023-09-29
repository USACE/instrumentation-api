package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProjectStore interface {
	ProjectSearch(ctx context.Context, searchInput string, limit int) ([]model.SearchResult, error)
	ListDistricts(ctx context.Context) ([]model.District, error)
	ListProjectSlugs(ctx context.Context) ([]string, error)
	ListProjects(ctx context.Context) ([]model.Project, error)
	ListProjectsByFederalID(ctx context.Context, federalID string) ([]model.Project, error)
	ListProjectsForProfile(ctx context.Context, profileID uuid.UUID) ([]model.Project, error)
	ListProjectInstruments(ctx context.Context, projectID uuid.UUID) ([]model.Instrument, error)
	ListProjectInstrumentNames(ctx context.Context, projectID uuid.UUID) ([]string, error)
	ListProjectInstrumentGroups(ctx context.Context, projectID uuid.UUID) ([]model.InstrumentGroup, error)
	GetProjectCount(ctx context.Context) (int, error)
	GetProject(ctx context.Context, projectID uuid.UUID) (model.Project, error)
	CreateProject(ctx context.Context, p model.Project) (model.IDAndSlug, error)
	CreateProjectBulk(ctx context.Context, projects []model.Project) ([]model.IDAndSlug, error)
	UpdateProject(ctx context.Context, p model.Project) (model.Project, error)
	DeleteFlagProject(ctx context.Context, projectID uuid.UUID) error
	CreateProjectTimeseries(ctx context.Context, projectID, timeseriesID uuid.UUID) error
	DeleteProjectTimeseries(ctx context.Context, projectID, timeseriesID uuid.UUID) error
}

type projectStore struct {
	db *model.Database
	*model.Queries
}

func NewProjectStore(db *model.Database, q *model.Queries) *projectStore {
	return &projectStore{db, q}
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

	qtx := s.WithTx(tx)

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

	qtx := s.WithTx(tx)

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
