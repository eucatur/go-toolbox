package card

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMask(t *testing.T) {
	card, _ := Mask("5353 1607 6798 7690")

	assert.Equal(t, card, "5353********7690", "The two words should be the same.")
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
