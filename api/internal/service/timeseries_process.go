package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type ProcessTimeseriesService interface {
	SelectMeasurements(ctx context.Context, f model.ProcessMeasurementFilter) (model.ProcessTimeseriesResponseCollection, error)
	SelectInclinometerMeasurements(ctx context.Context, f model.ProcessMeasurementFilter) (model.ProcessInclinometerTimeseriesResponseCollection, error)
}

type processTimeseriesService struct {
	db *model.Database
	*model.Queries
}

func NewProcessTimeseriesService(db *model.Database, q *model.Queries) *processTimeseriesService {
	return &processTimeseriesService{db, q}
}
