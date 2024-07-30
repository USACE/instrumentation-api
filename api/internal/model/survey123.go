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
	TimeKey   string                           `json:"time_key" db:"time_key"`
	Mappings  dbJSONSlice[EquivalencyTableRow] `json:"mappings,omitempty" db:"mappings"`
}

type Survey123Payload struct {
	// EventType  string `json:"eventType,omitempty"`
	// PortalInfo any    `json:"portalInfo,omitempty"`
	// SurveyInfo any    `json:"surveyInfo,omitempty"`
	// UserInfo   any    `json:"userInfo,omitempty"`
	// Response   any    `json:"response,omitempty"`
	// Feature    any    `json:"feature,omitempty"`
	ApplyEdits struct {
		Adds []struct {
			Attributes map[string]interface{} `json:"attributes,omitempty"`
			ObjectID   int                    `json:"objectId,omitempty"`
			Geometry   any                    `json:"geometry,omitempty"`
		} `json:"adds,omitempty"`
		Updates []any `json:"updates,omitempty"`
	} `json:"applyEdits,omitempty"`
}

const createSurvey123 = `
	INSERT INTO survey123 (project_id, name) VALUES ($1,$2) RETURNING id
`

func (q *Queries) CreateSurvey123(ctx context.Context, s Survey123) (uuid.UUID, error) {
	var newID uuid.UUID
	err := q.db.GetContext(ctx, newID, createSurvey123, s)
	return newID, err
}

const createOrUpdateSurvey123EquivalencyTable = `
	INSERT INTO survey123_equivalency_table (survey123_id, field_name, instrument_id, timeseries_id) VALUES ($1,$2,$3,$4)
	ON CONFLICT ON CONSTRAINT survey123_equivalency_table_survey123_id_field_name_key DO UPDATE SET
	instrument_id=EXCLUDED.instrument_id, timeseries_id=EXCLUDED.timeseries_id
`

func (q *Queries) CreateOrUpdateSurvey123EquivalencyTable(ctx context.Context, survey123ID uuid.UUID, r EquivalencyTableRow) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateSurvey123EquivalencyTable, survey123ID, r.FieldName, r.InstrumentID, r.TimeseriesID)
	return err
}

const createOrUpdateSurvey123Preview = `
	INSERT INTO survey123_preview (survey123_id, preview) VALUES ($1,$2)
	ON CONFLICT ON CONSTRAINT survey123_id_key DO UPDATE SET preview=EXCLUDED.preview
`

func (q *Queries) CreateOrUpdateSurvey123Preview(ctx context.Context, survey123ID uuid.UUID, preview pgtype.JSON) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateSurvey123Preview, survey123ID, preview)
	return err
}
