package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Survey123 struct {
	ID        uuid.UUID                        `json:"id" db:"id"`
	ProjectID uuid.UUID                        `json:"project_id" db:"project_id"`
	Name      string                           `json:"name" db:"name"`
	Mappings  dbJSONSlice[EquivalencyTableRow] `json:"mappings,omitempty" db:"mappings"`
}

type Survey123Payload struct {
	EventType string                `json:"eventType"`
	Adds      []Survey123ApplyEdits `json:"adds,omitempty"`
	Updates   []Survey123ApplyEdits `json:"updates,omitempty"`
}

type Survey123ApplyEdits struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	ObjectID   int                    `json:"objectId,omitempty"`
	Geometry   any                    `json:"geometry,omitempty"`
}

const listProjectSurvey123s = `
	SELECT * FROM survey123 WHERE project_id = $1
`

func (q *Queries) ListProjectSurvey123s(ctx context.Context, projectID uuid.UUID) ([]Survey123, error) {
	ss := make([]Survey123, 0)
	err := q.db.SelectContext(ctx, &ss, listProjectSurvey123s, projectID)
	return ss, err
}

const createSurvey123 = `
	INSERT INTO survey123 (project_id, name) VALUES ($1,$2) RETURNING id
`

func (q *Queries) CreateSurvey123(ctx context.Context, s Survey123) (uuid.UUID, error) {
	var newID uuid.UUID
	err := q.db.GetContext(ctx, newID, createSurvey123, s)
	return newID, err
}

const createOrUpdateSurvey123Preview = `
	INSERT INTO survey123_preview (survey123_id, preview) VALUES ($1,$2)
	ON CONFLICT ON CONSTRAINT survey123_id_key DO UPDATE SET preview=EXCLUDED.preview
`

func (q *Queries) CreateOrUpdateSurvey123Preview(ctx context.Context, survey123ID uuid.UUID, preview pgtype.JSON) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateSurvey123Preview, survey123ID, preview)
	return err
}

const createOrUpdateSurvey123EquivalencyTableRow = `
	INSERT INTO survey123_equivalency_table (survey123_id, field_name, display_name, instrument_id, timeseries_id) VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT ON CONSTRAINT survey123_equivalency_table_survey123_id_field_name_key DO UPDATE SET
	display_name=EXCLUDED.display_name, instrument_id=EXCLUDED.instrument_id, timeseries_id=EXCLUDED.timeseries_id
`

func (q *Queries) CreateOrUpdateSurvey123EquivalencyTableRow(ctx context.Context, survey123ID uuid.UUID, r EquivalencyTableRow) error {
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

func (q *Queries) ListSurvey123EquivalencyTableRows(ctx context.Context, survey123ID uuid.UUID) ([]EquivalencyTableRow, error) {
	eqq := make([]EquivalencyTableRow, 0)
	err := q.db.SelectContext(ctx, &eqq, listSurvey123EquivalencyTableRows, survey123ID)
	return eqq, err
}

const deleteSurvey123EquivalencyTableRow = `
	DELETE FROM survey123_equivalency_table WHERE id = $1
`

func (q *Queries) DeleteSurvey123EquivalencyTableRow(ctx context.Context, survey123RowID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSurvey123EquivalencyTableRow, survey123RowID)
	return err
}
