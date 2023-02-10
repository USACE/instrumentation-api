package models

import "github.com/jmoiron/sqlx"

func GetDataLoggerBySN(db *sqlx.DB, sn string) (*DataLogger, error) {
	var dl DataLogger
	if err := db.Get(&dl, `SELECT * FROM v_datalogger WHERE sn = $1`, sn); err != nil {
		return nil, err
	}
	return &dl, nil
}

func GetDataLoggerHashBySN(db *sqlx.DB, sn string) (string, error) {
	var hash string

	if err := db.Get(&hash, `SELECT "hash" FROM v_datalogger_hash WHERE sn = $1`, sn); err != nil {
		return "", err
	}

	return hash, nil
}

func GetEquivalencyTableBySN(db *sqlx.DB, sn string) (*EquivalencyTable, error) {
	var eq EquivalencyTable

	if err := db.Get(
		&eq, `SELECT * FROM v_datalogger_equivalency_table WHERE datalogger_id = $1`, sn,
	); err != nil {
		return nil, err
	}

	return &eq, nil
}

func UpdateDataLoggerPreviewBySN(db *sqlx.DB, dlp *DataLoggerPreview) error {
	if _, err := db.Exec(
		`UPDATE datalogger_preview SET payload = $2 WHERE sn = $1`,
		&dlp.SN, &dlp.Payload,
	); err != nil {
		return err
	}

	return nil
}
