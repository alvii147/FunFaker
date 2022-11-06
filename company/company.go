package company

import (
	"github.com/alvii147/FunFaker/data"
)

// GET company request URL parameters
type CompanyRequest struct {
	Group data.CompanyGroup `schema:"group"`
}

// GET company request response body
type CompanyResponse struct {
	Name   string `json:"name"`
	Trivia string `json:"trivia"`
}

// update company response using company
func (companyResponse *CompanyResponse) FromCompany(company data.Company) {
	companyResponse.Name = company.Name
	companyResponse.Trivia = company.Trivia
}
