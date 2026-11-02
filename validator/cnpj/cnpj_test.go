package cnpj

import "testing"

func TestDefinirMascara(t *testing.T) {
	type args struct {
		cnpj string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Definir máscara",
			args: args{
				cnpj: "11111111000123",
			},
			want: "11.111.111/0001-23",
		},
		{
			name: "CNPJ com máscara deve-se manter",
			args: args{
				cnpj: "11.111.111/0001-11",
			},
			want: "11.111.111/0001-11",
		},
		{
			name: "Nada informado, deve-se retornar nada",
			args: args{
				cnpj: "",
			},
			want: "",
		},
		{
			name: "CNPJ inválido irá retornar o mesmo conteúdo passado para definir a máscara",
			args: args{
				cnpj: "22222222222",
			},
			want: "22222222222",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefinirMascara(tt.args.cnpj); got != tt.want {
				t.Errorf("DefinirMascara() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValido(t *testing.T) {
	tests := []struct {
		name string
		cnpj string
		want bool
	}{
		{
			name: "CNPJ alfanumérico válido com máscara",
			cnpj: "MJ.R98.ZVY/0001-75",
			want: true,
		},
		{
			name: "CNPJ alfanumérico válido sem máscara",
			cnpj: "7948DBKX000140",
			want: true,
		},
		{
			name: "CNPJ numérico válido com máscara",
			cnpj: "05.174.502/0001-52",
			want: true,
		},
		{
			name: "CNPJ numérico válido sem máscara",
			cnpj: "05174502000152",
			want: true,
		},
		{
			name: "CNPJ alfanumérico inválido com máscara",
			cnpj: "MJ.R98.ZVY/0001-76",
			want: false,
		},
		{
			name: "CNPJ alfanumérico inválido sem máscara",
			cnpj: "7948DBKX000141",
			want: false,
		},
		{
			name: "CNPJ numérico inválido com máscara",
			cnpj: "05.174.502/0001-53",
			want: false,
		},
		{
			name: "CNPJ numérico inválido sem máscara",
			cnpj: "05174502000151",
			want: false,
		},
		{
			name: "CNPJ vazio",
			cnpj: "",
			want: false,
		},
		{
			name: "CNPJ com caracteres especiais inválido",
			cnpj: "05.174.502/0001-5@",
			want: false,
		},
		{
			name: "CNPJ incompleto",
			cnpj: "05.174.502/0001",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Valido(tt.cnpj)
			if got != tt.want {
				t.Errorf("Valido() | Test: %s | CNPJ: %s = %v, want %v", tt.name, tt.cnpj, got, tt.want)
			}
		})
	}
}
