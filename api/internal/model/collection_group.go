package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// CollectionGroup holds information for entity collection_group
type CollectionGroup struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ProjectID uuid.UUID `json:"project_id" db:"project_id"`
	Slug      string    `json:"slug" db:"slug"`
	Name      string    `json:"name" db:"name"`
	AuditInfo
}

// CollectionGroupDetails holds same information as a CollectionGroup
// In Addition, contains array of structs; Each struct contains
// all fields for Timeseries AND additional latest_value, latest_time
type CollectionGroupDetails struct {
	CollectionGroup
	Timeseries []collectionGroupDetailsTimeseries `json:"timeseries"`
}

// collectionGroupDetailsTimeseriesItem is a Timeseries with a little bit of extra information
type collectionGroupDetailsTimeseries struct {
	Timeseries
	LatestTime  *time.Time `json:"latest_time" db:"latest_time"`
	LatestValue *float32   `json:"latest_value" db:"latest_value"`
}

const listCollectionGroups = `
	SELECT id, project_id, slug, name, creator, create_date, updater, update_date
	FROM collection_group
	WHERE project_id = $1
`

// ListCollectionGroups lists all collection groups for a project
func (q *Queries) ListCollectionGroups(ctx context.Context, projectID uuid.UUID) ([]CollectionGroup, error) {
	aa := make([]CollectionGroup, 0)
	if err := q.db.SelectContext(ctx, &aa, listCollectionGroups, projectID); err != nil {
		return nil, err
	}
	return aa, nil
}

const getCollectionGroupDetails = listCollectionGroups + `
	AND id = $2
`

// GetCollectionGroupDetails returns details for a single CollectionGroup
func (q *Queries) GetCollectionGroupDetails(ctx context.Context, projectID, collectionGroupID uuid.UUID) (CollectionGroupDetails, error) {
	var a CollectionGroupDetails
	if err := q.db.GetContext(ctx, &a, getCollectionGroupDetails, projectID, collectionGroupID); err != nil {
		return a, err
	}
	return a, nil
}

const getCollectionGroupDetailsTimeseries = `
	SELECT t.*, tm.time as latest_time, tm.value as latest_value 
	FROM collection_group_timeseries cgt 
	INNER JOIN collection_group cg on cg.id = cgt.collection_group_id 
	INNER JOIN v_timeseries t on t.id = cgt.timeseries_id 
	LEFT JOIN timeseries_measurement tm on tm.timeseries_id = t.id and tm.time = (
		SELECT time FROM timeseries_measurement 
		WHERE timeseries_id = t.id 
		ORDER BY time DESC LIMIT 1
	) 
	WHERE t.instrument_id = ANY(
		SELECT instrument_id
		FROM project_instrument
		WHERE project_id = $1
	)
	AND cgt.collection_group_id = $2 
`

// GetCollectionGroupDetails returns details for a single CollectionGroup
func (q *Queries) GetCollectionGroupDetailsTimeseries(ctx context.Context, projectID, collectionGroupID uuid.UUID) ([]collectionGroupDetailsTimeseries, error) {
	aa := make([]collectionGroupDetailsTimeseries, 0)
	if err := q.db.SelectContext(ctx, &aa, getCollectionGroupDetailsTimeseries, projectID, collectionGroupID); err != nil {
		return nil, err
	}
	return aa, nil
}

const createCollectionGroup = `
	INSERT INTO collection_group (project_id, name, slug, creator, create_date, updater, update_date)
		VALUES ($1, $2::varchar, slugify($2::varchar, 'collection_group'), $3, $4, $5, $6)
	RETURNING id, project_id, name, slug, creator, create_date, updater, update_date
`

// CreateCollectionGroup creates a new collection group
func (q *Queries) CreateCollectionGroup(ctx context.Context, cg CollectionGroup) (CollectionGroup, error) {
	var cgNew CollectionGroup
	if err := q.db.GetContext(ctx, &cgNew, createCollectionGroup, cg.ProjectID, cg.Name, cg.CreatorID, cg.CreateDate, cg.UpdaterID, cg.UpdateDate); err != nil {
		return cgNew, err
	}
	return cgNew, nil
}

const updateCollectionGroup = `
	UPDATE collection_group SET name=$3, updater=$4, update_date=$5
	WHERE project_id=$1 AND id=$2
	RETURNING id, project_id, name, slug, creator, create_date, updater, update_date
`

// UpdateCollectionGroup updates an existing collection group's metadata
func (q *Queries) UpdateCollectionGroup(ctx context.Context, cg CollectionGroup) (CollectionGroup, error) {
	var cgUpdated CollectionGroup
	if err := q.db.GetContext(ctx, &cgUpdated, updateCollectionGroup, cg.ProjectID, cg.ID, cg.Name, cg.UpdaterID, cg.UpdateDate); err != nil {
		return cgUpdated, err
	}
	return cgUpdated, nil
}

const deleteCollectionGroup = `
	DELETE FROM collection_group WHERE project_id=$1 AND id=$2
`

// DeleteCollectionGroup deletes a collection group and associated timeseries relationships
// using the id of the collection group
func (q *Queries) DeleteCollectionGroup(ctx context.Context, projectID, collectionGroupID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCollectionGroup, projectID, collectionGroupID)
	return err
}

const addTimeseriesToCollectionGroup = `
	INSERT INTO collection_group_timeseries (collection_group_id, timeseries_id) VALUES ($1, $2)
	ON CONFLICT ON CONSTRAINT collection_group_unique_timeseries DO NOTHING
`

// AddTimeseriesToCollectionGroup adds a timeseries to a collection group
func (q *Queries) AddTimeseriesToCollectionGroup(ctx context.Context, collectionGroupID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, addTimeseriesToCollectionGroup, collectionGroupID, timeseriesID)
	return err
}

const removeTimeseriesFromCollectionGroup = `
	DELETE FROM collection_group_timeseries WHERE collection_group_id=$1 AND timeseries_id = $2
`

// RemoveTimeseriesFromCollectionGroup removes a timeseries from a collection group
func (q *Queries) RemoveTimeseriesFromCollectionGroup(ctx context.Context, collectionGroupID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, removeTimeseriesFromCollectionGroup, collectionGroupID, timeseriesID)
	return err
}
