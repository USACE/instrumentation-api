package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Survey123Service interface {
	ListProjectSurvey123s(ctx context.Context, projectID uuid.UUID) ([]model.Survey123, error)
	ListSurvey123EquivalencyTableRows(ctx context.Context, survey123ID uuid.UUID) ([]model.EquivalencyTableRow, error)
	CreateSurvey123(ctx context.Context, sv model.Survey123) (uuid.UUID, error)
	CreateOrUpdateSurvey123EquivalencyTable(ctx context.Context, survey123ID uuid.UUID, mappings []model.EquivalencyTableRow) error
	DeleteSurvey123EquivalencyTableRow(ctx context.Context, survey123RowID uuid.UUID) error
	CreateOrUpdateSurvey123Preview(ctx context.Context, survey123ID uuid.UUID, previewRawJson []byte) error
	CreateOrUpdateSurvey123Measurements(ctx context.Context, sp model.Survey123Payload, rr []model.EquivalencyTableRow) error
}

type survey123Service struct {
	db *model.Database
	*model.Queries
}

func NewSurvey123Service(db *model.Database, q *model.Queries) *survey123Service {
	return &survey123Service{db, q}
}

func (s survey123Service) CreateSurvey123(ctx context.Context, sv model.Survey123) (uuid.UUID, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	newID, err := qtx.CreateSurvey123(ctx, sv)
	if err != nil {
		return uuid.Nil, err
	}
	return newID, tx.Commit()
}

func (s survey123Service) CreateOrUpdateSurvey123EquivalencyTable(ctx context.Context, survey123ID uuid.UUID, mappings []model.EquivalencyTableRow) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, r := range mappings {
		if r.TimeseriesID != nil {
			if err := qtx.GetIsValidEquivalencyTableTimeseries(ctx, *r.TimeseriesID); err != nil {
				return err
			}
		}
		if err := qtx.CreateOrUpdateSurvey123EquivalencyTableRow(ctx, survey123ID, r); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (s survey123Service) CreateOrUpdateSurvey123Preview(ctx context.Context, survey123ID uuid.UUID, previewRawJson []byte) error {
	pgJSON := pgtype.JSON{}
	if err := pgJSON.Set(previewRawJson); err != nil {
		return err
	}

	return s.db.Queries().CreateOrUpdateSurvey123Preview(ctx, survey123ID, pgJSON)
}

func (s survey123Service) CreateOrUpdateSurvey123Measurements(ctx context.Context, sp model.Survey123Payload, rr []model.EquivalencyTableRow) error {
	eqt := make(map[string]model.EquivalencyTableRow)
	for _, r := range rr {
		eqt[r.FieldName] = model.EquivalencyTableRow{
			TimeseriesID: r.TimeseriesID,
			InstrumentID: r.InstrumentID,
		}
	}

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	switch sp.EventType {
	case "addData":
		for _, r := range sp.Adds {
			mm, err := parseAttributes(r.Attributes, eqt)
			if err != nil {
				return err
			}
			for _, m := range mm {
				if err := qtx.CreateTimeseriesMeasurement(ctx, m.TimeseriesID, m.Time, m.Value); err != nil {
					return err
				}
			}
		}
	case "editData":
		for _, r := range sp.Updates {
			mm, err := parseAttributes(r.Attributes, eqt)
			if err != nil {
				return err
			}
			for _, m := range mm {
				if err := qtx.CreateTimeseriesMeasurement(ctx, m.TimeseriesID, m.Time, m.Value); err != nil {
					return err
				}
			}
		}

		for _, r := range sp.Adds {
			mm, err := parseAttributes(r.Attributes, eqt)
			if err != nil {
				return err
			}
			for _, m := range mm {
				if err := qtx.CreateTimeseriesMeasurement(ctx, m.TimeseriesID, m.Time, m.Value); err != nil {
					return err
				}
			}
		}
	default:
		return errors.New("invalid value for 'eventType'")
	}

	return tx.Commit()
}

func parseAttributes(attr map[string]interface{}, eqt map[string]model.EquivalencyTableRow) ([]model.Measurement, error) {
	errs := make([]error, 0)
	mappings := make(map[string]map[string]interface{})
	// group by instrument prefix
	for k, v := range attr {
		if k == "" {
			errs = append(errs, errors.New("invalid key format, skipping fields with empty string key"))
			continue
		}
		before, after, found := strings.Cut(k, "__")
		if !found {
			mappings[""] = map[string]interface{}{k: v}
			continue
		}
		mappings[before] = map[string]interface{}{after: v}
	}

	mm := make([]model.Measurement, 0)
	for prefix, subMap := range mappings {
		dtKey := "datetime"
		commentKey := "comment"

		dtTmp, exists := subMap[dtKey]
		if !exists {
			if prefix != "" {
				dtKey = "__" + dtKey
			}
			errs = append(errs, fmt.Errorf("expected '%s%s' field to be present", prefix, dtKey))
			continue
		}
		dt, ok := dtTmp.(int64)
		if !ok {
			if prefix != "" {
				dtKey = "__" + dtKey
			}
			errs = append(errs, fmt.Errorf("expected '%s%s' field to have Unix timsestamp (epoch) format", prefix, dtKey))
			continue
		}

		t := time.UnixMilli(dt)

		commentTmp, _ := subMap[commentKey].(string)
		var comment *string
		if commentTmp != "" {
			comment = &commentTmp
		}

		delete(subMap, dtKey)
		delete(subMap, commentKey)

		for subKey, val := range subMap {
			if prefix != "" {
				subKey = "__" + subKey
			}
			var v float64
			switch _val := val.(type) {
			case float64:
				v = _val
			case string:
				var err error
				v, err = strconv.ParseFloat(_val, 64)
				if err != nil {
					errs = append(errs, fmt.Errorf("invalid value type for field '%s%s' (must be number)", prefix, subKey))
					continue
				}
			default:
				errs = append(errs, fmt.Errorf("invalid value type for field '%s%s' (must be number)", prefix, subKey))
				continue
			}
			r, exists := eqt[prefix+subKey]
			if !exists {
				if !exists || r.TimeseriesID == nil {
					errs = append(errs, fmt.Errorf("row %s%s does not exists in equivalency table or timeseries id not assigned", prefix, subKey))
					continue
				}
			}
			mm = append(mm, model.Measurement{TimeseriesID: *r.TimeseriesID, Time: t, Value: v, TimeseriesNote: model.TimeseriesNote{Annotation: comment}})
		}
	}
	return mm, errors.Join(errs...)
}
