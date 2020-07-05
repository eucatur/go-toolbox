package validator

import (
	"testing"
	"time"

	toolboxtime "github.com/eucatur/go-toolbox/time"
	"github.com/stretchr/testify/assert"
)

type (
	MyBool    bool
	MyInt     int
	MyInt64   int64
	MyFloat64 float64
)

var (
	expectedNil *VError
	expected    = &VError{
		Message: defaultMessage,
	}
)

func Test_fieldinterface(t *testing.T) {
	type ErrorTest struct {
		Code    interface{} `json:"Code"`
		Message string      `json:"Message"`
	}

	type Strct struct {
		Errors *[]ErrorTest `json:"errors" validate:"required"`
	}

	var strctFieldInterface Strct

	expected.Details = map[string]interface{}{
		"errors": "Informe ao menos um elemento para esse campo (array)",
	}

	test(t, expected, strctFieldInterface)
}

func Test_validator(t *testing.T) {
	// teste de funções
	vfield := &vField{}

	// checkCondition
	checkCondition(vfield, nil, "", "", "==")

	// getValidation
	vfield.getValidation("")

	// SetFuncMarshalJSON
	SetFuncMarshalJSON(funcMarshalJSON)

	// funcMarshalJSON
	funcMarshalJSON(&VError{})

	// VError
	vError := &VError{}
	_ = vError.Error()
	vError.MarshalJSON()
}

func Test_msg(t *testing.T) {
	structMsg := struct {
		Name string `json:"name" validate:"msg=This message will always be shown"`
	}{
		Name: "Jon",
	}

	expected.Details = map[string]interface{}{
		"name": "This message will always be shown",
	}

	test(t, expected, structMsg)
}

func Test_required(t *testing.T) {
	// required correto
	structRequired := struct {
		Name      string                `json:"name" validate:"required" errmsg:"Enter the name"`
		Names     []string              `json:"names" validate:"required" errmsg:"Enter the name or more"`
		TimeEUA   toolboxtime.TimeEUA   `json:"time_eua" validate:"required" errmsg:"Enter a timeEUA"`
		TimeCard  toolboxtime.TimeCard  `json:"time_card" validate:"required" errmsg:"Enter a timeCard"`
		Timestamp toolboxtime.Timestamp `json:"timestamp" validate:"required" errmsg:"Enter a timestamp"`
	}{
		Name:      "Jon",
		Names:     []string{"Jon", "Sam"},
		TimeEUA:   toolboxtime.TimeEUA{Time: time.Now()},
		TimeCard:  toolboxtime.TimeCard{Time: time.Now()},
		Timestamp: toolboxtime.Timestamp{Time: time.Now()},
	}

	test(t, expectedNil, structRequired)

	// required incorreto
	structRequired.Name = ""
	structRequired.Names = []string{}
	structRequired.TimeEUA = toolboxtime.TimeEUA{}
	structRequired.TimeCard = toolboxtime.TimeCard{}
	structRequired.Timestamp = toolboxtime.Timestamp{}

	expected.Details = map[string]interface{}{
		"name":      "Enter the name",
		"names":     "Enter the name or more",
		"time_eua":  "Enter a timeEUA",
		"time_card": "Enter a timeCard",
		"timestamp": "Enter a timestamp",
	}

	test(t, expected, structRequired)
}

func Test_uint(t *testing.T) {
	// uint correto
	structUInt := struct {
		IDInt   int     `json:"id_int" validate:"uint" errmsg:"Enter the id (uint)"`
		IDInt64 MyInt64 `json:"id_int64" validate:"uint" errmsg:"Enter the id (uint)"`
	}{
		IDInt:   10,
		IDInt64: 10,
	}

	test(t, expectedNil, structUInt)

	// uint incorreto
	structUInt.IDInt = 0
	structUInt.IDInt64 = 0

	expected.Details = map[string]interface{}{
		"id_int":   "Enter the id (uint)",
		"id_int64": "Enter the id (uint)",
	}

	test(t, expected, structUInt)
}

func Test_datebr(t *testing.T) {
	// datebr correta
	structDateBR := struct {
		DateBR      string `json:"datebr" validate:"datebr" errmsg:"Enter a valid date"`
		DateBREmpty string `json:"datebr_empty" validate:"datebr" errmsg:"Enter a valid date"`
	}{
		DateBR: "14/08/2017",
	}

	test(t, expectedNil, structDateBR)

	// datebr incorreta
	structDateBR.DateBR = "1/2/17"

	expected.Details = map[string]interface{}{
		"datebr": "Enter a valid date",
	}

	test(t, expected, structDateBR)
}

