package person

import (
	"github.com/alvii147/FunFaker/data"
)

// GET person request URL parameters
type PersonRequest struct {
	Sex          data.Sex         `schema:"sex"`
	Group        data.PersonGroup `schema:"group"`
	DomainName   string           `schema:"domain-name"`
	DomainSuffix string           `schema:"domain-suffix"`
}

// GET person request response body
type PersonResponse struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Trivia    string `json:"trivia"`
}

// update person response using person
func (personResponse *PersonResponse) FromPerson(person data.Person, personRequest PersonRequest) {
	// if domain name not specified, use person domain
	domainName := person.Domain
	if personRequest.DomainName != "" {
		domainName = personRequest.DomainName
	}

	// if domain suffix not specified, use .com
	domainSuffix := data.EMAIL_DEFAULT_DOMAIN_SUFFIX
	if personRequest.DomainSuffix != "" {
		domainSuffix = personRequest.DomainSuffix
	}

	personResponse.FirstName = person.FirstName
	personResponse.LastName = person.LastName
	personResponse.Email = data.GenerateEmail(person.FirstName, person.LastName, domainName, domainSuffix)
	personResponse.Trivia = person.Trivia
}
