package cpf

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func mascarado(cpf string) bool {
	r := regexp.MustCompile("^[0-9]{3}.[0-9]{3}.[0-9]{3}-[0-9]{2}$")
	return r.MatchString(cpf)
}

func desmascarado(cpf string) bool {
	r := regexp.MustCompile("^[0-9]{11}$")
	return r.MatchString(cpf)
}

// SomenteNumeros retorna somente os numeros da string
func SomenteNumeros(cpf string) string {
	r := regexp.MustCompile("[0-9]+")
	return strings.Join(r.FindAllString(cpf, -1), "")
}

// Valido verifica se um cpf é válido ou não
func Valido(cpf string) bool {

	if !mascarado(cpf) && !desmascarado(cpf) {
		return false
	}

	cpf = SomenteNumeros(cpf)

	var (
		nums         = make([]float64, 11)
		err          error
		seq1         = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2}
		seq2         = []float64{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
		soma1, soma2 float64
	)

	for i := 0; i < 11; i++ {
		nums[i], err = strconv.ParseFloat(cpf[i:i+1], 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	todosSaoIguais := true
	for i := 0; i < 10; i++ {
		if nums[i] != nums[i+1] {
			todosSaoIguais = false
		}
	}

	if todosSaoIguais {
		return false
	}

	for i := 0; i < 9; i++ {
		soma1 += nums[i] * seq1[i]
	}

	resto1 := math.Mod((soma1 * 10), 11)

	if resto1 == 10 {
		if nums[9] != 0 {
			return false
		}
	} else {
		if nums[9] != resto1 {
			return false
		}
	}

	for i := 0; i < 10; i++ {
		soma2 += nums[i] * seq2[i]
	}

	resto2 := math.Mod((soma2 * 10), 11)

	if resto2 == 10 {
		if nums[10] != 0 {
			return false
		}
	} else {
		if nums[10] != resto2 {
			return false
		}
	}

	return true
}

func DefinirMascara(cpf string) string {

	re := regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
	if !re.MatchString(cpf) {
		return cpf
	}

	return re.ReplaceAllString(cpf, "$1.$2.$3-$4")

}
