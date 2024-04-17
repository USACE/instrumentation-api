package service

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
)

type DcsLoaderService interface {
	ParseCsvMeasurementCollection(r io.Reader) ([]model.MeasurementCollection, int, error)
	PostMeasurementCollectionToApi(mcs []model.MeasurementCollection) error
}

type dcsLoaderService struct {
	apiClient *http.Client
	cfg       *config.DcsLoaderConfig
}

func NewDcsLoaderService(apiClient *http.Client, cfg *config.DcsLoaderConfig) *dcsLoaderService {
	return &dcsLoaderService{apiClient, cfg}
}

func (s dcsLoaderService) ParseCsvMeasurementCollection(r io.Reader) ([]model.MeasurementCollection, int, error) {
	mcs := make([]model.MeasurementCollection, 0)
	mCount := 0
	reader := csv.NewReader(r)

	rows := make([][]string, 0)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return mcs, mCount, err
		}
		rows = append(rows, row)
	}

	mcMap := make(map[uuid.UUID]*model.MeasurementCollection)
	for _, row := range rows {
		// 0=timeseries_id, 1=time, 2=value
		tsid, err := uuid.Parse(row[0])
		if err != nil {
			return mcs, mCount, err
		}
		t, err := time.Parse(time.RFC3339, row[1])
		if err != nil {
			return mcs, mCount, err
		}
		v, err := strconv.ParseFloat(row[2], 32)
		if err != nil {
			return mcs, mCount, err
		}

		if _, ok := mcMap[tsid]; !ok {
			mcMap[tsid] = &model.MeasurementCollection{
				TimeseriesID: tsid,
				Items:        make([]model.Measurement, 0),
			}
		}
		mcMap[tsid].Items = append(mcMap[tsid].Items, model.Measurement{TimeseriesID: tsid, Time: t, Value: v})
		mCount++
	}

	mcs = make([]model.MeasurementCollection, len(mcMap))
	idx := 0
	for _, v := range mcMap {
		mcs[idx] = *v
		idx++
	}

	return mcs, mCount, nil
}

func (s dcsLoaderService) PostMeasurementCollectionToApi(mcs []model.MeasurementCollection) error {
	requestBodyBytes, err := json.Marshal(mcs)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s?key=%s", s.cfg.PostURL, s.cfg.APIKey), bytes.NewReader(requestBodyBytes))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	req.Header.Add("Content-Type", "application/json")

	resp, err := s.apiClient.Do(req)
	if err != nil {
		urlErr := err.(*url.Error)
		urlRedact := util.RedactRequest{URL: urlErr.URL}
		if err := urlRedact.RedactQueryParam("key"); err != nil {
			return err
		}
		urlErr.URL = urlRedact.URL
		log.Printf("\n\t*** Error; unable to make request; %s", urlErr.Error())
		return urlErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		log.Printf("\n\t*** Error; Status Code: %d ***\n", resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body")
			return err
		}
		log.Printf("%s\n", body)
	}
	return nil
}
