package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProjectRoleStore interface {
	ListProjectMembers(ctx context.Context, projectID uuid.UUID) ([]model.ProjectMembership, error)
	GetProjectMembership(ctx context.Context, roleID uuid.UUID) (model.ProjectMembership, error)
	AddProjectMemberRole(ctx context.Context, projectID, profileID, roleID, grantedBy uuid.UUID) (model.ProjectMembership, error)
	RemoveProjectMemberRole(ctx context.Context, projectID, profileID, roleID uuid.UUID) error
}

type projectRoleStore struct {
	db *model.Database
	*model.Queries
}

func NewProjectRoleStore(db *model.Database, q *model.Queries) *projectRoleStore {
	return &projectRoleStore{db, q}
}

// AddProjectMemberRole adds a role to a user for a specific project
func (s projectRoleStore) AddProjectMemberRole(ctx context.Context, projectID, profileID, roleID, grantedBy uuid.UUID) (model.ProjectMembership, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.ProjectMembership{}, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.WithTx(tx)

	pprID, err := qtx.AddProjectMemberRole(ctx, projectID, profileID, roleID, grantedBy)
	if err != nil {
		return model.ProjectMembership{}, err
	}

	pm, err := qtx.GetProjectMembership(ctx, pprID)
	if err != nil {
		return model.ProjectMembership{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.ProjectMembership{}, err
	}

	return pm, nil
}
