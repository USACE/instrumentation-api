package handler_test

import (
	"net/http"
	"testing"
)

func TestHeartbeat(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "DoHeartbeat",
			URL:            "/heartbeat",
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "GetLatestHeartbeat",
			URL:            "/heartbeats/latest",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
