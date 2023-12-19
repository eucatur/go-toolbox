package commandredis

import (
	"bytes"
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {

	require.NotEmpty(t, List())

}

func TestCommandRedis_List(t *testing.T) {

	require.NotEmpty(t, CommandRedis(0).List())

}

func TestCommandRedis(t *testing.T) {

	type args struct {
		value string
	}

	tests := []struct {
		name string
		args args
		want CommandRedis
	}{
		{
			name: "Enum invalid",
			args: args{
				value: "2kl2j423lk4j23",
			},
			want: Undefined,
		},
		{
			name: "Enum valid",
			args: args{
				"Undefined",
			},
			want: Undefined,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := TryParseToEnum(tt.args.value); got != tt.want {

				t.Errorf("TryParseToEnum() = %v, want %v", got, tt.want)

			}

		})

	}

}

func TestCommandRedis_OperationsStatemented(t *testing.T) {

	tests := []struct {
		name                    string
		m                       CommandRedis
		wantOperationStatmented []CommandRedis
	}{
		{
			name:                    "Just undefined",
			m:                       Undefined,
			wantOperationStatmented: []CommandRedis{},
		},
		{
			name: "Some enums",
			m:    Delete | Expire,
			wantOperationStatmented: []CommandRedis{
				Delete,
				Expire,
			},
		},

		{
			name:                    "Just one enum",
			m:                       Delete,
			wantOperationStatmented: []CommandRedis{},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if gotOperationStatmented := tt.m.OperationsStatemented(); !reflect.DeepEqual(gotOperationStatmented, tt.wantOperationStatmented) {

				t.Errorf("CommandRedis.OperationsStatemented() = %v, want %v", gotOperationStatmented, tt.wantOperationStatmented)

			}

		})

	}

}

func TestCommandRedis_String(t *testing.T) {

	tests := []struct {
		name string
		c    CommandRedis
		want string
	}{
		{
			name: "Obter a string do enumerado indefinido",
			c:    Undefined,
			want: name_undefined,
		},
		{
			name: "Undefined para tipo não identificado",
			c:    95938475398475,
			want: name_undefined,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := tt.c.String(); got != tt.want {

				t.Errorf("CommandRedis.String() = %v, want %v", got, tt.want)

			}

		})

	}

}

func TestTypesAccepts(t *testing.T) {

	require.NotEmpty(t, TypesAccepts(), "Retornar a lista dos enumerados disponíveis")

}

func TestCommandRedis_TypesAccepts(t *testing.T) {

	enum := CommandRedis(0)

	require.NotEmpty(t, enum.TypesAccepts(), "Retornar a lista dos enumerados disponíveis")

}

func TestCommandRedis_MarshalJSON(t *testing.T) {

	tests := []struct {
		name    string
		c       CommandRedis
		want    []byte
		wantErr bool
	}{
		{
			name:    "Retornar undefined",
			c:       0,
			want:    []byte(fmt.Sprintf(`"%s"`, name_undefined)),
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.c.MarshalJSON()

			if (err != nil) != tt.wantErr {

				t.Errorf("CommandRedis.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)

				return

			}

			if !reflect.DeepEqual(got, tt.want) {

				t.Errorf("CommandRedis.MarshalJSON() = %v, want %v", got, tt.want)

			}

		})

	}

}

func TestCommandRedis_UnmarshalJSON(t *testing.T) {

	type args struct {
		possibilities []interface{}
	}
	tests := []struct {
		name         string
		c            *CommandRedis
		args         args
		wantErr      bool
		enumExpected CommandRedis
	}{
		{
			name: "Converter de JSON para enum tipo indefinido",
			c:    new(CommandRedis),
			args: args{
				possibilities: []interface{}{
					`"undefined"`,
					`"Undefined"`,
					`"UNDEFINED"`,
					"undefined",
					"Undefined",
					"UNDEFINED",
					`0`,
					`"0"`,
					`"0"`,
					int(0),
					CommandRedis(0),
				},
			},
			wantErr:      false,
			enumExpected: Undefined,
		},
		{
			name: "Converter de JSON para enum tipo não identificado",
			c:    new(CommandRedis),
			args: args{
				possibilities: []interface{}{
					`355345345`,
					`"345345345"`,
					`"345345345"`,
					int(345345345),
					[]uint8(`345345345`),
					float32(345345),
					float64(345345),
					int64(345345345),
					int32(345345345),
					`"kjfsdlfj"`,
					"kjfsdlfj",
					`234`,
					`"756"`,
					`"9732"`,
					int(34534534),
					[]uint8(`478`),
					float32(853),
					float64(841),
					int64(326),
					int32(765),
				},
			},
			wantErr:      true,
			enumExpected: Undefined,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			for _, possibility := range tt.args.possibilities {

				if err := tt.c.UnmarshalJSON([]byte(fmt.Sprint(possibility))); (err != nil) != tt.wantErr {

					t.Errorf("CommandRedis.UnmarshalJSON() error = %v, wantErr %v when unmarshal the value %v", err, tt.wantErr, possibility)

				}

				require.Equal(t, tt.enumExpected, *tt.c)

			}

		})

	}

}

