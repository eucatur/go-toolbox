package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Hmac256 gera um hash HMAC256 para o parâmetro content usando como chave
// o parâmetro key.
//
//  hash := Hmac256("texto", "chave-secreta")
//
func Hmac256(content, key string) string {
	k := []byte(key)
	h := hmac.New(sha256.New, k)
	h.Write([]byte(content))

	return hex.EncodeToString(h.Sum(nil))
}
