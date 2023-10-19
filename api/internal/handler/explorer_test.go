package handler_test

import (
	"fmt"
	"net/http"
	"testing"
)

const (
	testExplorerTimeAfter  = "2020-03-03T00:00:00Z"
	testExplorerTimeBefore = "2021-05-27T00:32:45Z"
)

const postExplorerBody = `["9e8f2ca4-4037-45a4-aaca-d9e598877439", "a7540f69-c41e-43b3-b655-6e44097edb7e"]`

func TestExplorer(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "PostExplorer Query By InstrumetID",
			URL:            "/explorer",
			Method:         http.MethodPost,
			Body:           postExplorerBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "PostExplorer Query By InstrumetID With Date Filters",
			URL:            fmt.Sprintf("/explorer?after=%s&before=%s", testExplorerTimeAfter, testExplorerTimeBefore),
			Method:         http.MethodPost,
			Body:           postExplorerBody,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
