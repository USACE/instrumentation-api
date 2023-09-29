package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/passwords"
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
	ID              uuid.UUID        `json:"id" db:"id"`
	Name            string           `json:"name" db:"name"`
	SN              string           `json:"sn" db:"sn"`
	ProjectID       uuid.UUID        `json:"project_id" db:"project_id"`
	Creator         uuid.UUID        `json:"creator" db:"creator"`
	CreatorUsername string           `json:"creator_username" db:"creator_username"`
	CreateDate      time.Time        `json:"create_date" db:"create_date"`
	Updater         *uuid.UUID       `json:"updater" db:"updater"`
	UpdaterUsername string           `json:"updater_username" db:"updater_username"`
	UpdateDate      *time.Time       `json:"update_date" db:"update_date"`
	Slug            string           `json:"slug" db:"slug"`
	ModelID         uuid.UUID        `json:"model_id" db:"model_id"`
	Model           *string          `json:"model" db:"model"`
	Deleted         bool             `json:"-" db:"deleted"`
	Errors          []string         `json:"errors" db:"-"`
	PgErrors        pgtype.TextArray `json:"-" db:"errors"`
}

type DataloggerWithKey struct {
	*Datalogger
	Key string `json:"key"`
}

type DataloggerPreview struct {
	DataloggerID uuid.UUID   `json:"datalogger_id" db:"datalogger_id"`
	UpdateDate   time.Time   `json:"update_date" db:"update_date"`
	Preview      pgtype.JSON `json:"preview" db:"preview"`
	Model        *string     `json:"model,omitempty"`
	SN           *string     `json:"sn,omitempty"`
}

type DataloggerError struct {
	DataloggerID uuid.UUID `json:"datalogger_id" db:"datalogger_id"`
	Errors       []string  `json:"errors" db:"errors"`
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
	SELECT EXISTS (SELECT id FROM v_datalogger WHERE id = $1)::int
`

// VerifyDataloggerExists checks if datalogger with sn already exists and is not deleted
func (q *Queries) VerifyDataloggerExists(ctx context.Context, dlID uuid.UUID) error {
	var dlExists bool
	if err := q.db.GetContext(ctx, &dlExists, verifyDataloggerExists, dlID); err != nil {
		return err
	}
	if !dlExists {
		return fmt.Errorf("active data logger with id %s not found", dlID)
	}
	return nil
}

const createDataloggerHash = `
	INSERT INTO datalogger_hash (datalogger_id, "hash") VALUES ($1, $2)
`

func (q *Queries) CreateDataloggerHash(ctx context.Context, dataloggerID uuid.UUID) (string, error) {
	key := passwords.GenerateRandom(40)
	if _, err := q.db.ExecContext(ctx, createDataloggerHash, dataloggerID, passwords.MustCreateHash(key, passwords.DefaultParams)); err != nil {
		return "", err
	}
	return key, nil
}

const createDataloggerPreview = `
	INSERT INTO datalogger_preview (datalogger_id) VALUES ($1)
`

func (q *Queries) CreateDataloggerPreview(ctx context.Context, dataloggerID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createDataloggerPreview, dataloggerID)
	return err
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
	VALUES ($1, $2, $3, $4, $4, $5, $6)
	RETURNING id
`

func (q *Queries) CreateDatalogger(ctx context.Context, dl Datalogger) (uuid.UUID, error) {
	var dlID uuid.UUID
	err := q.db.GetContext(ctx, &dlID, createDatalogger, dl.Name, dl.SN, dl.ProjectID, dl.Creator, dl.Slug, dl.ModelID)
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
	_, err := q.db.ExecContext(ctx, updateDatalogger, dl.ID, dl.Name, dl.Updater, dl.UpdateDate)
	return err
}

const updateDataloggerHash = `
	UPDATE datalogger_hash SET "hash" = $2 WHERE datalogger_id = $1
`

func (q *Queries) UpdateDataloggerHash(ctx context.Context, dataloggerID uuid.UUID) (string, error) {
	key := passwords.GenerateRandom(40)
	if _, err := q.db.ExecContext(ctx, updateDataloggerHash, dataloggerID, passwords.MustCreateHash(key, passwords.DefaultParams)); err != nil {
		return "", err
	}
	return key, nil
}

const updateDataloggerUpdater = `
	UPDATE datalogger SET updater = $2, update_date = $3 WHERE id = $1
`

func (q *Queries) UpdateDataloggerUpdater(ctx context.Context, dl Datalogger) error {
	_, err := q.db.ExecContext(ctx, updateDataloggerUpdater, dl.ID, dl.Updater, dl.UpdateDate)
	return err
}

const deleteDatalogger = `
	UPDATE datalogger SET deleted = true, updater = $2, update_date = $3  WHERE id = $1
`

func (q *Queries) DeleteDatalogger(ctx context.Context, dl Datalogger) error {
	_, err := q.db.ExecContext(ctx, deleteDatalogger, dl.ID, dl.Updater, dl.UpdateDate)
	return err
}

const getDataloggerPreview = `
	SELECT * FROM v_datalogger_preview WHERE datalogger_id = $1
`

func (q *Queries) GetDataloggerPreview(ctx context.Context, dlID uuid.UUID) (DataloggerPreview, error) {
	var dlp DataloggerPreview
	if err := q.db.GetContext(ctx, &dlp, getDataloggerPreview, dlID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dlp, fmt.Errorf("preview not found")
		}
		return dlp, err
	}
	return dlp, nil
}

const createUniquSlugDatalogger = `
	SELECT slug FROM datalogger
`

func (q *Queries) CreateUniqueSlugDatalogger(ctx context.Context, dataloggerName string) (string, error) {
	return q.CreateUniqueSlug(ctx, createUniquSlugDatalogger, dataloggerName)
}
