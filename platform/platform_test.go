package platform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromUserAgent(t *testing.T) {
	p, _ := FromUserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1")

	assert.Equal(t, Mobile, p)

	p, _ = FromUserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")

	assert.Equal(t, Desktop, p)
}
