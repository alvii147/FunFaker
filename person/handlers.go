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
	var err error
	defer func() {
		utils.LogHTTPTraffic(r, statusCode, err)
	}()

	switch r.Method {
	case "GET":
		statusCode = http.StatusOK
		// set random seed
		rand.Seed(time.Now().Unix())

		// decode incoming request URL parameters
		var nameRequest NameRequest
		decoder := schema.NewDecoder()
		err = decoder.Decode(&nameRequest, r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			utils.HTTPError(statusCode, err, w)

			return
		}

		// get list of names
		names, err := data.GetNames()
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}

		// filter list of names by decoded name request
		filteredNames := FilterNames(
			names,
			"",
			"",
			nameRequest.Sex,
			nameRequest.Group,
			"",
		)
		// if filtering returned no results, respond with "no content"
		if len(filteredNames) < 1 {
			statusCode = http.StatusNoContent
			w.WriteHeader(statusCode)
			return
		}

		// choose random name
		randomName := filteredNames[rand.Intn(len(filteredNames))]
		// update name response using random name
		var nameResponse NameResponse
		nameResponse.FromName(randomName)

		// encode response
		err = json.NewEncoder(w).Encode(nameResponse)
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}
	default:
		statusCode = http.StatusMethodNotAllowed
		w.WriteHeader(statusCode)
	}
}

func HandleEmail(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var err error
	defer func() {
		utils.LogHTTPTraffic(r, statusCode, err)
	}()

	switch r.Method {
	case "GET":
		statusCode = http.StatusOK
		// set random seed
		rand.Seed(time.Now().Unix())

		// decode incoming request URL parameters
		var emailRequest EmailRequest
		decoder := schema.NewDecoder()
		err = decoder.Decode(&emailRequest, r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			utils.HTTPError(statusCode, err, w)

			return
		}

		// get list of names
		names, err := data.GetNames()
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}

		// filter list of names by decoded email request
		filteredNames := FilterNames(
			names,
			"",
			"",
			emailRequest.Sex,
			emailRequest.Group,
			"",
		)
		// if filtering returned no results, respond with "no content"
		if len(filteredNames) < 1 {
			statusCode = http.StatusNoContent
			w.WriteHeader(statusCode)
			return
		}

		// choose random name
		randomName := filteredNames[rand.Intn(len(filteredNames))]
		var email EmailResponse
		// update email response using random name and email request
		email.FromNameAndEmailRequest(randomName, emailRequest)

		// encode response
		err = json.NewEncoder(w).Encode(email)
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}
	default:
		statusCode = http.StatusMethodNotAllowed
		w.WriteHeader(statusCode)
	}
}
