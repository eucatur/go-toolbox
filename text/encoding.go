package text

import (
	"golang.org/x/text/encoding/charmap"
)

// UTF8ToISO8859_1 transforma uma string UTF8 numa string ISO8859_1.
//
//  strISO, err := UTF8ToISO8859_1(strUFT8)
//
func UTF8ToISO8859_1(strUTF8 string) (string, error) {
	return charmap.ISO8859_1.NewEncoder().String(strUTF8)
}

// ISO8859_1ToUTF8 transforma uma string ISO8859_1 numa string UTF8.
//
//  strUTF8, err := ISO8859_1ToUTF8(strISO)
//
func ISO8859_1ToUTF8(strISO string) (string, error) {
	return charmap.ISO8859_1.NewDecoder().String(strISO)
}
