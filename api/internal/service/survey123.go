package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type Survey123Service interface {
	ListSurvey123sForProject(ctx context.Context, projectID uuid.UUID) ([]model.Survey123, error)
	ListSurvey123EquivalencyTableRows(ctx context.Context, survey123ID uuid.UUID) ([]model.Survey123EquivalencyTableRow, error)
	CreateSurvey123(ctx context.Context, sv model.Survey123) (uuid.UUID, error)
	UpdateSurvey123(ctx context.Context, sv model.Survey123) error
	SoftDeleteSurvey123(ctx context.Context, survey123ID uuid.UUID) error
	GetSurvey123Preview(ctx context.Context, survey123ID uuid.UUID) (model.Survey123Preview, error)
	CreateOrUpdateSurvey123Preview(ctx context.Context, pv model.Survey123Preview) error
	CreateOrUpdateSurvey123Measurements(ctx context.Context, survey123ID uuid.UUID, sp model.Survey123Payload, rr []model.Survey123EquivalencyTableRow) error
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

	tsIDs := make([]uuid.UUID, 0)
	for _, r := range sv.Rows {
		if r.TimeseriesID != nil {
			tsIDs = append(tsIDs, *r.TimeseriesID)
		}
	}

	if err = qtx.GetIsValidEquivalencyTableTimeseriesBatch(ctx, tsIDs); err != nil {
		return uuid.Nil, err
	}

	for _, r := range sv.Rows {
		if err := qtx.CreateOrUpdateSurvey123EquivalencyTableRow(ctx, newID, r); err != nil {
			return uuid.Nil, err
		}
	}

	return newID, tx.Commit()
}

func (s survey123Service) UpdateSurvey123(ctx context.Context, sv model.Survey123) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateSurvey123(ctx, sv); err != nil {
		return err
	}

	if err := qtx.DeleteAllSurvey123EquivalencyTableRows(ctx, sv.ID); err != nil {
		return err
	}

	tsIDs := make([]uuid.UUID, 0)
	for _, r := range sv.Rows {
		if r.TimeseriesID != nil {
			tsIDs = append(tsIDs, *r.TimeseriesID)
		}
	}

	if err = qtx.GetIsValidEquivalencyTableTimeseriesBatch(ctx, tsIDs); err != nil {
		return err
	}

	for _, r := range sv.Rows {
		if err := qtx.CreateOrUpdateSurvey123EquivalencyTableRow(ctx, sv.ID, r); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s survey123Service) createOrUpdateSurvey123PayloadError(ctx context.Context, survey123ID uuid.UUID, errMsgs []string) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.DeleteAllSurvey123PayloadErrors(ctx, survey123ID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	for _, errMsg := range errMsgs {
		if err := qtx.CreateSurvey123PayloadError(ctx, survey123ID, errMsg); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s survey123Service) CreateOrUpdateSurvey123Measurements(ctx context.Context, survey123ID uuid.UUID, sp model.Survey123Payload, rr []model.Survey123EquivalencyTableRow) error {
	eqt := make(map[string]model.Survey123EquivalencyTableRow)
	for _, r := range rr {
		eqt[r.FieldName] = model.Survey123EquivalencyTableRow{
			TimeseriesID: r.TimeseriesID,
			InstrumentID: r.InstrumentID,
		}
	}

	em := make([]string, 0)
	defer func() {
		if err := s.createOrUpdateSurvey123PayloadError(ctx, survey123ID, em); err != nil {
			log.Printf(err.Error())
		}
	}()

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	switch sp.EventType {
	case "addData":
		for _, edit := range sp.Edits {
			for _, r := range edit.Adds {
				mm, err := parseAttributes(r.Attributes, eqt)
				if err != nil {
					em = append(em, err.Error())
					continue
				}
				for _, m := range mm {
					if err := qtx.CreateTimeseriesMeasurement(ctx, m.TimeseriesID, m.Time, m.Value); err != nil {
						return err
					}
				}
			}
		}
	case "editData":
		for _, edit := range sp.Edits {
			for _, r := range edit.Updates {
				mm, err := parseAttributes(r.Attributes, eqt)
				if err != nil {
					em = append(em, err.Error())
					continue
				}
				for _, m := range mm {
					if err := qtx.CreateTimeseriesMeasurement(ctx, m.TimeseriesID, m.Time, m.Value); err != nil {
						return err
					}
				}
			}
			for _, r := range edit.Adds {
				mm, err := parseAttributes(r.Attributes, eqt)
				if err != nil {
					em = append(em, err.Error())
					continue
				}
				for _, m := range mm {
					if err := qtx.CreateTimeseriesMeasurement(ctx, m.TimeseriesID, m.Time, m.Value); err != nil {
						return err
					}
				}
			}
		}
	default:
		return errors.New("invalid value for 'eventType'")
	}

	return tx.Commit()
}

func parseAttributes(attr map[string]interface{}, eqt map[string]model.Survey123EquivalencyTableRow) ([]model.Measurement, error) {
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
