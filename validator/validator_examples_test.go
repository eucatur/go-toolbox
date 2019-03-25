package validator

import "log"

func Example() {

	// msg
	// Sempre retorna a mensagem de erro passada no parâmetro
	// Geralmente usada junto a uma condição
	type Msg struct {
		MyString string `json:"my_string" validate:"Msg=Enter message here"`
	}

	// {
	// 	"my_string": "Enter message here"
	// }

	// required
	// String: um ou mais caracteres
	// Slice um ou mais elementos
	type Required struct {
		MyString string        `json:"my_string" validate:"required"`
		MySlice  []interface{} `json:"my_slice" validate:"required"`
	}

	// uint
	// Int e int64: maior ou igual a um
	type UInt struct {
		MyInt   int   `json:"my_int" validate:"uint"`
		MyInt64 int64 `json:"my_int64" validate:"uint"`
	}

	// datebr
	// String: data em formato brasileiro ex: 18/09/2017
	type DateBR struct {
		MyString string `json:"my_string" validate:"datebr"`
	}

	// dateeua
	// String: data em formato americano ex: 2017-09-18
	type DateEUA struct {
		MyString string `json:"my_string" validate:"dateeua"`
	}

	// datetime
	// String: data no formato do parametro informado conforme layout da lib time
	// da golang
	type DateTime struct {
		MyString string `json:"my_string" validate:"datetime=2006-01-02"`
	}

	// len
	// Int: numero de casas decimais
	// String: numero de caracteres
	// Slice: numero de elementos
	type Len struct {
		MyInt    int           `json:"my_int" validate:"len=1"`
		MyString string        `json:"my_string" validate:"len=1"`
		MySlice  []interface{} `json:"my_slice" validate:"len=1"`
	}

	// min
	// Float64, int e int64: valor mínimo aceitável
	// String: numero mínimo de caracteres
	// Slice: numero mínimo de elementos
	type Min struct {
		MyFloat64 float64       `json:"my_float64" validate:"min=1"`
		MyInt     int           `json:"my_int" validate:"min=1"`
		MyInt64   int64         `json:"my_int64" validate:"min=1"`
		MyString  string        `json:"my_string" validate:"min=1"`
		MySlice   []interface{} `json:"my_slice" validate:"min=1"`
	}

	// max
	// Float64, int e int64: valor máximo aceitável
	// String: numero máximo de caracteres
	// Slice: numero máximo de elementos
	type Max struct {
		MyFloat64 float64       `json:"my_float64" validate:"max=1"`
		MyInt     int           `json:"my_int" validate:"max=1"`
		MyInt64   int64         `json:"my_int64" validate:"max=1"`
		MyString  string        `json:"my_string" validate:"max=1"`
		MySlice   []interface{} `json:"my_slice" validate:"max=1"`
	}

	// enum
	// Int, int64, string: grupo de valores aceitáveis
	type Enum struct {
		MyInt    int    `json:"my_int" validate:"enum=1|2"`
		MyInt64  int64  `json:"my_int64" validate:"enum=1|2"`
		MyString string `json:"my_string" validate:"enum=A|B"`
	}

	// email
	// String: email em formato válido
	type Email struct {
		MyString string `json:"my_string" validate:"email"`
	}

	// numeric
	// String: apenas numeros [0-9]
	type Numeric struct {
		MyString string `json:"my_string" validate:"numeric"`
	}

	// cpf
	// String: CPF válido com ou sem pontos e traço
	type CPF struct {
		MyString string `json:"my_string" validate:"cpf"`
	}

	// cnpj
	// String: CNPJ válidos com ou sem ponto e traço
	type CNPJ struct {
		MyString string `json:"my_string" validate:"cnpj"`
	}

	// novalidate
	// Ignora as validações do campo (todos os tipos: int, string, slice e etc).
	// No exemplo abaixo temos uma struct 'NoValidateElemento' que possui um campo
	// do tipo string com a validação 'required', temos também a struct 'NoValidate'
	// com um slice da primeira struct. Usando a novalidate:"true" no slice, as
	// validações dos elementos serão ignoradas. (Ajuda na reutilização de código)
	// Podemos ainda utilizar condições igual aos outros exemplos.

	type NoValidateElemento struct {
		MyString string `json:"my_string" validate:"required"`
	}

	type NoValidate struct {
		MySlice []NoValidate `json:"my_slice" novalidate:"true"`
	}

	type NoValidateCondicional struct {
		Nome     string `json:"nome" validate:"required"`
		Telefone string `json:"telefone" novalidate:"{{nome}}==''"`
		Email    string `json:"email" novalidate:"{{nome}}==teste"`
		Endereco *struct {
			Logradouro  string `json:"logradouro"`
			Numero      string `json:"numero"`
			Bairro      string `json:"bairro"`
			Complemento string `json:"complemento" novalidate:"isSuccessor(endereco)"`
		} `json:"endereco" novalidate:"isNil(endereco)"`
	}

	// regex
	// Verifica se o conteúdo do campo (convertido para string) é compatível com a
	// regex informada.
	// No exemplo a baixo temos a regex de placa de veículo.
	type Regex struct {
		MyString string `json:"my_string" regex:"^[A-Z]{3}[0-9]{4}$" errmsg:"Please provide a valid badge"`
	}

	// {
	// 	"my_string": "Please provide a valid badge"
	// }

	// grupo de validação
	// Permite criar grupo de validação onde:
	// 1) Campo sem grupo sempre é validado
	// 2) Grupo não solicitado não é validado
	// 3) Grupo solicitado é validado (e verificado a condição caso tenha)
	type Grupo struct {
		MyInt    int           `json:"my_int" validate:"uint"`           // Sem grupo
		MyInt64  int64         `json:"my_int64" validate:"uint{A}"`      // Grupo A
		MyString string        `json:"my_string" validate:"required{B}"` // Grupo B
		MySlice  []interface{} `json:"my_slice" validate:"max{A|B}=1"`   // Grupo A e B
	}

	var grupo = Grupo{
		MyInt:    0,
		MyInt64:  0,
		MyString: "",
		MySlice:  []interface{}{1, 2},
	}

	var vError, vErrorA, vErrorB *VError

	vError = Validate(grupo)
	vErrorA = Validate(grupo, "A")
	vErrorB = Validate(grupo, "B")

	log.Println(vError.Details)
	log.Println(vErrorA.Details)
	log.Println(vErrorB.Details)

	// JSON vError
	// {
	// 	"my_int": ""
	// }

	// JSON vErrorA
	// {
	// 	"my_int": "",
	// 	"my_int64": "",
	// 	"my_slice": ""
	// }

	// JSON vErrorB
	// {
	// 	"my_int": "",
	// 	"my_string": "",
	// 	"my_slice": ""
	// }
}
