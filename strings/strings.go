// Package strings é um pacote para funções que trabalham com strings.
package strings

import (
	"regexp"
	"strings"
)

var regexFirstUpper = regexp.MustCompile("(.)([A-Z][a-z]+)")
var regexAllUpper = regexp.MustCompile("([a-z0-9])([A-Z])")

// ExistsInSlice retorna se a string 'str' existe no slice 'list'.
//
//  existe := ExistsInSlice("A", []string{"A", "B", "C"})
//
func ExistsInSlice(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}
	return false
}

// SnakeCase converte uma string para o formato snake case.
//
//  str := SnakeCase("MeuTexto") // meu_texto
//
func SnakeCase(str string) string {
	str = regexFirstUpper.ReplaceAllString(str, "${1}_${2}")
	str = regexAllUpper.ReplaceAllString(str, "${1}_${2}")
	str = strings.ToLower(str)
	return strings.Replace(str, "__", "_", -1)
}

// Coalesce se a primeira string(firstStr) não for vazia, retorna a primeira string.
// Se a primeira string for vazia, retorna a segunda string(secondStr).
func Coalesce(firstStr, secondStr string) string {
	if strings.TrimSpace(firstStr) != "" {
		return firstStr
	}
	return secondStr
}