func Test_dateeua(t *testing.T) {
	// dateeua correta
	structDateEUA := struct {
		DateEUA      string `json:"dateeua" validate:"dateeua" errmsg:"Enter a valid date"`
		DateEUAEmpty string `json:"dateeua_empty" validate:"dateeua" errmsg:"Enter a valid date"`
	}{
		DateEUA: "2017-09-13",
	}

	test(t, expectedNil, structDateEUA)

	// dateeua incorreta
	structDateEUA.DateEUA = "17-09-13"

	expected.Details = map[string]interface{}{
		"dateeua": "Enter a valid date",
	}

	test(t, expected, structDateEUA)
}

func Test_datetime(t *testing.T) {
	// datetime correta
	structDateTime := struct {
		DateTime string `json:"datetime" validate:"datetime=2006-01-02 15:04" errmsg:"Enter a valid date"`
	}{
		DateTime: "2017-09-13 23:59",
	}

	test(t, expectedNil, structDateTime)

	// datetime incorreta
	structDateTime.DateTime = "17-09-13 29:59"

	expected.Details = map[string]interface{}{
		"datetime": "Enter a valid date",
	}

	test(t, expected, structDateTime)
}

func Test_len(t *testing.T) {
	// len correto
	structLen := struct {
		LenString      string   `json:"lenstring" validate:"len=10" errmsg:"Enter ten characters"`
		LenStringAcute string   `json:"lenstringacute" validate:"len=9" errmsg:"Enter nine characters"`
		LenStringEmpty string   `json:"lenstringempty" validate:"len=10" errmsg:"Enter ten characters"`
		LenInt         int      `json:"lenint" validate:"len=5" errmsg:"Enter five characters"`
		LenInt64       int      `json:"lenint64" validate:"len=5" errmsg:"Enter five characters"`
		LenSlice       []string `json:"lenslice" validate:"len=3" errmsg:"Enter three characters"`
		LenInvalid     string   `json:"leninvalid" validate:"len=len_test_invalid_int" errmsg:"Message"`
	}{
		LenString:      "Tecnologia",
		LenStringAcute: "Tecnólogo",
		LenInt:         12345,
		LenInt64:       12345,
		LenSlice:       []string{"A", "B", "C"},
	}

	test(t, expectedNil, structLen)

	// len incorreto
	structLen.LenString = "Tecnologias"
	structLen.LenStringAcute = "Tecnólogos"
	structLen.LenInt = 123456
	structLen.LenSlice = []string{"A", "B", "C", "D"}

	expected.Details = map[string]interface{}{
		"lenstring":      "Enter ten characters",
		"lenstringacute": "Enter nine characters",
		"lenint":         "Enter five characters",
		"lenslice":       "Enter three characters",
	}

	test(t, expected, structLen)
}

func Test_min(t *testing.T) {
	// min correto
	structMin := struct {
		MinString         string   `json:"minstring" validate:"min=10" errmsg:"Enter at least ten characters"`
		MinInt            int      `json:"minint" validate:"min=5" errmsg:"The minimum value for the field is five"`
		MinInt64          int64    `json:"minint64" validate:"min=5" errmsg:"The minimum value for the field is five"`
		MinFloat64        float64  `json:"minfloat64" validate:"min=5.5" errmsg:"The minimum value for the field is 5.5"`
		MinSlice          []string `json:"minslice" validate:"min=3" errmsg:"Report at least three elements"`
		MinInvalid        string   `json:"mininvalid" validate:"min=min_invalid_int" errmsg:"Message"`
		MinInvalidFloat64 float64  `json:"mininvalidfloat64" validate:"min=min_invalid_float64" errmsg:"Message"`
	}{
		MinString:  "Tecnologia",
		MinInt:     5,
		MinInt64:   5,
		MinFloat64: 5.5,
		MinSlice:   []string{"A", "B", "C"},
	}

	test(t, expectedNil, structMin)

	// min incorreto
	structMin.MinString = "Tec"
	structMin.MinInt = 4
	structMin.MinInt64 = 4
	structMin.MinFloat64 = 5.4
	structMin.MinSlice = []string{"A", "B"}

	expected.Details = map[string]interface{}{
		"minstring":  "Enter at least ten characters",
		"minint":     "The minimum value for the field is five",
		"minint64":   "The minimum value for the field is five",
		"minfloat64": "The minimum value for the field is 5.5",
		"minslice":   "Report at least three elements",
	}

	test(t, expected, structMin)
}

