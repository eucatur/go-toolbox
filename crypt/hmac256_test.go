package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHmac256(t *testing.T) {
	assert.Equal(t, "9d2bb116e1df997ffe8a5139fc1d187f976c19579a138414a112bc2e39020eba", Hmac256("test", "123456"))
	assert.Equal(t, "5d45ab14682a4d29a6f0531904dd9c23ef452d3b83e18d7f04853f9850e85e40", Hmac256("string to convert", "123456"))
}
