package website

import (
	"github.com/alvii147/FunFaker/data"
)

// GET website request URL parameters
type WebsiteRequest struct {
	Group data.WebsiteGroup `schema:"group"`
}

// GET website request response body
type WebsiteResponse struct {
	URL    string `json:"url"`
	Trivia string `json:"trivia"`
}

func (websiteResponse *WebsiteResponse) FromWebsite(website data.Website) {
	websiteResponse.URL = website.URL
	websiteResponse.Trivia = website.Trivia
}
