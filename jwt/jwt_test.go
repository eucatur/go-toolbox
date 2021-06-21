package jwt_test

import (
	"testing"

	jwtgotoolbox "github.com/eucatur/go-toolbox/jwt"
	"github.com/stretchr/testify/assert"
)

func Test_GetClaims(t *testing.T) {

	claims, err := jwtgotoolbox.GetClaims("eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJnZXRfY2xhaW1zIjoiMTIzNDU2Nzg5MCJ9.SGBVVfkS56Mtfx51mVjcbplCFDWmSBwqTHOEMmUae0rfhl5PbK-ypsDox72EqBs2")

	assert.Nil(t, err)

	assert.NotNil(t, claims)

	value, ok := claims["get_claims"]

	assert.Equal(t, true, ok)

	assert.Equal(t, "1234567890", value)
}
