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

// path to addresses.json
const ADDRESSES_FILE_NAME = "addresses.json"

// enum representing address group
type AddressGroup string

const (
	AddressGroupBatman           AddressGroup = "Batman"
	AddressGroupBreakingBad      AddressGroup = "Breaking-Bad"
	AddressGroupFamilyGuy        AddressGroup = "Family-Guy"
	AddressGroupSeinfield        AddressGroup = "Seinfeld"
	AddressGroupTheBigBangTheory AddressGroup = "The-Big-Bang-Theory"
	AddressGroupTheFlintstones   AddressGroup = "The-Flintstones"
)

// check address group enum is valid
func (group *AddressGroup) IsValid() bool {
	return utils.StringSoftEqual(string(*group), string(AddressGroupBatman)) ||
		utils.StringSoftEqual(string(*group), string(AddressGroupBreakingBad)) ||
		utils.StringSoftEqual(string(*group), string(AddressGroupFamilyGuy)) ||
		utils.StringSoftEqual(string(*group), string(AddressGroupSeinfield)) ||
		utils.StringSoftEqual(string(*group), string(AddressGroupTheBigBangTheory)) ||
		utils.StringSoftEqual(string(*group), string(AddressGroupTheFlintstones))
}

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

	// read addresses.json
	addressesBytes, err := os.ReadFile(addressesFilePath)
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
	err = os.WriteFile(addressesFilePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
