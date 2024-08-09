package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/model"
)

const (
	emailWarning  = "Warning"
	emailAlert    = "Alert"
	emailReminder = "Reminder"
)

type AlertCheckService interface {
	DoAlertSchedulerChecks(ctx context.Context) error
	DoAlertAfterRequestChecks(ctx context.Context, mcc []model.MeasurementCollection) error
}

type alertCheckEmailer interface {
	DoEmail(string, config.AlertCheckConfig) error
}

type alertCheckService struct {
	db *model.Database
	*model.Queries
	cfg *config.AlertCheckConfig
}

func NewAlertCheckService(db *model.Database, q *model.Queries, cfg *config.AlertCheckConfig) *alertCheckService {
	return &alertCheckService{db, q, cfg}
}
