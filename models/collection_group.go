package models

import (
	"time"

	ts "github.com/USACE/instrumentation-api/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CollectionGroup holds information for entity collection_group
type CollectionGroup struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id" db:"project_id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	AuditInfo
}

// CollectionGroupDetails holds same information as a CollectionGroup
// In Addition, contains array of structs; Each struct contains
// all fields for ts.Timeseries AND additional latest_value, latest_time
type CollectionGroupDetails struct {
	CollectionGroup
	Timeseries []cgdTsItem `json:"timeseries"`
}

// collectionGroupDetailsTimeseriesItem is a ts.Timeseries with a little bit of extra information
type cgdTsItem struct {
	ts.Timeseries
	LatestTime  *time.Time `json:"latest_time" db:"latest_time"`
	LatestValue *float32   `json:"latest_value" db:"latest_value"`
}

var listCollectionGroupsSQL = `
	SELECT id, project_id, slug, name, creator, create_date, updater, update_date
	FROM collection_group
`

// ListCollectionGroups lists all collection groups for a project
func ListCollectionGroups(db *sqlx.DB, projectID *uuid.UUID) ([]CollectionGroup, error) {
	cc := make([]CollectionGroup, 0)
	if err := db.Select(&cc, listCollectionGroupsSQL+" WHERE project_id = $1", projectID); err != nil {
		return make([]CollectionGroup, 0), err
	}
	return cc, nil
}

// ListCollectionGroupSlugs lists all collection group slugs for a project
func ListCollectionGroupSlugs(db *sqlx.DB, projectID *uuid.UUID) ([]string, error) {
	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM instrument_group WHERE project_id=$1", projectID); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

// GetCollectionGroupDetails returns details for a single CollectionGroup
func GetCollectionGroupDetails(db *sqlx.DB, projectID *uuid.UUID, collectionGroupID *uuid.UUID) (*CollectionGroupDetails, error) {
	var d CollectionGroupDetails
	// Query 1
	if err := db.Get(&d, listCollectionGroupsSQL+" WHERE project_id=$1 AND id=$2", projectID, collectionGroupID); err != nil {
		return nil, err
	}
	// Query 2
	d.Timeseries = make([]cgdTsItem, 0)
	if err := db.Select(
		&d.Timeseries,
		`SELECT t.*, tm.time as latest_time, tm.value as latest_value 
		FROM collection_group_timeseries cgt 
		INNER JOIN collection_group cg on cg.id = cgt.collection_group_id 
		INNER JOIN v_timeseries t on t.id = cgt.timeseries_id 
		INNER JOIN timeseries_measurement tm on tm.timeseries_id = t.id and tm.time = (
			select time from timeseries_measurement 
			where timeseries_id = t.id 
			order by time desc limit 1) 
		WHERE cgt.collection_group_id = $2 and t.project_id = $1
		 `, projectID, collectionGroupID,
	); err != nil {
		return nil, err
	}

	return &d, nil
}

// CreateCollectionGroup creates a new collection group
func CreateCollectionGroup(db *sqlx.DB, cg *CollectionGroup) (*CollectionGroup, error) {
	var cgNew CollectionGroup
	if err := db.Get(
		&cgNew,
		`INSERT INTO collection_group (project_id, name, slug, creator, create_date, updater, update_date) VALUES
			 ($1, $2, $3, $4, $5, $6, $7)
		 RETURNING id, project_id, name, slug, creator, create_date, updater, update_date
		 `, cg.ProjectID, cg.Name, cg.Slug, cg.Creator, cg.CreateDate, cg.Updater, cg.UpdateDate,
	); err != nil {
		return nil, err
	}
	return &cgNew, nil
}

// UpdateCollectionGroup updates an existing collection group's metadata
func UpdateCollectionGroup(db *sqlx.DB, cg *CollectionGroup) (*CollectionGroup, error) {
	var cgUpdated CollectionGroup
	if err := db.Get(
		&cgUpdated,
		`UPDATE collection_group SET name=$3, updater=$4, update_date=$5 WHERE project_id=$1 AND id=$2
		 RETURNING id, project_id, name, slug, creator, create_date, updater, update_date
		`, cg.ProjectID, cg.ID, cg.Name, cg.Updater, cg.UpdateDate,
	); err != nil {
		return nil, err
	}
	return &cgUpdated, nil
}

// DeleteCollectionGroup deletes a collection group and associated timeseries relationships
// using the id of the collection group
func DeleteCollectionGroup(db *sqlx.DB, projectID *uuid.UUID, collectionGroupID *uuid.UUID) error {
	if _, err := db.Exec(
		`DELETE FROM collection_group WHERE project_id=$1 AND id=$2`, projectID, collectionGroupID,
	); err != nil {
		return err
	}
	return nil
}

// AddTimeseriesToCollectionGroup adds a timeseries to a collection group
func AddTimeseriesToCollectionGroup(db *sqlx.DB, collectionGroupID *uuid.UUID, timeseriesID *uuid.UUID) error {
	if _, err := db.Exec(
		`INSERT INTO collection_group_timeseries (collection_group_id, timeseries_id) VALUES ($1, $2)
		 ON CONFLICT ON CONSTRAINT collection_group_unique_timeseries DO NOTHING;
		 `, collectionGroupID, timeseriesID,
	); err != nil {
		return err
	}
	return nil
}

// RemoveTimeseriesFromCollectionGroup removes a timeseries from a collection group
func RemoveTimeseriesFromCollectionGroup(db *sqlx.DB, collectionGroupID *uuid.UUID, timeseriesID *uuid.UUID) error {
	if _, err := db.Exec(
		`DELETE FROM collection_group_timeseries WHERE collection_group_id=$1 AND timeseries_id=$2`,
		collectionGroupID, timeseriesID,
	); err != nil {
		return err
	}
	return nil
}