func Test_max(t *testing.T) {
	// max correto
	structMax := struct {
		MaxString         string   `json:"maxstring" validate:"max=10" errmsg:"Enter a maximum of ten characters"`
		MaxStringAcute    string   `json:"maxstringacute" validate:"max=9" errmsg:"Enter a maximum of nine characters"`
		MaxInt            int      `json:"maxint" validate:"max=5" errmsg:"The maximum value for the field is five"`
		MaxInt64          int64    `json:"maxint64" validate:"max=5" errmsg:"The maximum value for the field is five"`
		MaxFloat64        float64  `json:"maxfloat64" validate:"max=5.5" errmsg:"The maximum value for the field is 5.5"`
		MaxSlice          []string `json:"maxslice" validate:"max=3" errmsg:"Enter at most three elements"`
		MaxInvalid        string   `json:"maxinvalid" validate:"max=max_invalid_int" errmsg:"Message"`
		MaxInvalidFloat64 float64  `json:"maxinvalidfloat64" validate:"max=max_invalid_float64" errmsg:"Message"`
	}{
		MaxString:      "Tecnologia",
		MaxStringAcute: "Tecnólogo",
		MaxInt:         5,
		MaxInt64:       5,
		MaxFloat64:     5.5,
		MaxSlice:       []string{"A", "B", "C"},
	}

	test(t, expectedNil, structMax)

	// max incorreto
	structMax.MaxString = "Tecnologias"
	structMax.MaxStringAcute = "Tecnólogos"
	structMax.MaxInt = 6
	structMax.MaxInt64 = 6
	structMax.MaxFloat64 = 5.6
	structMax.MaxSlice = []string{"A", "B", "C", "D"}

	expected.Details = map[string]interface{}{
		"maxstring":      "Enter a maximum of ten characters",
		"maxstringacute": "Enter a maximum of nine characters",
		"maxint":     "The maximum value for the field is five",
		"maxint64":   "The maximum value for the field is five",
		"maxfloat64": "The maximum value for the field is 5.5",
		"maxslice":   "Enter at most three elements",
	}

	test(t, expected, structMax)
}

func Test_enum(t *testing.T) {
	// enum correta
	structEnum := struct {
		EnumString string `json:"enumstring" validate:"enum=A|B|C" errmsg:"Enter one of the items: A, B or C"`
		EnumInt    int    `json:"enumint" validate:"enum=1|2|3" errmsg:"Enter one of the items: 1, 2 or 3"`
		EnumInt64  int64  `json:"enumint64" validate:"enum=1|2|3" errmsg:"Enter one of the items: 1, 2 or 3"`
	}{
		EnumString: "A",
		EnumInt:    1,
		EnumInt64:  2,
	}

	test(t, expectedNil, structEnum)

	structEnum.EnumString = "B"
	structEnum.EnumInt = 2
	structEnum.EnumInt64 = 2

	test(t, expectedNil, structEnum)

	structEnum.EnumString = "C"
	structEnum.EnumInt = 3
	structEnum.EnumInt64 = 3

	test(t, expectedNil, structEnum)

	// enum incorreta
	structEnum.EnumString = "D"
	structEnum.EnumInt = 4
	structEnum.EnumInt64 = 4

	expected.Details = map[string]interface{}{
		"enumstring": "Enter one of the items: A, B or C",
		"enumint":    "Enter one of the items: 1, 2 or 3",
		"enumint64":  "Enter one of the items: 1, 2 or 3",
	}

	test(t, expected, structEnum)
}

func Test_email(t *testing.T) {
	// email correto
	structEmail := struct {
		Email1 string `json:"email1" validate:"email" errmsg:"Enter a valid email address"`
		Email2 string `json:"email2" validate:"email" errmsg:"Enter a valid email address"`
		Email3 string `json:"email3" validate:"email" errmsg:"Enter a valid email address"`
		Email4 string `json:"email4" validate:"email" errmsg:"Enter a valid email address"`
		Email5 string `json:"email5" validate:"email" errmsg:"Enter a valid email address"`
	}{
		Email1: "eucatur@eucatur.com",
		Email2: "eucatur@eucatur.com.br",
		Email3: "eucatur_1@eucatur.com.br",
		Email4: "eucatur@eucatur-1.com.br",
		Email5: "EUCATUR@EUCATUR.COM.BR",
	}

	test(t, expectedNil, structEmail)

	// email incorreto
	structEmail.Email1 = "eucatur@"
	structEmail.Email2 = "@eucatur.com.br"
	structEmail.Email3 = "eucatur@eucatur"
	structEmail.Email4 = "eucatur@eucatur."

	expected.Details = map[string]interface{}{
		"email1": "Enter a valid email address",
		"email2": "Enter a valid email address",
		"email3": "Enter a valid email address",
		"email4": "Enter a valid email address",
	}

	test(t, expected, structEmail)
}

func Test_numeric(t *testing.T) {
	// numeric correto
	structNumeric := struct {
		Numeric string `json:"numeric" validate:"numeric" errmsg:"Enter a numeric value"`
	}{
		Numeric: "00001",
	}

	test(t, expectedNil, structNumeric)

	// numeric incorreto
	structNumeric.Numeric = "a0001"

	expected.Details = map[string]interface{}{
		"numeric": "Enter a numeric value",
	}

	test(t, expected, structNumeric)
}

