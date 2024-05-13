package model

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type ReasonCode int

const (
	None ReasonCode = iota
	Unauthorized
	InvalidName
	InvalidUnassign
)

type InstrumentsValidation struct {
	ReasonCode ReasonCode `json:"-"`
	IsValid    bool       `json:"is_valid"`
	Errors     []string   `json:"errors"`
}

const assignInstrumentToProject = `
	INSERT INTO project_instrument (project_id, instrument_id) VALUES ($1, $2)
	ON CONFLICT ON CONSTRAINT project_instrument_project_id_instrument_id_key DO NOTHING
`

func (q *Queries) AssignInstrumentToProject(ctx context.Context, projectID, instrumentID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, assignInstrumentToProject, projectID, instrumentID)
	return err
}

const unassignInstrumentFromProject = `
	DELETE FROM project_instrument WHERE project_id = $1 AND instrument_id = $2
`

func (q *Queries) UnassignInstrumentFromProject(ctx context.Context, projectID, instrumentID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, unassignInstrumentFromProject, projectID, instrumentID)
	return err
}

const validateInstrumentNamesProjectUnique = ` 
	SELECT i.name
	FROM project_instrument pi
	INNER JOIN instrument i ON pi.instrument_id = i.id
	WHERE pi.project_id = ?
	AND i.name IN (?)
	AND NOT i.deleted
`

// ValidateInstrumentNamesProjectUnique checks that the provided instrument names do not already belong to a project
func (q *Queries) ValidateInstrumentNamesProjectUnique(ctx context.Context, projectID uuid.UUID, instrumentNames []string) (InstrumentsValidation, error) {
	var v InstrumentsValidation
	query, args, err := sqlIn(validateInstrumentNamesProjectUnique, projectID, instrumentNames)
	if err != nil {
		return v, err
	}
	var nn []string
	if err := q.db.SelectContext(ctx, &nn, q.db.Rebind(query), args...); err != nil {
		return v, err
	}
	if len(nn) != 0 {
		vErrors := make([]string, len(nn))
		for idx := range nn {
			vErrors[idx] = fmt.Sprintf(
				"Instrument name '%s' is already taken. Instrument names must be unique within associated projects",
				nn[idx],
			)
		}
		v.Errors = vErrors
		v.ReasonCode = InvalidName
	} else {
		v.IsValid = true
		v.Errors = make([]string, 0)
	}
	return v, nil
}

const validateProjectsInstrumentNameUnique = ` 
	SELECT p.name, i.name
	FROM project_instrument pi
	INNER JOIN instrument i ON pi.instrument_id = i.id
	INNER JOIN project p ON pi.project_id = p.id
	WHERE i.name = ?
	AND pi.instrument_id IN (?)
	AND NOT i.deleted
	ORDER BY pi.project_id
`

// ValidateProjectsInstrumentNameUnique checks that the provided instrument name does not already belong to one of the provided projects
func (q *Queries) ValidateProjectsInstrumentNameUnique(ctx context.Context, instrumentName string, projectIDs []uuid.UUID) (InstrumentsValidation, error) {
	var v InstrumentsValidation
	query, args, err := sqlIn(validateProjectsInstrumentNameUnique, instrumentName, projectIDs)
	if err != nil {
		return v, err
	}
	var nn []string
	if err := q.db.SelectContext(ctx, &nn, q.db.Rebind(query), args...); err != nil {
		return v, err
	}
	if len(nn) != 0 {
		vErrors := make([]string, len(nn))
		for idx := range nn {
			vErrors[idx] = fmt.Sprintf(
				"Instrument name '%s' is already taken. Instrument names must be unique within associated projects",
				nn[idx],
			)
		}
		v.Errors = vErrors
		v.ReasonCode = InvalidName
	} else {
		v.IsValid = true
		v.Errors = make([]string, 0)
	}
	return v, nil
}

// case where service provides slice of instrument ids for single project
const validateInstrumentsAssignerAuthorized = `
	SELECT p.name AS project_name, i.name AS instrument_name
	FROM project_instrument pi
	INNER JOIN project p ON pi.project_id = p.id
	INNER JOIN instrument i ON pi.instrument_id = i.id
	WHERE pi.instrument_id IN (?)
	AND NOT EXISTS (
		SELECT 1 FROM v_profile_project_roles ppr
		WHERE ppr.profile_id = ?
		AND (ppr.is_admin OR (ppr.project_id = pi.project_id AND ppr.role = 'ADMIN'))
	)
	AND NOT i.deleted
`

func (q *Queries) ValidateInstrumentsAssignerAuthorized(ctx context.Context, profileID uuid.UUID, instrumentIDs []uuid.UUID) (InstrumentsValidation, error) {
	var v InstrumentsValidation
	query, args, err := sqlIn(validateInstrumentsAssignerAuthorized, instrumentIDs, profileID)
	if err != nil {
		return v, err
	}
	var nn []struct {
		ProjectName    string `db:"project_name"`
		InstrumentName string `db:"instrument_name"`
	}
	if err := q.db.SelectContext(ctx, &nn, q.db.Rebind(query), args...); err != nil {
		return v, err
	}
	if len(nn) != 0 {
		vErrors := make([]string, len(nn))
		for idx := range nn {
			vErrors[idx] = fmt.Sprintf(
				"Cannot assign instrument '%s' because is assigned to another project '%s' which the user is not an ADMIN of",
				nn[idx].InstrumentName, nn[idx].ProjectName,
			)
		}
		v.Errors = vErrors
		v.ReasonCode = Unauthorized
	} else {
		v.IsValid = true
		v.Errors = make([]string, 0)
	}
	return v, err
}

// case where service provides slice of project ids for single instrument
const validateProjectsAssignerAuthorized = `
	SELECT p.name
	FROM project_instrument pi
	INNER JOIN project p ON pi.project_id = p.id
	INNER JOIN instrument i ON pi.instrument_id = i.id
	WHERE pi.instrument_id = ?
	AND pi.project_id IN (?)
	AND NOT EXISTS (
		SELECT 1 FROM v_profile_project_roles ppr
		WHERE profile_id = ? AND (ppr.is_admin OR (ppr.project_id = pi.project_id AND ppr.role = 'ADMIN'))
	)
	AND NOT i.deleted
	ORDER BY p.name
`

func (q *Queries) ValidateProjectsAssignerAuthorized(ctx context.Context, profileID, instrumentID uuid.UUID, projectIDs []uuid.UUID) (InstrumentsValidation, error) {
	var v InstrumentsValidation
	query, args, err := sqlIn(validateProjectsAssignerAuthorized, instrumentID, projectIDs, profileID)
	if err != nil {
		return v, err
	}
	var nn []string
	if err := q.db.SelectContext(ctx, &nn, q.db.Rebind(query), args...); err != nil {
		return v, err
	}
	if len(nn) != 0 {
		vErrors := make([]string, len(nn))
		for idx := range nn {
			vErrors[idx] = fmt.Sprintf(
				"Cannot assign instrument to project '%s' because the user is not an ADMIN of this project",
				nn[idx],
			)
		}
		v.Errors = vErrors
		v.ReasonCode = Unauthorized
	} else {
		v.IsValid = true
		v.Errors = make([]string, 0)
	}
	return v, err
}
