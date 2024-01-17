package model

import (
	"context"

	"github.com/google/uuid"
)

// ProjectMembership holds
type ProjectMembership struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ProfileID uuid.UUID `json:"profile_id" db:"profile_id"`
	Username  *string   `json:"username"`
	Email     string    `json:"email"`
	RoleID    uuid.UUID `json:"role_id" db:"role_id"`
	Role      string    `json:"role"`
}

const listProjectMembers = `
	SELECT id, profile_id, username, email, role_id, role
	FROM v_profile_project_roles
	WHERE project_id = $1
	ORDER BY email
`

// ListProjectMembers lists users (profiles) who have permissions on a project and their role info
func (q *Queries) ListProjectMembers(ctx context.Context, projectID uuid.UUID) ([]ProjectMembership, error) {
	rr := make([]ProjectMembership, 0)
	if err := q.db.SelectContext(ctx, &rr, listProjectMembers, projectID); err != nil {
		return nil, err
	}
	return rr, nil
}

const getProjectMembership = `
	SELECT id, profile_id, username, email, role_id, role
	FROM v_profile_project_roles
	WHERE id = $1
`

func (q *Queries) GetProjectMembership(ctx context.Context, roleID uuid.UUID) (ProjectMembership, error) {
	var pm ProjectMembership
	err := q.db.GetContext(ctx, &pm, getProjectMembership, roleID)
	return pm, err
}

const addProjectMemberRole = `
	INSERT INTO profile_project_roles (project_id, profile_id, role_id, granted_by)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT ON CONSTRAINT unique_profile_project_role DO UPDATE SET project_id = EXCLUDED.project_id
	RETURNING id
`

// AddProjectMemberRole adds a role to a user for a specific project
func (q *Queries) AddProjectMemberRole(ctx context.Context, projectID, profileID, roleID, grantedBy uuid.UUID) (uuid.UUID, error) {
	var roleIDNew uuid.UUID
	err := q.db.GetContext(ctx, &roleIDNew, addProjectMemberRole, projectID, profileID, roleID, grantedBy)
	return roleIDNew, err
}

const removeProjectMemberRole = `
	DELETE FROM profile_project_roles WHERE project_id = $1 AND profile_id = $2 AND role_id = $3
`

// RemoveProjectMemberRole removes a role from a user for a specific project
func (q *Queries) RemoveProjectMemberRole(ctx context.Context, projectID, profileID, roleID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, removeProjectMemberRole, projectID, profileID, roleID)
	return err
}

const isProjectAdmin = `
	SELECT EXISTS (
		SELECT 1 FROM profile_project_roles pr
		INNER JOIN role r ON r.id = pr.role_id
		WHERE pr.profile_id = $1
		AND pr.project_id = $2
		AND r.name = 'ADMIN'
	)
`

func (q *Queries) IsProjectAdmin(ctx context.Context, profileID, projectID uuid.UUID) (bool, error) {
	var isAdmin bool
	err := q.db.GetContext(ctx, &isAdmin, isProjectAdmin, projectID)
	return isAdmin, err
}

const isProjectMember = `
	SELECT EXISTS (
		SELECT 1 FROM profile_project_roles pr
		INNER JOIN role r ON r.id = pr.role_id
		WHERE pr.profile_id = $1
		AND pr.project_id = $2
		AND (r.name = 'MEMBER' OR r.name = 'ADMIN')
	)
`

func (q *Queries) IsProjectMember(ctx context.Context, profileID, projectID uuid.UUID) (bool, error) {
	var isMember bool
	err := q.db.GetContext(ctx, &isMember, isProjectMember, projectID)
	return isMember, err
}
