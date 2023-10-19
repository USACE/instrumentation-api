package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

const testCollectionGroupID = "30b32cb1-0936-42c4-95d1-63a7832a57db"

func TestCollectionGroups(t *testing.T) {
	detailTests := []HTTPTest[model.CollectionGroupDetails]{{
		Name:                 "GetCollectionGroupDetails",
		URL:                  fmt.Sprintf("/projects/%s/collection_groups/%s", testProjectID, testCollectionGroupID),
		Method:               http.MethodGet,
		ExpectedStatus:       http.StatusOK,
		ExpectedResponseType: jsonObj,
	}}
	RunAll(t, detailTests)

	tests := []HTTPTest[model.CollectionGroup]{
		{
			Name:                 "ListCollectionGroups",
			URL:                  fmt.Sprintf("/projects/%s/collection_groups", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:           "DeleteCollectionGroup",
			URL:            fmt.Sprintf("/projects/%s/collection_groups/%s", testProjectID, testCollectionGroupID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
