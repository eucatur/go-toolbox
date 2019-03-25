package cnpj

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func mascarado(cnpj string) bool {
	r := regexp.MustCompile("^[0-9]{2}.[0-9]{3}.[0-9]{3}/[0-9]{4}-[0-9]{2}$")
	return r.MatchString(cnpj)
}

func desmascarado(cnpj string) bool {
	r := regexp.MustCompile("^[0-9]{14}$")
	return r.MatchString(cnpj)
}

// SomenteNumeros retorna somente os numeros da string
func SomenteNumeros(cnpj string) string {
	r := regexp.MustCompile("[0-9]+")
	return strings.Join(r.FindAllString(cnpj, -1), "")
}

// Valido verifica se um cpf é válido ou não
func Valido(cnpj string) bool {

	if !mascarado(cnpj) && !desmascarado(cnpj) {
		return false
	}

	cnpj = SomenteNumeros(cnpj)

	var (
		nums         = make([]float64, 14)
		err          error
		seq1         = []float64{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		seq2         = []float64{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		soma1, soma2 float64
	)

	for i := 0; i < 14; i++ {
		nums[i], err = strconv.ParseFloat(cnpj[i:i+1], 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	todosSaoIguais := true
	for i := 0; i < 13; i++ {
		if nums[i] != nums[i+1] {
			todosSaoIguais = false
		}
	}

	if todosSaoIguais {
		return false
	}

	for i := 0; i < 12; i++ {
		soma1 += nums[i] * seq1[i]
	}

	resto1 := math.Mod(soma1, 11)

	if resto1 < 2 {
		if nums[12] != 0 {
			return false
		}
	} else {
		if nums[12] != 11-resto1 {
			return false
		}
	}

	for i := 0; i < 13; i++ {
		soma2 += nums[i] * seq2[i]
	}

	resto2 := math.Mod(soma2, 11)

	if resto2 < 2 {
		if nums[13] != 0 {
			return false
		}
	} else {
		if nums[13] != 11-resto2 {
			return false
		}
	}

	return true
}
