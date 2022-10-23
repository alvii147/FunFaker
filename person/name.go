package person

import (
	"github.com/alvii147/FunFaker/data"
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

// update name response using person
func (nameResponse *NameResponse) FromPerson(name data.Person) {
	nameResponse.FirstName = name.FirstName
	nameResponse.LastName = name.LastName
	nameResponse.Trivia = name.Trivia
}
