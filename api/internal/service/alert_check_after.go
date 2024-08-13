package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AlertCheckAfterService interface {
	DoAlertAfterRequestChecks(mcc []model.MeasurementCollection)
}

type alertCheckAfterService struct {
	db *model.Database
	*model.Queries
	cfg *config.EmailConfig
}

func NewAlertCheckAfterService(db *model.Database, q *model.Queries, cfg *config.EmailConfig) *alertCheckAfterService {
	return &alertCheckAfterService{db, q, cfg}
}

var (
	ThresholdAlertTypeID    = uuid.MustParse("bb15e7c2-8eae-452c-92f7-e720dc5c9432")
	RateOfChangeAlertTypeID = uuid.MustParse("c37effee-6b48-4436-8d72-737ed78c1fb7")
)

const (
	alertLow  = "Alert Low"
	warnLow   = "Warn Low"
	warnHigh  = "Warn High"
	alertHigh = "Alert High"
)

func (s alertCheckAfterService) DoAlertAfterRequestChecks(mcc []model.MeasurementCollection) {
	var cause error
	actx, cancel := context.WithTimeoutCause(context.Background(), time.Second*10, cause)
	defer func() {
		cancel()
		if cause != nil {
			log.Println(cause.Error())
		}
	}()

	if err := s.doAlertAfterRequestChecks(actx, mcc); err != nil {
		log.Println(err.Error())
	}
}

