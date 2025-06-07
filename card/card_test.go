package card

import (
	"testing"
)

func TestMask(t *testing.T) {
	type args struct {
		cardNumber string
	}
	tests := []struct {
		name         string
		args         args
		wantCardMask string
		wantErr      bool
		debug        bool
	}{
		{
			name: "Número não informado",
			args: args{
				cardNumber: "",
			},
			wantCardMask: "",
			wantErr:      true,
		},
		{
			name: "Número com 13",
			args: args{
				cardNumber: "1234567890123",
			},
			wantCardMask: "123456XXX0123",
			wantErr:      false,
		},
		{
			name: "Número com 19",
			args: args{
				cardNumber: "1234567890123456789",
			},
			wantCardMask: "123456XXXXXXXXX6789",
			wantErr:      false,
		},
		{
			name: "Número inválido",
			args: args{
				cardNumber: "1234567890",
			},
			wantCardMask: "",
			wantErr:      true,
		},
		{
			name: "já mascarado",
			args: args{
				cardNumber: "400000XXXXXX1091",
			},
			wantCardMask: "400000XXXXXX1091",
			wantErr:      false,
		},
		{
			name: "Mascarar",
			args: args{
				cardNumber: "4000000000001091",
			},
			wantCardMask: "400000XXXXXX1091",
			wantErr:      false,
		},
		{
			name: "já mascarado com outro padrão",
			args: args{
				cardNumber: "400000******1091",
			},
			wantCardMask: "400000******1091",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.debug {
				print("")
			}
			gotCardMask, err := Mask(tt.args.cardNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCardMask != tt.wantCardMask {
				t.Errorf("Mask() = %v, want %v", gotCardMask, tt.wantCardMask)
			}
		})
	}
}

func TestGetInicialBin(t *testing.T) {
	type args struct {
		cardNumber string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito",
			args: args{
				cardNumber: "4024017150337433",
			},
			want: "402401",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito contendo espaços entre os números",
			args: args{
				cardNumber: "4024 0171 5033 7433",
			},
			want: "402401",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito contendo espaços entre os números e caracteres de espaços",
			args: args{
				cardNumber: "\t 4024 0171 5033 7433 \n   ",
			},
			want: "402401",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 18 números",
			args: args{
				cardNumber: "402401715033743367",
			},
			want: "402401",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 18 números e contendo espaços entre os números",
			args: args{
				cardNumber: "402 401 715 033 743 367",
			},
			want: "402401",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 18 número e contendo espaços entre os números e caracteres de espaços",
			args: args{
				cardNumber: "\t 402 401 715 033 743 367 \n   ",
			},
			want: "402401",
		},

		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 19 números",
			args: args{
				cardNumber: "4024017150337433679",
			},
			want: "402401",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 19 números e contendo espaços entre os números",
			args: args{
				cardNumber: "402401 7150 33743 3679",
			},
			want: "402401",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 19 número e contendo espaços entre os números e caracteres de espaços",
			args: args{
				cardNumber: "\t 402401 7150 33743 3679 \n   ",
			},
			want: "402401",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInicialBin(tt.args.cardNumber); got != tt.want {
				t.Errorf("GetInicialBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFinalBin(t *testing.T) {
	type args struct {
		cardNumber string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito",
			args: args{
				cardNumber: "4024017150337433",
			},
			want: "7433",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito contendo espaços entre os números",
			args: args{
				cardNumber: "4024 0171 5033 7433",
			},
			want: "7433",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito contendo espaços entre os números e caracteres de espaços",
			args: args{
				cardNumber: "\t 4024 0171 5033 7433 \n   ",
			},
			want: "7433",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 18 números",
			args: args{
				cardNumber: "402401715033743367",
			},
			want: "3367",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 18 números e contendo espaços entre os números",
			args: args{
				cardNumber: "402 401 715 033 743 367",
			},
			want: "3367",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 18 número e contendo espaços entre os números e caracteres de espaços",
			args: args{
				cardNumber: "\t 402 401 715 033 743 367 \n   ",
			},
			want: "3367",
		},

		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 19 números",
			args: args{
				cardNumber: "4024017150337433679",
			},
			want: "3679",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 19 números e contendo espaços entre os números",
			args: args{
				cardNumber: "402401 7150 33743 3679",
			},
			want: "3679",
		},
		{
			name: "Retornar os 6 primeiros dígitos do cartão de crédito sendo com 19 número e contendo espaços entre os números e caracteres de espaços",
			args: args{
				cardNumber: "\t 402401 7150 33743 3679 \n   ",
			},
			want: "3679",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFinalBin(tt.args.cardNumber); got != tt.want {
				t.Errorf("GetFinalBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
