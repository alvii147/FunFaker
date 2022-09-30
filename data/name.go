package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// path to names.json file
const NAMES_FILE_NAME = "names.json"

// enum representing person sex
type Sex string

const (
	SexMale   Sex = "Male"
	SexFemale Sex = "Female"
	SexOther  Sex = "Other"
)

// enum representing person group
type PersonGroup string

const (
	PersonGroupComics PersonGroup = "Comics"
)

// struct representing name from names.json
type Name struct {
	FirstName string      `json:"first-name"`
	LastName  string      `json:"last-name"`
	Sex       Sex         `json:"sex"`
	Group     PersonGroup `json:"group"`
	Trivia    string      `json:"trivia"`
}

// read names from names.json
func GetNames() ([]Name, error) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("unable to get current directory")
	}

	// open names.json
	namesFilePath := filepath.Join(path.Dir(filename), NAMES_FILE_NAME)
	namesFile, err := os.Open(namesFilePath)
	if err != nil {
		return nil, err
	}

	// read names.json
	namesBytes, err := ioutil.ReadAll(namesFile)
	if err != nil {
		return nil, err
	}

	// get names from bytes read
	var names []Name
	err = json.Unmarshal(namesBytes, &names)
	if err != nil {
		return nil, err
	}

	return names, nil
}

// write list of names to names.json
func WriteNames(names []Name) error {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get current directory")
	}

	namesFilePath := filepath.Join(path.Dir(filename), NAMES_FILE_NAME)

	// convert to bytes with indentation
	file, err := json.MarshalIndent(names, "", "    ")
	if err != nil {
		return err
	}

	// write to names.json
	err = ioutil.WriteFile(namesFilePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
