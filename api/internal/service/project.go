package service

import (
	"context"
	"image"
	"io"
	"mime/multipart"
	"os"

	"github.com/USACE/instrumentation-api/api/internal/img"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProjectService interface {
	SearchProjects(ctx context.Context, searchInput string, limit int) ([]model.SearchResult, error)
	ListDistricts(ctx context.Context) ([]model.District, error)
	ListProjects(ctx context.Context) ([]model.Project, error)
	ListProjectsByFederalID(ctx context.Context, federalID string) ([]model.Project, error)
	ListProjectsForProfile(ctx context.Context, profileID uuid.UUID) ([]model.Project, error)
	ListProjectsForProfileRole(ctx context.Context, profileID uuid.UUID, role string) ([]model.Project, error)
	ListProjectInstruments(ctx context.Context, projectID uuid.UUID) ([]model.Instrument, error)
	ListProjectInstrumentGroups(ctx context.Context, projectID uuid.UUID) ([]model.InstrumentGroup, error)
	GetProjectCount(ctx context.Context) (model.ProjectCount, error)
	GetProject(ctx context.Context, projectID uuid.UUID) (model.Project, error)
	CreateProject(ctx context.Context, p model.Project) (model.IDSlugName, error)
	CreateProjectBulk(ctx context.Context, projects []model.Project) ([]model.IDSlugName, error)
	UpdateProject(ctx context.Context, p model.Project) (model.Project, error)
	UploadProjectImage(ctx context.Context, projectID uuid.UUID, file multipart.FileHeader, u uploader) error
	DeleteFlagProject(ctx context.Context, projectID uuid.UUID) error
}

type projectService struct {
	db *model.Database
	*model.Queries
}

func NewProjectService(db *model.Database, q *model.Queries) *projectService {
	return &projectService{db, q}
}

type uploader func(ctx context.Context, r io.Reader, rawPath, bucketName string) error

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

func (s projectService) UploadProjectImage(ctx context.Context, projectID uuid.UUID, file multipart.FileHeader, u uploader) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	p, err := qtx.GetProject(ctx, projectID)
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if err := img.Resize(src, dst, image.Rect(0, 0, 480, 480)); err != nil {
		return err
	}

	if err := qtx.UpdateProjectImage(ctx, file.Filename, projectID); err != nil {
		return err
	}

	if err := u(ctx, src, "/projects/"+p.Slug+"/"+file.Filename, ""); err != nil {
		return err
	}

	return tx.Commit()
}
