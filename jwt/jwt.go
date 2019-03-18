// It's a wrapper to [jwt-go](https://github.com/dgrijalva/jwt-go) for facilitate use for jwt in golang
package jwt

import (
	jwt_go "github.com/dgrijalva/jwt-go"
)

// Functions for create token with some claims and one secret
func CreateTokenWithClaims(claims jwt_go.MapClaims, secret string) (token string, err error) {
	t := jwt_go.NewWithClaims(jwt_go.SigningMethodHS384, claims)

	token, err = t.SignedString([]byte(secret))

	return
}

// Verify if is token is valid
func VerifyTokenAndGetClaims(tokenString, secret string) (map[string]interface{}, error) {
	token, err := jwt_go.Parse(tokenString, func(token *jwt_go.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt_go.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
