package validator

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

const (
	defaultMessage         = "invalid parameters"
	regexCondition         = `([^<>!=]+)([<>!=]{0,2})(.*)`
	regexValidation        = `(?:(?:\()(.*)(?:\)))?([\w\d]+)(?:(?:{)([\w|]+)(?:}))?(?:(?:=)(.*))?`
	stringEmpty            = "''"
	tagErrMsg              = "errmsg"
	tagJSON                = "json"
	tagJSONIgnore          = "-"
	tagNoValidate          = "novalidate"
	tagQuery               = "query"
	tagRegex               = "regex"
	tagValidate            = "validate"
	tagValidationUnique    = "unique"
	tagValidationSeparator = ","
	tagGroupSeparator      = "|"
)

var errCondition = errors.New("the condition could not be processed")

var validator = vType{
	validations: map[string]vFunc{},
}

func contains(s string, slc []string) bool {
	for i := range slc {
		if s == slc[i] {
			return true
		}
	}
	return false
}

func warning(i ...interface{}) {
	log.Println("[LIB/VALIDATOR][WARNING]:", i)
}

func isFunc(s string) bool {
	return regexp.MustCompile(`(.+)\((.*)\)`).MatchString(s)
}

func isVar(s string) bool {

	l := len(s)

	if l <= 4 {
		return false
	}

	return s[0:2] == "{{" && s[l-2:] == "}}"
}

// antes de usar essa função verifique a string usando a função isVar
func removeVarDelimiters(s string) string {
	return s[2 : len(s)-2]
}

// checkOperationWithVFieldAndValue faz uma operação entre um vField e um valor
func checkOperationWithVFieldAndValue(vfield *vField, value, operator string) (bool, error) {

	// verifica se vai comparar com string vazia
	if value == stringEmpty {
		value = ""
	}

	switch {

	case vfield.isBool():
		valueBool, err := strconv.ParseBool(value)
		if err != nil {
			return false, err
		}
		switch operator {
		case "==":
			return vfield.toBool() == valueBool, nil
		case "!=":
			return vfield.toBool() != valueBool, nil
		}

	case vfield.isFloat64():
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false, err
		}
		switch operator {
		case "==":
			return vfield.toFloat64() == valueFloat64, nil
		case "!=":
			return vfield.toFloat64() != valueFloat64, nil
		case ">":
			return vfield.toFloat64() > valueFloat64, nil
		case "<":
			return vfield.toFloat64() < valueFloat64, nil
		case ">=":
			return vfield.toFloat64() >= valueFloat64, nil
		case "<=":
			return vfield.toFloat64() <= valueFloat64, nil
		}

	case vfield.isInt():
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return false, err
		}
		switch operator {
		case "==":
			return vfield.toInt() == valueInt, nil
		case "!=":
			return vfield.toInt() != valueInt, nil
		case ">":
			return vfield.toInt() > valueInt, nil
		case "<":
			return vfield.toInt() < valueInt, nil
		case ">=":
			return vfield.toInt() >= valueInt, nil
		case "<=":
			return vfield.toInt() <= valueInt, nil
		}
	case vfield.isInt64():
		valueInt64, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false, err
		}
		switch operator {
		case "==":
			return vfield.toInt64() == valueInt64, nil
		case "!=":
			return vfield.toInt64() != valueInt64, nil
		case ">":
			return vfield.toInt64() > valueInt64, nil
		case "<":
			return vfield.toInt64() < valueInt64, nil
		case ">=":
			return vfield.toInt64() >= valueInt64, nil
		case "<=":
			return vfield.toInt64() <= valueInt64, nil
		}

	case vfield.isString():
		switch operator {
		case "==":
			return vfield.toString() == value, nil
		case "!=":
			return vfield.toString() != value, nil
		}
	}

	return false, errCondition
}

// checkOperationWithVFieldAndVfield faz um operação entre dois vFields
func checkOperationWithVFieldAndVfield(vfield1, vfield2 *vField, operator string) (bool, error) {

	switch {
	case vfield1.isBool() && vfield2.isBool():
		switch operator {
		case "==":
			return vfield1.toBool() == vfield2.toBool(), nil
		case "!=":
			return vfield1.toBool() != vfield2.toBool(), nil
		}

	case vfield1.isFloat64() && vfield2.isFloat64():
		switch operator {
		case "==":
			return vfield1.toFloat64() == vfield2.toFloat64(), nil
		case "!=":
			return vfield1.toFloat64() != vfield2.toFloat64(), nil
		case ">":
			return vfield1.toFloat64() > vfield2.toFloat64(), nil
		case "<":
			return vfield1.toFloat64() < vfield2.toFloat64(), nil
		case ">=":
			return vfield1.toFloat64() >= vfield2.toFloat64(), nil
		case "<=":
			return vfield1.toFloat64() <= vfield2.toFloat64(), nil
		}

	case vfield1.isInt() && vfield2.isInt():
		switch operator {
		case "==":
			return vfield1.toInt() == vfield2.toInt(), nil
		case "!=":
			return vfield1.toInt() != vfield2.toInt(), nil
		case ">":
			return vfield1.toInt() > vfield2.toInt(), nil
		case "<":
			return vfield1.toInt() < vfield2.toInt(), nil
		case ">=":
			return vfield1.toInt() >= vfield2.toInt(), nil
		case "<=":
			return vfield1.toInt() <= vfield2.toInt(), nil
		}

	case vfield1.isString() && vfield2.isString():
		switch operator {
		case "==":
			return vfield1.toString() == vfield2.toString(), nil
		case "!=":
			return vfield1.toString() != vfield2.toString(), nil
		}
	}
	return false, errCondition
}

