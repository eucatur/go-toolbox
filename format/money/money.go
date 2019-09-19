package money

import (
	"bytes"
	"fmt"
	"math"
)

func Reais(valor int64) string {
	var buff bytes.Buffer

	reais := valor / 100
	centavos := valor % 100

	if reais <= 999 {
		buff.WriteString(fmt.Sprintf("%d", reais))
	} else {
		var total, milhares = reais, reais

		for {
			milhares = total / 1000
			reais = total % 1000

			if milhares <= 0 {
				buff.WriteString(fmt.Sprintf("%03d", reais))
				break
			}

			buff.WriteString(fmt.Sprintf("%d", milhares))
			buff.WriteString(".")

			total = reais
		}
	}

	buff.WriteString(",")
	buff.WriteString(fmt.Sprintf("%02d", centavos))

	return buff.String()
}

// Round faz o arredondamento das casas decimais do parâmetro 'value' de
// acordo com a quantidade de casas informado em 'precision'.
//
// Exemplos:
//
//  dinheiro := Round(150.141, 2)  // 150.14
//  dinheiro := Round(150.145, 2)  // 150.15
//  dinheiro := Round(150.146, 2)  // 150.15
//
func Round(value float64, precision int) float64 {
	var round float64

	pow := math.Pow(10, float64(precision))
	digit := pow * value
	_, div := math.Modf(digit)

	if div >= .5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	return round / pow
}

// ToInt - Realiza a conversão de valor em float para int considerando casas deciamis
// exemplo
// value = 18.9 irá retonar 1890
func ToInt(value float64) int {
	if value < 0 {
		return int(value*100 - 0.5)
	}
	return int(value*100 + 0.5)
}
