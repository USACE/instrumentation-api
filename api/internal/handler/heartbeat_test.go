package handler_test

import (
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

func TestHeartbeat(t *testing.T) {
	tests := []HTTPTest[model.Heartbeat]{
		{
			Name:           "DoHeartbeat",
			URL:            "/heartbeat?key=appkey",
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "GetLatestHeartbeat",
			URL:            "/heartbeat/latest",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
