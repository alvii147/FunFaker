package website

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/utils"
	"github.com/gorilla/schema"
)

// GET website handler
func HandleWebsite(w http.ResponseWriter, r *http.Request) {
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
		var websiteRequest WebsiteRequest
		decoder := schema.NewDecoder()
		err := decoder.Decode(&websiteRequest, r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			utils.HTTPError(statusCode, err, w)

			return
		}

		// get list of websites
		websites, err := data.GetWebsites()
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}

		// filter list of websites by decoded name request
		filteredWebsites := data.FilterWebsites(
			websites,
			"",
			websiteRequest.Group,
			"",
		)

		// if filtering returned no results, respond with "no content"
		if len(filteredWebsites) < 1 {
			statusCode = http.StatusNoContent
			w.WriteHeader(statusCode)
			return
		}

		// choose random website
		randomWebsite := filteredWebsites[rand.Intn(len(filteredWebsites))]
		// update website response using random website
		var websiteResponse WebsiteResponse
		websiteResponse.FromWebsite(randomWebsite)

		// encode response
		err = json.NewEncoder(w).Encode(websiteResponse)
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
