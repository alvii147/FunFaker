package date

import (
	"time"

	"github.com/alvii147/FunFaker/data"
)

// GET date request URL parameters
type DateRequest struct {
	After  time.Time      `schema:"after"`
	Before time.Time      `schema:"before"`
	Group  data.DateGroup `schema:"group"`
}

// GET date request response body
type DateResponse struct {
	Day    int    `json:"day"`
	Month  int    `json:"month"`
	Year   int    `json:"year"`
	Trivia string `json:"trivia"`
}

// update date response using date
func (DateResponse *DateResponse) FromDate(date data.Date) {
	DateResponse.Day = date.Day
	DateResponse.Month = date.Day
	DateResponse.Year = date.Year
	DateResponse.Trivia = date.Trivia
}
