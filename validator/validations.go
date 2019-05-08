package validator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/eucatur/go-toolbox/validator/cnpj"
	"github.com/eucatur/go-toolbox/validator/cpf"
)

var bpeRgx = map[string]string{
	"ER12": `0|0\.[0-9]{2}|[1-9]{1}[0-9]{0,2}(\.[0-9]{2})?`,
	"ER14": `0\.[0-9]{1}[1-9]{1}|0\.[1-9]{1}[0-9]{1}|[1-9]{1}[0-9]{0,2}(\.[0-9]{2})?`,
	"ER27": `0|0\.[0-9]{2}|[1-9]{1}[0-9]{0,12}(\.[0-9]{2})?`,
	"ER29": `[0-9]{0,14}|ISENTO|PR[0-9]{4,8}`,
	"ER35": `[!-ÿ]{1}[ -ÿ]{0,}[!-ÿ]{1}|[!-ÿ]{1}`,
	"ER43": `([!-ÿ]{0}|[!-ÿ]{5,20})?`,
	"ER50": `[^@]+@[^\.]+\..+`,
}

var bpeEnum = map[string][]string{
	"D1":  []string{"11", "12", "13", "14", "15", "16", "17", "21", "22", "23", "24", "25", "26", "27", "28", "29", "31", "32", "33", "35", "41", "42", "43", "50", "51", "52", "53"},
	"D7":  []string{"1", "2"},
	"D8":  []string{"1", "2", "3"},
	"D9":  []string{"00", "01"},
	"D10": []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	"D11": []string{"1", "2", "3", "4", "5"},
	"D13": []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "99"},
	"D16": []string{"00"},
	"D17": []string{"20"},
	"D18": []string{"40", "41", "51"},
	"D19": []string{"90"},
	"D24": []string{"1", "2", "3", "4", "5", "9"},
}

var (
	ufs       = []string{"AC", "AL", "AM", "AP", "BA", "CE", "DF", "ES", "GO", "MA", "MG", "MS", "MT", "PA", "PB", "PE", "PI", "PR", "RJ", "RN", "RO", "RR", "RS", "SC", "SE", "SP", "TO"}
	ufsWithEX = append(ufs, "EX")
)

// registro das validações
func init() {
	validator.addValidation("msg", vMsg)
	validator.addValidation("required", vRequired)
	validator.addValidation("uint", vUint)
	validator.addValidation("datebr", vDateBR)
	validator.addValidation("dateeua", vDateEUA)
	validator.addValidation("datetime", vDateTime)
	validator.addValidation("len", vLen)
	validator.addValidation("min", vMin)
	validator.addValidation("max", vMax)
	validator.addValidation("enum", vEnum)
	validator.addValidation("email", vEmail)
	validator.addValidation("numeric", vNumeric)
	validator.addValidation("cpf", vCPF)
	validator.addValidation("cnpj", vCNPJ)
	validator.addValidation("bpe", vBPE)
	validator.addValidation("uf", vUF)
	validator.addValidation("ufwithex", vUFWithEX)
	validator.addValidation("ipv4", vIPv4)
}

func vMsg(vf vField) string {
	return vf.getValidation("msg").Param
}

func vRequired(vf vField) string {
	if vf.isString() && len(vf.toString()) == 0 {
		return "Este campo é obrigatório"
	}
	if vf.isSlice() && vf.len() == 0 {
		return "Informe ao menos um elemento para esse campo (array)"
	}
	return ""
}

func vUint(vf vField) string {

	msg := "Informe um valor positivo (maior que zero)"

	if vf.isInt() && vf.toInt() < 1 {
		return msg
	}

	if vf.isInt64() && vf.toInt64() < 1 {
		return msg
	}

	return ""
}

func vDateBR(vf vField) string {

	layout := "02/01/2006"

	if vf.isString() && vf.toString() != "" {
		if _, err := time.Parse(layout, vf.toString()); err != nil {
			return "Informe uma data em formato brasileiro (ex: " + layout + ")"
		}
	}

	return ""
}

func vDateEUA(vf vField) string {

	layout := "2006-01-02"

	if vf.isString() && vf.toString() != "" {
		if _, err := time.Parse(layout, vf.toString()); err != nil {
			return "Informe uma data em formato EUA (ex: " + layout + ")"
		}
	}

	return ""
}

func vDateTime(vf vField) string {

	if vf.isString() && vf.toString() != "" {
		layout := vf.getValidation("datetime").Param
		if _, err := time.Parse(layout, vf.toString()); err != nil {
			return "Informe uma data no formato ex: " + layout
		}
	}

	return ""
}

func vLen(vf vField) string {

	var (
		msg              = "O parametro deve ter tamanho de %s caracteres"
		validationString = vf.getValidation("len").Param
	)

	length, err := strconv.Atoi(validationString)

	if err != nil {
		warning(err)
	}

	if vf.len() > 0 && vf.len() != length {
		if vf.isSlice() {
			return fmt.Sprintf("O campo deve ter %s elemmento(s)", validationString)
		}
		return fmt.Sprintf(msg, validationString)
	}

	return ""
}