func Test_cpf(t *testing.T) {
	// cpf correto
	structCPF := struct {
		CPF1 string `json:"cpf1" validate:"cpf" errmsg:"Enter a valid CPF"`
		CPF2 string `json:"cpf2" validate:"cpf" errmsg:"Enter a valid CPF"`
		CPF3 string `json:"cpf3" validate:"cpf" errmsg:"Enter a valid CPF"`
	}{
		CPF1: "782.854.588-65",
		CPF2: "78285458865",
		CPF3: "124.429.281-83",
	}

	test(t, expectedNil, structCPF)

	// cpf incorreto
	structCPF.CPF1 = "111.111.111-11"
	structCPF.CPF2 = "11111111111"
	structCPF.CPF3 = "a12442928183"

	expected.Details = map[string]interface{}{
		"cpf1": "Enter a valid CPF",
		"cpf2": "Enter a valid CPF",
		"cpf3": "Enter a valid CPF",
	}

	test(t, expected, structCPF)
}

func Test_cnpj(t *testing.T) {
	// cnpj correto
	structCNPJ := struct {
		CNPJ1 string `json:"cnpj1" validate:"cnpj" errmsg:"Enter a valid CNPJ"`
		CNPJ2 string `json:"cnpj2" validate:"cnpj" errmsg:"Enter a valid CNPJ"`
		CNPJ3 string `json:"cnpj3" validate:"cnpj" errmsg:"Enter a valid CNPJ"`
	}{
		CNPJ1: "64.767.833/0001-65",
		CNPJ2: "64767833000165",
		CNPJ3: "19.673.538/0001-95",
	}

	test(t, expectedNil, structCNPJ)

	// cnpj incorreto
	structCNPJ.CNPJ1 = "11.111.111/111-11"
	structCNPJ.CNPJ2 = "11111111111111"
	structCNPJ.CNPJ3 = "a19673538000195"

	expected.Details = map[string]interface{}{
		"cnpj1": "Enter a valid CNPJ",
		"cnpj2": "Enter a valid CNPJ",
		"cnpj3": "Enter a valid CNPJ",
	}

	test(t, expected, structCNPJ)
}

func Test_bpe(t *testing.T) {
	// bpe correto
	structBpe := struct {
		Regex string `json:"regex" validate:"bpe=ER35" errmsg:"Message"`
		Enum  string `json:"enum" validate:"bpe=D9" errmsg:"Message"`
	}{
		Regex: "regex",
		Enum:  "00",
	}

	test(t, expectedNil, structBpe)

	// bpe incorreto
	structBpe.Regex = " "
	structBpe.Enum = ""

	expected.Details = map[string]interface{}{
		"regex": "Message",
		"enum":  "Message",
	}

	test(t, expected, structBpe)
}

func Test_uf(t *testing.T) {
	// uf correto
	structUF := struct {
		UF       string `json:"uf" validate:"uf" errmsg:"Message"`
		UFWithEX string `json:"uf_with_ex" validate:"ufwithex" errmsg:"Message"`
	}{
		UF:       "RO",
		UFWithEX: "EX",
	}

	test(t, expectedNil, structUF)

	// uf incorreto
	structUF.UF = "EX"
	structUF.UFWithEX = ""

	expected.Details = map[string]interface{}{
		"uf":         "Message",
		"uf_with_ex": "Message",
	}

	test(t, expected, structUF)
}

func Test_ip(t *testing.T) {
	// ip correto
	structIP := struct {
		IPV4 string `json:"ipv4" validate:"ipv4" errmsg:"Message"`
	}{
		IPV4: "192.168.10.1",
	}

	test(t, expectedNil, structIP)

	// uf incorreto
	structIP.IPV4 = "192.168.10.256"

	expected.Details = map[string]interface{}{
		"ipv4": "Message",
	}

	test(t, expected, structIP)
}

func Test_func(t *testing.T) {
	// função correto
	type Address struct {
		PublicPlace string `json:"public_place" validate:"(isSuccessor(seller))required" errmsg:"Message"`
	}

	type Seller struct {
		Address `json:"address"`
	}

	type Buyer struct {
		Address `json:"address"`
	}

	type Sale struct {
		Seller `json:"seller"`
		Buyer  `json:"buyer"`
	}

	structFunc := Sale{
		Seller: Seller{
			Address: Address{
				PublicPlace: "Public place",
			},
		},
	}

	test(t, expectedNil, structFunc)

	// função incorreto
	structFunc.Seller.Address.PublicPlace = ""

	expected.Details = map[string]interface{}{
		"seller": map[string]interface{}{
			"address": map[string]interface{}{
				"public_place": "Message",
			},
		},
	}

	test(t, expected, structFunc)
}

