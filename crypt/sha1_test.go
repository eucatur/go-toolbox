package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha1(t *testing.T) {
	assert.Equal(t, "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3", Sha1("test"))
	assert.Equal(t, "2c00c8d78bc28d3ffb26a3751aae13028c9f0e67", Sha1("string to convert"))
}
