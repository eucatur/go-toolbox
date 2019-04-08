package text

import (
	"fmt"
	"strings"
)

// LPad preenche a string com o caracter separador à esquerda.
//
//  str := LPad("A", 3, "0")  // "00A"
//  str := LPad("A", 3, " ")  // "  A"
//  str := LPad("A", 3, "-")  // "--A"
//
func LPad(str string, size int, sep string) string {
	if sep == " " || sep == "0" {
		format := fmt.Sprintf("%ds", size)

		if sep == " " {
			return fmt.Sprintf("%"+format, str)
		}

		return fmt.Sprintf("%0"+format, str)
	}

	if len(str) >= size {
		return str[:size]
	}

	strFull := strings.Repeat(sep, size) + str
	return strFull[len(strFull)-size:]
}

// RPad preenche a string com o caracter separador à direita.
//
//  str := RPad("A", 3, "0")  // "A00"
//  str := RPad("A", 3, " ")  // "A  "
//  str := RPad("A", 3, "-")  // "A--"
//
func RPad(str string, size int, sep string) string {
	if sep == " " {
		format := fmt.Sprintf("%ds", size)

		if sep == " " {
			return fmt.Sprintf("%-"+format, str)
		}
	}

	if len(str) >= size {
		return str[:size]
	}

	strFull := str + strings.Repeat(sep, size)
	return strFull[0:size]
}
