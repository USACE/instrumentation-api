package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type TimeseriesCwmsService interface {
	ListTimeseriesCwmsForProject(ctx context.Context, projectID uuid.UUID) ([]model.TimeseriesCwms, error)
	ListTimeseriesCwmsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.TimeseriesCwms, error)
	ListTimeseriesCwmsMeasurements(ctx context.Context, timeseriesID uuid.UUID) (model.MeasurementCollection, error)
	CreateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) (model.TimeseriesCwms, error)
	UpdateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) error
}

type timeseriesCwmsService struct {
	cwmsClient  *http.Client
	cwmsDataUrl string
	db          *model.Database
	*model.Queries
}

func NewTimeseriesCwmsService(cwmsClient *http.Client, cwmsDataUrl string, db *model.Database, q *model.Queries) *timeseriesCwmsService {
	return &timeseriesCwmsService{cwmsClient, cwmsDataUrl, db, q}
}

// If using external timeseries measurement in formula, measurements need to be queried and processed,
// otherwise they can be requested directly by the client
func (s timeseriesCwmsService) ListTimeseriesCwmsMeasurements(ctx context.Context, timeseriesID uuid.UUID, threshold int) (model.MeasurementCollection, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.cwmsDataUrl, nil)
	if err != nil {
		return model.MeasurementCollection{}, err
	}
	res, err := s.cwmsClient.Do(req)
	if err != nil {
		return model.MeasurementCollection{}, err
	}
	defer res.Body.Close()

	var cm model.CwmsMeasurementsRaw
	if err := json.NewDecoder(req.Body).Decode(&cm); err != nil {
		return model.MeasurementCollection{}, err
	}

	var timeIdx, valIdx int
	for idx := range cm.ValueColumns {
		if cm.ValueColumns[idx].Name == "date-time" {
			timeIdx = idx
		}
		if cm.ValueColumns[idx].Name == "value" {
			valIdx = idx
		}
	}

	mm := make([]model.Measurement, len(cm.Values))
	for idx := range cm.Values {
		msEpoch, ok := cm.Values[idx][timeIdx].(int64)
		if !ok {
			continue
		}
		v, ok := cm.Values[idx][valIdx].(float64)
		if !ok {
			continue
		}
		mm[idx] = model.Measurement{Time: time.UnixMilli(msEpoch), Value: v}
	}

	return model.MeasurementCollection{TimeseriesID: timeseriesID, Items: model.LTTB(mm, threshold)}, nil
}

func (s timeseriesCwmsService) CreateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) (model.TimeseriesCwms, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return tsCwms, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	tsNew, err := qtx.CreateTimeseries(ctx, tsCwms.Timeseries)
	if err != nil {
		return tsCwms, err
	}

	if err := qtx.CreateTimeseriesCwms(ctx, tsCwms); err != nil {
		return tsCwms, err
	}

	if err := tx.Commit(); err != nil {
		return tsCwms, err
	}

	tsCwms.Timeseries = tsNew
	return tsCwms, nil
}

func (s timeseriesCwmsService) UpdateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if _, err := qtx.UpdateTimeseries(ctx, tsCwms.Timeseries); err != nil {
		return err
	}

	if err := qtx.UpdateTimeseriesCwms(ctx, tsCwms); err != nil {
		return err
	}

	return tx.Commit()
}