func (s alertCheckAfterService) doAlertAfterRequestChecks(ctx context.Context, mcc []model.MeasurementCollection) error {
	if len(mcc) == 0 {
		return errors.New("error measurement collection is empty")
	}

	fnmms := make([]model.TimeseriesMeasurement, 0)
	mmMap := make(map[uuid.UUID][]model.Measurement)

	for idx := range mcc {
		if len(mcc[idx].Items) == 0 {
			continue
		}

		slices.SortFunc(mcc[idx].Items, func(a, b model.Measurement) int { return a.Time.Compare(b.Time) })

		if mcc[idx].TimeseriesID == uuid.Nil {
			tsID := mcc[idx].Items[0].TimeseriesID
			if tsID == uuid.Nil {
				continue
			}
			mcc[idx].TimeseriesID = tsID
		}

		firstNewMmt := model.TimeseriesMeasurement{
			TimeseriesID: mcc[idx].TimeseriesID,
			Measurement: model.Measurement{
				Time:  mcc[idx].Items[0].Time,
				Value: mcc[idx].Items[0].Value,
			},
		}
		fnmms = append(fnmms, firstNewMmt)
		mmMap[mcc[idx].TimeseriesID] = mcc[idx].Items
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

	acc, err := qtx.GetTimeseriesAlertConfigsForTimeseriesAndAlertTypes(ctx, string(b), []uuid.UUID{ThresholdAlertTypeID, RateOfChangeAlertTypeID})
	if err != nil {
		return err
	}

	if len(acc) == 0 {
		return nil
	}

	for _, ac := range acc {
		var sev string
		var vv []string
		var lastMmt *model.Measurement
		if ac.LastMeasurementTime != nil && ac.LastMeasurementValue != nil {
			lastMmt = &model.Measurement{TimeseriesID: ac.TimeseriesID, Time: *ac.LastMeasurementTime, Value: *ac.LastMeasurementValue}
		}

		switch ac.AlertTypeID {
		case ThresholdAlertTypeID:
			opts, err := model.MapToStruct[model.AlertConfigThresholdOpts](ac.Opts)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			mm, exists := mmMap[ac.TimeseriesID]
			if !exists {
				continue
			}

			sev, vv, err = doAlertCheckThresholds(opts, mm, lastMmt)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		case RateOfChangeAlertTypeID:
			opts, err := model.MapToStruct[model.AlertConfigChangeOpts](ac.Opts)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			mm, exists := mmMap[ac.TimeseriesID]
			if !exists {
				continue
			}

			sev, vv, err = doAlertCheckChanges(opts, mm, lastMmt)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		default:
			log.Printf("error alert type not supported: %s", ac.AlertType)
			continue
		}

		ac.Violations = vv

		if err := ac.DoEmail(sev, *s.cfg); err != nil {
			log.Println(err.Error())
			continue
		}
	}
	return tx.Commit()
}

func alertCheckThresholdMessage(alertType string, m model.Measurement, thresholdName string, thresholdValue float64) string {
	return fmt.Sprintf("(%s) %s %.3f voilates \"%s\" threshold of %.3f", alertType, m.Time, m.Value, thresholdName, thresholdValue)
}

func doAlertCheckThresholds(opts model.AlertConfigThresholdOpts, mm []model.Measurement, lastMmt *model.Measurement) (string, []string, error) {
	vv := make([]string, 0)
	var highestEmailSev string
	var prevSev string

	if lastMmt != nil {
		sev, _, err := handleThresholdRange(opts, *lastMmt, prevSev)
		if err != nil {
			return highestEmailSev, vv, err
		}
		prevSev = sev
	}

	for idx := range mm {
		sev, msg, err := handleThresholdRange(opts, mm[idx], prevSev)
		if err != nil {
			continue
		}

		switch sev {
		case alertLow, alertHigh:
			highestEmailSev = emailAlert

		case warnLow, warnHigh:
			if highestEmailSev != emailAlert {
				highestEmailSev = emailWarning
			}
		case "":
			continue
		}

		prevSev = sev
		vv = append(vv, msg)
	}

	return highestEmailSev, vv, nil
}

func handleThresholdRange(opts model.AlertConfigThresholdOpts, m model.Measurement, prevSev string) (string, string, error) {
	v := m.Value
	switch {
	case opts.IgnoreLowValue != nil && v <= *opts.IgnoreLowValue:
		return "", "", nil

	case opts.AlertLowValue != nil && v <= *opts.AlertLowValue:
		return alertLow, alertCheckThresholdMessage(emailAlert, m, alertLow, *opts.AlertLowValue), nil

	case opts.WarnLowValue != nil && v <= *opts.WarnLowValue:
		if prevSev == alertLow && opts.AlertLowValue != nil && v <= *opts.AlertLowValue+opts.Variance {
			return alertLow, alertCheckThresholdMessage(emailAlert, m, alertLow+" + Variance", *opts.AlertLowValue), nil
		}
		return warnLow, alertCheckThresholdMessage(emailWarning, m, warnLow, *opts.WarnLowValue), nil

	case opts.IgnoreHighValue != nil && opts.AlertHighValue != nil && v >= *opts.AlertHighValue && v >= *opts.IgnoreHighValue:
		return "", "", nil

	case opts.AlertHighValue != nil && v >= *opts.AlertHighValue:
		return alertHigh, alertCheckThresholdMessage(emailAlert, m, alertHigh, *opts.AlertHighValue), nil

	case opts.WarnHighValue != nil && v >= *opts.WarnHighValue:
		if prevSev == alertHigh && opts.AlertLowValue != nil && v >= *opts.AlertLowValue-opts.Variance {
			return alertHigh, alertCheckThresholdMessage(emailAlert, m, alertHigh+" - Variance", *opts.AlertHighValue), nil
		}
		return warnHigh, alertCheckThresholdMessage(emailWarning, m, warnHigh, *opts.WarnHighValue), nil

	default:
		if prevSev == warnLow && opts.WarnLowValue != nil && v <= *opts.WarnLowValue+opts.Variance {
			return warnLow, alertCheckThresholdMessage(emailWarning, m, warnLow+" + Variance", *opts.WarnLowValue), nil

		} else if prevSev == warnHigh && opts.WarnHighValue != nil && v >= *opts.AlertLowValue-opts.Variance {
			return warnHigh, alertCheckThresholdMessage(emailWarning, m, warnHigh+" - Variance", *opts.WarnHighValue), nil
		}
		return "", "", nil
	}
}

func alertCheckChangeMessage(alertType string, last, next model.Measurement, change, threshold float64) string {
	return fmt.Sprintf("(%s) time %s, value %.3f (before); time %s value %.3f (after); %3f voilates rate of change %.3f",
		alertType, last.Time, last.Value, next.Time, next.Value, change, threshold,
	)
}

func doAlertCheckChanges(opts model.AlertConfigChangeOpts, mm []model.Measurement, lastMmt *model.Measurement) (string, []string, error) {
	vv := make([]string, 0)
	var sev string

	if len(mm) == 0 {
		return sev, vv, fmt.Errorf("error no measurements found in uploaded collection")
	}

	if lastMmt != nil {
		skip := false
		if opts.LocfBackfillMs != nil {
			if t := mm[0].Time.Add(-time.Duration(*opts.LocfBackfillMs)); t.After(lastMmt.Time) {
				skip = true
			}
		}

		if !skip {
			change := math.Abs(mm[0].Value - math.Abs(lastMmt.Value))
			switch {
			case opts.IgnoreRateOfChange != nil && change >= *opts.IgnoreRateOfChange:
				// oob

			case change >= opts.AlertRateOfChange:
				sev = emailAlert
				vv = append(vv, alertCheckChangeMessage(emailAlert, *lastMmt, mm[0], change, opts.AlertRateOfChange))

			case opts.WarnRateOfChange != nil && change >= *opts.WarnRateOfChange:
				if sev != emailAlert {
					sev = emailWarning
				}
				vv = append(vv, alertCheckChangeMessage(emailWarning, *lastMmt, mm[0], change, *opts.WarnRateOfChange))
			}
		}
	}

	for idx := 0; idx < len(mm)-1; idx++ {
		if opts.LocfBackfillMs != nil {
			if t := mm[idx+1].Time.Add(-time.Duration(*opts.LocfBackfillMs)); t.After(mm[idx].Time) {
				continue
			}
		}
		change := math.Abs(mm[idx].Value - math.Abs(mm[idx+1].Value))
		switch {
		case opts.IgnoreRateOfChange != nil && change >= *opts.IgnoreRateOfChange:
			// oob
			continue

		case change >= opts.AlertRateOfChange:
			sev = emailAlert
			vv = append(vv, alertCheckChangeMessage(emailAlert, mm[idx], mm[idx+1], change, opts.AlertRateOfChange))

		case opts.WarnRateOfChange != nil && change >= *opts.WarnRateOfChange:
			if sev != emailAlert {
				sev = emailWarning
			}
			vv = append(vv, alertCheckChangeMessage(emailWarning, mm[idx], mm[idx+1], change, *opts.WarnRateOfChange))
		}
	}

	return sev, vv, nil
}
