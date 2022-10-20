package money

import (
	"bytes"
	"fmt"
	"math"
	"reflect"

	"github.com/eucatur/decimal"
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
//	dinheiro := Round(150.141, 2)  // 150.14
//	dinheiro := Round(150.145, 2)  // 150.15
//	dinheiro := Round(150.146, 2)  // 150.15
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

func ToFloat(value int) float64 {
	v := float64(value)
	v = v / 100
	return v
}

func Format(valor interface{}) string {
	v := reflect.ValueOf(valor)

	switch v.Kind() {
	case reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Float64:
		return fmt.Sprintf("%.2f", v.Float())
	}

	return ""
}

// Truncate irá pegar a quantidade de casas definidas pelo parâmetro decimals
// de acordo com a valor passado no parâmetro value.
//
// Exemplos:
//
//	amount := Truncate(150.141, 2)  // 150.14
//	amount := Truncate(150.145234234, 3)  // 150.145
//	amount := Truncate(100.998, 2)  // 100.99
// package used github.com/shopspring/decimal
func Truncate(value float64, decimals int) float64 {

	d := decimal.NewFromFloat(value)
	if value > 0 {
		dd := d.RoundFloor(int32(decimals))
		v, _ := dd.Float64()

		return v
	} else {
		dd := d.RoundCeil(int32(decimals))
		v, _ := dd.Float64()

		return v
	}

}
