package router

import (
	"net/http"

	"github.com/alvii147/FunFaker/address"
	"github.com/alvii147/FunFaker/company"
	"github.com/alvii147/FunFaker/person"
	"github.com/alvii147/FunFaker/utils"
)

const (
	NAME_URL      = "/name"
	EMAIL_URL     = "/email"
	ADDRESS_URL   = "/address"
	COMPANY_URL   = "/company"
	CATCH_ALL_URL = "/"
)

// set up routing
func Routing() {
	// GET /name
	http.HandleFunc(NAME_URL, person.HandleName)
	// GET /email
	http.HandleFunc(EMAIL_URL, person.HandleEmail)
	// GET /address
	http.HandleFunc(ADDRESS_URL, address.HandleAddress)
	// GET /company
	http.HandleFunc(COMPANY_URL, company.HandleCompany)
	// invalid URL
	http.HandleFunc(CATCH_ALL_URL, utils.HandleNotFound)
}
