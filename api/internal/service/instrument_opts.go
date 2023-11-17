package service

import (
	"context"
	"fmt"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

func handleOpts(ctx context.Context, q *model.Queries, inst model.Instrument, rt requestType) error {
	switch inst.TypeID {
	case saaTypeID:
		opts, err := model.MapToStruct[model.SaaOpts](inst.Opts)
		if err != nil {
			return err
		}
		if rt == create {
			for i := 1; i <= opts.NumSegments; i++ {
				tsConstant := model.Timeseries{
					InstrumentID: inst.ID,
					Slug:         inst.Slug + fmt.Sprintf("segment-%d-length", i),
					Name:         inst.Slug + fmt.Sprintf("segment-%d-length", i),
					ParameterID:  model.SaaParameterID,
					UnitID:       model.MeterUnitID,
				}

				tsNew, err := q.CreateTimeseries(ctx, tsConstant)
				if err != nil {
					return err
				}
				if err := q.CreateInstrumentConstant(ctx, inst.ID, tsNew.ID); err != nil {
					return err
				}
				if err := q.CreateSaaSegment(ctx, model.SaaSegment{ID: i, InstrumentID: inst.ID, LengthTimeseriesID: tsNew.ID}); err != nil {
					return err
				}
			}

			tsConstant := model.Timeseries{
				InstrumentID: inst.ID,
				Slug:         inst.Slug + "-bottom-elevation",
				Name:         inst.Slug + "-bottom-elevation",
				ParameterID:  model.SaaParameterID,
				UnitID:       model.MeterUnitID,
			}

			tsNew, err := q.CreateTimeseries(ctx, tsConstant)
			if err != nil {
				return err
			}
			if err := q.CreateInstrumentConstant(ctx, inst.ID, tsNew.ID); err != nil {
				return err
			}
			opts.BottomElevationTimeseriesID = tsNew.ID
			if err := q.CreateSaaOpts(ctx, inst.ID, opts); err != nil {
				return err
			}
		}
		if rt == update {
			if err := q.UpdateSaaOpts(ctx, inst.ID, opts); err != nil {
				return err
			}
		}
		if err := q.CreateTimeseriesMeasurement(ctx, opts.BottomElevationTimeseriesID, time.Now(), opts.BottomElevation); err != nil {
			return err
		}
	case ipiTypeID:
		opts, err := model.MapToStruct[model.IpiOpts](inst.Opts)
		if err != nil {
			return err
		}
		if rt == create {
			for i := 1; i <= opts.NumSegments; i++ {
				tsConstant := model.Timeseries{
					InstrumentID: inst.ID,
					Slug:         inst.Slug + fmt.Sprintf("segment-%d-length", i),
					Name:         inst.Slug + fmt.Sprintf("segment-%d-length", i),
					ParameterID:  model.IpiParameterID,
					UnitID:       model.MeterUnitID,
				}

				tsNew, err := q.CreateTimeseries(ctx, tsConstant)
				if err != nil {
					return err
				}
				if err := q.CreateInstrumentConstant(ctx, inst.ID, tsNew.ID); err != nil {
					return err
				}
				if err := q.CreateIpiSegment(ctx, model.IpiSegment{ID: i, InstrumentID: inst.ID, LengthTimeseriesID: tsNew.ID}); err != nil {
					return err
				}
			}

			tsConstant := model.Timeseries{
				InstrumentID: inst.ID,
				Slug:         inst.Slug + "-bottom-elevation",
				Name:         inst.Slug + "-bottom-elevation",
				ParameterID:  model.IpiParameterID,
				UnitID:       model.MeterUnitID,
			}

			tsNew, err := q.CreateTimeseries(ctx, tsConstant)
			if err != nil {
				return err
			}
			if err := q.CreateInstrumentConstant(ctx, inst.ID, tsNew.ID); err != nil {
				return err
			}
			opts.BottomElevationTimeseriesID = tsNew.ID
			if err := q.CreateIpiOpts(ctx, inst.ID, opts); err != nil {
				return err
			}
		}
		if rt == update {
			if err := q.UpdateIpiOpts(ctx, inst.ID, opts); err != nil {
				return err
			}
		}
		if err := q.CreateTimeseriesMeasurement(ctx, opts.BottomElevationTimeseriesID, time.Now(), opts.BottomElevation); err != nil {
			return err
		}
	default:
	}
	return nil
}
