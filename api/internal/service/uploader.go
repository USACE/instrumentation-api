package service

import (
	"context"
	"encoding/csv"
	"io"
	"math"
	"strconv"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type UploaderService interface {
	CreateTimeseriesMeasurementsFromDuxFile(ctx context.Context, r io.Reader) error
	CreateTimeseriesMeasurementsFromTOA5File(ctx context.Context, r io.Reader) error
}

type uploaderService struct {
	db *model.Database
	*model.Queries
}

func NewUploaderService(db *model.Database, q *model.Queries) *uploaderService {
	return &uploaderService{db, q}
}

// TODO: transition away from datalogger equivalency table to different parser that's uploader specific
func (s uploaderService) CreateTimeseriesMeasurementsFromTOA5File(ctx context.Context, r io.Reader) error {
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
				continue
			}

			v, err := strconv.ParseFloat(cell, 64)
			if err != nil || math.IsNaN(v) || math.IsInf(v, 0) {
				continue
			}

			if err := qtx.CreateOrUpdateTimeseriesMeasurement(ctx, tsID, t, v); err != nil {
				return err
			}
		}
	}
	return nil
}
