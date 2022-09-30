package router

import (
	"net/http"

	"github.com/alvii147/FunFaker/address"
	"github.com/alvii147/FunFaker/person"
)

const (
	NAME_URL    = "/name"
	EMAIL_URL   = "/email"
	ADDRESS_URL = "/address"
)

// set up routing
func Routing() {
	// GET /name
	http.HandleFunc(NAME_URL, person.HandleName)
	// GET /email
	http.HandleFunc(EMAIL_URL, person.HandleEmail)
	// GET /address
	http.HandleFunc(ADDRESS_URL, address.HandleAddress)
}
