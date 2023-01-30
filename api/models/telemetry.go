package models

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jmoiron/sqlx"
)

// Telemetry struct
type Telemetry struct {
	ID       uuid.UUID
	TypeID   string
	TypeSlug string
	TypeName string
}

type DataLogger struct {
	Name      string
	SN        string
	ProjectID uuid.UUID
	CreatorID uuid.UUID
}

type DataLoggerPayload struct {
	Head Head    `json:"head"`
	Data []Datum `json:"data"`
}

type Datum struct {
	Time string    `json:"time"`
	No   int64     `json:"no"`
	Vals []float64 `json:"vals"`
}

type Head struct {
	Transaction int64       `json:"transaction"`
	Signature   int64       `json:"signature"`
	Environment Environment `json:"environment"`
	Fields      []Field     `json:"fields"`
}

type Environment struct {
	StationName string `json:"station_name"`
	TableName   string `json:"table_name"`
	Model       string `json:"model"`
	SerialNo    string `json:"serial_no"`
	OSVersion   string `json:"os_version"`
	ProgName    string `json:"prog_name"`
}

type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Units    string `json:"units"`
	Process  string `json:"process"`
	Settable bool   `json:"settable"`
}

type EquivalencyTable struct {
	DataLoggerSN string
	FieldMap     map[string]EquivalencyTableRow
}

type EquivalencyTableRow struct {
	InstrumentID uuid.UUID
	TimeseriesID uuid.UUID
}

type DataLoggerPreview struct {
	SN      string
	Payload pgtype.JSON `json:"payload"`
}

func GetDataLoggerHash(db *sqlx.DB, sn string) (string, error) {
	var hash string

	if err := db.Get(
		&hash,
		`SELECT hash
		FROM data_logger_hash
		WHERE serial_number = $1`,
		sn,
	); err != nil {
		return "", err
	}

	return hash, nil
}

func GetDataLoggerBySerialNumber(db *sqlx.DB, sn string) (*DataLogger, error) {
	var dl DataLogger

	if err := db.Get(
		&dl,
		`SELECT name, serial_number, project_id, creator_id
		FROM v_data_logger
		WHERE serial_number = $1`,
		sn,
	); err != nil {
		return nil, err
	}

	return &dl, nil
}

func GetEquivalencyTable(db *sqlx.DB, sn string) (*EquivalencyTable, error) {
	var eq EquivalencyTable

	if err := db.Get(
		&eq, ``, sn,
	); err != nil {
		return nil, err
	}

	return &eq, nil
}

func GetDataLoggerPreview(db *sqlx.DB, sn string) (*DataLoggerPreview, error) {
	var dlp DataLoggerPreview

	if err := db.Get(
		&dlp,
		`SELECT sn, payload FROM telemetry_preview WHERE sn = $1`,
		sn,
	); err != nil {
		return nil, err
	}

	return &dlp, nil
}

func UpdateDataLoggerPreview(db *sqlx.DB, dlp *DataLoggerPreview) error {
	if _, err := db.Exec(
		`UPDATE TABLE telemetry_preview SET payload = $2 WHERE sn = $1`,
		&dlp.SN, &dlp.Payload,
	); err != nil {
		return err
	}

	return nil
}

// ParseTOA5 parses a Campbell Scientific TOA5 data file that is simlar to a csv.
// The unique properties of TOA5 are that the meatdata are stored in header of file (first 4 lines of csv)
func ParseTOA5(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	// read headers
	// LINE 1: information about the data logger (e.g. serial number and program name)
	header1, err := r.Read()
	if err != nil {
		return nil, err
	}
	// LINE 2: data header (names of the variables stored in the table)
	header2, err := r.Read()
	if err != nil {
		return nil, err
	}
	// LINE 3: units for the variables if they have been defined in the data logger
	header3, err := r.Read()
	if err != nil {
		return nil, err
	}
	// LINE 4: abbreviation for processing data logger performed
	// (e.g. sample, average, standard deviation, maximum, minimum, etc.)
	header4, err := r.Read()
	if err != nil {
		return nil, err
	}
	log.Printf("header1: %v", header1)
	log.Printf("header2: %v", header2)
	log.Printf("header3: %v", header3)
	log.Printf("header4: %v", header4)

	// continue read until EOF
	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	log.Printf("data: %v", data)

	return data, nil
}
