package money

import (
	"testing"
)

func TestToFloat(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "with zero",
			args: args{value: 1800},
			want: 18,
		},
		{
			name: "negative with zero",
			args: args{value: -1800},
			want: -18,
		},
		{
			name: "with decimal",
			args: args{value: 1990},
			want: 19.90,
		},
		{
			name: "negative with decimal",
			args: args{value: -1990},
			want: -19.90,
		},
		{
			name: "another with decimal",
			args: args{value: 1690},
			want: 16.90,
		},
		{
			name: "another negative with decimal",
			args: args{value: -1690},
			want: -16.90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat(tt.args.value); got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with zero",
			args: args{value: 18.00},
			want: 1800,
		},
		{
			name: "negative with zero",
			args: args{value: -18.00},
			want: -1800,
		},
		{
			name: "with decimal",
			args: args{value: 19.90},
			want: 1990,
		},
		{
			name: "negative with decimal",
			args: args{value: -19.90},
			want: -1990,
		},
		{
			name: "another with decimal",
			args: args{value: 16.90},
			want: 1690,
		},
		{
			name: "another negative with decimal",
			args: args{value: -16.90},
			want: -1690,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.value); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		value     float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "round to down",
			args: args{value: 150.141, precision: 2},
			want: 150.14,
		},
		{
			name: "round to up",
			args: args{value: 150.145, precision: 2},
			want: 150.15,
		},
		{
			name: "round to up",
			args: args{value: 150.146, precision: 2},
			want: 150.15,
		},
		{
			name: "roud value",
			args: args{value: 100.998, precision: 2},
			want: 101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.value, tt.args.precision); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReais(t *testing.T) {
	type args struct {
		valor int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "centavos",
			args: args{valor: 50},
			want: "0,50",
		},
		{
			name: "dezena",
			args: args{valor: 2400},
			want: "24,00",
		},
		{
			name: "dezena e centavos",
			args: args{valor: 1990},
			want: "19,90",
		},
		{
			name: "centena",
			args: args{valor: 12300},
			want: "123,00",
		},
		{
			name: "centena, dezena e cetavos",
			args: args{valor: 15043},
			want: "150,43",
		},
		{
			name: "milhar, centena, dezena e centavos",
			args: args{valor: 155763},
			want: "1.557,63",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reais(tt.args.valor); got != tt.want {
				t.Errorf("Reais() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		value    float64
		decimals int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test with value 16.90",
			args: args{
				value:    16.90,
				decimals: 2,
			},
			want: 16.90,
		},
		{
			name: "test value 100.998 must be returned 100.99",
			args: args{
				value:    100.998,
				decimals: 2,
			},
			want: 100.99,
		},
		{
			name: "test with value -16.90",
			args: args{
				value:    -16.90,
				decimals: 2,
			},
			want: -16.90,
		},
		{
			name: "test value -100.998 must be returned -100.99",
			args: args{
				value:    -100.998,
				decimals: 2,
			},
			want: -100.99,
		},
		{
			name: "test with value 16.899999991 must be returned 16.89",
			args: args{
				value:    16.899999991,
				decimals: 2,
			},
			want: 16.89,
		},
		{
			name: "test with value -16.899999991 must be returned -16.89",
			args: args{
				value:    -16.899999991,
				decimals: 2,
			},
			want: -16.89,
		},
		{
			name: "test value 100.994999999999 must be returned 100.99",
			args: args{
				value:    100.994999999999,
				decimals: 2,
			},
			want: 100.99,
		},
		{
			name: "test value -100.994999999999 must be returned -100.99",
			args: args{
				value:    -100.994999999999,
				decimals: 2,
			},
			want: -100.99,
		},
		{
			name: "test value 100.994999999999 must be returned 100.994",
			args: args{
				value:    100.994999999999,
				decimals: 3,
			},
			want: 100.994,
		},
		{
			name: "test value -100.994999999999 must be returned -100.994",
			args: args{
				value:    -100.994999999999,
				decimals: 3,
			},
			want: -100.994,
		},
		{
			name: "test value 19.90 must be returned 19.90",
			args: args{
				value:    19.90,
				decimals: 2,
			},
			want: 19.90,
		},
		{
			name: "test value -19.90 must be returned -190.90",
			args: args{
				value:    -19.90,
				decimals: 2,
			},
			want: -19.90,
		},
		{
			name: "test with value 150.141 must be returned 150.14",
			args: args{
				value:    150.141,
				decimals: 2,
			},
			want: 150.14,
		},
		{
			name: "test with value 150.145234234 must be returned 150.145",
			args: args{
				value:    150.145234234,
				decimals: 3,
			},
			want: 150.145,
		},
		{
			name: "test with value -150.141 must be returned -150.14",
			args: args{
				value:    -150.141,
				decimals: 2,
			},
			want: -150.14,
		},
		{
			name: "test with value 150.145234234 must be returned 150.145",
			args: args{
				value:    -150.145234234,
				decimals: 3,
			},
			want: -150.145,
		},
		{
			name: "test with value 1000000.998 must be returned 1000000.998",
			args: args{
				value:    1000000.998,
				decimals: 2,
			},
			want: 1000000.99,
		},
		{
			name: "test with value -1000000.998 must be returned -1000000.998",
			args: args{
				value:    -1000000.998,
				decimals: 2,
			},
			want: -1000000.99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.args.value, tt.args.decimals); got != tt.want {
				t.Errorf("Trucate() = %v, want %v", got, tt.want)
			}
		})
	}
}