func prepareAndCheckCondition(vf *vField, condition string) (bool, error) {
	// regex para separar o campo1, operador e campo2 em grupos.
	// rgx := regexp.MustCompile(regexCondition)
	groups := regexp.MustCompile(regexCondition).FindStringSubmatch(condition)

	if strings.ToUpper(condition) == "TRUE" {
		return true, nil
	}

	if strings.ToUpper(condition) == "FALSE" {
		return false, nil
	}

	if isFunc(groups[1]) {
		groupsFunc := regexp.MustCompile(`(.+)\((.*)\)`).FindStringSubmatch(groups[1])
		switch groupsFunc[1] {
		case "isSuccessor":
			return vf.isSuccessor(groupsFunc[2]), nil
		case "isNil":
			return vf.isNil(groupsFunc[2])
		}
	} else {

		var (
			vfield1, vfield2 *vField
			err              error
		)

		// busca o primeiro campo (se for uma variável)
		if isVar(groups[1]) {
			if vfield1, err = vf.find(groups[1]); err != nil {
				return false, err
			}
		}

		// busca o segundo campo (se for uma variável)
		if isVar(groups[3]) {
			if vfield2, err = vf.find(groups[3]); err != nil {
				return false, err
			}
		}

		// verifica a condição de validação
		return checkCondition(vfield1, vfield2, groups[1], groups[3], groups[2])
	}
	return false, errors.New("a condição informada é inválida")
}

// checkCondition verifica um possível condição
func checkCondition(vfield1, vfield2 *vField, value1, value2, operator string) (bool, error) {

	switch {
	case vfield1 == nil && vfield2 == nil:
		return false, errors.New("enter a vfield")
	case vfield1 != nil && vfield2 != nil:
		return checkOperationWithVFieldAndVfield(vfield1, vfield2, operator)
	case vfield1 != nil && value2 != "":
		return checkOperationWithVFieldAndValue(vfield1, value2, operator)
	case vfield2 != nil && value1 != "":
		return checkOperationWithVFieldAndValue(vfield2, value1, operator)
	case vfield1 != nil && value2 == "" && operator == "":
		if vfield1.isBool() {
			return vfield1.toBool(), nil
		}
	}
	return false, errCondition
}

// toVField gera a estrutura de vFields "árvore"
func toVField(previous *vField, a interface{}, tags *reflect.StructTag, name string, isNil bool) *vField {

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	// remove o ponteiro
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// remove o ponteiro
	if v.Kind() == reflect.Ptr {
		// verifica se é nil antes para armazenar
		if !isNil {
			isNil = v.IsNil()
		}
		v = v.Elem()
	}

	vf := &vField{
		Name:     name,
		Kind:     t.Kind(),
		Previous: previous,
		IsNil:    isNil,
	}

	if tags != nil {
		tagJSONSplited := strings.Split(tags.Get(tagJSON), ",")
		vf.TagJSON = tagJSONSplited[0]
		vf.TagQuery = tags.Get(tagQuery)
		vf.TagRegex = tags.Get(tagRegex)
		vf.TagErrMsg = tags.Get(tagErrMsg)
		vf.TagNoValidate = tags.Get(tagNoValidate)
	}

	_, isTimeEUA := a.(gotoolboxtime.TimeEUA)
	_, isTimeCard := a.(gotoolboxtime.TimeCard)
	_, isTimestamp := a.(gotoolboxtime.Timestamp)

	switch {
	case isTimeEUA, isTimeCard, isTimestamp:
		vf.Kind = reflect.Invalid
		vf.Interface = a

	// quando é struct faz recursão em todos os campos e para a função
	case t.Kind() == reflect.Struct:

		vfs := []*vField{}

		for i := 0; i < t.NumField(); i++ {

			// ignora os campos onde não podemos obter a interface
			if !v.Field(i).CanInterface() {
				continue
			}

			fieldTags := t.Field(i).Tag

			if v.Field(i).Kind() == reflect.Ptr && v.Field(i).IsNil() {
				vfs = append(vfs, toVField(vf, reflect.New(v.Field(i).Type().Elem()).Elem().Interface(), &fieldTags, t.Field(i).Name, true))
				continue
			}

			vfs = append(vfs, toVField(vf, v.Field(i).Interface(), &fieldTags, t.Field(i).Name, false))
		}

		vf.Interface = vfs
		return vf

	// quando é um slice faz recursão em todos o elementos e salva como o Interface
	// dele (passam as ser os "filhos")
	case t.Kind() == reflect.Slice:
		vfs := []*vField{}

		if v.Len() != 0 {
			for i := 0; i < v.Len(); i++ {
				vfs = append(vfs, toVField(vf, v.Index(i).Interface(), nil, "", false))
			}
		} else {
			vf.IsSliceExample = true
			// instância um elemento do slice para testar as validações
			vfs = append(vfs, toVField(vf, reflect.New(v.Type().Elem()).Elem().Interface(), nil, "", false))
		}

		vf.Interface = vfs

	default:
		// define o interface com o próprio valor
		vf.Interface = a
	}

	if tags != nil {

		tagv := tags.Get(tagValidate)

		if tagv != "" {

			// separa todas a validações informadas
			v := strings.Split(tagv, tagValidationSeparator)

			// regex para separar a condição, nome e parametro da validação.
			rgx := regexp.MustCompile(regexValidation)

			// aplica a regex em todas as validação informadas e salva no vField
			for i := range v {
				ss := rgx.FindStringSubmatch(v[i])
				groups := []string{}

				if ss[3] != "" {
					groups = strings.Split(ss[3], tagGroupSeparator)
				}

				vf.setValidation(vValidation{
					Condition: ss[1],
					Name:      ss[2],
					Groups:    groups,
					Param:     ss[4],
				})
			}
		}
	}

	return vf
}

