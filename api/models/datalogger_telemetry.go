package models

import "github.com/jmoiron/sqlx"

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
