package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLPad(t *testing.T) {
	assert.Equal(t, "00A", LPad("A", 3, "0"))
	assert.Equal(t, "  A", LPad("A", 3, " "))
	assert.Equal(t, "--A", LPad("A", 3, "-"))

	assert.Equal(t, "0A", LPad("A", 2, "0"))
	assert.Equal(t, " A", LPad("A", 2, " "))
	assert.Equal(t, "-A", LPad("A", 2, "-"))

	assert.Equal(t, "AB", LPad("AB", 2, "0"))
	assert.Equal(t, "AB", LPad("AB", 2, " "))
	assert.Equal(t, "AB", LPad("AB", 2, "-"))
}

func TestRPad(t *testing.T) {
	assert.Equal(t, "A00", RPad("A", 3, "0"))
	assert.Equal(t, "A  ", RPad("A", 3, " "))
	assert.Equal(t, "A--", RPad("A", 3, "-"))

	assert.Equal(t, "A0", RPad("A", 2, "0"))
	assert.Equal(t, "A ", RPad("A", 2, " "))
	assert.Equal(t, "A-", RPad("A", 2, "-"))

	assert.Equal(t, "AB", RPad("AB", 2, "0"))
	assert.Equal(t, "AB", RPad("AB", 2, " "))
	assert.Equal(t, "AB", RPad("AB", 2, "-"))
}
