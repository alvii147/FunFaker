package data_test

import (
	"testing"

	"github.com/alvii147/FunFaker/data"
)

// test validation of all data files
func TestValidate(t *testing.T) {
	err := data.Validate(false)
	if err != nil {
		t.Error("error validating data:", err)
	}
}
