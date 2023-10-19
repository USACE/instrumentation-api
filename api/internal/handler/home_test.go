package handler_test

import (
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

func TestHome(t *testing.T) {
	tests := []HTTPTest[model.Home]{{
		Name:                 "GetHome",
		URL:                  "/home",
		Method:               http.MethodGet,
		ExpectedStatus:       http.StatusOK,
		ExpectedResponseType: jsonObj,
	}}

	RunAll(t, tests)
}
