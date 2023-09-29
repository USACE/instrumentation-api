package handler

import (
	"context"
)

func (h *AlertCheckHandler) DoAlertChecks() error {
	ctx := context.Background()

	return h.AlertCheckStore.DoAlertChecks(ctx)
}
