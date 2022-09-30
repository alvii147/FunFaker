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

// path to addresses.json
const ADDRESSES_FILE_NAME = "addresses.json"

// enum representing address group
type AddressGroup string

const (
	AddressGroupBatman         AddressGroup = "Batman"
	AddressGroupBreakingBad    AddressGroup = "Breaking-Bad"
	AddressGroupFamilyGuy      AddressGroup = "Family-Guy"
	AddressGroupSeinfield      AddressGroup = "Seinfeld"
	AddressGroupTheFlintstones AddressGroup = "The-Flintstones"
)

// struct representing address from addresses.json
type Address struct {
	StreetName string       `json:"street-name"`
	City       string       `json:"city"`
	State      string       `json:"state"`
	Country    string       `json:"country"`
	PostalCode string       `json:"postal-code"`
	Group      AddressGroup `json:"group"`
	Valid      bool         `json:"valid"`
	Trivia     string       `json:"trivia"`
}

// read addresses from addresses.json
func GetAddresses() ([]Address, error) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("unable to get current directory")
	}

	// open addresses.json
	addressesFilePath := filepath.Join(path.Dir(filename), ADDRESSES_FILE_NAME)
	addressesFile, err := os.Open(addressesFilePath)
	if err != nil {
		return nil, err
	}

	// read addresses.json
	addressesBytes, err := ioutil.ReadAll(addressesFile)
	if err != nil {
		return nil, err
	}

	// get addresses from bytes read
	var addresses []Address
	err = json.Unmarshal(addressesBytes, &addresses)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

// write list of addresses to addresses.json
func WriteAddresses(addresses []Address) error {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get current directory")
	}

	addressesFilePath := filepath.Join(path.Dir(filename), ADDRESSES_FILE_NAME)

	// convert to bytes with indentation
	file, err := json.MarshalIndent(addresses, "", "    ")
	if err != nil {
		return err
	}

	// write to addresses.json
	err = ioutil.WriteFile(addressesFilePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
