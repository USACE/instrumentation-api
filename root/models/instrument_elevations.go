package models

// Work in progress

// import (
// 	// pq library
// 	"api/root/timeseries"
// 	_ "api/root/timeseries"

// 	"github.com/google/uuid"
// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/lib/pq"
// )

// func ListInstrumentElevations(db *sqlx.DB, instrumentID *uuid.UUID) ([]TimeseriesMeasurement, error) {

// 	// Get Instrument ZReference
// 	zz, err := ListInstrumentZReference(db, instrumentID)
// 	if err != nil {
// 		return make([]TimeseriesMeasurement, 0), err
// 	}

// 	// Get Instrument Measurements for Time Window
// 	mm, err := ListTimeseriesMeasurements(db, instrument)

// 	// Transform to Elevations & Return
// 	return timeseries.Shifter(mm, zz)

// }
