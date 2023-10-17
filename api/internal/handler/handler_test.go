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

const host = "http://localhost:8080"

// HTTPTest contains parameters for HTTP Integration Tests
type HTTPTest struct {
	Name           string
	URL            string
	Method         string
	Body           string
	ExpectedStatus int
	ExpectedSchema *gojsonschema.JSONLoader
}

// RunHTTPTest accepts a HTTPTest type to execute the HTTP request
func RunHTTPTest(test HTTPTest) (*http.Response, error) {
	req, err := http.NewRequest(test.Method, host+test.URL, strings.NewReader(test.Body))
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()

	cfg := config.NewApiConfig()
	h := handler.NewApi(cfg)
	s := server.NewApiServer(cfg, h)

	s.ServeHTTP(rr, req)
	return rr.Result(), err
}

func RunAll(t *testing.T, tests []HTTPTest) {
	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			run, err := RunHTTPTest(v)
			assert.Nil(t, err)
			body, err := io.ReadAll(run.Body)
			assert.Nil(t, err)
			assert.Equal(t, v.ExpectedStatus, run.StatusCode)
			if v.ExpectedSchema != nil {
				loader, err := gojsonschema.NewReaderLoader(run.Body)
				assert.Nil(t, err)
				result, validationErr := gojsonschema.Validate(*v.ExpectedSchema, loader)
				assert.Nil(t, validationErr)
				assert.True(t, result.Valid())
			}
			t.Logf("Test %v got: %v\n", v.Name, string(body))
		})
	}
}
