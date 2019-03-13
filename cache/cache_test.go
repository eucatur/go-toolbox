package cache

import (
	"errors"
	"testing"
)

const KEY = "key"
const VALUE = "value"

func TestSet(t *testing.T) {
	Set(KEY, VALUE, DefaultExpiration)
}
func TestGet(t *testing.T) {
	TestSet(t)

	value, found := Get(KEY)

	if !found || value != VALUE {
		t.Error(errors.New("Cache not found"))
	}
}
