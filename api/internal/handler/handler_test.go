package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	truncateLinesBody = 30
	host              = "http://localhost:8080"
	mockJwt           = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ikw0YXFVRmd6YV9RVjhqc1ZOa281OW5GVzl6bGh1b0JGX3RxdlpkTUZkajQifQ.eyJzdWIiOiJmOGRjYWZlYS0yNDNlLTRiODktOGQ3ZC1mYTAxOTE4MTMwZjQiLCJ0eXAiOiJCZWFyZXIiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cDovL2xvY2FsaG9zdDozMDAwIl0sIm5hbWUiOiJBbnRob255IExhbWJlcnQiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJ0ZXN0IiwiZ2l2ZW5fbmFtZSI6IkFudGhvbnkiLCJmYW1pbHlfbmFtZSI6IkxhbWJlcnQiLCJlbWFpbCI6ImFudGhvbnkubS5sYW1iZXJ0QGZha2UudXNhY2UuYXJteS5taWwiLCJzdWJqZWN0RE4iOiJsYW1iZXJ0LmFudGhvbnkubS4yIiwiY2FjVUlEIjoiMiJ9.8CjeifD51ZEZZOx9eeMd7RPanvtgkQQus-R19aU91Rw`
	mockAppKey        = "appkey"
)

const IDSlugNameArrSchema = `{
  "type": "array",
  "items": {
    "type": "object",
    "properties": {
      "id": { "type": "string" },
      "slug": { "type": "string" },
      "name": { "type": "string" }
    }
  }
}`

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
			run, httpErr := RunHTTPTest(v)
			assert.Nil(t, httpErr, "error calling RunHTTPTest(v)")
			if httpErr != nil {
				t.Log(httpErr.Error())
			}

			assert.Equal(t, v.ExpectedStatus, run.StatusCode)

			body, err := io.ReadAll(run.Body)
			assert.Nil(t, err, "error calling io.ReadAll(run.Body)")
			if err != nil {
				t.Log(err.Error())
				return
			}

			// truncate verbose -v output
			var dst bytes.Buffer
			if err := json.Indent(&dst, body, "", "  "); err != nil {
				s := string(body)
				if len(s) > 500 {
					s = fmt.Sprintf("%s\n...", s[:500])
				}
				t.Logf("could not format json response body: %s", s)
			} else {
				s := dst.String()
				ss := strings.Split(s, "\n")
				if len(ss) > truncateLinesBody {
					s = fmt.Sprintf("%s\n...", strings.Join(ss[:truncateLinesBody], "\n"))
				}
				t.Logf("response body: %s", s)
			}

			if v.ExpectedStatus != run.StatusCode {
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
