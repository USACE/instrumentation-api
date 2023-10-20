package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AwareParameterService interface {
	ListAwareParameters(ctx context.Context) ([]model.AwareParameter, error)
	ListAwarePlatformParameterConfig(ctx context.Context) ([]model.AwarePlatformParameterConfig, error)
}

type awareParameterService struct {
	db *model.Database
	*model.Queries
}

func NewAwareParameterService(db *model.Database, q *model.Queries) *awareParameterService {
	return &awareParameterService{db, q}
}

// ListAwarePlatformParameterConfig returns aware platform parameter configs
func (s awareParameterService) ListAwarePlatformParameterConfig(ctx context.Context) ([]model.AwarePlatformParameterConfig, error) {
	aa := make([]model.AwarePlatformParameterConfig, 0)
	ee, err := s.ListAwarePlatformParameterEnabled(ctx)
	if err != nil {
		return aa, err
	}
	// reorganize aware_parameter_key, timeseries_id into map for each instrument
	// Map of aware parameters to timeseries
	m1 := make(map[uuid.UUID]model.AwarePlatformParameterConfig)
	for _, e := range ee {
		if _, ok := m1[e.InstrumentID]; !ok {
			m1[e.InstrumentID] = model.AwarePlatformParameterConfig{
				ProjectID:       e.ProjectID,
				InstrumentID:    e.InstrumentID,
				AwareID:         e.AwareID,
				AwareParameters: make(map[string]*uuid.UUID),
			}
		}
		m1[e.InstrumentID].AwareParameters[e.AwareParameterKey] = e.TimeseriesID
	}

	for k := range m1 {
		aa = append(aa, m1[k])
	}
	return aa, nil
}
