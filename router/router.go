package router

import (
	"net/http"

	"github.com/alvii147/FunFaker/address"
	"github.com/alvii147/FunFaker/company"
	"github.com/alvii147/FunFaker/date"
	"github.com/alvii147/FunFaker/person"
	"github.com/alvii147/FunFaker/utils"
	"github.com/alvii147/FunFaker/website"
)

const (
	ADDRESS_URL   = "/address"
	COMPANY_URL   = "/company"
	DATE_URL      = "/date"
	PERSON_URL    = "/person"
	WEBSITE_URL   = "/website"
	CATCH_ALL_URL = "/"
)

// set up routing
func Routing() {
	// GET /address
	http.HandleFunc(ADDRESS_URL, address.HandleAddress)
	// GET /company
	http.HandleFunc(COMPANY_URL, company.HandleCompany)
	// GET /date
	http.HandleFunc(DATE_URL, date.HandleDate)
	// GET /person
	http.HandleFunc(PERSON_URL, person.HandlePerson)
	// GET /website
	http.HandleFunc(WEBSITE_URL, website.HandleWebsite)
	// invalid URL
	http.HandleFunc(CATCH_ALL_URL, utils.HandleNotFound)
}
