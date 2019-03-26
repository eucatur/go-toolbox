package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type vField struct {
	Name           string
	TagJSON        string
	TagQuery       string
	TagRegex       string
	TagErrMsg      string
	TagNoValidate  string
	TagNoRegex     string
	Kind           reflect.Kind
	Previous       *vField
	Interface      interface{}
	Validations    []vValidation
	IsSliceExample bool
	IsNil          bool
}

func (v vField) len() int {
	if v.isSlice() {
		if v.IsSliceExample {
			return 0
		}
		return len(v.toSlice())
	}
	return len(v.toString())
}

func (v vField) getNameForIndex() string {

	if v.TagJSON != "" {
		return v.TagJSON
	}

	if v.TagQuery != "" {
		return v.TagQuery
	}

	return v.Name
}

func (v *vField) findInInterface(s string) (*vField, error) {

	// regex para separar a string no primeiro "."
	// exemplo: "venda.vendedor.nome"
	// groups[1] = "venda"
	// groups[2] = "vendedor.nome"
	rgx := regexp.MustCompile(`(\w+)(?:\.?)(.*)`)
	groups := rgx.FindStringSubmatch(s)

	// if v.Interface == nil {
	// 	return nil, errors.New("there is not any next")
	// }

	if v.isSlice() || v.isStruct() {
		vfs := v.toSlice()
		for i := range vfs {
			// identifica o campo procurado
			if vfs[i].getNameForIndex() == groups[1] {
				// verifica se a busca se extende ao
				// proximo elemento (filho) por recursão
				if groups[2] != "" {
					return vfs[i].findInInterface(groups[2])
				}
				return vfs[i], nil
			}
		}
		return nil, errors.New("not found")
	}

	return nil, errors.New("type mismatch")
}

func (v *vField) findInPrevious(s string) (*vField, error) {

	if v.Previous == nil {
		return nil, errors.New("there is no previous")
	}

	// quando o anterior é um slice ele é pulado via recursão
	if v.Previous.isSlice() {
		return v.Previous.findInPrevious(s)
	}

	// quando o primeiro caracter é um "." busca no anterior via recursão
	if s[0:1] == "." {
		return v.Previous.findInPrevious(s[1:])
	}

	return v.Previous, nil
}

func (v *vField) find(s string) (*vField, error) {

	s = removeVarDelimiters(s)

	// retrocede cada "." até o primeiro elemento
	vFieldPrevious, err := v.findInPrevious(s)

	if err != nil {
		return nil, err
	}

	rgx := regexp.MustCompile(`(?:\.?)(.*)`)
	groups := rgx.FindStringSubmatch(s)

	// caso mudar a expressão regular acima essa validação precisar ser usada
	// if len(groups) != 2 {
	// 	return nil, errors.New("could not search")
	// }

	// avança cada "." até o elemento desejado
	vFieldNext, err := vFieldPrevious.findInInterface(groups[1])

	if err != nil {
		return nil, err
	}

	return vFieldNext, nil
}

func (v vField) isSuccessor(s string) bool {
	if v.Previous == nil {
		return false
	}
	if v.getNameForIndex() == s {
		return true
	}
	return v.Previous.isSuccessor(s)
}

func (v vField) isNil(s string) (bool, error) {
	var (
		vf  *vField
		err error
	)

	if vf, err = v.find("{{" + s + "}}"); err != nil {
		return false, err
	}

	if vf == nil {
		return false, fmt.Errorf("func isNil: elemento %s não encontrado", s)
	}

	return vf.IsNil, nil
}

func (v *vField) setValidation(vv vValidation) {
	v.Validations = append(v.Validations, vv)
}

func (v vField) getValidation(name string) vValidation {
	for i := range v.Validations {
		if v.Validations[i].Name == name {
			return v.Validations[i]
		}
	}
	return vValidation{}
}

func (v vField) toBool() bool {
	var (
		value, ok bool
		err       error
	)

	if value, ok = v.Interface.(bool); !ok {
		if value, err = strconv.ParseBool(v.toString()); err != nil {
			panic(err)
		}
	}

	return value
}

func (v vField) toInt() int {

	var (
		value int
		ok    bool
		err   error
	)

	if value, ok = v.Interface.(int); !ok {
		if value, err = strconv.Atoi(v.toString()); err != nil {
			panic(err)
		}
	}

	return value
}

func (v vField) toInt64() int64 {
	var (
		value int64
		ok    bool
		err   error
	)

	if value, ok = v.Interface.(int64); !ok {
		if value, err = strconv.ParseInt(v.toString(), 10, 64); err != nil {
			panic(err)
		}
	}

	return value
}

func (v vField) toSlice() []*vField {
	return v.Interface.([]*vField)
}

func (v vField) toString() string {
	return fmt.Sprint(v.Interface)
}

func (v vField) toFloat64() float64 {
	var (
		value float64
		ok    bool
		err   error
	)

	if value, ok = v.Interface.(float64); !ok {
		if value, err = strconv.ParseFloat(v.toString(), 64); err != nil {
			panic(err)
		}
	}

	return value
}

func (v vField) isBool() bool {
	return v.Kind == reflect.Bool
}

func (v vField) isInt() bool {
	return v.Kind == reflect.Int
}

func (v vField) isInt64() bool {
	return v.Kind == reflect.Int64
}

func (v vField) isSlice() bool {
	return v.Kind == reflect.Slice
}

func (v vField) isString() bool {
	return v.Kind == reflect.String
}

func (v vField) isStruct() bool {
	return v.Kind == reflect.Struct
}

func (v vField) isFloat64() bool {
	return v.Kind == reflect.Float64
}
