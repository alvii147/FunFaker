package address

import (
	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/utils"
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

// filter addresses by properties
func FilterAddresses(
	addresses []data.Address,
	streetName string,
	city string,
	state string,
	country string,
	postalCode string,
	group data.AddressGroup,
	trivia string,
) []data.Address {
	filteredAddresses := []data.Address{}
	for _, address := range addresses {
		if !utils.StringSoftEqual(streetName, address.StreetName) {
			continue
		}

		if !utils.StringSoftEqual(city, address.City) {
			continue
		}

		if !utils.StringSoftEqual(state, address.State) {
			continue
		}

		if !utils.StringSoftEqual(country, address.Country) {
			continue
		}

		if !utils.StringSoftEqual(postalCode, address.PostalCode) {
			continue
		}

		if !utils.StringSoftEqual(string(group), string(address.Group)) {
			continue
		}

		if !utils.StringSoftEqual(trivia, address.Trivia) {
			continue
		}

		filteredAddresses = append(filteredAddresses, address)
	}

	return filteredAddresses
}

// filter addresses by validity
func FilterValidAddresses(addresses []data.Address) []data.Address {
	validAddresses := []data.Address{}
	for _, address := range addresses {
		if !address.Valid {
			continue
		}

		validAddresses = append(validAddresses, address)
	}

	return validAddresses
}

func (addressResponse *AddressResponse) FromAddress(address data.Address) {
	addressResponse.StreetName = address.StreetName
	addressResponse.City = address.City
	addressResponse.State = address.State
	addressResponse.Country = address.Country
	addressResponse.PostalCode = address.PostalCode
	addressResponse.Valid = address.Valid
	addressResponse.Trivia = address.Trivia
}
