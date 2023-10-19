package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentGroupService interface {
	ListInstrumentGroupSlugs(ctx context.Context) ([]string, error)
	ListInstrumentGroups(ctx context.Context) ([]model.InstrumentGroup, error)
	GetInstrumentGroup(ctx context.Context, instrumentGroupID uuid.UUID) (model.InstrumentGroup, error)
	CreateInstrumentGroup(ctx context.Context, groups []model.InstrumentGroup) ([]model.InstrumentGroup, error)
	UpdateInstrumentGroup(ctx context.Context, group model.InstrumentGroup) (model.InstrumentGroup, error)
	DeleteFlagInstrumentGroup(ctx context.Context, instrumentGroupID uuid.UUID) error
	ListInstrumentGroupInstruments(ctx context.Context, groupID uuid.UUID) ([]model.Instrument, error)
	CreateInstrumentGroupInstruments(ctx context.Context, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) error
	DeleteInstrumentGroupInstruments(ctx context.Context, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) error
}

type instrumentGroupService struct {
	db *model.Database
	*model.Queries
}

func NewInstrumentGroupService(db *model.Database, q *model.Queries) *instrumentGroupService {
	return &instrumentGroupService{db, q}
}

// CreateInstrumentGroup creates many instruments from an array of instruments
func (s instrumentGroupService) CreateInstrumentGroup(ctx context.Context, groups []model.InstrumentGroup) ([]model.InstrumentGroup, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	q := s.WithTx(tx)

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
