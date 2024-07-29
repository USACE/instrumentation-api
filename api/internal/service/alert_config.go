package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AlertConfigService interface {
	GetAllAlertConfigsForProject(ctx context.Context, projectID uuid.UUID) ([]model.AlertConfig, error)
	GetAllAlertConfigsForProjectAndAlertType(ctx context.Context, projectID, alertTypeID uuid.UUID) ([]model.AlertConfig, error)
	GetAllAlertConfigsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.AlertConfig, error)
	GetOneAlertConfig(ctx context.Context, alertConfigID uuid.UUID) (model.AlertConfig, error)
	DeleteAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error
	alertConfigSchedulerService
}

type alertConfigService struct {
	db *model.Database
	*model.Queries
}

func NewAlertConfigService(db *model.Database, q *model.Queries) *alertConfigService {
	return &alertConfigService{db, q}
}
