package address

import (
	"github.com/alvii147/FunFaker/data"
)

// GET address request URL parameters
type AddressRequest struct {
	Group     data.AddressGroup `schema:"group"`
	ValidOnly bool              `schema:"valid-only"`
}

// GET address request response body
type AddressResponse struct {
	StreetName string `json:"street-name"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal-code"`
	Valid      bool   `json:"valid"`
	Trivia     string `json:"trivia"`
}

// update address response using address
func (addressResponse *AddressResponse) FromAddress(address data.Address) {
	addressResponse.StreetName = address.StreetName
	addressResponse.City = address.City
	addressResponse.State = address.State
	addressResponse.Country = address.Country
	addressResponse.PostalCode = address.PostalCode
	addressResponse.Valid = address.Valid
	addressResponse.Trivia = address.Trivia
}
