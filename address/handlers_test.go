package address_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvii147/FunFaker/address"
	"github.com/alvii147/FunFaker/data"
)

func TestHandleAddress(t *testing.T) {
	// set up test table
	testcases := []struct {
		name               string
		url                string
		method             string
		expectedStatusCode int
		expectBody         bool
		expectedGroup      data.AddressGroup
		expectedValidOnly  bool
	}{
		{
			name:               "HandleAddress returns random address",
			url:                "/name",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      "",
			expectedValidOnly:  false,
		},
		{
			name:               "HandleAddress returns 405 on POST request",
			url:                "/name",
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectBody:         false,
			expectedGroup:      "",
			expectedValidOnly:  false,
		},
		{
			name:               "HandleAddress returns random address of Breaking Bad group",
			url:                "/name?group=breaking-bad",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      data.AddressGroupBreakingBad,
			expectedValidOnly:  false,
		},
		{
			name:               "HandleAddress returns random valid address",
			url:                "/name?valid-only=true",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedGroup:      "",
			expectedValidOnly:  false,
		},
		{
			name:               "HandleAddress returns 400 on invalid URL parameters",
			url:                "/name?invalid=parameter",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusBadRequest,
			expectBody:         false,
			expectedGroup:      "",
			expectedValidOnly:  false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// create request and response objects
			req := httptest.NewRequest(testcase.method, testcase.url, nil)
			res := httptest.NewRecorder()

			// send request to handler and record response
			address.HandleAddress(res, req)
			if res.Code != testcase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testcase.expectedStatusCode, res.Code)
			}

			// if body is expected to have contents
			if testcase.expectBody {
				// parse response body
				var addressResponse address.AddressResponse
				err := json.Unmarshal(res.Body.Bytes(), &addressResponse)
				if err != nil {
					t.Error("error parsing response body:", err)
				}

				// get list of addresses
				addresses, err := data.GetAddresses()
				if err != nil {
					t.Error("error getting addresses:", err)
				}

				// filter list of addresses by group
				filteredAddresses := address.FilterAddresses(
					addresses,
					addressResponse.StreetName,
					addressResponse.City,
					addressResponse.State,
					addressResponse.Country,
					addressResponse.PostalCode,
					testcase.expectedGroup,
					addressResponse.Trivia,
				)
				// filter by validity if expected
				if testcase.expectedValidOnly {
					filteredAddresses = address.FilterValidAddresses(filteredAddresses)
				}

				// log error if there isn't exactly a single entry after filtering
				if len(filteredAddresses) != 1 {
					t.Errorf("expected 1 name match, got %d", len(filteredAddresses))
				}
			}
		})
	}
}