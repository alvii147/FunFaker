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

// update email response using name and email request
func (emailResponse *EmailResponse) FromNameAndEmailRequest(name data.Name, emailRequest EmailRequest) {
	// if domain name not specified, use name group
	domainName := strings.ToLower(string(name.Group))
	if emailRequest.DomainName != "" {
		domainName = emailRequest.DomainName
	}

	// if domain suffix not specified, use .com
	domainSuffix := "com"
	if emailRequest.DomainSuffix != "" {
		domainSuffix = emailRequest.DomainSuffix
	}

	emailResponse.Email = strings.ToLower(name.FirstName) +
		"." +
		strings.ToLower(name.LastName) +
		"@" +
		domainName +
		"." +
		domainSuffix
	emailResponse.Trivia = name.Trivia
}
