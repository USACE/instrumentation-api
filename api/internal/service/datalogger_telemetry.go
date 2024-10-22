package service

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DataloggerTelemetryService interface {
	GetDataloggerByModelSN(ctx context.Context, modelName, sn string) (model.Datalogger, error)
	GetDataloggerHashByModelSN(ctx context.Context, modelName, sn string) (string, error)
	CreateDataloggerTablePreview(ctx context.Context, prv model.DataloggerTablePreview) error
	UpdateDataloggerTablePreview(ctx context.Context, dataloggerID uuid.UUID, tableName string, prv model.DataloggerTablePreview) (uuid.UUID, error)
	UpdateDataloggerTableError(ctx context.Context, dataloggerID uuid.UUID, tableName *string, e *model.DataloggerError) error
	CreateOrUpdateDataloggerTOA5MeasurementCollection(ctx context.Context, r io.Reader) error
}

type dataloggerTelemetryService struct {
	db *model.Database
	*model.Queries
}

func NewDataloggerTelemetryService(db *model.Database, q *model.Queries) *dataloggerTelemetryService {
	return &dataloggerTelemetryService{db, q}
}

// UpdateDataloggerTablePreview attempts to update a table preview by datalogger_id and table_name, creates the
// datalogger table and corresponding preview if it doesn't exist
func (s dataloggerTelemetryService) UpdateDataloggerTablePreview(ctx context.Context, dataloggerID uuid.UUID, tableName string, prv model.DataloggerTablePreview) (uuid.UUID, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	// replace empty datalogger table name with most recent payload
	if err := qtx.RenameEmptyDataloggerTableName(ctx, dataloggerID, tableName); err != nil {
		return uuid.Nil, err
	}

	tableID, err := qtx.GetOrCreateDataloggerTable(ctx, dataloggerID, tableName)
	if err != nil {
		return uuid.Nil, err
	}
	if err := qtx.UpdateDataloggerTablePreview(ctx, dataloggerID, tableName, prv); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, err
		}
		prv.DataloggerTableID = tableID
		if err := qtx.CreateDataloggerTablePreview(ctx, prv); err != nil {
		}
	}

	return tableID, tx.Commit()
}

func (s dataloggerTelemetryService) UpdateDataloggerTableError(ctx context.Context, dataloggerID uuid.UUID, tableName *string, e *model.DataloggerError) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.DeleteDataloggerTableError(ctx, dataloggerID, tableName); err != nil {
		return err
	}

	for _, m := range e.Errors {
		if err := qtx.CreateDataloggerTableError(ctx, dataloggerID, tableName, m); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// ParseTOA5 parses a Campbell Scientific TOA5 data file that is simlar to a csv.
// The unique properties of TOA5 are that the meatdata are stored in header of file (first 4 lines of csv)
func (s dataloggerTelemetryService) CreateOrUpdateDataloggerTOA5MeasurementCollection(ctx context.Context, r io.Reader) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	reader := csv.NewReader(r)

	envHeader, err := reader.Read()
	if err != nil {
		return err
	}
	fieldHeader, err := reader.Read()
	if err != nil {
		return err
	}
	unitsHeader, err := reader.Read()
	if err != nil {
		return err
	}
	processHeader, err := reader.Read()
	if err != nil {
		return err
	}

	meta := model.Environment{
		StationName: envHeader[1],
		Model:       envHeader[2],
		SerialNo:    envHeader[3],
		OSVersion:   envHeader[4],
		ProgName:    envHeader[5],
		TableName:   envHeader[6],
	}

	dl, err := qtx.GetDataloggerByModelSN(ctx, meta.Model, meta.SerialNo)
	if err != nil {
		return err
	}

	tableID, err := qtx.GetOrCreateDataloggerTable(ctx, dl.ID, meta.TableName)
	if err != nil {
		return err
	}

	em := make([]string, 0)
	defer func() {
		s.UpdateDataloggerTableError(ctx, dl.ID, &meta.TableName, &model.DataloggerError{Errors: em})
	}()

	// first two columns are timestamp and record number
	// we only want to collect the measurement fields here
	fields := make([]model.Field, len(fieldHeader)-2)
	for i := 2; i < len(fieldHeader); i++ {
		fields[i] = model.Field{
			Name:    fieldHeader[i],
			Units:   unitsHeader[i],
			Process: processHeader[i],
		}
	}

	eqt, err := qtx.GetEquivalencyTable(ctx, tableID)
	if err != nil {
		return err
	}

	fieldNameTimeseriesIDMap := make(map[string]uuid.UUID)
	for _, eqtRow := range eqt.Rows {
		fieldNameTimeseriesIDMap[eqtRow.FieldName] = *eqtRow.TimeseriesID
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		t, err := time.Parse(record[0], time.RFC3339)
		if err != nil {
			return err
		}

		for idx, cell := range record[2:] {
			fieldName := fields[idx].Name
			tsID, ok := fieldNameTimeseriesIDMap[fieldName]
			if !ok {
				// key error, field_name does not exist for equivalency table
				// add error to Measurement payload to report back to user
				em = append(em, fmt.Sprintf(
					"key error: field_name %s does not exist for equivalency table %s",
					fieldName, meta.TableName,
				))
				continue
			}

			v, err := strconv.ParseFloat(cell, 64)
			if err != nil || math.IsNaN(v) || math.IsInf(v, 0) {
				// could not parse float
				// add error to Measurement payload to report back to user
				em = append(em, fmt.Sprintf(
					"value error: field_name %s contains invalid value entry at %s",
					fieldName, t,
				))
				continue
			}

			if err := qtx.CreateOrUpdateTimeseriesMeasurement(ctx, tsID, t, v); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}
