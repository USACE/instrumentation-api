package model

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const getDataloggerByModelSN = `
	SELECT * FROM v_datalogger
	WHERE model = $1 AND sn = $2
`

func (q *Queries) GetDataloggerByModelSN(ctx context.Context, modelName, sn string) (Datalogger, error) {
	var dl Datalogger
	err := q.db.GetContext(ctx, &dl, getDataloggerByModelSN, modelName, sn)
	return dl, err
}

const getDataloggerHashByModelSN = `
	SELECT "hash" FROM v_datalogger_hash
	WHERE model = $1 AND sn = $2
`

func (q *Queries) GetDataloggerHashByModelSN(ctx context.Context, modelName, sn string) (string, error) {
	var hash string
	if err := q.db.GetContext(ctx, &hash, getDataloggerHashByModelSN, modelName, sn); err != nil {
		return "", err
	}
	return hash, nil
}

const createDataloggerTablePreview = `
	INSERT INTO datalogger_preview (datalogger_table_id, preview, update_date) VALUES ($1, $2, $3)
`

func (q *Queries) CreateDataloggerTablePreview(ctx context.Context, prv DataloggerTablePreview) error {
	_, err := q.db.ExecContext(ctx, createDataloggerTablePreview, prv.DataloggerTableID, prv.Preview, prv.UpdateDate)
	return err
}

const updateDataloggerTablePreview = `
	UPDATE datalogger_preview SET preview = $3, update_date = $4
	WHERE datalogger_table_id IN (SELECT id FROM datalogger_table WHERE datalogger_id = $1 AND table_name = $2)
`

func (q *Queries) UpdateDataloggerTablePreview(ctx context.Context, dataloggerID uuid.UUID, tableName string, prv DataloggerTablePreview) error {
	result, err := q.db.ExecContext(ctx, updateDataloggerTablePreview, dataloggerID, tableName, prv.Preview, prv.UpdateDate)
	r, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if r == 0 {
		return sql.ErrNoRows
	}
	return err
}

const deleteDataloggerTableError = `
	DELETE FROM datalogger_error
	WHERE datalogger_table_id IN (SELECT id FROM datalogger_table WHERE datalogger_id = $1 AND table_name = $2)
`

func (q *Queries) DeleteDataloggerTableError(ctx context.Context, dataloggerID uuid.UUID, tableName *string) error {
	_, err := q.db.ExecContext(ctx, deleteDataloggerTableError, dataloggerID, tableName)
	return err
}

const createDataloggerError = `
	INSERT INTO datalogger_error (datalogger_table_id, error_message)
	SELECT id, $3 FROM datalogger_table
	WHERE datalogger_id = $1 AND table_name = $2
	AND NOT EXISTS (
	   SELECT 1 FROM datalogger_table WHERE datalogger_id = $1 AND table_name = $2
	);
`

func (q *Queries) CreateDataloggerTableError(ctx context.Context, dataloggerID uuid.UUID, tableName *string, errMessage string) error {
	_, err := q.db.ExecContext(ctx, createDataloggerError, dataloggerID, tableName, errMessage)
	return err
}
