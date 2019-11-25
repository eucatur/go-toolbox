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

// Coalesce retorna a primeira string que não for vazia conforme a order informada.
// Se todas forem vazias, o retorno também será vazio.
func Coalesce(expressions ...string) string {
	for _, expression := range expressions {
		if strings.TrimSpace(expression) != "" {
			return expression
		}
	}

	return ""
}

// ExistsValueInt verifica se o valor existe no slice de inteiros
//
//  existe := ExistsValueInt(200, []int{200, 204, 422})
//
func ExistsValueInt(value int, values []int) bool {
	for _, v := range values {
		if value == v {
			return true
		}
	}
	return false
}
