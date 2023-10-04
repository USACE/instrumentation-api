package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentNoteService interface {
	ListInstrumentNotes(ctx context.Context) ([]model.InstrumentNote, error)
	ListInstrumentInstrumentNotes(ctx context.Context, instrumentID uuid.UUID) ([]model.InstrumentNote, error)
	GetInstrumentNote(ctx context.Context, noteID uuid.UUID) (model.InstrumentNote, error)
	CreateInstrumentNote(ctx context.Context, notes []model.InstrumentNote) ([]model.InstrumentNote, error)
	UpdateInstrumentNote(ctx context.Context, n model.InstrumentNote) (model.InstrumentNote, error)
	DeleteInstrumentNote(ctx context.Context, noteID uuid.UUID) error
}

type instrumentNoteService struct {
	db *model.Database
	*model.Queries
}

func NewInstrumentNoteService(db *model.Database, q *model.Queries) *instrumentNoteService {
	return &instrumentNoteService{db, q}
}

// CreateInstrumentNote creates many instrument notes from an array of instrument notes
func (s instrumentNoteService) CreateInstrumentNote(ctx context.Context, notes []model.InstrumentNote) ([]model.InstrumentNote, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	nn := make([]model.InstrumentNote, len(notes))
	for idx, n := range notes {
		noteNew, err := qtx.CreateInstrumentNote(ctx, n)
		if err != nil {
			return nil, err
		}
		nn[idx] = noteNew
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nn, nil
}
