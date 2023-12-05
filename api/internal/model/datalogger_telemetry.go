package model

import (
	"context"

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

const updateDataloggerTablePreview = `
	UPDATE datalogger_preview SET preview = $3, update_date = $4
	WHERE datalogger_table_id IN (SELECT id FROM datalogger_table WHERE datalogger_id = $1 AND table_name = $2)
`

func (q *Queries) UpdateDataloggerTablePreview(ctx context.Context, dataloggerID uuid.UUID, tableName string, dlp DataloggerPreview) error {
	_, err := q.db.ExecContext(ctx, updateDataloggerTablePreview, dataloggerID, tableName, dlp.Preview, dlp.UpdateDate)
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

const renameEmptyDataloggerTableName = `
	UPDATE datalogger_table
	SET table_name = $2
	WHERE table_name = '' AND datalogger_id = $1
	AND NOT EXISTS (
	   SELECT 1 FROM datalogger_table WHERE datalogger_id = $1 AND table_name = $2
	);
`

func (q *Queries) RenameEmptyDataloggerTableName(ctx context.Context, dataloggerID uuid.UUID, tableName string) error {
	_, err := q.db.ExecContext(ctx, renameEmptyDataloggerTableName, dataloggerID, tableName)
	return err
}

const getOrCreateDataloggerTable = `
	WITH dt AS (
		INSERT INTO datalogger_table (datalogger_id, table_name) VALUES ($1, $2)
		ON CONFLICT ON CONSTRAINT datalogger_table_datalogger_id_table_name_key DO NOTHING
		RETURNING id
	)
	SELECT id FROM dt
	UNION
	SELECT id FROM datalogger_table WHERE datalogger_id = $1 AND table_name = $2
`

func (q *Queries) GetOrCreateDataloggerTable(ctx context.Context, dataloggerID uuid.UUID, tableName string) (uuid.UUID, error) {
	var tID uuid.UUID
	err := q.db.GetContext(ctx, &tID, getOrCreateDataloggerTable, dataloggerID, tableName)
	return tID, err
}
