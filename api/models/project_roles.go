package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

// ListProjectMembers lists users (profiles) who have permissions on a project and their role info
func ListProjectMembers(db *sqlx.DB, projectID *uuid.UUID) ([]ProjectMembership, error) {

	rr := make([]ProjectMembership, 0)
	if err := db.Select(
		&rr,
		`SELECT id, profile_id, username, email, role_id, role
		 FROM v_profile_project_roles
		 WHERE project_id = $1
		 ORDER BY email`,
		projectID,
	); err != nil {
		return make([]ProjectMembership, 0), err
	}
	return rr, nil
}

// AddProjectMemberRole adds a role to a user for a specific project
func AddProjectMemberRole(db *sqlx.DB, projectID, profileID, roleID, grantedBy *uuid.UUID) (*ProjectMembership, error) {
	var id uuid.UUID
	// NOTE: DO UPDATE does not change underlying value;
	// UPDATE is needed so `RETURNING id` works under all conditions
	// https://stackoverflow.com/questions/34708509/how-to-use-returning-with-on-conflict-in-postgresql
	if err := db.Get(
		&id,
		`INSERT INTO profile_project_roles (project_id, profile_id, role_id, granted_by)
	     VALUES ($1, $2, $3, $4)
		 ON CONFLICT ON CONSTRAINT unique_profile_project_role DO UPDATE SET project_id = EXCLUDED.project_id
		 RETURNING id`,
		projectID, profileID, roleID, grantedBy,
	); err != nil {
		return nil, err
	}

	var pm ProjectMembership
	if err := db.Get(
		&pm,
		`SELECT id, profile_id, username, email, role_id, role
		 FROM v_profile_project_roles
		 WHERE id = $1`,
		id,
	); err != nil {
		return nil, err
	}
	return &pm, nil
}

// RemoveProjectMemberRole removes a role from a user for a specific project
func RemoveProjectMemberRole(db *sqlx.DB, projectID, profileID, roleID *uuid.UUID) error {

	if _, err := db.Exec(
		`DELETE FROM profile_project_roles WHERE project_id = $1 AND profile_id = $2 AND role_id = $3`,
		projectID, profileID, roleID,
	); err != nil {
		return err
	}
	return nil
}
