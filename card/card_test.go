package card

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMask(t *testing.T) {
	card, _ := Mask("5353 1607 6798 7690")

	assert.Equal(t, card, "5353********7690", "The two words should be the same.")
}