func Test_conditions(t *testing.T) {
	// condições
	structCondition := struct {
		Nome1 string `json:"nome1"`
		Nome2 string `json:"nome2"`
		Int1  int    `json:"int1"`
		Int2  int    `json:"int2"`
		Int3  int    `json:"int3"`
		Int64 int64  `json:"int64"`

		StringVazia      string `json:"string_vazia" validate:"({{nome1}}=='')msg=Nome1 is empty"`
		StringPreenchida string `json:"string_preenchida" validate:"({{nome2}}!='')msg=Nome2 is filled in"`
		StringDiferente  string `json:"string_diferente" validate:"({{nome1}}!=Eucatur)msg=Nome1 is different from Eucatur"`
		StringIgual      string `json:"string_igual" validate:"({{nome2}}==Eucatur)msg=Nome2 is equal from Eucatur"`

		Diferente  string `json:"diferente" validate:"({{int1}}!={{int2}})msg=Int1 is different from Int2"`
		Igual      string `json:"igual" validate:"({{int2}}=={{int3}})msg=Int2 is equal from Int3"`
		Maior      string `json:"maior" validate:"({{int2}}>{{int1}})msg=Int2 is greater than Int1"`
		Menor      string `json:"menor" validate:"({{int1}}<{{int2}})msg=Int1 is less then Int2"`
		MaiorIgual string `json:"maior_igual" validate:"({{int2}}>={{int1}})msg=Int2 is greater or equal than Int1"`
		MenorIgual string `json:"menor_igual" validate:"({{int1}}<={{int2}})msg=Int1 is less or equal then Int2"`

		ValorDiferente  string `json:"valor_diferente" validate:"({{int1}}!=0)msg=Int1 is different from zero"`
		ValorIgual      string `json:"valor_igual" validate:"({{int1}}==1)msg=Int1 is equal from one"`
		ValorMaior      string `json:"valor_maior" validate:"({{int1}}>0)msg=Int1 is greater than zero"`
		ValorMenor      string `json:"valor_menor" validate:"({{int1}}<2)msg=Int1 is less then two"`
		ValorMaiorIgual string `json:"valor_maior_igual" validate:"({{int1}}>=1)msg=Int1 is greater or equal than zero"`
		ValorMenorIgual string `json:"valor_menor_igual" validate:"({{int1}}<=1)msg=Int1 is less or equal then zero"`

		ValorDiferente64  string `json:"valor_diferente_64" validate:"({{int64}}!=0)msg=Int64 is different from zero"`
		ValorIgual64      string `json:"valor_igual_64" validate:"({{int64}}==1)msg=Int64 is equal from one"`
		ValorMaior64      string `json:"valor_maior_64" validate:"({{int64}}>0)msg=Int64 is greater than zero"`
		ValorMenor64      string `json:"valor_menor_64" validate:"({{int64}}<2)msg=Int64 is less then two"`
		ValorMaiorIgual64 string `json:"valor_maior_igual_64" validate:"({{int64}}>=1)msg=Int64 is greater or equal than zero"`
		ValorMenorIgual64 string `json:"valor_menor_igual_64" validate:"({{int64}}<=1)msg=Int64 is less or equal then zero"`
	}{
		Nome1: "",
		Nome2: "Eucatur",
		Int1:  1,
		Int2:  2,
		Int3:  2,
		Int64: 1,
	}

	expected.Details = map[string]interface{}{
		"string_vazia":      "Nome1 is empty",
		"string_preenchida": "Nome2 is filled in",
		"string_diferente":  "Nome1 is different from Eucatur",
		"string_igual":      "Nome2 is equal from Eucatur",

		"diferente":   "Int1 is different from Int2",
		"igual":       "Int2 is equal from Int3",
		"maior":       "Int2 is greater than Int1",
		"menor":       "Int1 is less then Int2",
		"maior_igual": "Int2 is greater or equal than Int1",
		"menor_igual": "Int1 is less or equal then Int2",

		"valor_diferente":   "Int1 is different from zero",
		"valor_igual":       "Int1 is equal from one",
		"valor_maior":       "Int1 is greater than zero",
		"valor_menor":       "Int1 is less then two",
		"valor_maior_igual": "Int1 is greater or equal than zero",
		"valor_menor_igual": "Int1 is less or equal then zero",

		"valor_diferente_64":   "Int64 is different from zero",
		"valor_igual_64":       "Int64 is equal from one",
		"valor_maior_64":       "Int64 is greater than zero",
		"valor_menor_64":       "Int64 is less then two",
		"valor_maior_igual_64": "Int64 is greater or equal than zero",
		"valor_menor_igual_64": "Int64 is less or equal then zero",
	}

	test(t, expected, structCondition)
}

func Test_groups(t *testing.T) {
	// grupos de validação
	structGroup := struct {
		Nome          string `json:"nome" validate:"min=2,max=10"`
		CargoID       int    `json:"cargo_id" validate:"uint{ADD}" errmsg:"Report the cargo_id to register"`
		DataDasFerias string `json:"data_das_ferias" validate:"required{UPDATE},datebr{UPDATE}" errmsg:"Enter the data_das_ferias to update"`
	}{
		Nome:          "Eucatur",
		CargoID:       0,
		DataDasFerias: "15/9/17",
	}

	expected.Details = map[string]interface{}{
		"cargo_id": "Report the cargo_id to register",
	}

	assert.Equal(t, expected, Validate(structGroup, "ADD"))

	expected.Details = map[string]interface{}{
		"data_das_ferias": "Enter the data_das_ferias to update",
	}

	assert.Equal(t, expected, Validate(structGroup, "UPDATE"))
}

