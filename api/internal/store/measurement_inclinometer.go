package store

import (
	"context"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InclinometerMeasurementStore interface {
	ListInclinometerMeasurements(ctx context.Context, timeseriesID uuid.UUID, tw model.TimeWindow) (*model.InclinometerMeasurementCollection, error)
	ListInclinometerMeasurementValues(ctx context.Context, timeseriesID uuid.UUID, time time.Time, inclConstant float64) ([]*model.InclinometerMeasurementValues, error)
	DeleteInclinometerMeasurement(ctx context.Context, timeseriesID uuid.UUID, time time.Time) error
	CreateOrUpdateInclinometerMeasurements(ctx context.Context, im []model.InclinometerMeasurementCollection, p model.Profile, createDate time.Time) ([]model.InclinometerMeasurementCollection, error)
	ListInstrumentIDsFromTimeseriesID(ctx context.Context, timeseriesID uuid.UUID) ([]uuid.UUID, error)
	CreateTimeseriesConstant(ctx context.Context, timeseriesID uuid.UUID, parameterName string, unitName string, value float64) error
}

type inclinometerMeasurementStore struct {
	db *model.Database
	*model.Queries
}

func NewInclinometerMeasurementStore(db *model.Database, q *model.Queries) *inclinometerMeasurementStore {
	return &inclinometerMeasurementStore{db, q}
}

// CreateInclinometerMeasurements creates many inclinometer from an array of inclinometer
// If a inclinometer measurement already exists for a given timeseries_id and time, the values is updated
func (s inclinometerMeasurementStore) CreateOrUpdateInclinometerMeasurements(ctx context.Context, im []model.InclinometerMeasurementCollection, p model.Profile, createDate time.Time) ([]model.InclinometerMeasurementCollection, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.WithTx(tx)

	// Iterate All inclinometer Measurements
	for idx := range im {
		for i := range im[idx].Inclinometers {
			im[idx].Inclinometers[i].Creator = p.ID
			im[idx].Inclinometers[i].CreateDate = createDate
			if err := qtx.CreateOrUpdateInclinometerMeasurement(ctx, im[idx].TimeseriesID, im[idx].Inclinometers[i].Time, im[idx].Inclinometers[i].Values, p.ID, createDate); err != nil {
				return nil, err
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return im, nil
}

// CreateTimeseriesConstant creates timeseries constant
func (s inclinometerMeasurementStore) CreateTimeseriesConstant(ctx context.Context, timeseriesID uuid.UUID, parameterName string, unitName string, value float64) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.WithTx(tx)

	instrumentIDs, err := qtx.ListInstrumentIDsFromTimeseriesID(ctx, timeseriesID)
	if err != nil {
		return err
	}

	parameterIDs, err := qtx.ListParameterIDsFromParameterName(ctx, parameterName)
	if err != nil {
		return err
	}

	unitIDs, err := qtx.ListUnitIDsFromUnitName(ctx, unitName)
	if err != nil {
		return err
	}

	if len(instrumentIDs) > 0 && len(parameterIDs) > 0 && len(unitIDs) > 0 {
		t := model.Timeseries{}
		measurement := model.Measurement{}
		measurements := []model.Measurement{}
		mc := model.MeasurementCollection{}
		mcs := []model.MeasurementCollection{}
		ts := []model.Timeseries{}

		t.InstrumentID = instrumentIDs[0]
		t.Slug = parameterName
		t.Name = parameterName
		t.ParameterID = parameterIDs[0]
		t.UnitID = unitIDs[0]
		ts = append(ts, t)

		// Create timeseries for constant
		tsNew, err := qtx.CreateTimeseries(ctx, t)
		if err != nil {
			return err
		}
		// Assign timeseries
		if err := qtx.CreateInstrumentConstant(ctx, t.InstrumentID, t.ID); err != nil {
			return err
		}

		measurement.Time = time.Now()
		measurement.Value = value
		measurements = append(measurements, measurement)
		mc.TimeseriesID = tsNew.ID
		mc.Items = measurements
		mcs = append(mcs, mc)

		if err = createMeasurements(ctx, mcs, qtx.CreateOrUpdateTimeseriesMeasurement, qtx.CreateOrUpdateTimeseriesNote); err != nil {
			return err
		}
	}

	return nil
}
