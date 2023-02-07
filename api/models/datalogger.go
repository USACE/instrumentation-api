package models

import (
	"time"

	"github.com/USACE/instrumentation-api/api/passwords"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Telemetry struct
type Telemetry struct {
	ID       uuid.UUID
	TypeID   string
	TypeSlug string
	TypeName string
}

type DataLogger struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	SN         string     `json:"sn" db:"sn"`
	ProjectID  uuid.UUID  `json:"project_id" db:"project_id"`
	Creator    uuid.UUID  `json:"creator" db:"creator"`
	CreateDate time.Time  `json:"create_date" db:"create_date"`
	Updater    *uuid.UUID `json:"updater" db:"updater"`
	UpdateDate *time.Time `json:"update_date" db:"update_date"`
	Slug       string     `json:"slug" db:"slug"`
	Model      string     `json:"model" db:"model"`
	Deleted    bool       `json:"-" db:"deleted"`
}

type DataLoggerWithKey struct {
	DataLogger
	Key string `json:"key"`
}

type EquivalencyTable struct {
	DataLoggerID uuid.UUID `json:"datalogger_id"`
	Rows         []EquivalencyTableRow
}

type EquivalencyTableRow struct {
	DataLoggerID uuid.UUID  `json:"-"`
	FieldName    string     `json:"field_name"`
	DisplayName  string     `json:"display_name"`
	InstrumentID *uuid.UUID `json:"instrument_id"`
	TimeseriesID *uuid.UUID `json:"timeseries_id"`
}

type DataLoggerPreview struct {
	SN      string
	Payload pgtype.JSON `json:"payload"`
}

func ListProjectDataLoggers(db *sqlx.DB, projectID *uuid.UUID) ([]DataLogger, error) {
	dls := make([]DataLogger, 0)

	if err := db.Select(
		&dls, `SELECT * FROM datalogger WHERE project_id = $1`, projectID,
	); err != nil {
		return make([]DataLogger, 0), err
	}
	return dls, nil
}

func ListAllDataLoggers(db *sqlx.DB) ([]DataLogger, error) {
	dls := make([]DataLogger, 0)

	if err := db.Select(&dls, `SELECT * FROM datalogger`); err != nil {
		return make([]DataLogger, 0), err
	}
	return dls, nil
}

func VerifyUniqueSN(db *sqlx.DB, sn string) error {
	// check if datalogger with sn already exists and is not deleted
	var snExists bool
	if err := db.Get(&snExists, `SELECT EXISTS (SELECT sn FROM v_datalogger WHERE sn = $1)::int`, sn); err != nil {
		return err
	}
	if snExists {
		return errors.Errorf("active data logger with serial number %s already exists", sn)
	}
	return nil
}

func CreateDataLogger(db *sqlx.DB, n *DataLogger) (*DataLoggerWithKey, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`
		INSERT INTO datalogger (name, sn, project_id, creator, updater, slug, model)
			VALUES ($1, $2, $3, $4, $4, $5, $6) RETURNING *
	`)
	if err != nil {
		return nil, err
	}

	stmt2, err := txn.Preparex(`INSERT INTO datalogger_hash (datalogger_id, "hash") VALUES ($1, $2)`)
	if err != nil {
		return nil, err
	}

	var dl DataLogger
	if err := stmt1.Get(&dl, n.Name, n.SN, n.ProjectID, n.Creator, n.Slug, n.Model); err != nil {
		return nil, err
	}

	key := passwords.GenerateRandom(40)
	hash := passwords.MustCreateHash(key, passwords.DefaultParams)
	_, err = stmt2.Exec(&dl.ID, hash)
	if err != nil {
		return nil, err
	}

	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}

	dk := DataLoggerWithKey{
		DataLogger: dl,
		Key:        key,
	}

	return &dk, nil
}

