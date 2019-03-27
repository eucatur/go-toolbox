package money

import (
	"bytes"
	"fmt"
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
