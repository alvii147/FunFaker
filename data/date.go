package data

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/alvii147/FunFaker/utils"
)

// path to dates.json
const DATES_FILE_NAME = "dates.json"

// enum representing date group
type DateGroup string

const (
	DateGroupGames   DateGroup = "Games"
	DateGroupMovies  DateGroup = "Movies"
	DateGroupTVShows DateGroup = "TV-Shows"
)

// check date group enum is valid
func (group *DateGroup) IsValid() bool {
	return utils.StringSoftEqual(string(*group), string(DateGroupGames)) ||
		utils.StringSoftEqual(string(*group), string(DateGroupMovies)) ||
		utils.StringSoftEqual(string(*group), string(DateGroupTVShows))
}

// struct representing website from websites.json
type Date struct {
	Day    int       `json:"day"`
	Month  int       `json:"month"`
	Year   int       `json:"year"`
	Group  DateGroup `json:"group"`
	Trivia string    `json:"trivia"`
}

// convert date to time object
func (date *Date) ToTimeObj() time.Time {
	return time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, time.UTC)
}

// read dates from dates.json
func GetDates() ([]Date, error) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("unable to get current directory")
	}

	// open dates.json
	datesFilePath := filepath.Join(path.Dir(filename), DATES_FILE_NAME)

	// read dates.json
	datesBytes, err := os.ReadFile(datesFilePath)
	if err != nil {
		return nil, err
	}

	// get dates from bytes read
	var dates []Date
	err = json.Unmarshal(datesBytes, &dates)
	if err != nil {
		return nil, err
	}

	return dates, nil
}

// write list of dates to dates.json
func WriteDates(dates []Date) error {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get current directory")
	}

	// open dates.json
	datesFilePath := filepath.Join(path.Dir(filename), DATES_FILE_NAME)

	// convert to bytes with indentation
	file, err := json.MarshalIndent(dates, "", "    ")
	if err != nil {
		return err
	}

	// write to dates.json
	err = os.WriteFile(datesFilePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// filter dates by properties
func FilterDates(
	dates []Date,
	startDate time.Time,
	endDate time.Time,
	group DateGroup,
	trivia string,
) []Date {
	filteredDates := []Date{}
	for _, date := range dates {
		timeObj := date.ToTimeObj()
		if !startDate.IsZero() && startDate.After(timeObj) {
			continue
		}

		if !endDate.IsZero() && endDate.Before(timeObj) {
			continue
		}

		if !utils.StringSoftEqual(string(group), string(date.Group)) {
			continue
		}

		if !utils.StringSoftEqual(trivia, date.Trivia) {
			continue
		}

		filteredDates = append(filteredDates, date)
	}

	return filteredDates
}
