package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

var (
	ThresholdAlertTypeID    = uuid.MustParse("bb15e7c2-8eae-452c-92f7-e720dc5c9432")
	RateOfChangeAlertTypeID = uuid.MustParse("c37effee-6b48-4436-8d72-737ed78c1fb7")
)

type TimeseriesTimeMeasurement struct {
	FirstNewMeasurement model.Measurement
	Measurements        []model.Measurement
}

func (s alertCheckService) DoAlertAfterRequestChecks(ctx context.Context, mcc []model.MeasurementCollection) error {
	if len(mcc) == 0 {
		return errors.New("error measurement collection is empty")
	}

	fnmms := make([]model.Measurement, 0)
	mmMap := make(map[uuid.UUID]TimeseriesTimeMeasurement)

	for idx := range mcc {
		if len(mcc[idx].Items) != 0 {
			continue
		}
		firstNewMmt := model.Measurement{
			TimeseriesID: mcc[idx].TimeseriesID,
			Time:         mcc[idx].Items[0].Time,
			Value:        mcc[idx].Items[0].Value,
		}
		fnmms = append(fnmms, firstNewMmt)
		mmMap[mcc[idx].TimeseriesID] = TimeseriesTimeMeasurement{
			FirstNewMeasurement: firstNewMmt,
			Measurements:        mcc[idx].Items,
		}
	}

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	b, err := json.Marshal(fnmms)
	if err != nil {
		return err
	}

	acc, err := qtx.GetTimeseriesAlertConfigsForTimeseriesAndAlertTypes(ctx, b, []uuid.UUID{ThresholdAlertTypeID, RateOfChangeAlertTypeID})
	if err != nil {
		return err
	}

	for _, ac := range acc {
		switch ac.AlertTypeID {
		case ThresholdAlertTypeID:
			opts, err := model.MapToStruct[model.AlertConfigThresholdOpts](ac.Opts)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			if err := doAlertCheckThresholds(opts, mmMap[ac.TimeseriesID]); err != nil {
				log.Println(err.Error())
				continue
			}
		case RateOfChangeAlertTypeID:
			opts, err := model.MapToStruct[model.AlertConfigChangeOpts](ac.Opts)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			if err := doAlertCheckChanges(opts, mmMap[ac.TimeseriesID]); err != nil {
				log.Println(err.Error())
				continue
			}
		default:
			log.Printf("error alert type not supported: %s", ac.AlertType)
		}
	}

	return tx.Commit()
}

func doAlertCheckThresholds(opts model.AlertConfigThresholdOpts, mm TimeseriesTimeMeasurement) error {
	for idx := range mm.Measurements {
		v := mm.Measurements[idx].Value
		switch {
		case opts.IgnoreLowValue != nil && opts.AlertLowValue != nil && v <= *opts.AlertLowValue && v >= *opts.IgnoreLowValue:
			// case where ignore low is set
		case opts.AlertLowValue != nil && v <= *opts.AlertLowValue:
			// case where ignore low is not set
		case opts.WarnLowValue != nil && v <= *opts.WarnLowValue:
			// case where warn low
		case opts.WarnHighValue != nil && v <= *opts.WarnHighValue:
			// safe zone
		case opts.WarnHighValue != nil && v >= *opts.WarnHighValue:
			// case where warn high
		case opts.IgnoreHighValue != nil && opts.AlertHighValue != nil && v >= *opts.AlertHighValue && v <= *opts.IgnoreHighValue:
			// case where ignore high is set
		case opts.AlertHighValue != nil && v >= *opts.AlertHighValue:
			// case where ignore high is not set
		}
	}
	return nil
}

func doAlertCheckChanges(opts model.AlertConfigChangeOpts, mm TimeseriesTimeMeasurement) error {
	changes := make([]model.Measurement, 0)
	fnm := mm.FirstNewMeasurement

	change := math.Abs(mm.Measurements[0].Value - fnm.Value)
	if change >= opts.RateOfChange {
		changes = append(changes, model.Measurement{Time: fnm.Time, Value: fnm.Value})
	}

	for idx := 0; idx < len(mm.Measurements)-1; idx++ {
		change := math.Abs(mm.Measurements[idx].Value - mm.Measurements[idx+1].Value)
		if change >= opts.RateOfChange {
			changes = append(changes, model.Measurement{Time: fnm.Time, Value: fnm.Value})
		}
	}

	return nil
}