func CycleDataLoggerKey(db *sqlx.DB, u *DataLogger) (*DataLoggerWithKey, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`UPDATE datalogger_hash SET "hash" = $2 WHERE datalogger_id = $1`)
	if err != nil {
		return nil, err
	}
	stmt2, err := txn.Preparex(`UPDATE datalogger SET updater = $2, update_date = $3 WHERE id = $1 RETURNING *`)
	if err != nil {
		return nil, err
	}

	key := passwords.GenerateRandom(40)
	hash := passwords.MustCreateHash(key, passwords.DefaultParams)
	_, err = stmt1.Exec(&u.ID, hash)
	if err != nil {
		return nil, err
	}

	var dl DataLogger
	if err := stmt2.Get(&dl, &u.ID, &u.Updater, &u.UpdateDate); err != nil {
		return nil, err
	}

	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}

	dk := DataLoggerWithKey{
		DataLogger: dl,
		Key:        key,
	}

	return &dk, nil
}

func GetDataLogger(db *sqlx.DB, dlID *uuid.UUID) (*DataLogger, error) {
	var dl DataLogger
	if err := db.Get(&dl, `SELECT * FROM v_datalogger WHERE id = $1`, dlID); err != nil {
		return nil, err
	}
	return &dl, nil
}

func UpdateDataLogger(db *sqlx.DB, u *DataLogger) (*DataLogger, error) {
	var dl DataLogger
	err := db.Get(&dl, `
		UPDATE datalogger SET
			name = $2,
			sn = $3,
			model = $4,
			updater = $5,
			update_date = $6
		WHERE id = $1
		RETURNING *
		`, &u.ID, &u.Name, &u.SN, &u.Model, &u.Updater, &u.UpdateDate,
	)
	if err != nil {
		return nil, err
	}
	return &dl, nil
}

func DeleteDataLogger(db *sqlx.DB, d *DataLogger) error {
	if _, err := db.Exec(
		`UPDATE datalogger SET deleted = true, updater = $2, update_date = $3  WHERE id = $1`,
		&d.ID, &d.Updater, &d.UpdateDate,
	); err != nil {
		return err
	}

	return nil
}

func GetEquivalencyTable(db *sqlx.DB, dlID *uuid.UUID) (*EquivalencyTable, error) {
	var t EquivalencyTable
	if err := db.Get(&t, `
		SELECT * FROM v_datalogger_field_instrument_timeseries WHERE datalogger_id = $1
	`, &dlID); err != nil {
		return &EquivalencyTable{
			DataLoggerID: *dlID,
			Rows:         make([]EquivalencyTableRow, 0),
		}, err
	}

	return &t, nil
}

func CreateOrUpdateEquivalencyTableRow(db *sqlx.DB, u *EquivalencyTableRow) (*EquivalencyTableRow, error) {
	var r EquivalencyTableRow
	if err := db.Get(&r, `
		INSERT INTO datalogger_field_instrument_timeseries
		(datalogger_id, field_name, display_name, instrument_id, timeseries_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
		ON CONFLICT ON CONSTRAINT unique_datalogger_field DO UPDATE SET
			display_name = EXCLUDED.display_name,
			instrument_id = EXCLUDED.instrument_id,
			timeseries_id = EXCLUDED.timeseries_id
	`, &u.DataLoggerID); err != nil {
		return nil, err
	}

	return &r, nil
}

func DeleteEquivalencyTable(db *sqlx.DB, dlID *uuid.UUID) error {
	_, err := db.Exec(`DELETE FROM datalogger_field_instrument_timeseries WHERE datalogger_id = $1`, &dlID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEquivalencyTableRow(db *sqlx.DB, dlID *uuid.UUID, field string) error {
	if _, err := db.Exec(`
		DELETE FROM datalogger_field_instrument_timeseries WHERE datalogger_id = $1 AND field_name = $2
	`, &dlID, field); err != nil {
		return err
	}
	return nil
}

func GetDataLoggerPreview(db *sqlx.DB, sn string) (*DataLoggerPreview, error) {
	var dlp DataLoggerPreview

	if err := db.Get(&dlp, `SELECT * FROM v_datalogger_preview WHERE sn = $1`, sn); err != nil {
		return nil, err
	}

	return &dlp, nil
}
