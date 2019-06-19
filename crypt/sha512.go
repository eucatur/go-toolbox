package crypt

import (
	"crypto/sha512"
	"encoding/hex"
)

// Sha512 gera um hash SHA512 para o parâmetro content.
//
//  hash := Sha512("texto")
//
func Sha512(content string) string {
	h := sha512.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
