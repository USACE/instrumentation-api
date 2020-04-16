package models

import (
	"api/root/dbutils"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/geojson"
)

// Instrument is an instrument data structure
type Instrument struct {
	ID         uuid.UUID        `json:"id"`
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
}

// ListInstruments returns an array of instruments from the database
func ListInstruments(db *sqlx.DB) []Instrument {
	sql := `SELECT instrument.id,
				   instrument.slug,
				   instrument.NAME,
				   instrument.INSTRUMENT_TYPE_ID,
	        	   instrument_type.NAME              AS instrument_type, 
	               instrument.height, 
				   ST_AsBinary(instrument.geometry) AS geometry,
				   instrument.creator,
				   instrument.create_date,
				   instrument.updater,
				   instrument.update_date
            FROM   instrument
	               INNER JOIN instrument_type
	               		   ON instrument_type.id = instrument.instrument_type_id
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
			&i.ID, &i.Slug, &i.Name, &i.TypeID, &i.Type, &i.Height, wkb.Scanner(&p),
			&i.Creator, &i.CreateDate, &i.Updater, &i.UpdateDate,
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
func GetInstrument(db *sqlx.DB, id uuid.UUID) Instrument {
	sql := `SELECT instrument.id,
	               instrument.slug,
	        	   instrument.NAME,
				   instrument.INSTRUMENT_TYPE_ID,
	        	   instrument_type.NAME              AS instrument_type,
	               instrument.height,
				   ST_AsBinary(instrument.geometry) AS geometry,
				   instrument.creator,
				   instrument.create_date,
				   instrument.updater,
				   instrument.update_date
            FROM   instrument
	               INNER JOIN instrument_type
							  ON instrument_type.id = instrument.instrument_type_id
			WHERE instrument.id = $1
			`

	var i Instrument
	var p orb.Point
	err := db.QueryRow(sql, id).Scan(&i.ID, &i.Slug, &i.Name, &i.TypeID, &i.Type, &i.Height, wkb.Scanner(&p),
		&i.Creator, &i.CreateDate, &i.Updater, &i.UpdateDate,
	)
	if err != nil {
		log.Printf("Fail to query and scan row with ID %s; %s", id, err)
	}
	i.Geometry = *geojson.NewGeometry(p)

	return i
}

// CreateInstrument creates a single instrument
func CreateInstrument(db *sqlx.DB, i *Instrument) error {

	// unique slug
	slug, err := dbutils.NextUniqueSlug(db, i.Name, "instrument", "slug")
	if err != nil {
		return err
	}
	i.Slug = slug

	if _, err := db.Exec(
		`INSERT INTO instrument (id, slug, name, height, instrument_type_id, geometry, creator, create_date, updater, update_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		i.ID, i.Slug, i.Name, i.Height, i.TypeID, wkb.Value(i.Geometry.Geometry()),
		i.Creator, i.CreateDate, i.Updater, i.UpdateDate,
	); err != nil {
		return err
	}

	return nil
}

// UpdateInstrument updates a single instrument
func UpdateInstrument(db *sqlx.DB, i *Instrument) error {

	if _, err := db.Exec(
		`UPDATE instrument
		SET name = $1, height = $2, instrument_type_id = $3, geometry = ST_GeomFromWKB($4), creator = $5, create_date = $6, updater = $7, update_date = $8
		WHERE id = $9`,
		i.Name, i.Height, i.TypeID, i.Creator, i.CreateDate, i.Updater, i.UpdateDate, wkb.Value(i.Geometry.Geometry()), i.ID,
	); err != nil {
		return err
	}

	return nil
}

// DeleteInstrument deletes a single instrument
func DeleteInstrument(db *sqlx.DB, id uuid.UUID) error {

	if _, err := db.Exec(`DELETE FROM instrument WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}
