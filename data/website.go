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

// path to websites.json
const WEBSITES_FILE_NAME = "websites.json"

// enum representing website group
type WebsiteGroup string

const (
	WebsiteGroupTVShows WebsiteGroup = "TV-Shows"
)

// check website group enum is valid
func (group *WebsiteGroup) IsValid() bool {
	return utils.StringSoftEqual(string(*group), string(WebsiteGroupTVShows))
}

// struct representing website from websites.json
type Website struct {
	URL    string       `json:"url"`
	Group  CompanyGroup `json:"group"`
	Trivia string       `json:"trivia"`
}

// read websites from websites.json
func GetWebsites() ([]Website, error) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("unable to get current directory")
	}

	// open websites.json
	websitesFilePath := filepath.Join(path.Dir(filename), WEBSITES_FILE_NAME)

	// read websites.json
	websitesBytes, err := os.ReadFile(websitesFilePath)
	if err != nil {
		return nil, err
	}

	// get websites from bytes read
	var websites []Website
	err = json.Unmarshal(websitesBytes, &websites)
	if err != nil {
		return nil, err
	}

	return websites, nil
}

// write list of websites to websites.json
func WriteWebsites(websites []Website) error {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get current directory")
	}

	// open websites.json
	websitesFilePath := filepath.Join(path.Dir(filename), WEBSITES_FILE_NAME)

	// convert to bytes with indentation
	file, err := json.MarshalIndent(websites, "", "    ")
	if err != nil {
		return err
	}

	// write to websites.json
	err = os.WriteFile(websitesFilePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// filter websites by properties
func FilterWebsites(
	websites []Website,
	url string,
	group WebsiteGroup,
	trivia string,
) []Website {
	filteredWebsites := []Website{}
	for _, website := range websites {
		if !utils.StringSoftEqual(url, website.URL) {
			continue
		}

		if !utils.StringSoftEqual(string(group), string(website.Group)) {
			continue
		}

		if !utils.StringSoftEqual(trivia, website.Trivia) {
			continue
		}

		filteredWebsites = append(filteredWebsites, website)
	}

	return filteredWebsites
}
