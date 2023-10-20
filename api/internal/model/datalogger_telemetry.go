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

const updateDataloggerPreview = `
	UPDATE datalogger_preview SET preview = $2, update_date = $3
	WHERE datalogger_id = $1
`

func (q *Queries) UpdateDataloggerPreview(ctx context.Context, dlp DataloggerPreview) error {
	_, err := q.db.ExecContext(ctx, updateDataloggerPreview, dlp.DataloggerID, dlp.Preview, dlp.UpdateDate)
	return err
}

const deleteDataloggerError = `
	DELETE FROM datalogger_error WHERE datalogger_id = $1
`

func (q *Queries) DeleteDataloggerError(ctx context.Context, dataloggerID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDataloggerError, dataloggerID)
	return err
}

const createDataloggerError = `
	INSERT INTO datalogger_error (datalogger_id, error_message) VALUES ($1, $2)
`

func (q *Queries) CreateDataloggerError(ctx context.Context, dataloggerID uuid.UUID, errMessage string) error {
	_, err := q.db.ExecContext(ctx, createDataloggerError, dataloggerID, errMessage)
	return err
}
