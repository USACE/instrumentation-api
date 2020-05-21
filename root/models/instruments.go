package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/encoding/wkt"
	"github.com/paulmach/orb/geojson"
)

// Instrument is an instrument data structure
type Instrument struct {
	ID         uuid.UUID        `json:"id"`
	Active     bool             `json:"active"`
	Deleted    bool             `json:"-"`
	Slug       string           `json:"slug"`
	Name       string           `json:"name"`
	TypeID     string           `json:"type_id"`
	Type       string           `json:"type"`
	Height     float32          `json:"height"`
	Geometry   geojson.Geometry `json:"geometry,omitempty"`
	Creator    int              `json:"creator"`
	CreateDate time.Time        `json:"create_date"`
	Updater    int              `json:"updater"`
	UpdateDate time.Time        `json:"update_date"`
	ProjectID  *string          `json:"project_id" db:"project_id"`
}

// ListInstrumentSlugs lists used instrument slugs in the database
func ListInstrumentSlugs(db *sqlx.DB) ([]string, error) {

	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM instrument"); err != nil {
		return make([]string, 0), err
	}

	return ss, nil

}

// ListInstruments returns an array of instruments from the database
func ListInstruments(db *sqlx.DB) []Instrument {
	sql := `SELECT instrument.id,
				   instrument.active,
				   instrument.slug,
				   instrument.NAME,
				   instrument.INSTRUMENT_TYPE_ID,
	        	   instrument_type.NAME              AS instrument_type, 
	               instrument.height, 
				   ST_AsBinary(instrument.geometry) AS geometry,
				   instrument.creator,
				   instrument.create_date,
				   instrument.updater,
				   instrument.update_date,
				   instrument.project_id
			FROM   instrument
	               INNER JOIN instrument_type
							  ON instrument_type.id = instrument.instrument_type_id
		   WHERE NOT instrument.deleted
			`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]Instrument, 0)
	for rows.Next() {
		var p orb.Point
		var i Instrument
		err := rows.Scan(
			&i.ID, &i.Active, &i.Slug, &i.Name, &i.TypeID, &i.Type, &i.Height, wkb.Scanner(&p),
			&i.Creator, &i.CreateDate, &i.Updater, &i.UpdateDate, &i.ProjectID,
		)
		i.Geometry = *geojson.NewGeometry(p)

		if err != nil {
			panic(err)
		}

		result = append(result, i)
	}
	return result
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB, id uuid.UUID) (*Instrument, error) {
	sql := `SELECT instrument.id,
	               instrument.active,
	               instrument.slug,
	        	   instrument.NAME,
				   instrument.INSTRUMENT_TYPE_ID,
	        	   instrument_type.NAME              AS instrument_type,
	               instrument.height,
				   ST_AsBinary(instrument.geometry) AS geometry,
				   instrument.creator,
				   instrument.create_date,
				   instrument.updater,
				   instrument.update_date,
				   instrument.project_id
            FROM   instrument
	               INNER JOIN instrument_type
							  ON instrument_type.id = instrument.instrument_type_id
			WHERE instrument.id = $1
			`

	var i Instrument
	var p orb.Point
	err := db.QueryRow(sql, id).Scan(&i.ID, &i.Active, &i.Slug, &i.Name, &i.TypeID, &i.Type, &i.Height, wkb.Scanner(&p),
		&i.Creator, &i.CreateDate, &i.Updater, &i.UpdateDate, &i.ProjectID,
	)
	if err != nil {
		return nil, err
	}
	i.Geometry = *geojson.NewGeometry(p)

	return &i, nil
}

// CreateInstrumentBulk creates many instruments from an array of instruments
func CreateInstrumentBulk(db *sqlx.DB, instruments []Instrument) error {

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn(
		"instrument",
		"id", "active", "slug", "name", "height", "instrument_type_id",
		"geometry", "creator", "create_date", "updater", "update_date", "project_id",
	))

	if err != nil {
		log.Fatal(err)
	}

	for _, i := range instruments {

		_, err = stmt.Exec(
			i.ID, i.Active, i.Slug, i.Name, i.Height, i.TypeID,
			wkt.MarshalString(i.Geometry.Geometry()), i.Creator, i.CreateDate, i.Updater, i.UpdateDate, i.ProjectID,
		)

		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

// UpdateInstrument updates a single instrument
func UpdateInstrument(db *sqlx.DB, i *Instrument) (*Instrument, error) {

	var iUpdated Instrument
	if err := db.QueryRowx(
		`UPDATE instrument
		 SET    name = $2,
			    active = $3,
			    height = $4,
			    instrument_type_id = $5,
			    geometry = ST_GeomFromWKB($6),
			    updater = $7,
				update_date = $8,
				project_id = $9
		 WHERE id = $1
		 RETURNING *
		`, i.ID, i.Name, i.Active, i.Height, i.TypeID, wkb.Value(i.Geometry.Geometry()), i.Updater, i.UpdateDate, i.ProjectID,
	).StructScan(&iUpdated); err != nil {
		return nil, err
	}

	return &iUpdated, nil
}

// DeleteFlagInstrument changes delete flag to true
func DeleteFlagInstrument(db *sqlx.DB, id uuid.UUID) error {

	if _, err := db.Exec(`UPDATE instrument SET deleted = true WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}
