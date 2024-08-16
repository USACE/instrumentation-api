package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Survey123 struct {
	ID        uuid.UUID                                 `json:"id" db:"id"`
	ProjectID uuid.UUID                                 `json:"project_id" db:"project_id"`
	Name      string                                    `json:"name" db:"name"`
	Slug      string                                    `json:"slug" db:"slug"`
	Rows      dbJSONSlice[Survey123EquivalencyTableRow] `json:"rows" db:"fields"`
	Errors    dbSlice[string]                           `json:"errors" db:"errors"`
	AuditInfo
}

type Survey123EquivalencyTableRow struct {
	FieldName    string     `json:"field_name" db:"field_name"`
	DisplayName  string     `json:"display_name" db:"display_name"`
	InstrumentID *uuid.UUID `json:"instrument_id" db:"instrument_id"`
	TimeseriesID *uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
}

type Survey123Payload struct {
	EventType string
	Edits     []Survey123Edits
}

type Survey123Edits struct {
	Adds    []Survey123ApplyEdits `json:"adds,omitempty"`
	Updates []Survey123ApplyEdits `json:"updates,omitempty"`
}

type Survey123ApplyEdits struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	ObjectID   any                    `json:"objectId,omitempty"`
	Geometry   any                    `json:"geometry,omitempty"`
}

type Survey123Preview struct {
	Survey123ID uuid.UUID  `json:"survey123_id" db:"survey123_id"`
	Preview     string     `json:"preview" db:"preview"`
	UpdateDate  *time.Time `json:"update_date" db:"update_date"`
}

const listSurvey123sForProject = `
	SELECT * FROM v_survey123 WHERE project_id = $1
`

func (q *Queries) ListSurvey123sForProject(ctx context.Context, projectID uuid.UUID) ([]Survey123, error) {
	ss := make([]Survey123, 0)
	err := q.db.SelectContext(ctx, &ss, listSurvey123sForProject, projectID)
	return ss, err
}

const createSurvey123 = `
	INSERT INTO survey123 (project_id, name, slug, creator) VALUES ($1, $2, slugify($2, 'survey123'), $3) RETURNING id
`

func (q *Queries) CreateSurvey123(ctx context.Context, sv Survey123) (uuid.UUID, error) {
	var newID uuid.UUID
	err := q.db.GetContext(ctx, &newID, createSurvey123, sv.ProjectID, sv.Name, sv.CreatorID)
	return newID, err
}

const updateSurvey123 = `
	UPDATE survey123 SET name=$2, updater=$3, update_date=$4 WHERE id=$1
`

func (q *Queries) UpdateSurvey123(ctx context.Context, sv Survey123) error {
	_, err := q.db.ExecContext(ctx, updateSurvey123, sv.ID, sv.Name, sv.UpdaterID, sv.UpdateDate)
	return err
}

const createOrUpdateSurvey123Preview = `
	INSERT INTO survey123_preview (survey123_id, preview, update_date) VALUES ($1,$2,$3)
	ON CONFLICT ON CONSTRAINT survey123_id_key DO UPDATE SET preview=EXCLUDED.preview, update_date=EXCLUDED.update_date
`

func (q *Queries) CreateOrUpdateSurvey123Preview(ctx context.Context, pv Survey123Preview) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateSurvey123Preview, pv.Survey123ID, pv.Preview, pv.UpdateDate)
	return err
}

const createOrUpdateSurvey123EquivalencyTableRow = `
	INSERT INTO survey123_equivalency_table (survey123_id, field_name, display_name, instrument_id, timeseries_id) VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT ON CONSTRAINT survey123_equivalency_table_survey123_id_survey123_deleted_field_name_key DO UPDATE SET
	display_name=EXCLUDED.display_name, instrument_id=EXCLUDED.instrument_id, timeseries_id=EXCLUDED.timeseries_id
`

func (q *Queries) CreateOrUpdateSurvey123EquivalencyTableRow(ctx context.Context, survey123ID uuid.UUID, r Survey123EquivalencyTableRow) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateSurvey123EquivalencyTableRow, survey123ID, r.FieldName, r.DisplayName, r.InstrumentID, r.TimeseriesID)
	return err
}

const softDeleteSurvey123 = `
	UPDATE survey123 SET deleted = true WHERE id = $1
`

func (q *Queries) SoftDeleteSurvey123(ctx context.Context, survey123ID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, softDeleteSurvey123, survey123ID)
	return err
}

const listSurvey123EquivalencyTableRows = `
	SELECT * FROM survey123_equivalency_table WHERE survey123_id = $1
`

func (q *Queries) ListSurvey123EquivalencyTableRows(ctx context.Context, survey123ID uuid.UUID) ([]Survey123EquivalencyTableRow, error) {
	eqq := make([]Survey123EquivalencyTableRow, 0)
	err := q.db.SelectContext(ctx, &eqq, listSurvey123EquivalencyTableRows, survey123ID)
	return eqq, err
}

const deleteAllSurvey123PayloadErrors = `
	DELETE FROM survey123_payload_error WHERE survey123_id=$1
`

func (q *Queries) DeleteAllSurvey123PayloadErrors(ctx context.Context, survey123ID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllSurvey123PayloadErrors, survey123ID)
	return err
}

const createSurvey123PayloadError = `
	INSERT INTO survey123_payload_error (survey123_id, error_message) VALUES ($1, $2)
`

func (q *Queries) CreateSurvey123PayloadError(ctx context.Context, survey123ID uuid.UUID, errMsg string) error {
	_, err := q.db.ExecContext(ctx, createSurvey123PayloadError, survey123ID, errMsg)
	return err
}

const deleteAllSurvey123EquivalencyTableRows = `
	DELETE FROM survey123_equivalency_table WHERE survey123_id=$1
`

func (q *Queries) DeleteAllSurvey123EquivalencyTableRows(ctx context.Context, survey123ID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllSurvey123EquivalencyTableRows, survey123ID)
	return err
}

const getSurvey123Preview = `
    SELECT p.survey123_id, p.preview, p.update_date
    FROM survey123_preview p
    INNER JOIN survey123 s ON p.survey123_id = s.id
    WHERE p.survey123_id = $1
    AND NOT s.deleted
`

func (q *Queries) GetSurvey123Preview(ctx context.Context, survey123ID uuid.UUID) (Survey123Preview, error) {
	var pv Survey123Preview
	err := q.db.GetContext(ctx, &pv, getSurvey123Preview, survey123ID)
	if errors.Is(err, sql.ErrNoRows) {
		pv.Survey123ID = survey123ID
		pv.Preview = "null"
		return pv, nil
	}
	return pv, err
}
