package address

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/utils"

	"github.com/gorilla/schema"
)

// GET address handler
func HandleAddress(w http.ResponseWriter, r *http.Request) {
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
		var addressRequest AddressRequest
		decoder := schema.NewDecoder()
		err = decoder.Decode(&addressRequest, r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			utils.HTTPError(statusCode, err, w)

			return
		}

		// get list of addresses
		addresses, err := data.GetAddresses()
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}

		// filter list of addresses by decoded name request
		filteredAddresses := FilterAddresses(
			addresses,
			"",
			"",
			"",
			"",
			"",
			addressRequest.Group,
			"",
		)
		// filter list of addresses by validity if true in request
		if addressRequest.ValidOnly {
			filteredAddresses = FilterValidAddresses(filteredAddresses)
		}
		// if filtering returned no results, respond with "no content"
		if len(filteredAddresses) < 1 {
			statusCode = http.StatusNoContent
			w.WriteHeader(statusCode)
			return
		}

		// choose random address
		randomAddress := filteredAddresses[rand.Intn(len(filteredAddresses))]
		// update address response using random address
		var addressResponse AddressResponse
		addressResponse.Valid = false
		addressResponse.FromAddress(randomAddress)

		// encode response
		err = json.NewEncoder(w).Encode(addressResponse)
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
