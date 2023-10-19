package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
	"github.com/USACE/instrumentation-api/api/internal/server"
	"github.com/stretchr/testify/assert"
)

const (
	host    = "http://localhost:8080"
	mockJwt = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwibmFtZSI6IlVzZXIuQWRtaW4iLCJpYXQiOjE1MTYyMzkwMjIsImV4cCI6MjAwMDAwMDAwMCwicm9sZXMiOlsiUFVCTElDLlVTRVIiXX0.4VAMamtH92GiIb5CpGKpP6LKwU6IjIfw5wS4qc8O8VM`
)

type responseType int

const (
	none responseType = iota
	jsonObj
	jsonArr
)

// HTTPTest contains parameters for HTTP Integration Tests
type HTTPTest[T any] struct {
	Name                 string
	URL                  string
	Method               string
	Body                 string
	ExpectedStatus       int
	ExpectedResponseType responseType
	authHeader           string
}

// singleton api server since database is used in integration tests
var testApi *server.ApiServer

func testApiServer() *server.ApiServer {
	if testApi == nil {
		cfg := config.NewApiConfig()
		h := handler.NewApi(cfg)
		testApi = server.NewApiServer(cfg, h)
	}
	return testApi
}

// RunHTTPTest accepts a HTTPTest type to execute the HTTP request
func RunHTTPTest[T any](v HTTPTest[T]) (*http.Response, error) {
	req, err := http.NewRequest(v.Method, host+v.URL, strings.NewReader(v.Body))
	if err != nil {
		return nil, err
	}

	if v.authHeader == "" {
		v.authHeader = mockJwt
	}

	req.Header.Set("Authorization", v.authHeader)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	s := testApiServer()

	s.ServeHTTP(rr, req)
	return rr.Result(), err
}

func RunAll[T any](t *testing.T, tests []HTTPTest[T]) {
	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			run, err := RunHTTPTest(v)
			assert.Nil(t, err)
			if err != nil {
				t.Log(err.Error())
			}

			assert.Equal(t, v.ExpectedStatus, run.StatusCode)

			body, err := io.ReadAll(run.Body)
			assert.Nil(t, err)
			if err != nil {
				t.Log(err.Error())
				return
			}

			if v.ExpectedStatus != run.StatusCode {
				t.Log(string(body))
			}

			if v.ExpectedResponseType == jsonObj {
				var res T
				err := json.Unmarshal(body, &res)
				assert.Nil(t, err)
				if err != nil {
					t.Log(err.Error())
				}
			}

			if v.ExpectedResponseType == jsonArr {
				var res []T
				err := json.Unmarshal(body, &res)
				assert.Nil(t, err)
				if err != nil {
					t.Log(err.Error())
				}
			}

		})
	}
}
