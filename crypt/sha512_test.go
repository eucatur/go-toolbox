package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha512(t *testing.T) {
	assert.Equal(t, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", Sha512("test"))
	assert.Equal(t, "6cde6c16f8c67b4ce2597ed2ed52979feffcabbb8e10d99be281305f0dd3ab5b", Sha512("string to convert"))
}
