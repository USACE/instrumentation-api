package model

import (
	"context"
	"encoding/json"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
)

// InstrumentStatus is an instrument status
type InstrumentStatus struct {
	ID       uuid.UUID `json:"id"`
	Time     time.Time `json:"time"`
	StatusID uuid.UUID `json:"status_id" db:"status_id"`
	Status   string    `json:"status"`
}

// InstrumentStatusCollection is a collection of instrument status
type InstrumentStatusCollection struct {
	Items []InstrumentStatus
}

// UnmarshalJSON implements the UnmarshalJSONinterface
func (c *InstrumentStatusCollection) UnmarshalJSON(b []byte) error {
	switch util.JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var s InstrumentStatus
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		c.Items = []InstrumentStatus{s}
	default:
		c.Items = make([]InstrumentStatus, 0)
	}
	return nil
}

const listInstrumentStatusSQL = `
	SELECT
		S.id,
		S.status_id,
		D.name         AS status,
		S.time
	FROM instrument_status S
	INNER JOIN status D
		ON D.id = S.status_id
`

const listInstrumentStatus = listInstrumentStatusSQL + `
	WHERE S.instrument_id = $1 ORDER BY time DESC
`

// ListInstrumentStatus returns all status values for an instrument
func (q *Queries) ListInstrumentStatus(ctx context.Context, instrumentID uuid.UUID) ([]InstrumentStatus, error) {
	ss := make([]InstrumentStatus, 0)
	if err := q.db.SelectContext(ctx, &ss, listInstrumentStatus, instrumentID); err != nil {
		return nil, err
	}
	return ss, nil
}

const getInstrumentStatus = listInstrumentStatusSQL + `
	WHERE S.id = $1
`

// GetInstrumentStatus gets a single status
func (q *Queries) GetInstrumentStatus(ctx context.Context, statusID uuid.UUID) (InstrumentStatus, error) {
	var s InstrumentStatus
	if err := q.db.GetContext(ctx, &s, getInstrumentStatus, statusID); err != nil {
		return s, err
	}
	return s, nil
}

const createOrUpdateInstrumentStatus = `
	INSERT INTO instrument_status (instrument_id, status_id, time) VALUES ($1, $2, $3)
	ON CONFLICT ON CONSTRAINT instrument_unique_status_in_time DO UPDATE SET status_id = EXCLUDED.status_id
`

// CreateOrUpdateInstrumentStatus creates a Instrument Status, updates value on conflict
func (q *Queries) CreateOrUpdateInstrumentStatus(ctx context.Context, instrumentID, statusID uuid.UUID, statusTime time.Time) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateInstrumentStatus, instrumentID, statusID, statusTime)
	return err
}

const deleteInstrumentStatus = `
	DELETE FROM instrument_status WHERE id = $1
`

// DeleteInstrumentStatus deletes a status for an instrument
func (q *Queries) DeleteInstrumentStatus(ctx context.Context, statusID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInstrumentStatus, statusID)
	return err
}
