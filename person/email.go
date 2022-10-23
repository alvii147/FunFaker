package person

import (
	"strings"

	"github.com/alvii147/FunFaker/data"
)

// GET email request URL parameters
type EmailRequest struct {
	Sex          data.Sex         `schema:"sex"`
	Group        data.PersonGroup `schema:"group"`
	DomainName   string           `schema:"domain-name"`
	DomainSuffix string           `schema:"domain-suffix"`
}

// GET email request response body
type EmailResponse struct {
	Email  string `json:"email"`
	Trivia string `json:"trivia"`
}

// update email response using person and email request
func (emailResponse *EmailResponse) FromPersonAndEmailRequest(person data.Person, emailRequest EmailRequest) {
	// if domain name not specified, use person domain
	domainName := person.Domain
	if emailRequest.DomainName != "" {
		domainName = emailRequest.DomainName
	}

	// if domain suffix not specified, use .com
	domainSuffix := "com"
	if emailRequest.DomainSuffix != "" {
		domainSuffix = emailRequest.DomainSuffix
	}

	emailResponse.Email = strings.ToLower(person.FirstName) +
		"." +
		strings.ToLower(person.LastName) +
		"@" +
		domainName +
		"." +
		domainSuffix
	emailResponse.Trivia = person.Trivia
}
