package date_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/date"
)

func TestHandleDate(t *testing.T) {
	// set up test table
	testcases := []struct {
		name               string
		url                string
		method             string
		expectedStatusCode int
		expectBody         bool
		expectedDateAfter  time.Time
		expectedDateBefore time.Time
		expectedGroup      data.DateGroup
	}{
		{
			name:               "HandleDate returns random date",
			url:                "/date",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      "",
		},
		{
			name:               "HandleDate returns 405 on POST request",
			url:                "/date",
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectBody:         false,
			expectedGroup:      "",
		},
		{
			name:               "HandleDate returns random date of Movies group",
			url:                "/date?group=movies",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      data.DateGroupMovies,
		},
		{
			name:               "HandleDate returns random date after 1985-01-01",
			url:                "/date?after=1985-01-01",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedDateAfter:  time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedGroup:      "",
		},
		{
			name:               "HandleDate returns random date after 1995-01-01",
			url:                "/date?before=1995-01-01",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedDateBefore: time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedGroup:      "",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// create request and response objects
			req := httptest.NewRequest(testcase.method, testcase.url, nil)
			res := httptest.NewRecorder()

			// send request to handler and record response
			date.HandleDate(res, req)

			// check status code
			if res.Code != testcase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testcase.expectedStatusCode, res.Code)
			}

			// check for CORS header
			corsHeader := res.Header().Get("Access-Control-Allow-Origin")
			if corsHeader != "*" {
				t.Errorf("expected CORS header to be set to \"*\", got %s", corsHeader)
			}

			// if body is expected to have contents
			if testcase.expectBody {
				// parse response body
				var dateResponse date.DateResponse
				err := json.Unmarshal(res.Body.Bytes(), &dateResponse)
				if err != nil {
					t.Error("error parsing response body:", err)
				}

				// get list of dates
				dates, err := data.GetDates()
				if err != nil {
					t.Error("error getting dates:", err)
				}

				// filter list of dates by group
				filteredDates := data.FilterDates(
					dates,
					testcase.expectedDateAfter,
					testcase.expectedDateBefore,
					testcase.expectedGroup,
					dateResponse.Trivia,
				)

				// throw error if there isn't exactly a single entry after filtering
				if len(filteredDates) != 1 {
					t.Errorf("expected 1 date match, got %d", len(filteredDates))
				}
			}
		})
	}
}
