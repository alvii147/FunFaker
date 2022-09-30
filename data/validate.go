package data

import (
	"errors"
	"reflect"
	"sort"

	"github.com/alvii147/FunFaker/utils"
)

// check if name1 is alphabetically lower than name2
func IsNameLess(name1 Name, name2 Name) bool {
	if name1.Group != name2.Group {
		return utils.IsStringAlphabeticallyLess(
			string(name1.Group),
			string(name2.Group),
		)
	}

	if name1.FirstName != name2.FirstName {
		return utils.IsStringAlphabeticallyLess(
			name1.FirstName,
			name2.FirstName,
		)
	}

	if name1.LastName != name2.LastName {
		return utils.IsStringAlphabeticallyLess(
			name1.LastName,
			name2.LastName,
		)
	}

	if name1.Sex != name2.Sex {
		return utils.IsStringAlphabeticallyLess(
			string(name1.Sex),
			string(name2.Sex),
		)
	}

	if name1.Trivia != name2.Trivia {
		return utils.IsStringAlphabeticallyLess(
			name1.Trivia,
			name2.Trivia,
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

// check if names are sorted alphabetically
func CheckNamesSorted(unsortedNames []Name, autoFix bool) bool {
	sortedNames := make([]Name, len(unsortedNames))
	copy(sortedNames, unsortedNames)

	// sort names
	sort.Slice(sortedNames, func(i int, j int) bool {
		return IsNameLess(sortedNames[i], sortedNames[j])
	})

	// check if sorted and unsorted slices are equal
	sorted := reflect.DeepEqual(sortedNames, unsortedNames)

	if !sorted && autoFix {
		// if not sorted, apply fix
		WriteNames(sortedNames)

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

// validate data in names.json
func ValidateNamesData(autofix bool) error {
	names, err := GetNames()
	if err != nil {
		return err
	}

	for _, name := range names {
		// check if person sex enum is valid
		if !name.Sex.IsValid() {
			return errors.New(
				"invalid sex " +
					string(name.Sex) +
					" for name " +
					name.FirstName +
					" " +
					name.LastName,
			)
		}

		// check if person group enum is valid
		if !name.Group.IsValid() {
			return errors.New(
				"invalid person group " +
					string(name.Group) +
					" for name " +
					name.FirstName +
					" " +
					name.LastName,
			)
		}
	}

	if !CheckNamesSorted(names, autofix) {
		return errors.New("names not sorted properly")
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
	err := ValidateNamesData(autofix)
	if err != nil {
		return err
	}

	err = ValidateAddressesData(autofix)
	if err != nil {
		return err
	}

	return nil
}