func Test_novalidate(t *testing.T) {
	// novalidate - usado para ignorar a validação de um campo de qualquer tipo
	structNoValidate := struct {
		Nome   string `json:"nome" validate:"required" errmsg:"Required field"`
		Filhos []struct {
			Nome string `json:"nome" validate:"required" errmsg:"Required field"`
		} `json:"filhos" novalidate:"true"`
		Endereco *struct {
			Logradouro string `json:"logradouro" validate:"required" errmsg:"Required field"`
			Numero     string `json:"numero" validate:"required" errmsg:"Required field"`
			Bairro     string `json:"bairro" validate:"required" errmsg:"Required field"`
		} `json:"endereco" novalidate:"isNil(endereco)"`
		Telefone *struct {
			DDD    string `json:"ddd" validate:"required"  errmsg:"Required field"`
			Numero string `json:"numero" validate:"required"  errmsg:"Required field"`
		} `json:"telefone" novalidate:"isNil(telefone)"`
		Permissoes []struct {
			Codigo   string `json:"codigo" validate:"required" errmsg:"Required field"`
			Detalhes struct {
				Tipo      string `json:"tipo"`
				Descricao string `json:"descricao"`
			} `json:"detalhes" novalidate:"isSuccessor(permissoes)"`
		} `json:"permissoes"`
		Obs struct {
			Descricao string `json:"descricao" validate:"required" errmsg:"Required field"`
			Data      string `json:"data" validate:"required,dateeua" errmsg:"Required field"`
		} `json:"obs" novalidate:"{{nome}}!=John Snow"`
	}{
		Nome: "",
		Filhos: []struct {
			Nome string `json:"nome" validate:"required" errmsg:"Required field"`
		}{
			{Nome: ""},
			{Nome: ""},
		},
		Telefone: &struct {
			DDD    string `json:"ddd" validate:"required"  errmsg:"Required field"`
			Numero string `json:"numero" validate:"required"  errmsg:"Required field"`
		}{
			DDD:    "",
			Numero: "",
		},
		Permissoes: []struct {
			Codigo   string `json:"codigo" validate:"required" errmsg:"Required field"`
			Detalhes struct {
				Tipo      string `json:"tipo"`
				Descricao string `json:"descricao"`
			} `json:"detalhes" novalidate:"isSuccessor(permissoes)"`
		}{
			{
				Codigo: "001",
			},
		},
		Obs: struct {
			Descricao string `json:"descricao" validate:"required" errmsg:"Required field"`
			Data      string `json:"data" validate:"required,dateeua" errmsg:"Required field"`
		}{
			Descricao: "",
			Data:      "",
		},
	}

	expected.Details = map[string]interface{}{
		"nome": "Required field",
		"telefone": map[string]interface{}{
			"ddd":    "Required field",
			"numero": "Required field",
		},
	}

	test(t, expected, structNoValidate)
}

func Test_regex(t *testing.T) {
	// correct regex
	structRegex := struct {
		Name  string `json:"name" noregex:"{{name}}!=''" regex:"^[a-zA-Z]{4,}$" errmsg:"The name must have more than four letters"`
		Phone string `json:"phone" regex:"^[+]?[0-9]{8,}$" errmsg:"Please provide a valid phone number"`
	}{
		Name:  "Michael",
		Phone: "+5569999999999",
	}

	test(t, expectedNil, structRegex)

	// incorrect regex
	structRegex.Name = "Ana"
	structRegex.Phone = "9999999"

	expected.Details = map[string]interface{}{
		"name":  "The name must have more than four letters",
		"phone": "Please provide a valid phone number",
	}

	test(t, expected, structRegex)
}

func Test_bool_operations(t *testing.T) {
	structBoolOperations := struct {
		MyBool1           MyBool `json:"my_bool_1"` // teste do tipo MyBool ao mesmo tempo
		MyBool2           bool   `json:"my_bool_2"`
		Equal             string `json:"equal" validate:"({{my_bool_1}}==true)msg=MyBool1 equals to true"`
		Different         string `json:"different" validate:"({{my_bool_1}}!=false)msg=MyBool1 is different from false"`
		BoolEqualBool     string `json:"bool_equal_bool" validate:"({{my_bool_1}}=={{my_bool_2}})msg=MyBool1 equal MyBool2"`
		BoolDifferentBool string `json:"bool_different_bool" validate:"({{my_bool_1}}!={{my_bool_2}})msg=MyBool1 is different MyBool2"`
		BoolOnly          string `json:"bool_only" validate:"({{my_bool_1}})msg=MyBool1 equals to true"`
		Invalid           string `json:"invalid" validate:"({{my_bool_1}}!=fals)msg=Message"`
	}{
		MyBool1: true,
		MyBool2: true,
	}

	expected.Details = map[string]interface{}{
		"equal":           "MyBool1 equals to true",
		"different":       "MyBool1 is different from false",
		"bool_equal_bool": "MyBool1 equal MyBool2",
		"bool_only":       "MyBool1 equals to true",
	}

	test(t, expected, structBoolOperations)
}

