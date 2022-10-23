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

// path to persons.json file
const PERSONS_FILE_NAME = "persons.json"

// enum representing person sex
type Sex string

const (
	SexMale   Sex = "Male"
	SexFemale Sex = "Female"
	SexOther  Sex = "Other"
)

// check if sex enum is valid
func (sex *Sex) IsValid() bool {
	return utils.StringSoftEqual(string(*sex), string(SexMale)) ||
		utils.StringSoftEqual(string(*sex), string(SexFemale)) ||
		utils.StringSoftEqual(string(*sex), string(SexOther))
}

// enum representing person group
type PersonGroup string

const (
	PersonGroupComics  PersonGroup = "Comics"
	PersonGroupMovies  PersonGroup = "Movies"
	PersonGroupTVShows PersonGroup = "TV-Shows"
)

// check person group enum is valid
func (group *PersonGroup) IsValid() bool {
	return utils.StringSoftEqual(string(*group), string(PersonGroupComics)) ||
		utils.StringSoftEqual(string(*group), string(PersonGroupMovies)) ||
		utils.StringSoftEqual(string(*group), string(PersonGroupTVShows))
}

// struct representing person from persons.json
type Person struct {
	FirstName string      `json:"first-name"`
	LastName  string      `json:"last-name"`
	Sex       Sex         `json:"sex"`
	Group     PersonGroup `json:"group"`
	Domain    string      `json:"domain"`
	Trivia    string      `json:"trivia"`
}

// read persons from persons.json
func GetPersons() ([]Person, error) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("unable to get current directory")
	}

	// open persons.json
	personsFilePath := filepath.Join(path.Dir(filename), PERSONS_FILE_NAME)

	// read persons.json
	personsBytes, err := os.ReadFile(personsFilePath)
	if err != nil {
		return nil, err
	}

	// get persons from bytes read
	var person []Person
	err = json.Unmarshal(personsBytes, &person)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// write list of persons to persons.json
func WritePersons(persons []Person) error {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get current directory")
	}

	// open persons.json
	personsFilePath := filepath.Join(path.Dir(filename), PERSONS_FILE_NAME)

	// convert to bytes with indentation
	personsBytes, err := json.MarshalIndent(persons, "", "    ")
	if err != nil {
		return err
	}

	// write to persons.json
	err = os.WriteFile(personsFilePath, personsBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// filter persons by properties
func FilterPersons(
	persons []Person,
	firstName string,
	lastName string,
	sex Sex,
	group PersonGroup,
	domain string,
	trivia string,
) []Person {
	filteredPersons := []Person{}
	for _, person := range persons {
		if !utils.StringSoftEqual(firstName, person.FirstName) {
			continue
		}

		if !utils.StringSoftEqual(lastName, person.LastName) {
			continue
		}

		if !utils.StringSoftEqual(string(sex), string(person.Sex)) {
			continue
		}

		if !utils.StringSoftEqual(string(group), string(person.Group)) {
			continue
		}

		if !utils.StringSoftEqual(domain, person.Domain) {
			continue
		}

		if !utils.StringSoftEqual(trivia, person.Trivia) {
			continue
		}

		filteredPersons = append(filteredPersons, person)
	}

	return filteredPersons
}
