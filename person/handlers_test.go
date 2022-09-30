package person_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/mail"
	"strings"
	"testing"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/person"
	"github.com/alvii147/FunFaker/utils"
)

// parse email into username, domain name, and domain suffix
func ParseEmail(email string) (string, string, string, error) {
	// verify email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", "", "", err
	}

	// split email by last occurrence of "@" symbol
	atIdx := strings.LastIndex(email, "@")
	if atIdx < 0 {
		return "", "", "", errors.New("Unable to find \"@\"")
	}
	username := email[:atIdx]
	domain := email[atIdx+1:]

	// split domain by last occurrence of "." symbol
	dotIdx := strings.LastIndex(domain, ".")
	if dotIdx < 0 {
		return "", "", "", errors.New("Unable to find \".\"")
	}
	domainName := domain[:dotIdx]
	domainSuffix := domain[dotIdx+1:]

	return username, domainName, domainSuffix, nil
}

func TestHandleName(t *testing.T) {
	// set up test table
	testcases := []struct {
		name               string
		url                string
		method             string
		expectedStatusCode int
		expectBody         bool
		expectedSex        data.Sex
		expectedGroup      data.PersonGroup
	}{
		{
			name:               "HandleName returns random name",
			url:                "/name",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedSex:        "",
			expectedGroup:      "",
		},
		{
			name:               "HandleName returns 405 on POST request",
			url:                "/name",
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectBody:         false,
			expectedSex:        "",
			expectedGroup:      "",
		},
		{
			name:               "HandleName returns random Male name",
			url:                "/name?sex=male",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedSex:        data.SexMale,
			expectedGroup:      "",
		},
		{
			name:               "HandleName returns random Female name",
			url:                "/name?sex=female",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedSex:        data.SexFemale,
			expectedGroup:      "",
		},
		{
			name:               "HandleName returns random name of Comics group",
			url:                "/name?group=comics",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedSex:        "",
			expectedGroup:      data.PersonGroupComics,
		},
		{
			name:               "HandleName returns random Male name of Comics group",
			url:                "/name?sex=male&group=comics",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectBody:         true,
			expectedSex:        data.SexMale,
			expectedGroup:      data.PersonGroupComics,
		},
		{
			name:               "HandleName returns 400 on invalid URL parameters",
			url:                "/name?invalid=parameter",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusBadRequest,
			expectBody:         false,
			expectedSex:        "",
			expectedGroup:      "",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// create request and response objects
			req := httptest.NewRequest(testcase.method, testcase.url, nil)
			res := httptest.NewRecorder()

			// send request to handler and record response
			person.HandleName(res, req)
			if res.Code != testcase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testcase.expectedStatusCode, res.Code)
			}

			// if body is expected to have contents
			if testcase.expectBody {
				// parse response body
				var nameResponse person.NameResponse
				err := json.Unmarshal(res.Body.Bytes(), &nameResponse)
				if err != nil {
					t.Error("error parsing response body:", err)
				}

				// get list of names
				names, err := data.GetNames()
				if err != nil {
					t.Error("error getting names:", err)
				}

				// filter list of names to get the returned entry
				filteredNames := person.FilterNames(
					names,
					nameResponse.FirstName,
					nameResponse.LastName,
					testcase.expectedSex,
					testcase.expectedGroup,
					nameResponse.Trivia,
				)

				// log error if there isn't exactly a single entry after filtering
				if len(filteredNames) != 1 {
					t.Errorf("expected 1 name match, got %d", len(filteredNames))
				}
			}
		})
	}
}

func TestHandleEmail(t *testing.T) {
	// set up test table
	testcases := []struct {
		name                    string
		url                     string
		method                  string
		expectedStatusCode      int
		expectBody              bool
		expectEmailDomainName   string
		expectEmailDomainSuffix string
	}{
		{
			name:                    "HandleEmail returns random email",
			url:                     "/email",
			method:                  http.MethodGet,
			expectedStatusCode:      http.StatusOK,
			expectBody:              true,
			expectEmailDomainName:   "",
			expectEmailDomainSuffix: "",
		},
		{
			name:                    "HandleEmail returns 405 on POST request",
			url:                     "/email",
			method:                  http.MethodPost,
			expectedStatusCode:      http.StatusMethodNotAllowed,
			expectBody:              false,
			expectEmailDomainName:   "",
			expectEmailDomainSuffix: "",
		},
		{
			name:                    "HandleEmail returns random Female email of Comics group",
			url:                     "/email?sex=female&group=comics",
			method:                  http.MethodGet,
			expectedStatusCode:      http.StatusOK,
			expectBody:              true,
			expectEmailDomainName:   "comics",
			expectEmailDomainSuffix: "",
		},
		{
			name:                    "HandleEmail returns random email of Gmail domain name and org domain suffix",
			url:                     "/email?domain-name=gmail&domain-suffix=org",
			method:                  http.MethodGet,
			expectedStatusCode:      http.StatusOK,
			expectBody:              true,
			expectEmailDomainName:   "gmail",
			expectEmailDomainSuffix: "org",
		},
		{
			name:                    "HandleEmail returns 400 on GET request with invalid URL parameters",
			url:                     "/email?invalid=parameter",
			method:                  http.MethodGet,
			expectedStatusCode:      http.StatusBadRequest,
			expectBody:              false,
			expectEmailDomainName:   "",
			expectEmailDomainSuffix: "",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// create request and response objects
			req := httptest.NewRequest(testcase.method, testcase.url, nil)
			res := httptest.NewRecorder()

			// send request to handler and record response
			person.HandleEmail(res, req)
			if res.Code != testcase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testcase.expectedStatusCode, res.Code)
			}

			// if body is expected to have contents
			if testcase.expectBody {
				// parse response body
				var emailResponse person.EmailResponse
				err := json.Unmarshal(res.Body.Bytes(), &emailResponse)
				if err != nil {
					t.Error("error parsing request body:", err)
				}

				// parse email into username, domain name, and domain suffix
				_, domainName, domainSuffix, err := ParseEmail(emailResponse.Email)
				if err != nil {
					t.Error("error parsing email:", err)
				}

				// check if domain name is correct
				if !utils.StringSoftEqual(testcase.expectEmailDomainName, domainName) {
					t.Errorf("expected domain name %s, got %s", testcase.expectEmailDomainName, domainName)
				}
				// check if domain suffix is correct
				if !utils.StringSoftEqual(testcase.expectEmailDomainSuffix, domainSuffix) {
					t.Errorf("expected domain suffix %s, got %s", testcase.expectEmailDomainSuffix, domainSuffix)
				}
			}
		})
	}
}
