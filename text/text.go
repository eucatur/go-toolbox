// Package text disponibiliza métodos para trabalhar com textos/strings.
package text

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// encontra todos os caracteres não-numéricos.
var regexExceptNumbers = regexp.MustCompile("[^0-9]")

// OnlyNumbers retorna apenas os números de uma string.
//
//   str := OnlyNumbers("A123BC")  // "123"
//
func OnlyNumbers(str string) string {
	return regexExceptNumbers.ReplaceAllString(str, "")
}

// OnlyNumbersToInt64 retorna apenas os números de uma string, removendo zeros a esquerda. Caso não tenha nenhum numero retorno é 0.
//
//   str := OnlyNumbersToInt64("A123BC")  // 123
//   str := OnlyNumbersToInt64("00123")  // 123
//   str := OnlyNumbersToInt64("ABC")  // 0
//
func OnlyNumbersToInt64(str string) int64 {
	number, _ := strconv.ParseInt(strings.TrimLeft(OnlyNumbers(str), "0"), 10, 64)
	return number
}

// Normalize substitui caracteres especiais de uma string por caracteres
// da tabela ASCII.
//
//  str := Normalize("JI-PARANÁ")  // "JI-PARANA"
//
func Normalize(str string) (string, error) {
	transf := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	data, _, err := transform.String(transf, str)
	return data, err
}

// RandomCharacters gerar caracteres alertórios com base na quantidade informada
//
// str := RandomCharacters(4) // XwpT
//
func RandomCharacters(numberOfCharacters int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, numberOfCharacters)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// RemoveAccents is a functions for remove accents
func RemoveAccents(s string) (r string) {
	r, _ = Normalize(s)
	r = regexp.MustCompile(`[^\w]|\s`).ReplaceAllString(r, "")
	return
}

// PadRight - Complementa com determinado valor a quantidade que falta para inteirar uma string desejada
// uma string deve ser abxx e só vem  ab com esta função irá inserir o caracter desejado para complementar
// Aplicação: PadRight("ab", "x", 4) irá retornar abxx
func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) == lenght {
			return str[0:lenght]
		}
	}
}

// PadLeft - Complementa com determinado valor a quantidade que falta para inteirar uma string desejada
// uma string deve ser xxab e só vem  ab com esta função irá inserir o caracter desejado para complementar
// Aplicação: PadLeft("ab", "x", 4) irá retornar xxab
func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) == lenght {
			return str[0:lenght]
		}
	}
}

// StringIsEmptyOrWhiteSpace função responsável por verificar string completamente vazia
func StringIsEmptyOrWhiteSpace(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
