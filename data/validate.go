package data

import (
	"errors"
	"reflect"
	"sort"

	"github.com/alvii147/FunFaker/utils"
)

// check if person1 is alphabetically lower than person2
func IsPersonLess(person1 Person, person2 Person) bool {
	if person1.Group != person2.Group {
		return utils.IsStringAlphabeticallyLess(
			string(person1.Group),
			string(person2.Group),
		)
	}

	if person1.FirstName != person2.FirstName {
		return utils.IsStringAlphabeticallyLess(
			person1.FirstName,
			person2.FirstName,
		)
	}

	if person1.LastName != person2.LastName {
		return utils.IsStringAlphabeticallyLess(
			person1.LastName,
			person2.LastName,
		)
	}

	if person1.Sex != person2.Sex {
		return utils.IsStringAlphabeticallyLess(
			string(person1.Sex),
			string(person2.Sex),
		)
	}

	if person1.Domain != person2.Domain {
		return utils.IsStringAlphabeticallyLess(
			person1.Domain,
			person2.Domain,
		)
	}

	if person1.Trivia != person2.Trivia {
		return utils.IsStringAlphabeticallyLess(
			person1.Trivia,
			person2.Trivia,
		)
	}

	return true
}

// check if address1 is alphabetically lower than address2
func IsAddressLess(address1 Address, address2 Address) bool {
	if address1.Group != address2.Group {
		return utils.IsStringAlphabeticallyLess(
			string(address1.Group),
			string(address2.Group),
		)
	}

	if address1.Country != address2.Country {
		return utils.IsStringAlphabeticallyLess(
			address1.Country,
			address2.Country,
		)
	}

	if address1.State != address2.State {
		return utils.IsStringAlphabeticallyLess(
			address1.State,
			address2.State,
		)
	}

	if address1.City != address2.City {
		return utils.IsStringAlphabeticallyLess(
			address1.City,
			address2.City,
		)
	}

	if address1.PostalCode != address2.PostalCode {
		return utils.IsStringAlphabeticallyLess(
			address1.PostalCode,
			address2.PostalCode,
		)
	}

	if address1.StreetName != address2.StreetName {
		return utils.IsStringAlphabeticallyLess(
			address1.StreetName,
			address2.StreetName,
		)
	}

	if address1.Valid != address2.Valid {
		return address1.Valid
	}

	if address1.Trivia != address2.Trivia {
		return utils.IsStringAlphabeticallyLess(
			address1.Trivia,
			address2.Trivia,
		)
	}

	return true
}

// check if persons are sorted alphabetically
func CheckPersonsSorted(unsortedPersons []Person, autoFix bool) bool {
	sortedPersons := make([]Person, len(unsortedPersons))
	copy(sortedPersons, unsortedPersons)

	// sort persons
	sort.Slice(sortedPersons, func(i int, j int) bool {
		return IsPersonLess(sortedPersons[i], sortedPersons[j])
	})

	// check if sorted and unsorted slices are equal
	sorted := reflect.DeepEqual(sortedPersons, unsortedPersons)

	if !sorted && autoFix {
		// if not sorted, apply fix
		WritePersons(sortedPersons)

		return true
	}

	return sorted
}

// check if addresses are sorted alphabetically
func CheckAddressesSorted(unsortedAddresses []Address, autoFix bool) bool {
	sortedAddresses := make([]Address, len(unsortedAddresses))
	copy(sortedAddresses, unsortedAddresses)

	// sort addresses
	sort.Slice(sortedAddresses, func(i int, j int) bool {
		return IsAddressLess(sortedAddresses[i], sortedAddresses[j])
	})

	// check if sorted and unsorted slices are equal
	sorted := reflect.DeepEqual(sortedAddresses, unsortedAddresses)

	if !sorted && autoFix {
		// if not sorted, apply fix
		WriteAddresses(sortedAddresses)

		return true
	}

	return sorted
}

// validate data in persons.json
func ValidatePersonsData(autofix bool) error {
	persons, err := GetPersons()
	if err != nil {
		return err
	}

	for _, person := range persons {
		// check if person sex enum is valid
		if !person.Sex.IsValid() {
			return errors.New(
				"invalid sex " +
					string(person.Sex) +
					" for name " +
					person.FirstName +
					" " +
					person.LastName,
			)
		}

		// check if person group enum is valid
		if !person.Group.IsValid() {
			return errors.New(
				"invalid person group " +
					string(person.Group) +
					" for name " +
					person.FirstName +
					" " +
					person.LastName,
			)
		}
	}

	if !CheckPersonsSorted(persons, autofix) {
		return errors.New("persons not sorted properly")
	}

	return nil
}

// validate data in addresses.json
func ValidateAddressesData(autofix bool) error {
	addresses, err := GetAddresses()
	if err != nil {
		return err
	}

	for _, address := range addresses {
		// check if address group enum is valid
		if !address.Group.IsValid() {
			return errors.New(
				"invalid person group " +
					string(address.Group) +
					" for name " +
					address.StreetName +
					", " +
					address.City +
					", " +
					address.State +
					", " +
					address.Country +
					", " +
					address.PostalCode,
			)
		}
	}

	if !CheckAddressesSorted(addresses, autofix) {
		return errors.New("addresses not sorted properly")
	}

	return nil
}

// validate all data files
func Validate(autofix bool) error {
	err := ValidatePersonsData(autofix)
	if err != nil {
		return err
	}

	err = ValidateAddressesData(autofix)
	if err != nil {
		return err
	}

	return nil
}
