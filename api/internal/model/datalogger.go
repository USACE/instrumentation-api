package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/password"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

// Telemetry struct
type Telemetry struct {
	ID       uuid.UUID
	TypeID   string
	TypeSlug string
	TypeName string
}

type Datalogger struct {
	ID        uuid.UUID                    `json:"id" db:"id"`
	Name      string                       `json:"name" db:"name"`
	SN        string                       `json:"sn" db:"sn"`
	ProjectID uuid.UUID                    `json:"project_id" db:"project_id"`
	Slug      string                       `json:"slug" db:"slug"`
	ModelID   uuid.UUID                    `json:"model_id" db:"model_id"`
	Model     *string                      `json:"model" db:"model"`
	Errors    []string                     `json:"errors" db:"-"`
	PgErrors  pgtype.TextArray             `json:"-" db:"errors"`
	Tables    dbJSONSlice[DataloggerTable] `json:"tables" db:"tables"`
	AuditInfo
}

type DataloggerWithKey struct {
	Datalogger
	Key string `json:"key"`
}

type DataloggerTable struct {
	ID        uuid.UUID `json:"id" db:"id"`
	TableName string    `json:"table_name" db:"table_name"`
}

type DataloggerTablePreview struct {
	DataloggerTableID uuid.UUID   `json:"datalogger_table_id" db:"datalogger_table_id"`
	UpdateDate        time.Time   `json:"update_date" db:"update_date"`
	Preview           pgtype.JSON `json:"preview" db:"preview"`
}

type DataloggerError struct {
	DataloggerTableID uuid.UUID `json:"datalogger_id" db:"datalogger_id"`
	Errors            []string  `json:"errors" db:"errors"`
}

const getDataloggerModelName = `
	SELECT model FROM datalogger_model WHERE id = $1
`

func (q *Queries) GetDataloggerModelName(ctx context.Context, modelID uuid.UUID) (string, error) {
	var modelName string
	if err := q.db.GetContext(ctx, &modelName, getDataloggerModelName, modelID); err != nil {
		return "", err
	}
	return modelName, nil
}

const listProjectDataloggers = `
	SELECT * FROM v_datalogger WHERE project_id = $1
`

func (q *Queries) ListProjectDataloggers(ctx context.Context, projectID uuid.UUID) ([]Datalogger, error) {
	dls := make([]Datalogger, 0)
	if err := q.db.SelectContext(ctx, &dls, listProjectDataloggers, projectID); err != nil {
		return make([]Datalogger, 0), err
	}
	for i := 0; i < len(dls); i++ {
		if err := dls[i].PgErrors.AssignTo(&dls[i].Errors); err != nil {
			return make([]Datalogger, 0), err
		}
	}
	return dls, nil
}

const listAllDataloggers = `
	SELECT * FROM v_datalogger
`

func (q *Queries) ListAllDataloggers(ctx context.Context) ([]Datalogger, error) {
	dls := make([]Datalogger, 0)
	if err := q.db.SelectContext(ctx, &dls, listAllDataloggers); err != nil {
		return make([]Datalogger, 0), err
	}
	for i := 0; i < len(dls); i++ {
		if err := dls[i].PgErrors.AssignTo(&dls[i].Errors); err != nil {
			return make([]Datalogger, 0), err
		}
	}
	return dls, nil
}

const getDataloggerIsActive = `
	SELECT EXISTS (SELECT * FROM v_datalogger WHERE model = $1 AND sn = $2)::int
`

// GetDataloggerIsActive checks if datalogger with sn already exists and is not deleted
func (q *Queries) GetDataloggerIsActive(ctx context.Context, modelName, sn string) (bool, error) {
	var isActive bool
	if err := q.db.GetContext(ctx, &isActive, getDataloggerIsActive, modelName, sn); err != nil {
		return false, err
	}
	return isActive, nil
}

const verifyDataloggerExists = `
	SELECT id FROM v_datalogger WHERE id = $1
`

// VerifyDataloggerExists checks if datalogger with sn already exists and is not deleted
func (q *Queries) VerifyDataloggerExists(ctx context.Context, dlID uuid.UUID) error {
	return q.db.GetContext(ctx, &uuid.UUID{}, verifyDataloggerExists, dlID)
}

const createDataloggerHash = `
	INSERT INTO datalogger_hash (datalogger_id, "hash") VALUES ($1, $2)
`

