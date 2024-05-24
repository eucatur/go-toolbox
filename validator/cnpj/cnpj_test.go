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
