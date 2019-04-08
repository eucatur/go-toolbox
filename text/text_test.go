package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnlyNumbers(t *testing.T) {
	assert.Equal(t, "1", OnlyNumbers("1"))
	assert.Equal(t, "1", OnlyNumbers("1A"))
	assert.Equal(t, "1", OnlyNumbers("A1"))
	assert.Equal(t, "1", OnlyNumbers("A1A"))
	assert.Equal(t, "123", OnlyNumbers("1A2B3C"))
	assert.Equal(t, "123", OnlyNumbers("123ABC"))
	assert.Equal(t, "123", OnlyNumbers("ABC123"))
}

func TestNormalize(t *testing.T) {
	var (
		text string
		err  error
	)

	text, err = Normalize("JI-PARANÁ")
	assert.Equal(t, "JI-PARANA", text)
	assert.Nil(t, err)

	text, err = Normalize("ji-paraná")
	assert.Equal(t, "ji-parana", text)
	assert.Nil(t, err)

	text, err = Normalize("áBéCíDóFú")
	assert.Equal(t, "aBeCiDoFu", text)
	assert.Nil(t, err)

	text, err = Normalize("àBèCìDòFù")
	assert.Equal(t, "aBeCiDoFu", text)
	assert.Nil(t, err)

	text, err = Normalize("ÁBÉCÍDÓFÚ")
	assert.Equal(t, "ABECIDOFU", text)
	assert.Nil(t, err)

	text, err = Normalize("ÀBÈCÌDÒFÙ")
	assert.Equal(t, "ABECIDOFU", text)
	assert.Nil(t, err)
}