func TestParseType(t *testing.T) {

	type args struct {
		f    reflect.Type
		t    reflect.Type
		data interface{}
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Aplicar verificação para tipo enumerado não definido",
			args: args{
				f:    reflect.TypeOf(234),
				t:    reflect.TypeOf(234),
				data: 234,
			},
			want:    234,
			wantErr: false,
		},
		{
			name: "Aplicar verificação para o tipo enumerado definido não válido",
			args: args{
				f:    reflect.TypeOf("234"),
				t:    reflect.TypeOf(Undefined),
				data: "2l3mklm23mkl",
			},
			want:    Undefined,
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got, err := ParseType(tt.args.f, tt.args.t, tt.args.data)

			if (err != nil) != tt.wantErr {

				t.Errorf("ParseType() error = %v, wantErr %v", err, tt.wantErr)

				return

			}

			if !reflect.DeepEqual(got, tt.want) {

				t.Errorf("ParseType() = %v, want %v", got, tt.want)

			}

		})

	}

}

func TestCommandRedis_Value(t *testing.T) {

	tests := []struct {
		name    string
		c       CommandRedis
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "Retorno valor do enumerado",
			c:       0,
			want:    name_undefined,
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.c.Value()

			if (err != nil) != tt.wantErr {

				t.Errorf("CommandRedis.Value() error = %v, wantErr %v", err, tt.wantErr)

				return

			}

			if !reflect.DeepEqual(got, tt.want) {

				t.Errorf("CommandRedis.Value() = %v, want %v", got, tt.want)

			}

		})

	}

}

func TestCommandRedis_Scan(t *testing.T) {

	type args struct {
		possibilities []interface{}
	}
	tests := []struct {
		name         string
		c            *CommandRedis
		args         args
		wantErr      bool
		enumExpected CommandRedis
	}{
		{
			name: "Sem informação deve ser retornado como indefinido",
			c:    new(CommandRedis),
			args: args{
				possibilities: []interface{}{
					"",
					`""`,
					[]uint8(``),
					[]uint8(""),
				},
			},
			wantErr:      false,
			enumExpected: Undefined,
		},
		{
			name: "Scan para enumerado indefinido",
			c:    new(CommandRedis),
			args: args{
				possibilities: []interface{}{
					`"undefined"`,
					`"Undefined"`,
					`"UNDEFINED"`,
					"undefined",
					"Undefined",
					"UNDEFINED",
					`0`,
					`"0"`,
					`"0"`,
					int(0),
					[]uint8(`0`),
					float32(0),
					float64(0),
					int64(0),
					int32(0),
					CommandRedis(0),
				},
			},
			wantErr:      false,
			enumExpected: Undefined,
		},
		{
			name: "Scan para enumerado não identificado",
			c:    new(CommandRedis),
			args: args{
				possibilities: []interface{}{
					`875644657`,
					`"875644657"`,
					`"875644657875644657"`,
					int(8756457),
					[]uint8(`875644657`),
					float32(564657),
					float64(8744657),
					int64(875644657),
					int32(874657),
					CommandRedis(875644657),
					`"kjfsdlfj"`,
					"kjfsdlfj",
					`234`,
					`"756"`,
					`"9732"`,
					int(123),
					[]uint8(`478`),
					float32(853),
					float64(841),
					int64(326),
					int32(765),
					CommandRedis(3405),
				},
			},
			wantErr:      true,
			enumExpected: Undefined,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			for _, possibility := range tt.args.possibilities {

				if err := tt.c.Scan(possibility); (err != nil) != tt.wantErr {

					t.Errorf("CommandRedis.Scan() error = %v, wantErr %v", err, tt.wantErr)

				}

				require.Equal(t, tt.enumExpected, *tt.c)

			}

		})

	}

}

func TestCommandRedis_IsValid(t *testing.T) {

	tests := []struct {
		name    string
		c       CommandRedis
		wantErr bool
	}{
		{
			name:    "Valor inespero",
			c:       234324,
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if err := tt.c.IsValid(); (err != nil) != tt.wantErr {

				t.Errorf("CommandRedis.IsValid() error = %v, wantErr %v", err, tt.wantErr)

			}

		})

	}

}

func TestCommandRedis_MarshalXML(t *testing.T) {

	type args struct {
		xmlEnc *xml.Encoder
		start  xml.StartElement
	}
	tests := []struct {
		name    string
		c       CommandRedis
		args    args
		wantErr bool
	}{
		{
			name: "Conversão para XML",
			c:    Undefined,
			args: args{
				xmlEnc: xml.NewEncoder(bytes.NewBuffer([]byte(""))),
				start:  xml.StartElement{Name: xml.Name{Space: "", Local: "enum_test"}},
			},
			wantErr: false,
		},
		{
			name: "falha ao converter para XML",
			c:    Undefined,
			args: args{
				xmlEnc: xml.NewEncoder(bytes.NewBuffer([]byte(""))),
				start:  xml.StartElement{Name: xml.Name{Space: "", Local: ""}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if err := tt.c.MarshalXML(tt.args.xmlEnc, tt.args.start); (err != nil) != tt.wantErr {

				t.Errorf("CommandRedis.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)

			}

		})

	}

}

func TestCommandRedis_UnmarshalXML(t *testing.T) {

	tests := []struct {
		name    string
		c       *CommandRedis
		data    []byte
		wantErr bool
	}{
		{
			name:    "Converter de XML para enumerado",
			c:       new(CommandRedis),
			data:    []byte("<enum_test>undefined</enum_test>"),
			wantErr: false,
		},
		{
			name:    "Converter de XML para enumerado com informação inválida",
			c:       new(CommandRedis),
			data:    []byte("<enum_test>32klj423l</enum_test>"),
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if err := xml.Unmarshal(tt.data, tt.c); (err != nil) != tt.wantErr {

				t.Errorf("CommandRedis.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)

			}

		})

	}

}
