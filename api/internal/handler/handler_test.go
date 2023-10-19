package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
	"github.com/USACE/instrumentation-api/api/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const (
	host    = "http://localhost:8080"
	mockJwt = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwibmFtZSI6IlVzZXIuQWRtaW4iLCJpYXQiOjE1MTYyMzkwMjIsImV4cCI6MjAwMDAwMDAwMCwicm9sZXMiOlsiUFVCTElDLlVTRVIiXX0.4VAMamtH92GiIb5CpGKpP6LKwU6IjIfw5wS4qc8O8VM`
)

// HTTPTest contains parameters for HTTP Integration Tests
type HTTPTest struct {
	Name           string
	URL            string
	Method         string
	Body           string
	ExpectedStatus int
	ExpectedSchema *gojsonschema.Schema
	authHeader     string
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
func RunHTTPTest(v HTTPTest) (*http.Response, error) {
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

func RunAll(t *testing.T, tests []HTTPTest) {
	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			run, err := RunHTTPTest(v)
			assert.Nil(t, err, "error calling RunHTTPTest(v)")
			if err != nil {
				t.Log(err.Error())
			}

			assert.Equal(t, v.ExpectedStatus, run.StatusCode)

			body, err := io.ReadAll(run.Body)
			assert.Nil(t, err, "error calling io.ReadAll(run.Body)")
			if err != nil {
				t.Log(err.Error())
				return
			}

			if v.ExpectedStatus != run.StatusCode {
				t.Log(string(body))
				return
			}

			if v.ExpectedSchema != nil {
				loader := gojsonschema.NewBytesLoader(body)

				result, err := v.ExpectedSchema.Validate(loader)
				assert.Nil(t, err, "error calling v.ExpectedSchema.Validate")
				if err != nil {
					return
				}

				valid := result.Valid()

				assert.Truef(t, valid, "response body did not match json schema:")
				if !valid {
					var errs string
					for _, err := range result.Errors() {
						errs += "\n"
						errs += err.String()
					}
					t.Log(errs)
				}
			}
		})
	}
}
