package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "098f6bcd4621d373cade4e832627b4f6", Md5("test"))
	assert.Equal(t, "75ffdc5159c4ac4df06394f8ccb2a65f", Md5("string to convert"))
}
