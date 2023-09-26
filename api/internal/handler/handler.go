package handler

import (
	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/store"
)

type Handler struct {
	// s3mediaStore         MediaStoreHandler
	AlertStore store.AlertStore
}

func New(cfg *config.DBConfig) *Handler {
	db := model.NewDatabase(cfg)
	q := db.Queries()

	// mediaStore := NewMediaStore(cfg)

	return &Handler{
		// s3mediaStore:         NewS3MediaStore(mediaStore),
		AlertStore: store.NewAlertStore(db, q),
	}
}
