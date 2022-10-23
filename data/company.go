package data

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/alvii147/FunFaker/utils"
)

// path to companies.json
const COMPANIES_FILE_NAME = "companies.json"

// enum representing companies group
type CompanyGroup string

const (
	CompanyGroupCartoons CompanyGroup = "Cartoons"
	CompanyGroupComics   CompanyGroup = "Comics"
	CompanyGroupMovies   CompanyGroup = "Movies"
	CompanyGroupTVShows  CompanyGroup = "TV-Shows"
)

// check company group enum is valid
func (group *CompanyGroup) IsValid() bool {
	return utils.StringSoftEqual(string(*group), string(CompanyGroupCartoons)) ||
		utils.StringSoftEqual(string(*group), string(CompanyGroupComics)) ||
		utils.StringSoftEqual(string(*group), string(CompanyGroupMovies)) ||
		utils.StringSoftEqual(string(*group), string(CompanyGroupTVShows))
}

// struct representing company from companies.json
type Company struct {
	Name   string       `json:"name"`
	Group  CompanyGroup `json:"group"`
	Trivia string       `json:"trivia"`
}

// read companies from companies.json
func GetCompanies() ([]Company, error) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("unable to get current directory")
	}

	// open companies.json
	companiesFilePath := filepath.Join(path.Dir(filename), COMPANIES_FILE_NAME)

	// read companies.json
	companiesBytes, err := os.ReadFile(companiesFilePath)
	if err != nil {
		return nil, err
	}

	// get companies from bytes read
	var companies []Company
	err = json.Unmarshal(companiesBytes, &companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

// write list of companies to companies.json
func WriteCompanies(companies []Company) error {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get current directory")
	}

	// open companies.json
	companiesFilePath := filepath.Join(path.Dir(filename), COMPANIES_FILE_NAME)

	// convert to bytes with indentation
	file, err := json.MarshalIndent(companies, "", "    ")
	if err != nil {
		return err
	}

	// write to companies.json
	err = os.WriteFile(companiesFilePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// filter companies by properties
func FilterCompanies(
	companies []Company,
	name string,
	group CompanyGroup,
	trivia string,
) []Company {
	filteredCompanies := []Company{}
	for _, company := range companies {
		if !utils.StringSoftEqual(name, company.Name) {
			continue
		}

		if !utils.StringSoftEqual(string(group), string(company.Group)) {
			continue
		}

		if !utils.StringSoftEqual(trivia, company.Trivia) {
			continue
		}

		filteredCompanies = append(filteredCompanies, company)
	}

	return filteredCompanies
}
