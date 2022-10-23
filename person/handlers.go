package person

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/utils"

	"github.com/gorilla/schema"
)

// GET name handler
func HandleName(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	defer func() {
		utils.LogHTTPTraffic(r, statusCode)
	}()

	// enable cross-origin resource sharing
	utils.SetCORSHeader(w, "*")

	switch r.Method {
	case "GET":
		statusCode = http.StatusOK
		// set random seed
		rand.Seed(time.Now().Unix())

		// decode incoming request URL parameters
		var nameRequest NameRequest
		decoder := schema.NewDecoder()
		err := decoder.Decode(&nameRequest, r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			utils.HTTPError(statusCode, err, w)

			return
		}

		// get list of persons
		persons, err := data.GetPersons()
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}

		// filter list of persons by decoded name request
		filteredPersons := data.FilterPersons(
			persons,
			"",
			"",
			nameRequest.Sex,
			nameRequest.Group,
			"",
			"",
		)
		// if filtering returned no results, respond with "no content"
		if len(filteredPersons) < 1 {
			statusCode = http.StatusNoContent
			w.WriteHeader(statusCode)
			return
		}

		// choose random person
		randomPerson := filteredPersons[rand.Intn(len(filteredPersons))]
		// update name response using random person
		var nameResponse NameResponse
		nameResponse.FromPerson(randomPerson)

		// encode response
		err = json.NewEncoder(w).Encode(nameResponse)
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}
	default:
		statusCode = http.StatusMethodNotAllowed
		utils.HTTPError(statusCode, nil, w)
	}
}

func HandleEmail(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	defer func() {
		utils.LogHTTPTraffic(r, statusCode)
	}()

	// enable cross-origin resource sharing
	utils.SetCORSHeader(w, "*")

	switch r.Method {
	case "GET":
		statusCode = http.StatusOK
		// set random seed
		rand.Seed(time.Now().Unix())

		// decode incoming request URL parameters
		var emailRequest EmailRequest
		decoder := schema.NewDecoder()
		err := decoder.Decode(&emailRequest, r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			utils.HTTPError(statusCode, err, w)

			return
		}

		// get list of names
		persons, err := data.GetPersons()
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}

		// filter list of persons by decoded email request
		filteredPersons := data.FilterPersons(
			persons,
			"",
			"",
			emailRequest.Sex,
			emailRequest.Group,
			"",
			"",
		)
		// if filtering returned no results, respond with "no content"
		if len(filteredPersons) < 1 {
			statusCode = http.StatusNoContent
			w.WriteHeader(statusCode)
			return
		}

		// choose random person
		randomPerson := filteredPersons[rand.Intn(len(filteredPersons))]
		var email EmailResponse
		// update email response using random person and email request
		email.FromPersonAndEmailRequest(randomPerson, emailRequest)

		// encode response
		err = json.NewEncoder(w).Encode(email)
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}
	default:
		statusCode = http.StatusMethodNotAllowed
		utils.HTTPError(statusCode, nil, w)
	}
}
