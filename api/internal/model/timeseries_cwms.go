package model

import (
	"context"
)

type TimeseriesCwms struct {
	Timeseries
	CwmsTimeseriesID string
	CwmsOfficeID     string
}

type CwmsMeasurementsRaw struct {
	Begin          string             `json:"begin"`
	End            string             `json:"end"`
	Interval       string             `json:"interval"`
	IntervalOffset int                `json:"interval-offset"`
	Name           string             `json:"name"`
	OfficeID       string             `json:"office-id"`
	Page           string             `json:"page"`
	PageSize       int                `json:"page-size"`
	TimeZone       string             `json:"time-zone"`
	Total          int                `json:"total"`
	Units          string             `json:"units"`
	ValueColumns   []CwmsValueColumns `json:"value-columns"`
	Values         [][]any            `json:"values"`
}

type CwmsValueColumns struct {
	Name     string `json:"name"`
	Ordinal  int    `json:"ordinal"`
	Datatype string `json:"datatype"`
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
