package viacep

import (
	"testing"
)

func TestFindByZipCode(t *testing.T) {
	zipcode := 76901042

	_, err := FindByZipCode(zipcode)

	if err != nil {
		t.Error(err)
	}
}
