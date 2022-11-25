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

// GET person handler
func HandlePerson(w http.ResponseWriter, r *http.Request) {
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
		var personRequest PersonRequest
		decoder := schema.NewDecoder()
		err := decoder.Decode(&personRequest, r.URL.Query())
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
			personRequest.Sex,
			personRequest.Group,
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
		var personResponse PersonResponse
		personResponse.FromPerson(randomPerson, personRequest)

		// encode response
		err = json.NewEncoder(w).Encode(personResponse)
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
