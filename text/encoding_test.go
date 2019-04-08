package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/charmap"
)

func TestUTF8ToISO8859_1(t *testing.T) {
	strUTF8 := "áBéCíDóFú"
	strISO8859_1Ref, _ := charmap.ISO8859_1.NewEncoder().String(strUTF8)

	strISO8859_1, err := UTF8ToISO8859_1(strUTF8)

	assert.Equal(t, strISO8859_1, strISO8859_1Ref)
	assert.Nil(t, err)
}

func TestISO8859_1ToUTF8(t *testing.T) {
	strUTF8Ref := "áBéCíDóFú"
	strISO8859_1, _ := charmap.ISO8859_1.NewEncoder().String(strUTF8Ref)

	strUTF8, err := ISO8859_1ToUTF8(strISO8859_1)

	assert.Equal(t, strUTF8, strUTF8Ref)
	assert.Nil(t, err)
}
