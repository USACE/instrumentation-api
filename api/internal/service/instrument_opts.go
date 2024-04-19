package service

import (
	"context"
	"fmt"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

func handleSaaOpts(ctx context.Context, q *model.Queries, inst model.Instrument, rt requestType) error {
	opts, err := model.MapToStruct[model.SaaOpts](inst.Opts)
	if err != nil {
		return err
	}
	switch rt {
	case create:
		for i := 1; i <= opts.NumSegments; i++ {
			tsConstant := model.Timeseries{
				InstrumentID: inst.ID,
				ParameterID:  model.SaaParameterID,
				UnitID:       model.FeetUnitID,
			}
			tsConstant.Slug = inst.Slug + fmt.Sprintf("segment-%d-length", i)
			tsConstant.Name = inst.Slug + fmt.Sprintf("segment-%d-length", i)

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
			ParameterID:  model.SaaParameterID,
			UnitID:       model.FeetUnitID,
		}
		tsConstant.Slug = inst.Slug + "-bottom-elevation"
		tsConstant.Name = inst.Slug + "-bottom-elevation"

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
	case update:
		if err := q.UpdateSaaOpts(ctx, inst.ID, opts); err != nil {
			return err
		}
	}
	if err := q.CreateTimeseriesMeasurement(ctx, opts.BottomElevationTimeseriesID, time.Now(), opts.BottomElevation); err != nil {
		return err
	}
	return nil
}

func handleIpiOpts(ctx context.Context, q *model.Queries, inst model.Instrument, rt requestType) error {
	opts, err := model.MapToStruct[model.IpiOpts](inst.Opts)
	if err != nil {
		return err
	}
	switch rt {
	case create:
		for i := 1; i <= opts.NumSegments; i++ {
			tsConstant := model.Timeseries{
				InstrumentID: inst.ID,
				ParameterID:  model.IpiParameterID,
				UnitID:       model.FeetUnitID,
			}
			tsConstant.Slug = inst.Slug + fmt.Sprintf("segment-%d-length", i)
			tsConstant.Name = inst.Slug + fmt.Sprintf("segment-%d-length", i)

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
			ParameterID:  model.IpiParameterID,
			UnitID:       model.FeetUnitID,
		}
		tsConstant.Slug = inst.Slug + "-bottom-elevation"
		tsConstant.Name = inst.Slug + "-bottom-elevation"

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
	case update:
		if err := q.UpdateIpiOpts(ctx, inst.ID, opts); err != nil {
			return err
		}
	}
	if err := q.CreateTimeseriesMeasurement(ctx, opts.BottomElevationTimeseriesID, time.Now(), opts.BottomElevation); err != nil {
		return err
	}
	return nil
}

func handleOpts(ctx context.Context, q *model.Queries, inst model.Instrument, rt requestType) error {
	switch inst.TypeID {
	case saaTypeID:
		return handleSaaOpts(ctx, q, inst, rt)
	case ipiTypeID:
		return handleIpiOpts(ctx, q, inst, rt)
	}
	return nil
}
