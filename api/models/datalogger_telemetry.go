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

func GetEquivalencyTableByModelSN(db *sqlx.DB, model, sn string) (*EquivalencyTable, error) {
	var eq EquivalencyTable

	if err := db.Get(&eq, `
		SELECT * FROM v_datalogger_equivalency_table
		WHERE model = $1 AND sn = $2
		`, model, sn,
	); err != nil {
		return nil, err
	}

	return &eq, nil
}

func UpdateDataLoggerPreviewByModelSN(db *sqlx.DB, dlp *DataLoggerPreview) error {
	if _, err := db.Exec(`
		UPDATE datalogger_preview SET payload = $2 WHERE model = $1 AND sn = $2
	`, &dlp.Model, &dlp.SN, &dlp.Payload,
	); err != nil {
		return err
	}

	return nil
}
