package text

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/eucatur/go-toolbox/slice"
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

func TestOnlyNumbersToInt64(t *testing.T) {
	assert.Equal(t, int64(1), OnlyNumbersToInt64("1"))
	assert.Equal(t, int64(12), OnlyNumbersToInt64("01"))
	assert.Equal(t, int64(1), OnlyNumbersToInt64("A1"))
	assert.Equal(t, int64(1), OnlyNumbersToInt64("A1A"))
	assert.Equal(t, int64(0), OnlyNumbersToInt64("ABC"))
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

func TestRandomCharacters(t *testing.T) {
	var (
		textGenerated []string
		text          string
		err           error
	)

	loop := 6
	for i := 0; i < loop; i++ {

		err = nil

		text = RandomCharacters(6)

		if strings.EqualFold(strings.ToUpper(text), strings.ToUpper("OJNNPG")) {
			err = errors.New("Ever is generated the 'OJNNPG' when run first")
		}

		assert.Nil(t, err)

		if strings.EqualFold(strings.ToUpper(text), strings.ToUpper("SIUZYT")) {
			err = errors.New("Ever is generated the 'SIUZYT' when run twice")
		}

		assert.Nil(t, err)

		if !slice.SliceExists(textGenerated, text) {
			textGenerated = append(textGenerated, text)
		} else {
			err = fmt.Errorf("Text already generated %s", text)
		}

		assert.Nil(t, err)

	}
}
