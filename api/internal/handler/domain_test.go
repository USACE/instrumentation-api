package handler_test

import (
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

func TestDomain(t *testing.T) {
	tests := []HTTPTest[model.Domain]{{
		Name:                 "GetDomains",
		URL:                  "/domains",
		Method:               http.MethodGet,
		ExpectedStatus:       http.StatusOK,
		ExpectedResponseType: jsonArr,
	}}

	RunAll(t, tests)
}
