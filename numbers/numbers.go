// Package numbers é um pacote para funções que trabalham com números.
package numbers

import (
	"math"
)

// Decimals ajusta a quantidade de casas decimais no parâmetro 'num' para a
// quantidade total definida no parâmetro 'places'.
//
// ATENÇÃO: esse método não faz round das casas decimais, ele apenas faz o
// "truncamento". Para utilização de round, utilize o pacote:
//  bitbucket.org/eucatur/go-toolbox/lib/money/
//
// Exemplo:
//  // converte o float para utilizar 2 casas decimais
//  numero := Decimais(50.124, 2) // 50.12
//
func Decimals(num float64, places int) float64 {
	ind := math.Pow(10, float64(places))

	return float64(int(num*ind)) / ind
}

// IntExistsInSlice retorna se o int 'num' existe no slice 'list'.
//
//  existe := IntExistsInSlice(1, []int{1, 2, 3})  // true
//
func IntExistsInSlice(num int, list []int) bool {
	for _, i := range list {
		if num == i {
			return true
		}
	}
	return false
}

// UniqueInts retorna um slice único de números presentes em 'list'. Se um número aparecer
// mais de uma vez em 'list', ele aparecerá somente uma vez no resultado.
//
//  unicos := UniqueInts([]int{1, 1, 2, 2, 3, 3})  // []int{1, 2, 3}
//
func UniqueInts(list []int) []int {
	result := []int{}
	present := map[int]bool{}

	for _, num := range list {
		if _, exists := present[num]; !exists {
			result = append(result, num)
			present[num] = true
		}
	}
	return result
}

// UniqueInts64 retorna um slice único de números presentes em 'list'. Se um número aparecer
// mais de uma vez em 'list', ele aparecerá somente uma vez no resultado.
//
//  unicos := UniqueInts64([]int{1, 1, 2, 2, 3, 3})  // []int64{1, 2, 3}
//
func UniqueInts64(list []int64) []int64 {
	result := []int64{}
	present := map[int64]bool{}

	for _, num := range list {
		if _, exists := present[num]; !exists {
			result = append(result, num)
			present[num] = true
		}
	}
	return result
}

// AbsInt retorna o valor absoluto do int informado: remove o sinal de menos
func AbsInt(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
