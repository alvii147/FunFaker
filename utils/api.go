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
	response := HTTPErrorResponse{
		Detail: err.Error(),
	}
	json.NewEncoder(w).Encode(response)
}
