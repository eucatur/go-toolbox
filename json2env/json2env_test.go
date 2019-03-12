package json2env

import (
	"testing"
)

func TestLoadFile(t *testing.T) {
	err := LoadFile("test.json")

	if err != nil {
		t.Error(err)
	}
}
