package utils

import (
	"encoding/json"
	"net/http"
)

// default HTTP response for errors
type HTTPErrorResponse struct {
	Detail string `json:"detail"`
}

// respond with HTTP error message
func HTTPError(statusCode int, err error, w http.ResponseWriter) {
	LogError(err)
	w.WriteHeader(statusCode)
	if err != nil {
		response := HTTPErrorResponse{
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
	}
}

func SetHeader(w http.ResponseWriter, header string, value string) {
	w.Header().Set(header, value)
}

// set CORS header
func SetCORSHeader(w http.ResponseWriter, url string) {
	SetHeader(w, "Access-Control-Allow-Origin", url)
}

// 404 not found handler
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	defer func() {
		LogHTTPTraffic(r, statusCode)
	}()

	// enable cross-origin resource sharing
	SetCORSHeader(w, "*")

	statusCode = http.StatusNotFound
	w.WriteHeader(statusCode)
}
