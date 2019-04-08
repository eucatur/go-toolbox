package crypt

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 gera um hash SHA256 para o parâmetro content.
//
//  hash := Sha256("texto")
//
func Sha256(content string) string {
	h := sha256.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