// validateVField faz a validação dos vFields (verificando as condições)
func validateVField(vf *vField, group string) interface{} {

	// pula quando o coringa é informado
	if vf.TagJSON == tagJSONIgnore {
		return nil
	}

	// check regex
	if !vf.isStruct() && vf.TagRegex != "" {
		var (
			condition = false
			err       error
		)

		if vf.TagNoRegex != "" {
			condition, err = prepareAndCheckCondition(vf, vf.TagNoRegex)
		}

		if err == nil && !condition {
			valueString := vf.toString()
			if len(valueString) > 0 && !regexp.MustCompile(vf.TagRegex).MatchString(valueString) {
				if vf.TagErrMsg != "" {
					return vf.TagErrMsg
				}
				return fmt.Sprintf(`Informe o conteúdo conforme a expressão regular: %s`, vf.TagRegex)
			}
		}
	}

	if vf.TagNoValidate != "" {
		condition, err := prepareAndCheckCondition(vf, vf.TagNoValidate)
		if err != nil || condition {
			return nil
		}
	}

	// quando é uma struct faz a recursão em todos os filhos
	// (campos da struct, outros vFields)
	if vf.isStruct() {
		r := map[string]interface{}{}
		vfs := vf.toSlice()
		for i := range vfs {
			if rr := validateVField(vfs[i], group); rr != nil {
				r[vfs[i].getNameForIndex()] = rr
			}
		}
		if len(r) == 0 {
			return nil
		}
		return r
	}

	// faz as validações
	for j := range vf.Validations {

		// verifica se a validação está no grupo informado
		if len(vf.Validations[j].Groups) > 0 && (group == "" || !contains(group, vf.Validations[j].Groups)) {
			continue
		}

		// verifica a condição informada para uma validação
		if vf.Validations[j].Condition != "" {

			condition, err := prepareAndCheckCondition(vf, vf.Validations[j].Condition)

			if err != nil {
				warning(vf.Name, vf.Kind, err)
			}

			// pula a validação se a condição for false
			if !condition {
				continue
			}
		}

		// verifica se todos os campos do slice são unicos
		if vf.Validations[j].Name == tagValidationUnique {
			if vf.Previous.Previous != nil && vf.Previous.Previous.Kind == reflect.Slice {
				vfs := vf.Previous.Previous.toSlice()
				var encontrado int
				for i := range vfs {
					if vfs[i].Kind == reflect.Struct {
						vfsSctruct := vfs[i].toSlice()
						for j := range vfsSctruct {
							if vf.Name == vfsSctruct[j].Name && vf.Interface == vfsSctruct[j].Interface {
								encontrado++
								if encontrado == 2 {
									return "Esse item não pode ser repetido"
								}
							}
						}
					}
				}
			}
		}

		// busca a validação para executar a função
		if vfunc, ok := validator.getValidation(vf.Validations[j].Name); ok {
			if s := vfunc(*vf); s != "" {
				if vf.TagErrMsg != "" {
					return vf.TagErrMsg
				}
				return s
			}
		}
	}

	// quando é um slice faz a recursão em todos os filho (outros vFields)
	if vf.Kind == reflect.Slice {
		r := map[int]interface{}{}
		vfs := vf.toSlice()
		for i := range vfs {
			if rr := validateVField(vfs[i], group); rr != nil {
				r[i] = rr
			}
		}
		if len(r) == 0 {
			return nil
		}
		return r
	}

	return nil
}

// Validate faz todo o processo de validação conforme os parâmetros informados:
// a 					-> struct ou slice qualquer
// params[0]  -> grupo de validação ex: required{1|2|3}
func Validate(a interface{}, options ...string) *VError {

	var grupo string

	if len(options) > 0 {
		grupo = options[0]
	}

	if details := validateVField(toVField(nil, a, nil, "", false), grupo); details != nil {
		return &VError{
			Message: defaultMessage,
			Details: details,
		}
	}

	return nil
}
