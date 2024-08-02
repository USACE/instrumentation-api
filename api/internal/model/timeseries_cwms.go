package model

import (
	"context"

	"github.com/google/uuid"
)

type TimeseriesCwms struct {
	Timeseries
	CwmsTimeseriesID string `json:"cwms_timeseries_id" db:"cwms_timeseries_id"`
	CwmsOfficeID     string `json:"cwms_office_id" db:"cwms_office_id"`
}

const listTimeseriesCwms = `
	SELECT * FROM v_timeseries_cwms
	WHERE instrument_id = $1
`

func (q *Queries) ListTimeseriesCwms(ctx context.Context, instrumentID uuid.UUID) ([]TimeseriesCwms, error) {
	tss := make([]TimeseriesCwms, 0)
	err := q.db.SelectContext(ctx, &tss, listTimeseriesCwms, instrumentID)
	return tss, err
}

const getTimeseriesCwms = `
	SELECT * FROM v_timeseries_cwms
	WHERE id = $1
`

func (q *Queries) GetTimeseriesCwms(ctx context.Context, timeseriesID uuid.UUID) (TimeseriesCwms, error) {
	var t TimeseriesCwms
	err := q.db.GetContext(ctx, &t, getTimeseriesCwms, timeseriesID)
	return t, err
}

const createTimeseriesCwms = `
	INSERT INTO timeseries_cwms (timeseries_id, cwms_timeseries_id, cwms_office_id) VALUES ($1, $2, $3)
`

func (q *Queries) CreateTimeseriesCwms(ctx context.Context, tsCwms TimeseriesCwms) error {
	_, err := q.db.ExecContext(ctx, createTimeseriesCwms, tsCwms.ID, tsCwms.CwmsTimeseriesID, tsCwms.CwmsOfficeID)
	return err
}

const updateTimeseriesCwms = `
	UPDATE timeseries_cwms SET cwms_timeseries_id=$2, cwms_office_id=$3 WHERE timeseries_id=$1
`

func (q *Queries) UpdateTimeseriesCwms(ctx context.Context, tsCwms TimeseriesCwms) error {
	_, err := q.db.ExecContext(ctx, updateTimeseriesCwms, tsCwms.ID, tsCwms.CwmsTimeseriesID, tsCwms.CwmsOfficeID)
	return err
}
