package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentGroupStore interface {
}

type instrumentGroupStore struct {
	db *model.Database
	q  *model.Queries
}

func NewInstrumentGroupStore(db *model.Database, q *model.Queries) *instrumentGroupStore {
	return &instrumentGroupStore{db, q}
}

// ListInstrumentGroupSlugs lists used instrument group slugs in the database
func (s instrumentGroupStore) ListInstrumentGroupSlugs(ctx context.Context) ([]string, error) {
	return s.q.ListInstrumentSlugs(ctx)
}

// ListInstrumentGroups returns a list of instrument groups
func (s instrumentGroupStore) ListInstrumentGroups(ctx context.Context) ([]model.InstrumentGroup, error) {
	return s.q.ListInstrumentGroups(ctx)
}

// GetInstrumentGroup returns a single instrument group
func (s instrumentGroupStore) GetInstrumentGroup(ctx context.Context, instrumentGroupID uuid.UUID) (model.InstrumentGroup, error) {
	return s.q.GetInstrumentGroup(ctx, instrumentGroupID)
}

// CreateInstrumentGroup creates many instruments from an array of instruments
func (s instrumentGroupStore) CreateInstrumentGroup(ctx context.Context, groups []model.InstrumentGroup) ([]model.InstrumentGroup, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	q := s.q.WithTx(tx)

	gg := make([]model.InstrumentGroup, len(groups))
	for idx, g := range groups {
		gNew, err := q.CreateInstrumentGroup(ctx, g)
		if err != nil {
			return nil, err
		}
		gg[idx] = gNew
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return gg, nil
}

// UpdateInstrumentGroup updates an instrument group
func (s instrumentGroupStore) UpdateInstrumentGroup(ctx context.Context, group model.InstrumentGroup) (model.InstrumentGroup, error) {
	return s.q.UpdateInstrumentGroup(ctx, group)
}

// DeleteFlagInstrumentGroup sets the deleted field to true
func (s instrumentGroupStore) DeleteFlagInstrumentGroup(ctx context.Context, instrumentGroupID uuid.UUID) error {
	return s.q.DeleteFlagInstrumentGroup(ctx, instrumentGroupID)
}

// ListInstrumentGroupInstruments returns a list of instrument group instruments for a given instrument
func (s instrumentGroupStore) ListInstrumentGroupInstruments(ctx context.Context, groupID uuid.UUID) ([]model.Instrument, error) {
	return s.q.ListInstrumentGroupInstruments(ctx, groupID)
}

// CreateInstrumentGroupInstruments adds an instrument to an instrument group
func (s instrumentGroupStore) CreateInstrumentGroupInstruments(ctx context.Context, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) error {
	return s.q.CreateInstrumentGroupInstruments(ctx, instrumentGroupID, instrumentID)
}

// DeleteInstrumentGroupInstruments adds an instrument to an instrument group
func (s instrumentGroupStore) DeleteInstrumentGroupInstruments(ctx context.Context, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) error {
	return s.q.DeleteInstrumentGroupInstruments(ctx, instrumentGroupID, instrumentID)
}
