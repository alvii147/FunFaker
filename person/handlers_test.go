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

func TestHandlePerson(t *testing.T) {
	// set up test table
	testcases := []struct {
		name                      string
		url                       string
		method                    string
		expectedStatusCode        int
		expectBody                bool
		expectedSex               data.Sex
		expectedGroup             data.PersonGroup
		expectedEmailDomainName   string
		expectedEmailDomainSuffix string
	}{
		{
			name:                      "HandlePerson returns random person",
			url:                       "/person",
			method:                    http.MethodGet,
			expectedStatusCode:        http.StatusOK,
			expectBody:                true,
			expectedSex:               "",
			expectedGroup:             "",
			expectedEmailDomainName:   "",
			expectedEmailDomainSuffix: "",
		},
		{
			name:                      "HandlePerson returns 405 on POST request",
			url:                       "/person",
			method:                    http.MethodPost,
			expectedStatusCode:        http.StatusMethodNotAllowed,
			expectBody:                false,
			expectedSex:               "",
			expectedGroup:             "",
			expectedEmailDomainName:   "",
			expectedEmailDomainSuffix: "",
		},
		{
			name:                      "HandlePerson returns random Male person",
			url:                       "/person?sex=male",
			method:                    http.MethodGet,
			expectedStatusCode:        http.StatusOK,
			expectBody:                true,
			expectedSex:               data.SexMale,
			expectedGroup:             "",
			expectedEmailDomainName:   "",
			expectedEmailDomainSuffix: "",
		},
		{
			name:                      "HandlePerson returns random Female person",
			url:                       "/person?sex=female",
			method:                    http.MethodGet,
			expectedStatusCode:        http.StatusOK,
			expectBody:                true,
			expectedSex:               data.SexFemale,
			expectedGroup:             "",
			expectedEmailDomainName:   "",
			expectedEmailDomainSuffix: "",
		},
		{
			name:                      "HandlePerson returns random person of Comics group",
			url:                       "/person?group=comics",
			method:                    http.MethodGet,
			expectedStatusCode:        http.StatusOK,
			expectBody:                true,
			expectedSex:               "",
			expectedGroup:             data.PersonGroupComics,
			expectedEmailDomainName:   "",
			expectedEmailDomainSuffix: "",
		},
		{
			name:                      "HandlePerson returns random Male person of Comics group",
			url:                       "/person?sex=male&group=comics",
			method:                    http.MethodGet,
			expectedStatusCode:        http.StatusOK,
			expectBody:                true,
			expectedSex:               data.SexMale,
			expectedGroup:             data.PersonGroupComics,
			expectedEmailDomainName:   "",
			expectedEmailDomainSuffix: "",
		},
		{
			name:                      "HandlePerson returns random email of gmail domain name and org domain suffix",
			url:                       "/person?domain-name=gmail&domain-suffix=org",
			method:                    http.MethodGet,
			expectedStatusCode:        http.StatusOK,
			expectBody:                true,
			expectedSex:               "",
			expectedGroup:             "",
			expectedEmailDomainName:   "gmail",
			expectedEmailDomainSuffix: "org",
		},
		{
			name:                      "HandlePerson returns 400 on invalid URL parameters",
			url:                       "/person?invalid=parameter",
			method:                    http.MethodGet,
			expectedStatusCode:        http.StatusBadRequest,
			expectBody:                false,
			expectedSex:               "",
			expectedGroup:             "",
			expectedEmailDomainName:   "",
			expectedEmailDomainSuffix: "",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// create request and response objects
			req := httptest.NewRequest(testcase.method, testcase.url, nil)
			res := httptest.NewRecorder()

			// send request to handler and record response
			person.HandlePerson(res, req)

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
				var personResponse person.PersonResponse
				err := json.Unmarshal(res.Body.Bytes(), &personResponse)
				if err != nil {
					t.Error("error parsing response body:", err)
				}

				// get list of persons
				persons, err := data.GetPersons()
				if err != nil {
					t.Error("error getting names:", err)
				}

				// filter list of persons to get the returned entry
				filteredPersons := data.FilterPersons(
					persons,
					personResponse.FirstName,
					personResponse.LastName,
					testcase.expectedSex,
					testcase.expectedGroup,
					"",
					personResponse.Trivia,
				)

				// throw error if there isn't exactly a single entry after filtering
				if len(filteredPersons) != 1 {
					t.Errorf("expected 1 name match, got %d", len(filteredPersons))
				}

				// parse email into username, domain name, and domain suffix
				_, domainName, domainSuffix, err := ParseEmail(personResponse.Email)
				if err != nil {
					t.Error("error parsing email:", err)
				}

				// check if domain name is correct
				if !utils.StringSoftEqual(testcase.expectedEmailDomainName, domainName) {
					t.Errorf("expected domain name %s, got %s", testcase.expectedEmailDomainName, domainName)
				}
				// check if domain suffix is correct
				if !utils.StringSoftEqual(testcase.expectedEmailDomainSuffix, domainSuffix) {
					t.Errorf("expected domain suffix %s, got %s", testcase.expectedEmailDomainSuffix, domainSuffix)
				}
			}
		})
	}
}
