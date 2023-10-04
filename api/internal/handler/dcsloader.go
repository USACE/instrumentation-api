package handler

import (
	"context"
	"io"
	"log"
	"time"
)

func (h *DcsLoaderHandler) Start() error {
	handler := func(ctx context.Context, r io.Reader) error {
		mcs, mCount, err := h.DcsLoaderService.ParseCsvMeasurementCollection(r)
		if err != nil {
			return err
		}

		startPostTime := time.Now()
		if err := h.DcsLoaderService.PostMeasurementCollectionToApi(mcs); err != nil {
			return err
		}
		log.Printf(
			"\n\tSUCCESS; POST %d measurements across %d timeseries in %f seconds\n",
			mCount, len(mcs), time.Since(startPostTime).Seconds(),
		)
		return nil
	}

	return h.Pubsub.ProcessMessages(context.Background(), handler)
}