func (q *Queries) CreateDataloggerHash(ctx context.Context, dataloggerID uuid.UUID) (string, error) {
	key := password.GenerateRandom(40)
	if _, err := q.db.ExecContext(ctx, createDataloggerHash, dataloggerID, password.MustCreateHash(key, password.DefaultParams)); err != nil {
		return "", err
	}
	return key, nil
}

const getOneDatalogger = `
	SELECT * FROM v_datalogger WHERE id = $1
`

func (q *Queries) GetOneDatalogger(ctx context.Context, dataloggerID uuid.UUID) (Datalogger, error) {
	var dl Datalogger
	if err := q.db.GetContext(ctx, &dl, getOneDatalogger, dataloggerID); err != nil {
		return dl, err
	}
	if err := dl.PgErrors.AssignTo(&dl.Errors); err != nil {
		return dl, err
	}
	return dl, nil
}

const createDatalogger = `
	INSERT INTO datalogger (name, sn, project_id, creator, updater, slug, model_id)
	VALUES ($1, $2, $3, $4, $4, slugify($1, 'datalogger'), $5)
	RETURNING id
`

func (q *Queries) CreateDatalogger(ctx context.Context, dl Datalogger) (uuid.UUID, error) {
	var dlID uuid.UUID
	err := q.db.GetContext(ctx, &dlID, createDatalogger, dl.Name, dl.SN, dl.ProjectID, dl.CreatorID, dl.ModelID)
	return dlID, err
}

const updateDatalogger = `
	UPDATE datalogger SET
		name = $2,
		updater = $3,
		update_date = $4
	WHERE id = $1
`

func (q *Queries) UpdateDatalogger(ctx context.Context, dl Datalogger) error {
	_, err := q.db.ExecContext(ctx, updateDatalogger, dl.ID, dl.Name, dl.UpdaterID, dl.UpdateDate)
	return err
}

const updateDataloggerHash = `
	UPDATE datalogger_hash SET "hash" = $2 WHERE datalogger_id = $1
`

func (q *Queries) UpdateDataloggerHash(ctx context.Context, dataloggerID uuid.UUID) (string, error) {
	key := password.GenerateRandom(40)
	if _, err := q.db.ExecContext(ctx, updateDataloggerHash, dataloggerID, password.MustCreateHash(key, password.DefaultParams)); err != nil {
		return "", err
	}
	return key, nil
}

const updateDataloggerUpdater = `
	UPDATE datalogger SET updater = $2, update_date = $3 WHERE id = $1
`

func (q *Queries) UpdateDataloggerUpdater(ctx context.Context, dl Datalogger) error {
	_, err := q.db.ExecContext(ctx, updateDataloggerUpdater, dl.ID, dl.UpdaterID, dl.UpdateDate)
	return err
}

const deleteDatalogger = `
	UPDATE datalogger SET deleted = true, updater = $2, update_date = $3  WHERE id = $1
`

func (q *Queries) DeleteDatalogger(ctx context.Context, dl Datalogger) error {
	_, err := q.db.ExecContext(ctx, deleteDatalogger, dl.ID, dl.UpdaterID, dl.UpdateDate)
	return err
}

const getDataloggerTablePreview = `
	SELECT * FROM v_datalogger_preview WHERE datalogger_table_id = $1
`

func (q *Queries) GetDataloggerTablePreview(ctx context.Context, dataloggerTableID uuid.UUID) (DataloggerTablePreview, error) {
	var dlp DataloggerTablePreview
	err := q.db.GetContext(ctx, &dlp, getDataloggerTablePreview, dataloggerTableID)
	if errors.Is(err, sql.ErrNoRows) {
		dlp.DataloggerTableID = dataloggerTableID
		if err := dlp.Preview.Set("null"); err != nil {
			return DataloggerTablePreview{}, err
		}
		return dlp, nil
	}
	return dlp, err
}

const resetDataloggerTableName = `
	UPDATE datalogger_table SET table_name = '' WHERE id = $1
`

func (q *Queries) ResetDataloggerTableName(ctx context.Context, dataloggerTableID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, resetDataloggerTableName, dataloggerTableID)
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

const deleteDataloggerTable = `
	DELETE FROM datalogger_table WHERE id = $1
`

func (q *Queries) DeleteDataloggerTable(ctx context.Context, dataloggerTableID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDataloggerTable, dataloggerTableID)
	return err
}