func Test_string_operations(t *testing.T) {
	structStringOperations := struct {
		MyString1         string `json:"my_string_1"`
		MyString2         string `json:"my_string_2"`
		Equal             string `json:"equal" validate:"({{my_string_1}}==Jon)msg=MyString1 equals to Jon"`
		Different         string `json:"different" validate:"({{my_string_1}}!=Snow)msg=MyString1 is different from Snow"`
		BoolEqualBool     string `json:"bool_equal_bool" validate:"({{my_string_1}}=={{my_string_2}})msg=MyString1 equal MyString2"`
		BoolDifferentBool string `json:"bool_different_bool" validate:"({{my_string_1}}!={{my_string_2}})msg=MyString1 is different MyString2"`
	}{
		MyString1: "Jon",
		MyString2: "Jon",
	}

	expected.Details = map[string]interface{}{
		"equal":           "MyString1 equals to Jon",
		"different":       "MyString1 is different from Snow",
		"bool_equal_bool": "MyString1 equal MyString2",
	}

	test(t, expected, structStringOperations)
}

func Test_float64_operations(t *testing.T) {
	structFloat64Operations := struct {
		MyFloat64One               MyFloat64 `json:"my_float64_one"` // teste do tipo MyFloat64 ao mesmo tempo
		MyFloat64Two               float64   `json:"my_float64_two"`
		Equal                      string    `json:"equal" validate:"(10.5=={{my_float64_one}})msg=MyFloat64One equals to 10.5"`
		Different                  string    `json:"different" validate:"({{my_float64_one}}!=10.6)msg=MyFloat64One is different from 10.6"`
		Larger                     string    `json:"larger" validate:"({{my_float64_one}}>10.4)msg=MyFloat64One is greater than 10.4"`
		LargerEqual                string    `json:"larger_equal" validate:"({{my_float64_one}}>=10.5)msg=MyFloat64One is greater than 10.4"`
		Smaller                    string    `json:"smaller" validate:"({{my_float64_one}}<10.6)msg=MyFloat64One is less than 10.6"`
		SmallerEqual               string    `json:"smaller_equal" validate:"({{my_float64_one}}<=10.5)msg=MyFloat64One is less than or equal to 10.5"`
		Float64EqualFloat64        string    `json:"float64_equal_float64" validate:"({{my_float64_one}}=={{my_float64_two}})msg=MyFloat64One equal MyFloat64Two"`
		Float64DifferentFloat64    string    `json:"float64_different_float64" validate:"({{my_float64_one}}!={{my_float64_two}})msg=MyFloat64One is different from MyFloat64Two"`
		Float64LargerFloat64       string    `json:"float64_larger_float64" validate:"({{my_float64_one}}>{{my_float64_two}})msg=MyFloat64One is greater than MyFloat64Two"`
		Float64LargerEqualFloat64  string    `json:"float64_larger_equal_float64" validate:"({{my_float64_one}}>={{my_float64_two}})msg=MyFloat64One is greater or equal to MyFloat64Two"`
		Float64SmallerFloat64      string    `json:"float64_smaller_float64" validate:"({{my_float64_one}}<{{my_float64_two}})msg=MyFloat64One is greater than MyFloat64Two"`
		Float64SmallerEqualFloat64 string    `json:"float64_smaller_equal_float64" validate:"({{my_float64_one}}<={{my_float64_two}})msg=MyFloat64One is greater or equal to MyFloat64Two"`
		Invalid                    string    `json:"invalid" validate:"({{my_float64_one}}!=fals)msg=Message"`
	}{
		MyFloat64One: 10.5,
		MyFloat64Two: 10.6,
	}

	expected.Details = map[string]interface{}{
		"equal":                         "MyFloat64One equals to 10.5",
		"different":                     "MyFloat64One is different from 10.6",
		"larger":                        "MyFloat64One is greater than 10.4",
		"larger_equal":                  "MyFloat64One is greater than 10.4",
		"smaller":                       "MyFloat64One is less than 10.6",
		"smaller_equal":                 "MyFloat64One is less than or equal to 10.5",
		"float64_different_float64":     "MyFloat64One is different from MyFloat64Two",
		"float64_smaller_float64":       "MyFloat64One is greater than MyFloat64Two",
		"float64_smaller_equal_float64": "MyFloat64One is greater or equal to MyFloat64Two",
	}

	test(t, expected, structFloat64Operations)
}

