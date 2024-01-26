package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProjectService interface {
	SearchProjects(ctx context.Context, searchInput string, limit int) ([]model.SearchResult, error)
	ListDistricts(ctx context.Context) ([]model.District, error)
	ListProjects(ctx context.Context) ([]model.Project, error)
	ListProjectsByFederalID(ctx context.Context, federalID string) ([]model.Project, error)
	ListProjectsByRole(ctx context.Context, profileID uuid.UUID) ([]model.Project, error)
	ListProjectsForProfile(ctx context.Context, profileID uuid.UUID) ([]model.Project, error)
	ListProjectInstruments(ctx context.Context, projectID uuid.UUID) ([]model.Instrument, error)
	ListProjectInstrumentGroups(ctx context.Context, projectID uuid.UUID) ([]model.InstrumentGroup, error)
	GetProjectCount(ctx context.Context) (model.ProjectCount, error)
	GetProject(ctx context.Context, projectID uuid.UUID) (model.Project, error)
	CreateProject(ctx context.Context, p model.Project) (model.IDSlugName, error)
	CreateProjectBulk(ctx context.Context, projects []model.Project) ([]model.IDSlugName, error)
	UpdateProject(ctx context.Context, p model.Project) (model.Project, error)
	DeleteFlagProject(ctx context.Context, projectID uuid.UUID) error
}

type projectService struct {
	db *model.Database
	*model.Queries
}

func NewProjectService(db *model.Database, q *model.Queries) *projectService {
	return &projectService{db, q}
}

// CreateProjectBulk creates one or more projects from an array of projects
func (s projectService) CreateProjectBulk(ctx context.Context, projects []model.Project) ([]model.IDSlugName, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	pp := make([]model.IDSlugName, len(projects))
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
func (s projectService) UpdateProject(ctx context.Context, p model.Project) (model.Project, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Project{}, err
	}
	defer model.TxDo(tx.Rollback)

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
