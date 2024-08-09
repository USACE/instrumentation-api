package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type alertConfigChangeService interface {
	CreateAlertConfigChange(ctx context.Context, ac model.AlertConfigChange) (model.AlertConfig, error)
	UpdateAlertConfigChange(ctx context.Context, ac model.AlertConfigChange) (model.AlertConfig, error)
}
