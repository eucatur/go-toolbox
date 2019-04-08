package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 gera um hash MD5 para o parâmetro content.
//
//  hash := Md5("texto")
//
func Md5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
