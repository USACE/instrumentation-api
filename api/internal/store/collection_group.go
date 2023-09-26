package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type CollectionGroupStore interface {
}

type collectionGroupStore struct {
	db *model.Database
	q  *model.Queries
}

func NewCollectionGroupStore(db *model.Database, q *model.Queries) *collectionGroupStore {
	return &collectionGroupStore{db, q}
}

// ListCollectionGroups lists all collection groups for a project
func (s collectionGroupStore) ListCollectionGroups(ctx context.Context, projectID uuid.UUID) ([]model.CollectionGroup, error) {
	return s.q.ListCollectionGroups(ctx, projectID)
}

// ListCollectionGroupSlugs lists all collection group slugs for a project
func (s collectionGroupStore) ListCollectionGroupSlugs(ctx context.Context, projectID uuid.UUID) ([]string, error) {
	return s.q.ListCollectionGroupSlugs(ctx, projectID)
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

	qtx := s.q.WithTx(tx)

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

// CreateCollectionGroup creates a new collection group
func (s collectionGroupStore) CreateCollectionGroup(ctx context.Context, cg model.CollectionGroup) (model.CollectionGroup, error) {
	return s.q.CreateCollectionGroup(ctx, cg)
}

// UpdateCollectionGroup updates an existing collection group's metadata
func (s collectionGroupStore) UpdateCollectionGroup(ctx context.Context, cg model.CollectionGroup) (model.CollectionGroup, error) {
	return s.q.UpdateCollectionGroup(ctx, cg)
}

// DeleteCollectionGroup deletes a collection group and associated timeseries relationships
// using the id of the collection group
func (s collectionGroupStore) DeleteCollectionGroup(ctx context.Context, projectID, collectionGroupID uuid.UUID) error {
	return s.q.DeleteCollectionGroup(ctx, projectID, collectionGroupID)
}

// AddTimeseriesToCollectionGroup adds a timeseries to a collection group
func (s collectionGroupStore) AddTimeseriesToCollectionGroup(ctx context.Context, collectionGroupID, timeseriesID uuid.UUID) error {
	return s.q.AddTimeseriesToCollectionGroup(ctx, collectionGroupID, timeseriesID)
}

// RemoveTimeseriesFromCollectionGroup removes a timeseries from a collection group
func (s collectionGroupStore) RemoveTimeseriesFromCollectionGroup(ctx context.Context, collectionGroupID, timeseriesID uuid.UUID) error {
	return s.q.RemoveTimeseriesFromCollectionGroup(ctx, collectionGroupID, timeseriesID)
}
