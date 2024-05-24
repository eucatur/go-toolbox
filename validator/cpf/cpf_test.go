package cpf

import "testing"

func TestDefinirMascara(t *testing.T) {
	type args struct {
		cpf string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Definir máscara",
			args: args{
				cpf: "11111111111",
			},
			want: "111.111.111-11",
		},
		{
			name: "CPF com máscara deve-se manter",
			args: args{
				cpf: "111.111.111-11",
			},
			want: "111.111.111-11",
		},
		{
			name: "Nada informado, deve-se retornar nada",
			args: args{
				cpf: "",
			},
			want: "",
		},
		{
			name: "CPF inválido irá retornar o mesmo conteúdo passado para definir a máscara",
			args: args{
				cpf: "222222222",
			},
			want: "222222222",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefinirMascara(tt.args.cpf); got != tt.want {
				t.Errorf("DefinirMascara() = %v, want %v", got, tt.want)
			}
		})
	}
}
