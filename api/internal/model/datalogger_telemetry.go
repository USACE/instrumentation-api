package model

import (
	"github.com/jmoiron/sqlx"
)

func GetDataLoggerByModelSN(db *sqlx.DB, model, sn string) (*DataLogger, error) {
	var dl DataLogger
	if err := db.Get(&dl, `
		SELECT * FROM v_datalogger
		WHERE model = $1 AND sn = $2
	`, model, sn); err != nil {
		return nil, err
	}
	return &dl, nil
}

func GetDataLoggerHashByModelSN(db *sqlx.DB, model, sn string) (string, error) {
	var hash string

	if err := db.Get(&hash, `
		SELECT "hash" FROM v_datalogger_hash
		WHERE model = $1 AND sn = $2
	`, model, sn); err != nil {
		return "", err
	}

	return hash, nil
}

func UpdateDataLoggerPreview(db *sqlx.DB, dlp *DataLoggerPreview) error {
	if _, err := db.Exec(`
		UPDATE datalogger_preview SET preview = $2, update_date = $3
		WHERE datalogger_id = $1
	`, &dlp.DataLoggerID, &dlp.Preview, &dlp.UpdateDate,
	); err != nil {
		return err
	}

	return nil
}

func UpdateDataLoggerError(db *sqlx.DB, e *DataLoggerError) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`DELETE FROM datalogger_error WHERE datalogger_id = $1`)
	if err != nil {
		return err
	}

	stmt2, err := txn.Preparex(`INSERT INTO datalogger_error (datalogger_id, error_message) VALUES ($1, $2)`)
	if err != nil {
		return err
	}

	_, err = stmt1.Exec(&e.DataLoggerID)
	if err != nil {
		return err
	}

	for _, m := range e.Errors {
		_, err = stmt2.Exec(&e.DataLoggerID, m)
		if err != nil {
			return err
		}
	}

	if err := stmt1.Close(); err != nil {
		return err
	}
	if err := stmt2.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}