func vMin(vf vField) string {

	var (
		msg           = "O valor minimo para o parametro é "
		msgNumeric    = msg + "%d"
		msgFloat64    = msg + "%g"
		validationInt int
		err           error
	)

	if vf.isString() || vf.isInt() || vf.isInt64() || vf.isSlice() {
		if validationInt, err = strconv.Atoi(vf.getValidation("min").Param); err != nil {
			warning(err)
			return ""
		}
	}

	if vf.isString() && vf.len() < validationInt {
		return fmt.Sprintf("O parametro deve ter no mínimo %d caracteres", validationInt)
	}

	if vf.isInt() && vf.toInt() < validationInt {
		return fmt.Sprintf(msgNumeric, validationInt)
	}

	if vf.isInt64() && vf.toInt64() < int64(validationInt) {
		return fmt.Sprintf(msgNumeric, validationInt)
	}

	if vf.isFloat64() {
		validationFloat64, err := strconv.ParseFloat(vf.getValidation("min").Param, 64)
		if err != nil {
			warning(err)
			return ""
		}
		if vf.toFloat64() < validationFloat64 {
			return fmt.Sprintf(msgFloat64, validationFloat64)
		}
	}

	if vf.isSlice() && (vf.IsSliceExample || vf.len() < validationInt) {
		return fmt.Sprintf("O campo deve no minimo %d elemmento(s)", validationInt)
	}

	return ""
}

func vMax(vf vField) string {

	var (
		msg        = "O valor máximo para o parametro é "
		msgNumeric = msg + "%d"
		msgFloat64 = msg + "%g"
		maxString  = vf.getValidation("max").Param
		maxInt     int
		err        error
	)

	if vf.isString() || vf.isInt() || vf.isInt64() || vf.isSlice() {
		if maxInt, err = strconv.Atoi(maxString); err != nil {
			warning(err)
			return ""
		}
	}

	if vf.isString() && vf.len() > maxInt {
		return fmt.Sprintf("O parametro deve ter no máximo %d caracteres", maxInt)
	}

	if vf.isInt() && vf.toInt() > maxInt {
		return fmt.Sprintf(msgNumeric, maxInt)
	}

	if vf.isInt64() && vf.toInt64() > int64(maxInt) {
		return fmt.Sprintf(msgNumeric, maxInt)
	}

	if vf.isFloat64() {
		validationFloat64, err := strconv.ParseFloat(maxString, 64)
		if err != nil {
			warning(err)
			return ""
		}
		if vf.toFloat64() > validationFloat64 {
			return fmt.Sprintf(msgFloat64, validationFloat64)
		}
	}

	if vf.isSlice() && vf.len() > maxInt {
		return fmt.Sprintf("O campo deve no máximo %d elemmento(s)", maxInt)
	}

	return ""
}

func vEnum(vf vField) string {

	enumsString := vf.getValidation("enum").Param
	enums := strings.Split(enumsString, "|")

	if (vf.isString() || vf.isInt() || vf.isInt64()) && !contains(vf.toString(), enums) {
		return fmt.Sprintf("Informe um dos valores: %s.", enumsString)
	}

	return ""
}

func vEmail(vf vField) string {
	if vf.isString() && vf.toString() != "" {
		rgx := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,6}$`)
		if !rgx.MatchString(vf.toString()) {
			return "Informe um e-mail válido"
		}
	}
	return ""
}

func vNumeric(vf vField) string {
	if vf.isString() && vf.toString() != "" {
		rgx := regexp.MustCompile(`^[0-9]+$`)
		if !rgx.MatchString(vf.toString()) {
			return "Informe um numero válido"
		}
	}
	return ""
}

func vCPF(vf vField) string {
	if vf.isString() && vf.toString() != "" && !cpf.Valido(vf.toString()) {
		return "Informe um CPF válido"
	}
	return ""
}

func vCNPJ(vf vField) string {
	if vf.isString() && vf.toString() != "" && !cnpj.Valido(vf.toString()) {
		return "Informe um CNPJ válido"
	}
	return ""
}

func vBPE(vf vField) string {

	if exp, ok := bpeRgx[vf.getValidation("bpe").Param]; ok {
		if !regexp.MustCompile(exp).MatchString(vf.toString()) {
			return fmt.Sprintf("Informe um valor conforme a expressão regular %s", exp)
		}
		return ""
	}

	if list, ok := bpeEnum[vf.getValidation("bpe").Param]; ok {
		if !contains(vf.toString(), list) {
			return fmt.Sprintf("Informe um dos valores: %s.", strings.Join(list, "|"))
		}
	}

	return ""
}

func vUF(vf vField) string {
	if !contains(vf.toString(), ufs) {
		return fmt.Sprintf("Informe uma das UFs: %s.", strings.Join(ufs, "|"))
	}
	return ""
}

func vUFWithEX(vf vField) string {
	if !contains(vf.toString(), ufsWithEX) {
		return fmt.Sprintf("Informe uma das UFs: %s.", strings.Join(ufsWithEX, "|"))
	}
	return ""
}

func vIPv4(vf vField) string {
	if value := vf.toString(); value != "" {
		if !regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`).MatchString(value) {
			return "Informe um ÌP válido"
		}
	}
	return ""
}
