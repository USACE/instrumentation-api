package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type alertConfigThresholdService interface {
	CreateAlertConfigScheduler(ctx context.Context, ac model.AlertConfigScheduler) (model.AlertConfig, error)
	UpdateAlertConfigScheduler(ctx context.Context, ac model.AlertConfigScheduler) (model.AlertConfig, error)
}
