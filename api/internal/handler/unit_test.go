package handler_test

import (
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

func TestUnits(t *testing.T) {
	tests := []HTTPTest[model.Unit]{{
		Name:                 "ListUnits",
		URL:                  "/units",
		Method:               http.MethodGet,
		ExpectedStatus:       http.StatusOK,
		ExpectedResponseType: jsonArr,
	}}

	RunAll(t, tests)
}
