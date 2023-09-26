package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentNoteStore interface {
}

type instrumentNoteStore struct {
	db *model.Database
	q  *model.Queries
}

func NewInstrumentNoteStore(db *model.Database, q *model.Queries) *instrumentNoteStore {
	return &instrumentNoteStore{db, q}
}

// ListInstrumentNotes returns an array of instruments from the database
func (s instrumentNoteStore) ListInstrumentNotes(ctx context.Context) ([]model.InstrumentNote, error) {
	return s.q.ListInstrumentNotes(ctx)
}

// ListInstrumentInstrumentNotes returns an array of instrument notes for a given instrument
func (s instrumentNoteStore) ListInstrumentInstrumentNotes(ctx context.Context, instrumentID uuid.UUID) ([]model.InstrumentNote, error) {
	return s.q.ListInstrumentInstrumentNotes(ctx, instrumentID)
}

// GetInstrumentNote returns a single instrument note
func (s instrumentNoteStore) GetInstrumentNote(ctx context.Context, noteID uuid.UUID) (model.InstrumentNote, error) {
	return s.q.GetInstrumentNote(ctx, noteID)
}

// CreateInstrumentNote creates many instrument notes from an array of instrument notes
func (s instrumentNoteStore) CreateInstrumentNote(ctx context.Context, notes []model.InstrumentNote) ([]model.InstrumentNote, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

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

// UpdateInstrumentNote updates a single instrument note
func (s instrumentNoteStore) UpdateInstrumentNote(ctx context.Context, n model.InstrumentNote) (model.InstrumentNote, error) {
	return s.q.UpdateInstrumentNote(ctx, n)
}

// DeleteInstrumentNote deletes an instrument note
func (s instrumentNoteStore) DeleteInstrumentNote(ctx context.Context, noteID uuid.UUID) error {
	return s.q.DeleteInstrumentNote(ctx, noteID)
}
