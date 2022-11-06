package date

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/utils"
	"github.com/gorilla/schema"
)

// GET date handler
func HandleDate(w http.ResponseWriter, r *http.Request) {
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
		var dateRequest DateRequest
		decoder := schema.NewDecoder()
		decoder.RegisterConverter(time.Time{}, func(value string) reflect.Value {
			t, err := time.Parse("2006-01-02", value)
			if err != nil {
				return reflect.Value{}
			}

			return reflect.ValueOf(t)
		})
		err := decoder.Decode(&dateRequest, r.URL.Query())
		if err != nil {
			statusCode = http.StatusBadRequest
			utils.HTTPError(statusCode, err, w)

			return
		}

		// parse dates
		// var startDate *time.Time = nil
		// var endDate *time.Time = nil
		// if dateRequest.StartDate != "" {
		// 	t, err := time.Parse("2006-01-02", dateRequest.StartDate)
		// 	if err != nil {
		// 		statusCode = http.StatusBadRequest
		// 		utils.HTTPError(statusCode, err, w)

		// 		return
		// 	}

		// 	startDate = &t
		// }
		// if dateRequest.EndDate != "" {
		// 	t, err := time.Parse("2006-01-02", dateRequest.EndDate)
		// 	if err != nil {
		// 		statusCode = http.StatusBadRequest
		// 		utils.HTTPError(statusCode, err, w)

		// 		return
		// 	}

		// 	endDate = &t
		// }

		// get list of dates
		dates, err := data.GetDates()
		if err != nil {
			statusCode = http.StatusInternalServerError
			utils.HTTPError(statusCode, err, w)

			return
		}

		// filter list of dates by decoded name request
		filteredDates := data.FilterDates(
			dates,
			dateRequest.After,
			dateRequest.Before,
			dateRequest.Group,
			"",
		)

		// if filtering returned no results, respond with "no content"
		if len(filteredDates) < 1 {
			statusCode = http.StatusNoContent
			w.WriteHeader(statusCode)
			return
		}

		// choose random date
		randomDate := filteredDates[rand.Intn(len(filteredDates))]
		// update date response using random date
		var dateResponse DateResponse
		dateResponse.FromDate(randomDate)

		// encode response
		err = json.NewEncoder(w).Encode(dateResponse)
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
