package data

import (
	"errors"
	"net/mail"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"time"

	"github.com/alvii147/FunFaker/utils"
)

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

// check if company1 is alphabetically lower than company2
func IsCompanyLess(company1 Company, company2 Company) bool {
	if company1.Group != company2.Group {
		return utils.IsStringAlphabeticallyLess(
			string(company1.Group),
			string(company2.Group),
		)
	}

	if company1.Name != company2.Name {
		return utils.IsStringAlphabeticallyLess(
			company1.Name,
			company2.Name,
		)
	}

	if company1.Trivia != company2.Trivia {
		return utils.IsStringAlphabeticallyLess(
			company1.Trivia,
			company2.Trivia,
		)
	}

	return true
}

// check if date1 is alphabetically lower than date2
func IsDateLess(date1 Date, date2 Date) bool {
	if date1.Group != date2.Group {
		return utils.IsStringAlphabeticallyLess(
			string(date1.Group),
			string(date2.Group),
		)
	}

	if date1.Year != date2.Year {
		return date1.Year < date2.Year
	}

	if date1.Month != date2.Month {
		return date1.Month < date2.Month
	}

	if date1.Day != date2.Day {
		return date1.Day < date2.Day
	}

	if date1.Trivia != date2.Trivia {
		return utils.IsStringAlphabeticallyLess(
			date1.Trivia,
			date2.Trivia,
		)
	}

	return true
}

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

// check if website1 is alphabetically lower than website2
func IsWebsiteLess(website1 Website, website2 Website) bool {
	if website1.Group != website2.Group {
		return utils.IsStringAlphabeticallyLess(
			string(website1.Group),
			string(website2.Group),
		)
	}

	if website1.URL != website2.URL {
		return utils.IsStringAlphabeticallyLess(
			website1.URL,
			website2.URL,
		)
	}

	if website1.Trivia != website2.Trivia {
		return utils.IsStringAlphabeticallyLess(
			website1.Trivia,
			website2.Trivia,
		)
	}

	return true
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

// check if companies are sorted alphabetically
func CheckCompaniesSorted(unsortedCompanies []Company, autoFix bool) bool {
	sortedCompanies := make([]Company, len(unsortedCompanies))
	copy(sortedCompanies, unsortedCompanies)

	// sort companies
	sort.Slice(sortedCompanies, func(i int, j int) bool {
		return IsCompanyLess(sortedCompanies[i], sortedCompanies[j])
	})

	// check if sorted and unsorted slices are equal
	sorted := reflect.DeepEqual(sortedCompanies, unsortedCompanies)

	if !sorted && autoFix {
		// if not sorted, apply fix
		WriteCompanies(sortedCompanies)

		return true
	}

	return sorted
}

// check if dates are sorted alphabetically
func CheckDatesSorted(unsortedDates []Date, autoFix bool) bool {
	sortedDates := make([]Date, len(unsortedDates))
	copy(sortedDates, unsortedDates)

	// sort dates
	sort.Slice(sortedDates, func(i int, j int) bool {
		return IsDateLess(sortedDates[i], sortedDates[j])
	})

	// check if sorted and unsorted slices are equal
	sorted := reflect.DeepEqual(sortedDates, unsortedDates)

	if !sorted && autoFix {
		// if not sorted, apply fix
		WriteDates(sortedDates)

		return true
	}

	return sorted
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

// check if websites are sorted alphabetically
func CheckWebsitesSorted(unsortedWebsites []Website, autoFix bool) bool {
	sortedWebsites := make([]Website, len(unsortedWebsites))
	copy(sortedWebsites, unsortedWebsites)

	// sort websites
	sort.Slice(sortedWebsites, func(i int, j int) bool {
		return IsWebsiteLess(sortedWebsites[i], sortedWebsites[j])
	})

	// check if sorted and unsorted slices are equal
	sorted := reflect.DeepEqual(sortedWebsites, unsortedWebsites)

	if !sorted && autoFix {
		// if not sorted, apply fix
		WriteWebsites(sortedWebsites)

		return true
	}

	return sorted
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
				"invalid group " +
					string(address.Group) +
					" for address " +
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

// validate data in companies.json
func ValidateCompaniesData(autofix bool) error {
	companies, err := GetCompanies()
	if err != nil {
		return err
	}

	for _, company := range companies {
		// check if company group enum is valid
		if !company.Group.IsValid() {
			return errors.New(
				"invalid group " +
					string(company.Group) +
					" for company " +
					company.Name,
			)
		}
	}

	if !CheckCompaniesSorted(companies, autofix) {
		return errors.New("companies not sorted properly")
	}

	return nil
}

// validate data in dates.json
func ValidateDatesData(autofix bool) error {
	dates, err := GetDates()
	if err != nil {
		return err
	}

	for _, date := range dates {
		dateString := strconv.Itoa(date.Year) +
			"-" +
			strconv.Itoa(date.Month) +
			"-" +
			strconv.Itoa(date.Day)

		// check if year is not older than 1970
		if date.Year < 1970 {
			return errors.New(
				"invalid year " +
					strconv.Itoa(date.Year) +
					" for date " +
					dateString +
					", " +
					"year must not be older than 1970",
			)
		}

		// check if date is valid
		_, err := time.Parse("2006-1-2", dateString)
		if err != nil {
			return err
		}

		// check if date group enum is valid
		if !date.Group.IsValid() {
			return errors.New(
				"invalid group " +
					string(date.Group) +
					" for date " +
					dateString,
			)
		}
	}

	if !CheckDatesSorted(dates, autofix) {
		return errors.New("dates not sorted properly")
	}

	return nil
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
					" for person " +
					person.FirstName +
					" " +
					person.LastName,
			)
		}

		// check if person group enum is valid
		if !person.Group.IsValid() {
			return errors.New(
				"invalid group " +
					string(person.Group) +
					" for person " +
					person.FirstName +
					" " +
					person.LastName,
			)
		}

		email := GenerateEmail(person.FirstName, person.LastName, person.Domain, EMAIL_DEFAULT_DOMAIN_SUFFIX)
		_, err := mail.ParseAddress(email)
		if err != nil {
			return errors.New(
				"invalid email " +
					email +
					" generated for person " +
					person.FirstName +
					" " +
					person.LastName +
					", " +
					err.Error(),
			)
		}
	}

	if !CheckPersonsSorted(persons, autofix) {
		return errors.New("persons not sorted properly")
	}

	return nil
}

// validate data in websites.json
func ValidateWebsitesData(autofix bool) error {
	websites, err := GetWebsites()
	if err != nil {
		return err
	}

	for _, website := range websites {
		// check if url is valid
		_, err := url.ParseRequestURI(website.URL)
		if err != nil {
			return err
		}

		// check if website group enum is valid
		if !website.Group.IsValid() {
			return errors.New(
				"invalid group " +
					string(website.Group) +
					" for website " +
					website.URL,
			)
		}
	}

	if !CheckWebsitesSorted(websites, autofix) {
		return errors.New("websites not sorted properly")
	}

	return nil
}

// validate all data files
func Validate(autofix bool) error {
	err := ValidateAddressesData(autofix)
	if err != nil {
		return err
	}

	err = ValidateCompaniesData(autofix)
	if err != nil {
		return err
	}

	err = ValidateDatesData(autofix)
	if err != nil {
		return err
	}

	err = ValidatePersonsData(autofix)
	if err != nil {
		return err
	}

	err = ValidateWebsitesData(autofix)
	if err != nil {
		return err
	}

	return nil
}
