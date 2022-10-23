package person

import (
	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/utils"
)

// GET name request URL parameters
type NameRequest struct {
	Sex   data.Sex         `schema:"sex"`
	Group data.PersonGroup `schema:"group"`
}

// GET name request response body
type NameResponse struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Trivia    string `json:"trivia"`
}

// filter names by properties
func FilterNames(
	names []data.Name,
	firstName string,
	lastName string,
	sex data.Sex,
	group data.PersonGroup,
	domain string,
	trivia string,
) []data.Name {
	filteredNames := []data.Name{}
	for _, name := range names {
		if !utils.StringSoftEqual(firstName, name.FirstName) {
			continue
		}

		if !utils.StringSoftEqual(lastName, name.LastName) {
			continue
		}

		if !utils.StringSoftEqual(string(sex), string(name.Sex)) {
			continue
		}

		if !utils.StringSoftEqual(string(group), string(name.Group)) {
			continue
		}

		if !utils.StringSoftEqual(domain, name.Domain) {
			continue
		}

		if !utils.StringSoftEqual(trivia, name.Trivia) {
			continue
		}

		filteredNames = append(filteredNames, name)
	}

	return filteredNames
}

// update name response using name
func (nameResponse *NameResponse) FromName(name data.Name) {
	nameResponse.FirstName = name.FirstName
	nameResponse.LastName = name.LastName
	nameResponse.Trivia = name.Trivia
}
