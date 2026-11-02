package cnpj

import (
	"regexp"
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

func RemoverMascara(cnpj string) string {
	re := regexp.MustCompile(`[^A-Za-z0-9]`)
	return re.ReplaceAllString(cnpj, "")
}

// Valido verifica se um cpf é válido ou não
func Valido(cnpj string) bool {

	cnpj = RemoverMascara(cnpj)

	clean := strings.TrimSpace(strings.ToUpper(cnpj))
	if len(clean) != 14 {
		return false
	}
	if len(clean) > 0 {
		first := clean[0]
		allSame := true
		for i := 1; i < len(clean); i++ {
			if clean[i] != first {
				allSame = false
				break
			}
		}
		if allSame {
			return false
		}
	}

	values := make([]int, 0, 14)
	for _, r := range clean {
		switch {
		case r >= '0' && r <= '9':
			values = append(values, int(r)-48)
		case r >= 'A' && r <= 'Z':
			values = append(values, int(r)-48)
		default:
			return false
		}
	}
	if len(values) != 14 {
		return false
	}

	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum1 := 0
	for i, weight := range weights1 {
		sum1 += values[i] * weight
	}
	first := sum1 % 11
	if first < 2 {
		first = 0
	} else {
		first = 11 - first
	}
	if first != values[12] {
		return false
	}

	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum2 := 0
	for i, weight := range weights2 {
		sum2 += values[i] * weight
	}
	second := sum2 % 11
	if second < 2 {
		second = 0
	} else {
		second = 11 - second
	}
	return second == values[13]
}

func DefinirMascara(cnpj string) string {

	re := regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`)
	if !re.MatchString(cnpj) {
		return cnpj
	}
	return re.ReplaceAllString(cnpj, "$1.$2.$3/$4-$5")

}
