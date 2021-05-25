// It's a wrapper to [jwt-go](https://github.com/dgrijalva/jwt-go) for facilitate use for jwt in golang
package jwt

import (
	"errors"

	jwt_go "github.com/dgrijalva/jwt-go"
)

// Functions for create token with some claims and one secret
func CreateTokenWithClaims(claims jwt_go.MapClaims, secret string) (token string, err error) {
	t := jwt_go.NewWithClaims(jwt_go.SigningMethodHS384, claims)

	token, err = t.SignedString([]byte(secret))

	return
}

// VerifyTokenAndGetClaims verify if is token is valid
func VerifyTokenAndGetClaims(tokenString, secret string) (claims map[string]interface{}, err error) {
	token, err := jwt_go.Parse(tokenString, func(token *jwt_go.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return
	}

	errInvalidToken := errors.New("token inválido")
	if !token.Valid {
		err = errInvalidToken
		return
	}

	claims, ok := token.Claims.(jwt_go.MapClaims)
	if !ok {
		err = errInvalidToken
		return
	}

	return
}

func VerifyTokenAndGetClaimsWithRSAPublicKey(tokenString string, public_key []byte) (claims map[string]interface{}, err error) {
	verifyKey, err := jwt_go.ParseRSAPublicKeyFromPEM(public_key)
	if err != nil {
		return
	}

	token, err := jwt_go.Parse(tokenString, func(token *jwt_go.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		return
	}

	errInvalidToken := errors.New("token inválido")
	if !token.Valid {
		err = errInvalidToken
		return
	}

	claims, ok := token.Claims.(jwt_go.MapClaims)
	if !ok {
		err = errInvalidToken
		return
	}

	return
}
