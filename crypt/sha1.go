package crypt

import (
	cryptosha1 "crypto/sha1"
	"encoding/hex"
)

// Sha1 gera um hash SHA1 para o parâmetro content.
//
//  hash := Sha1("texto")
//
func Sha1(content string) string {
	h := cryptosha1.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
