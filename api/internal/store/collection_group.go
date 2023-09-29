package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type CollectionGroupStore interface {
	ListCollectionGroups(ctx context.Context, projectID uuid.UUID) ([]model.CollectionGroup, error)
	ListCollectionGroupSlugs(ctx context.Context, projectID uuid.UUID) ([]string, error)
	GetCollectionGroupDetails(ctx context.Context, projectID, collectionGroupID uuid.UUID) (model.CollectionGroupDetails, error)
	CreateCollectionGroup(ctx context.Context, cg model.CollectionGroup) (model.CollectionGroup, error)
	UpdateCollectionGroup(ctx context.Context, cg model.CollectionGroup) (model.CollectionGroup, error)
	DeleteCollectionGroup(ctx context.Context, projectID, collectionGroupID uuid.UUID) error
	AddTimeseriesToCollectionGroup(ctx context.Context, collectionGroupID, timeseriesID uuid.UUID) error
	RemoveTimeseriesFromCollectionGroup(ctx context.Context, collectionGroupID, timeseriesID uuid.UUID) error
}

type collectionGroupStore struct {
	db *model.Database
	*model.Queries
}

func NewCollectionGroupStore(db *model.Database, q *model.Queries) *collectionGroupStore {
	return &collectionGroupStore{db, q}
}

// GetCollectionGroupDetails returns details for a single CollectionGroup
func (s collectionGroupStore) GetCollectionGroupDetails(ctx context.Context, projectID, collectionGroupID uuid.UUID) (model.CollectionGroupDetails, error) {
	var a model.CollectionGroupDetails
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.WithTx(tx)

	cg, err := qtx.GetCollectionGroupDetails(ctx, projectID, collectionGroupID)
	if err != nil {
		return a, err
	}
	ts, err := qtx.GetCollectionGroupDetailsTimeseries(ctx, projectID, collectionGroupID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	cg.Timeseries = ts

	return cg, nil
}
