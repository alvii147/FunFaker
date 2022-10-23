package company

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/utils"
	"github.com/gorilla/schema"
)

// GET company handler
func HandleCompany(w http.ResponseWriter, r *http.Request) {
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
		var companyRequest CompanyRequest
		decoder := schema.NewDecoder()
		err := decoder.Decode(&companyRequest, r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			utils.HTTPError(statusCode, err, w)

			return
		}

		// get list of companies
		companies, err := data.GetCompanies()
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}

		// filter list of companies by decoded name request
		filteredCompanies := data.FilterCompanies(
			companies,
			"",
			companyRequest.Group,
			"",
		)

		// if filtering returned no results, respond with "no content"
		if len(filteredCompanies) < 1 {
			statusCode = http.StatusNoContent
			w.WriteHeader(statusCode)
			return
		}

		// choose random company
		randomCompany := filteredCompanies[rand.Intn(len(filteredCompanies))]
		// update company response using random company
		var companyResponse CompanyResponse
		companyResponse.FromCompany(randomCompany)

		// encode response
		err = json.NewEncoder(w).Encode(companyResponse)
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
