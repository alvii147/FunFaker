package website_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/website"
)

func TestHandleWebsite(t *testing.T) {
	// set up test table
	testcases := []struct {
		name               string
		url                string
		method             string
		expectedStatusCode int
		expectBody         bool
		expectedGroup      data.WebsiteGroup
	}{
		{
			name:               "HandleWebsite returns random website",
			url:                "/website",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      "",
		},
		{
			name:               "HandleWebsite returns 405 on POST request",
			url:                "/website",
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectBody:         false,
			expectedGroup:      "",
		},
		{
			name:               "HandleWebsite returns random website of TV-shows group",
			url:                "/website?group=tv-shows",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      data.WebsiteGroupTVShows,
		},
		{
			name:               "HandleWebsite returns 400 on invalid URL parameters",
			url:                "/website?invalid=parameter",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusBadRequest,
			expectBody:         false,
			expectedGroup:      "",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// create request and response objects
			req := httptest.NewRequest(testcase.method, testcase.url, nil)
			res := httptest.NewRecorder()

			// send request to handler and record response
			website.HandleWebsite(res, req)

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
				var websiteResponse website.WebsiteResponse
				err := json.Unmarshal(res.Body.Bytes(), &websiteResponse)
				if err != nil {
					t.Error("error parsing response body:", err)
				}

				// get list of websites
				websites, err := data.GetWebsites()
				if err != nil {
					t.Error("error getting websites:", err)
				}

				// filter list of websites by group
				filteredWebsites := data.FilterWebsites(
					websites,
					websiteResponse.URL,
					testcase.expectedGroup,
					websiteResponse.Trivia,
				)

				// throw error if there isn't exactly a single entry after filtering
				if len(filteredWebsites) != 1 {
					t.Errorf("expected 1 name match, got %d", len(filteredWebsites))
				}
			}
		})
	}
}
