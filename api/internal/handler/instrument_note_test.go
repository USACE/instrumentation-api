package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

const (
	testInstrumentNoteID          = "90a3f8de-de65-48a7-8286-024c13162958"
	testInstrumentNoteIntrumentID = "a7540f69-c41e-43b3-b655-6e44097edb7e"
)

const putInstrumentNoteBody = `{
    "id": "90a3f8de-de65-48a7-8286-024c13162958",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "title": "Instrument Test Note 1",
    "body": "Updated instrument note body text.  This is example updated text.",
    "time": "2020-06-09T01:49:48.505713Z"
}`

const createInstrumentNoteArrayBody = `[
    {
        "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
        "title": "Instrument Test Note 101",
        "body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut\n labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris\n nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse\n cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui\n officia deserunt mollit anim id est laborum.\n",
        "time": "2020-06-09T01:49:48.505713Z"
    },
        {
        "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
        "title": "Instrument Test Note 102",
        "body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut\n labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris\n nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse\n cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui\n officia deserunt mollit anim id est laborum.\n",
        "time": "2020-06-09T01:49:48.505713Z"
    },
        {
        "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
        "title": "Instrument Test Note 103",
        "body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut\n labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris\n nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse\n cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui\n officia deserunt mollit anim id est laborum.\n",
        "time": "2020-06-09T01:49:48.505713Z"
    }
]`

const createInstrumentNoteObjectBody = `{
        "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
        "title": "Instrument Note from Object Upload",
        "body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut\n labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris\n nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse\n cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui\n officia deserunt mollit anim id est laborum.\n",
        "time": "2020-06-09T01:49:48.505713Z"
    }`

func TestInstrumentNotes(t *testing.T) {
	tests := []HTTPTest[model.InstrumentNote]{
		{
			Name:                 "GetInstrumentNote",
			URL:                  fmt.Sprintf("/instruments/%s/notes/%s", testInstrumentNoteIntrumentID, testInstrumentNoteID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "ListInstrumentNotes",
			URL:                  "/instruments/notes",
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListInstrumentInstrumentNotes",
			URL:                  fmt.Sprintf("/instruments/%s/notes", testInstrumentNoteIntrumentID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "PutInstrumentNote",
			URL:                  fmt.Sprintf("/instruments/%s/notes/%s", testInstrumentNoteIntrumentID, testInstrumentNoteID),
			Method:               http.MethodPut,
			Body:                 putInstrumentNoteBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "CreateInstrumentNote_Array",
			URL:                  "/instruments/notes",
			Method:               http.MethodPost,
			Body:                 createInstrumentNoteArrayBody,
			ExpectedStatus:       http.StatusCreated,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "CreateInstrumentNote_Object",
			URL:                  "/instruments/notes",
			Method:               http.MethodPost,
			Body:                 createInstrumentNoteObjectBody,
			ExpectedStatus:       http.StatusCreated,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:           "DeleteInstrumentNote",
			URL:            fmt.Sprintf("/instruments/notes/%s", testInstrumentNoteID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
