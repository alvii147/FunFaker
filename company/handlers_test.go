package company_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvii147/FunFaker/company"
	"github.com/alvii147/FunFaker/data"
)

func TestHandleCompany(t *testing.T) {
	// set up test table
	testcases := []struct {
		name               string
		url                string
		method             string
		expectedStatusCode int
		expectBody         bool
		expectedGroup      data.CompanyGroup
	}{
		{
			name:               "HandleCompany returns random company",
			url:                "/company",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      "",
		},
		{
			name:               "HandleCompany returns 405 on POST request",
			url:                "/company",
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectBody:         false,
			expectedGroup:      "",
		},
		{
			name:               "HandleCompany returns random address of Comics group",
			url:                "/company?group=comics",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      data.CompanyGroupComics,
		},
		{
			name:               "HandleCompany returns 400 on invalid URL parameters",
			url:                "/company?invalid=parameter",
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
			company.HandleCompany(res, req)

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
				var companyResponse company.CompanyResponse
				err := json.Unmarshal(res.Body.Bytes(), &companyResponse)
				if err != nil {
					t.Error("error parsing response body:", err)
				}

				// get list of companies
				companies, err := data.GetCompanies()
				if err != nil {
					t.Error("error getting companies:", err)
				}

				// filter list of companies by group
				filteredCompanies := data.FilterCompanies(
					companies,
					companyResponse.Name,
					testcase.expectedGroup,
					companyResponse.Trivia,
				)

				// throw error if there isn't exactly a single entry after filtering
				if len(filteredCompanies) != 1 {
					t.Errorf("expected 1 name match, got %d", len(filteredCompanies))
				}
			}
		})
	}
}
