package handler

import (
	"io"
	"log"
	"time"
)

// Entrypoint for Dcs Loader service queue
func (h *DcsLoaderHandler) Start() error {
	handler := func(r io.Reader) error {
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

	return h.Pubsub.ProcessMessages(handler)
}
