package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/paulmach/orb/geojson"
)

type InstrumentStore interface {
}

type instrumentStore struct {
	db *model.Database
	q  *model.Queries
}

func NewInstrumentStore(db *model.Database, q *model.Queries) *instrumentStore {
	return &instrumentStore{db, q}
}

// ListInstrumentSlugs lists used instrument slugs in the database
func (s instrumentStore) ListInstrumentSlugs(ctx context.Context) ([]string, error) {
	return s.q.ListInstrumentSlugs(ctx)
}

// ListInstruments returns an array of instruments from the database
func (s instrumentStore) ListInstruments(ctx context.Context) ([]model.Instrument, error) {
	return s.q.ListInstruments(ctx)
}

// GetInstrument returns a single instrument
func (s instrumentStore) GetInstrument(ctx context.Context, instrumentID uuid.UUID) (model.Instrument, error) {
	return s.q.GetInstrument(ctx, instrumentID)
}

// GetInstrumentCount returns the number of instruments in the database
func (s instrumentStore) GetInstrumentCount(ctx context.Context) (int, error) {
	return s.q.GetInstrumentCount(ctx)
}

func (s instrumentStore) CreateInstrument(ctx context.Context, i model.Instrument) (model.IDAndSlug, error) {
	return s.q.CreateInstrument(ctx, i)
}

func (s instrumentStore) CreateInstruments(ctx context.Context, instruments []model.Instrument) ([]model.IDAndSlug, error) {
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

	ii := make([]model.IDAndSlug, len(instruments))
	for idx, i := range instruments {
		newInstrument, err := qtx.CreateInstrument(ctx, i)
		if err != nil {
			return nil, err
		}
		ii[idx] = newInstrument

		if err := qtx.CreateOrUpdateInstrumentStatus(ctx, newInstrument.ID, i.StatusID, i.StatusTime); err != nil {
			return nil, err
		}

		if i.AwareID != nil {
			if err := qtx.CreateAwarePlatform(ctx, newInstrument.ID, *i.AwareID); err != nil {
				return nil, err
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return ii, nil
}

// ValidateCreateInstruments creates many instruments from an array of instruments
func (s instrumentStore) ValidateCreateInstruments(ctx context.Context, instruments []model.Instrument) (model.CreateInstrumentsValidationResult, error) {
	return s.q.ValidateCreateInstruments(ctx, instruments)
}

// UpdateInstrument updates a single instrument
func (s instrumentStore) UpdateInstruments(ctx context.Context, i model.Instrument) (model.Instrument, error) {
	e := model.Instrument{}
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return e, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	if err := qtx.UpdateInstrument(ctx, i); err != nil {
		return e, err
	}
	if err := qtx.CreateOrUpdateInstrumentStatus(ctx, i.ID, i.StatusID, i.StatusTime); err != nil {
		return e, err
	}

	aa, err := qtx.GetInstrument(ctx, i.ID)
	if err != nil {
		return e, err
	}

	if err := tx.Commit(); err != nil {
		return e, err
	}

	return aa, nil
}

func (s instrumentStore) UpdateInstrumentGeometry(ctx context.Context, projectID, instrumentID uuid.UUID, geom geojson.Geometry, p model.Profile) (model.Instrument, error) {
	e := model.Instrument{}
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return e, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	if err := qtx.UpdateInstrumentGeometry(ctx, projectID, instrumentID, geom, p); err != nil {
		return e, err
	}

	aa, err := qtx.GetInstrument(ctx, instrumentID)
	if err != nil {
		return e, err
	}

	if err := tx.Commit(); err != nil {
		return e, err
	}

	return aa, nil
}

// DeleteFlagInstrument changes delete flag to true
func (s instrumentStore) DeleteFlagInstrument(ctx context.Context, projectID, instrumentID uuid.UUID) error {
	return s.q.DeleteFlagInstrument(ctx, projectID, instrumentID)
}
