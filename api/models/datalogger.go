package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/USACE/instrumentation-api/api/passwords"
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
	ID              uuid.UUID  `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	SN              string     `json:"sn" db:"sn"`
	ProjectID       uuid.UUID  `json:"project_id" db:"project_id"`
	Creator         uuid.UUID  `json:"creator" db:"creator"`
	CreatorUsername string     `json:"creator_username" db:"creator_username"`
	CreateDate      time.Time  `json:"create_date" db:"create_date"`
	Updater         *uuid.UUID `json:"updater" db:"updater"`
	UpdaterUsername string     `json:"updater_username" db:"updater_username"`
	UpdateDate      *time.Time `json:"update_date" db:"update_date"`
	Slug            string     `json:"slug" db:"slug"`
	ModelID         uuid.UUID  `json:"model_id" db:"model_id"`
	Model           *string    `json:"model" db:"model"`
	Deleted         bool       `json:"-" db:"deleted"`
}

type DataLoggerWithKey struct {
	DataLogger
	Key string `json:"key"`
}

type DataLoggerPreview struct {
	DataLoggerID uuid.UUID               `json:"datalogger_id" db:"datalogger_id"`
	UpdateDate   time.Time               `json:"update_date" db:"update_date"`
	Preview      pgtype.JSON             `json:"preview" db:"preview"`
	Model        *string                 `json:"model,omitempty"`
	SN           *string                 `json:"sn,omitempty"`
	Errors       DataLoggerErrorMessages `json:"errors" db:"errors"`
}

type DataLoggerError struct {
	DataLoggerID uuid.UUID               `json:"datalogger_id" db:"datalogger_id"`
	Errors       DataLoggerErrorMessages `json:"errors" db:"errors"`
}

type DataLoggerErrorMessages []string

func (t *DataLoggerErrorMessages) Scan(v interface{}) error {
	if v == nil {
		*t = DataLoggerErrorMessages{}
		return nil
	}
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("scan expected string, got [%+v]", v)
	}
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	if len(s) == 0 {
		*t = make([]string, 0)
	} else {
		*t = strings.Split(s, ",")
	}
	return nil
}
func (t *DataLoggerErrorMessages) Value() (driver.Value, error) {
	s := fmt.Sprintf("{%v}", strings.Join(([]string)(*t), ","))
	return s, nil
}

func GetDataLoggerModel(db *sqlx.DB, modelID *uuid.UUID) (string, error) {
	var model string

	if err := db.Get(&model, `SELECT model FROM datalogger_model WHERE id = $1`, modelID); err != nil {
		return "", err
	}
	return model, nil
}

func ListProjectDataLoggers(db *sqlx.DB, projectID *uuid.UUID) ([]DataLogger, error) {
	dls := make([]DataLogger, 0)

	if err := db.Select(
		&dls, `SELECT * FROM v_datalogger WHERE project_id = $1`, projectID,
	); err != nil {
		return make([]DataLogger, 0), err
	}
	return dls, nil
}

func ListAllDataLoggers(db *sqlx.DB) ([]DataLogger, error) {
	dls := make([]DataLogger, 0)

	if err := db.Select(&dls, `SELECT * FROM v_datalogger`); err != nil {
		return make([]DataLogger, 0), err
	}
	return dls, nil
}

func DataLoggerActive(db *sqlx.DB, model, sn string) (bool, error) {
	// check if datalogger with sn already exists and is not deleted
	var exists bool
	if err := db.Get(&exists, `SELECT EXISTS (SELECT * FROM v_datalogger WHERE model = $1 AND sn = $2)::int`, model, sn); err != nil {
		return false, err
	}
	return exists, nil
}

func VerifyDataLoggerExists(db *sqlx.DB, dlID *uuid.UUID) error {
	// check if datalogger with sn already exists and is not deleted
	var dlExists bool
	if err := db.Get(&dlExists, `SELECT EXISTS (SELECT id FROM v_datalogger WHERE id = $1)::int`, &dlID); err != nil {
		return err
	}
	if !dlExists {
		return fmt.Errorf("active data logger with id %s not found", dlID)
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
		INSERT INTO datalogger (name, sn, project_id, creator, updater, slug, model_id)
			VALUES ($1, $2, $3, $4, $4, $5, $6) RETURNING id`)
	if err != nil {
		return nil, err
	}

	stmt2, err := txn.Preparex(`INSERT INTO datalogger_hash (datalogger_id, "hash") VALUES ($1, $2)`)
	if err != nil {
		return nil, err
	}

	stmt3, err := txn.Preparex(`INSERT INTO datalogger_preview (datalogger_id) VALUES ($1)`)
	if err != nil {
		return nil, err
	}

	stmt4, err := txn.Preparex(`SELECT * FROM v_datalogger WHERE id = $1`)
	if err != nil {
		return nil, err
	}

	// create datalogger
	var dlID uuid.UUID
	if err := stmt1.Get(&dlID, n.Name, n.SN, n.ProjectID, n.Creator, n.Slug, n.ModelID); err != nil {
		return nil, err
	}

	// store hash
	key := passwords.GenerateRandom(40)
	hash := passwords.MustCreateHash(key, passwords.DefaultParams)
	_, err = stmt2.Exec(&dlID, hash)
	if err != nil {
		return nil, err
	}

	// create preview
	_, err = stmt3.Exec(&dlID)
	if err != nil {
		return nil, err
	}

	// return datalogger view
	var dl DataLogger
	if err = stmt4.Get(&dl, &dlID); err != nil {
		return nil, err
	}

	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := stmt3.Close(); err != nil {
		return nil, err
	}
	if err := stmt4.Close(); err != nil {
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
	stmt2, err := txn.Preparex(`UPDATE datalogger SET updater = $2, update_date = $3 WHERE id = $1`)
	if err != nil {
		return nil, err
	}
	stmt3, err := txn.Preparex(`SELECT * FROM v_datalogger WHERE id = $1`)
	if err != nil {
		return nil, err
	}

	key := passwords.GenerateRandom(40)
	hash := passwords.MustCreateHash(key, passwords.DefaultParams)
	if _, err = stmt1.Exec(&u.ID, hash); err != nil {
		return nil, err
	}

	if _, err := stmt2.Exec(&u.ID, &u.Updater, &u.UpdateDate); err != nil {
		return nil, err
	}

	var dl DataLogger
	if err = stmt3.Get(&dl, &u.ID); err != nil {
		return nil, err
	}

	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := stmt3.Close(); err != nil {
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
	_, err := db.Exec(`
		UPDATE datalogger SET
			name = $2,
			updater = $3,
			update_date = $4
		WHERE id = $1
	`, &u.ID, &u.Name, &u.Updater, &u.UpdateDate)
	if err != nil {
		return nil, err
	}

	return GetDataLogger(db, &u.ID)
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

func GetDataLoggerPreview(db *sqlx.DB, dlID *uuid.UUID) (*DataLoggerPreview, error) {
	var dlp DataLoggerPreview

	if err := db.Get(&dlp, `
		SELECT * FROM v_datalogger_preview WHERE datalogger_id = $1
	`, &dlID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("preview not found")
		}
		return nil, err
	}

	return &dlp, nil
}
