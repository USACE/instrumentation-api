package dbutils

import (
	"database/sql"
)

// Domains
func createTableInstrumentType(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS public.instrument_type (
		id UUID PRIMARY KEY NOT NULL,
		name VARCHAR(120) UNIQUE NOT NULL
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func createTableUnit(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS public.unit (
		id UUID PRIMARY KEY NOT NULL,
		name VARCHAR(120) UNIQUE NOT NULL
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func createTableParameter(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS public.parameter (
		id UUID PRIMARY KEY NOT NULL,
		name VARCHAR(120) UNIQUE NOT NULL
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func createTableInstrumentGroup(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS public.instrument_group (
		id UUID PRIMARY KEY NOT NULL,
		name VARCHAR(120) UNIQUE NOT NULL
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func createTableInstrument(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS public.instrument (
		id UUID PRIMARY KEY NOT NULL,
		name VARCHAR(120) UNIQUE NOT NULL,
		height REAL,
		geometry geometry,
		instrument_type_id UUID REFERENCES instrument_type (id),
		instrument_group_id UUID REFERENCES instrument_group (id)
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func createTableInstrumentGroupInstruments(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS public.instrument_group_instruments (
		instrument_id UUID NOT NULL REFERENCES instrument (id),
		instrument_group_id UUID NOT NULL REFERENCES instrument_group (id)
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func createTableTimeseries(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS public.timeseries (
		id UUID PRIMARY KEY NOT NULL,
		name VARCHAR(240) UNIQUE NOT NULL,
		instrument_id UUID REFERENCES instrument (id),
		parameter_id UUID NOT NULL REFERENCES parameter (id),
		unit_id UUID NOT NULL REFERENCES unit (id)
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func createTableTimeseriesMeasurement(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS public.timeseries_measurement (
		id UUID PRIMARY KEY NOT NULL,
		time TIMESTAMPTZ NOT NULL,
		value REAL NOT NULL,
		timeseries_id UUID NOT NULL REFERENCES unit (id)
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

// CreateTables creates database tables and views if they do not exist
func CreateTables(db *sql.DB) {
	// Domain Tables
	createTableInstrumentType(db)
	createTableUnit(db)
	createTableParameter(db)
	// Tables
	createTableInstrumentGroup(db)
	createTableInstrument(db)
	createTableInstrumentGroupInstruments(db)
	createTableTimeseries(db)
	createTableTimeseriesMeasurement(db)
}