func Test_int_operations(t *testing.T) {
	structIntOperations := struct {
		MyInt        MyInt  `json:"my_int"` // teste do tipo MyInt ao mesmo tempo
		Equal        string `json:"equal" validate:"({{my_int}}==10)msg=MyInt equals to 10"`
		Different    string `json:"different" validate:"({{my_int}}!=9)msg=MyBool is different from 9"`
		Larger       string `json:"larger" validate:"({{my_int}}>9)msg=MyInt is greater than 9"`
		LargerEqual  string `json:"larger_equal" validate:"({{my_int}}>=10)msg=MyInt is greater than 10"`
		Smaller      string `json:"smaller" validate:"({{my_int}}<11)msg=MyInt is less than 11"`
		SmallerEqual string `json:"smaller_equal" validate:"({{my_int}}<=10)msg=MyInt is less than or equal to 10"`
		Invalid      string `json:"invalid" validate:"({{my_int}}!=fals)msg=Message"`
	}{
		MyInt: 10,
	}

	expected.Details = map[string]interface{}{
		"equal":         "MyInt equals to 10",
		"different":     "MyBool is different from 9",
		"larger":        "MyInt is greater than 9",
		"larger_equal":  "MyInt is greater than 10",
		"smaller":       "MyInt is less than 11",
		"smaller_equal": "MyInt is less than or equal to 10",
	}

	test(t, expected, structIntOperations)
}

func Test_unique(t *testing.T) {
	// correto
	structUnique := struct {
		Search1 string `json:"search_1" validate:"({{search_2.name}}==Hi)msg=Message"`
		Search2 string `json:"search_2" validate:"({{people.name.not_found}}==Hi)msg=Message"`
		People  []struct {
			Name string `json:"name" validate:"unique,({{.search_1}}!='')msg=Message"`
		} `json:"people"`
	}{
		People: []struct {
			Name string `json:"name" validate:"unique,({{.search_1}}!='')msg=Message"`
		}{
			{
				Name: "Jon",
			},
			{
				Name: "Snow",
			},
		},
	}

	test(t, expectedNil, structUnique)

	// incorreto
	structUnique.People[0].Name = "Ghost"
	structUnique.People[1].Name = "Ghost"

	expected.Details = map[string]interface{}{
		"people": map[int]interface{}{
			0: map[string]interface{}{
				"name": "Esse item não pode ser repetido"},
			1: map[string]interface{}{
				"name": "Esse item não pode ser repetido"},
		},
	}

	test(t, expected, structUnique)
}

func Test_pointers(t *testing.T) {
	strPointer := ""
	structPointer := struct {
		Pointer              *string `json:"pointer" validate:"({{pointer_nil_required}}=='')msg=Pointer"`
		PointerNilRequired   *string `json:"pointer_nil_required" validate:"required" errmsg:"Required"`
		PointerNilNoRequired *string `json:"pointer_nil_no_required"`
	}{
		Pointer: &strPointer,
	}

	expected.Details = map[string]interface{}{
		"pointer":              "Pointer",
		"pointer_nil_required": "Required",
	}

	test(t, expected, structPointer)
}

func Test_others(t *testing.T) {
	structOthersTests := struct {
		ValidationEmpty string `json:"validation_empty" validation:""`
		Ignored         string `json:"-"`
		Regex           string `json:"regex" regex:"[0-9]"`
		GroupsEmpty     string `json:"groups_empty" validate:"msg=Message"`
		VarNotFrond1    string `json:"var_not_found_1" validate:"({{not_found}}==Invalid)msg=Message"`
		VarNotFrond2    string `json:"var_not_found_2" validate:"(Invalid=={{.not_found.not_found}})msg=Message"`
		WithoutVar      string `json:"without_var" validate:"(a==a)msg=Message"`

		MyInt64One    int64  `json:"my_int64_one"` // teste do tipo MyInt64 ao mesmo tempo
		MyInt64Two    int64  `json:"my_int64_two"`
		NotValidated1 string `json:"not_validated_1" validate:"({{my_int64_one}}==10.1)msg=Message"`
		NotValidated2 string `json:"not_validated_2" validate:"({{my_int64_one}}=={{my_int64_two}})msg=Message"`

		WithoutTagJSON         string `query:"without_tag_json" validate:"({{WithoutTagJSONAndQuery}}=='')msg=Message"`
		WithoutTagJSONAndQuery string `validate:"({{without_tag_json}}=='')msg=Message"`
		privateField           string `validate:"required"` // é ignorado pois não pode ler o valor
	}{
		Regex: "A",
	}

	expected.Details = map[string]interface{}{
		"regex":                  "Informe o conteúdo conforme a expressão regular: [0-9]",
		"groups_empty":           "Message",
		"without_tag_json":       "Message",
		"WithoutTagJSONAndQuery": "Message",
	}

	test(t, expected, structOthersTests)
}

func test(t *testing.T, expected, body interface{}) {
	assert.Equal(t, expected, Validate(body))
}
