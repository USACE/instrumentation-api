package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type TimeseriesCwmsService interface {
	GetTimeseriesCwms(ctx context.Context, timeseriesID uuid.UUID) (model.TimeseriesCwms, error)
	CreateTimeseriesCwmsBatch(ctx context.Context, instrumentID uuid.UUID, tcc []model.TimeseriesCwms) ([]model.TimeseriesCwms, error)
	UpdateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) error
	ListTimeseriesCwmsMeasurements(ctx context.Context, timeseriesID uuid.UUID, tw model.TimeWindow, threshold int) (model.MeasurementCollection, error)
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

func (s timeseriesCwmsService) CreateTimeseriesCwmsBatch(ctx context.Context, instrumentID uuid.UUID, tcc []model.TimeseriesCwms) ([]model.TimeseriesCwms, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return tcc, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for idx := range tcc {
		tcc[idx].Type = model.CwmsTimeseriesType
		tcc[idx].InstrumentID = instrumentID
		tsNew, err := qtx.CreateTimeseries(ctx, tcc[idx].Timeseries)
		if err != nil {
			return tcc, err
		}
		tcc[idx].Timeseries = tsNew
		if err := qtx.CreateTimeseriesCwms(ctx, tcc[idx]); err != nil {
			return tcc, err
		}
	}

	if err := tx.Commit(); err != nil {
		return tcc, err
	}

	return tcc, nil
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

// If using external timeseries measurement in formula, measurements need to be queried and processed,
// otherwise they can be requested directly by the client
func (s timeseriesCwmsService) ListTimeseriesCwmsMeasurements(ctx context.Context, timeseriesID uuid.UUID, tw model.TimeWindow, threshold int) (model.MeasurementCollection, error) {
	tc, err := s.GetTimeseriesCwms(ctx, timeseriesID)
	if err != nil {
		return model.MeasurementCollection{}, err
	}

	url := fmt.Sprintf("%s?name=%s&office=%s&begin=%s&end=%s&page-size=500", s.cwmsDataUrl, tc.CwmsTimeseriesID, tc.CwmsOfficeID, tw.After, tw.Before)

	cm, err := downloadCwmsTimeseries(ctx, s.cwmsClient, url)
	if err != nil {
		return model.MeasurementCollection{}, err
	}
	if cm.Total == 0 {
		return model.MeasurementCollection{}, nil
	}

	downsamplePerPage := threshold
	if cm.Total > cm.PageSize {
		downsamplePerPage = (cm.PageSize / cm.Total) * threshold
	}

	items, err := parseCwmsTimeseriesRequest(cm, downsamplePerPage)
	if err != nil {
		return model.MeasurementCollection{}, err
	}

	for cm.NextPage != nil {
		nextPageUrl := fmt.Sprintf("%s&next-page=%s", url, *cm.NextPage)
		cm, err := downloadCwmsTimeseries(ctx, s.cwmsClient, nextPageUrl)
		if err != nil {
			return model.MeasurementCollection{}, err
		}
		nextItems, err := parseCwmsTimeseriesRequest(cm, downsamplePerPage)
		if err != nil {
			return model.MeasurementCollection{}, err
		}

		items = append(items, nextItems...)
	}

	return model.MeasurementCollection{TimeseriesID: timeseriesID, Items: items}, nil
}

func downloadCwmsTimeseries(ctx context.Context, client *http.Client, url string) (model.CwmsMeasurementsRaw, error) {
	var cm model.CwmsMeasurementsRaw

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return cm, err
	}
	res, err := client.Do(req)
	if err != nil {
		return cm, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&cm)
	return cm, err
}

func parseCwmsTimeseriesRequest(cm model.CwmsMeasurementsRaw, downsamplePerPage int) ([]model.Measurement, error) {
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
	return model.LTTB(mm, downsamplePerPage), nil
}
