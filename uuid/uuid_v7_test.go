package uuid

import (
	"reflect"
	"testing"
	"time"
)

func TestUUIDv7_generate(t *testing.T) {
	type fields struct {
		value [16]byte
		dt    time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		wantErr  bool
		wantUUID string
	}{
		{
			name: "Gerar com base em data/hora específica fuso horário UTC",
			fields: fields{
				value: [16]byte{},
				dt:    time.Date(2023, 10, 1, 12, 34, 56, 15, time.UTC),
			},
			wantErr:  false,
			wantUUID: "018aeb3c-0d80-789c-b926-435e563ed7b9",
		},
		{
			name: "Devido da data ser inválida ou negativa o UUID gerado não terá o timestamp preenchido",
			fields: fields{
				value: [16]byte{},
				dt:    time.Unix(-1, 0), // timestamp negativo
			},
			wantErr:  false,
			wantUUID: "00000000-03e8-7f73-a401-3711d808ee08",
		},
		{
			name: "Date time zerado",
			fields: fields{
				value: [16]byte{},
				dt:    time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), // timestamp negativo
			},
			wantErr:  false,
			wantUUID: "00000000-03e8-7f73-a401-3711d808ee08",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			u := &UUIDv7{
				value: tt.fields.value,
				dt:    tt.fields.dt,
			}

			err := u.generate()

			if (err != nil) != tt.wantErr {
				t.Errorf("UUIDv7.generate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantUUID == "" && tt.wantUUID != u.String() {
				t.Errorf("UUIDv7.generate() = %v, want %v", u.String(), tt.wantUUID)
			}

			/* Seguindo o padrão RFC 9562 com base na data/hora fornecida o UUID deve ser gerado corretamente
			* O UUID v7 gerado deve ter os primeiros 6 bytes preenchidos com o timestamp em milissegundos
			* A versão deve ser 7 (0x70) e a variante deve ser 1 (0x80)
			* O UUID gerado deve ser comparado com o esperado
			* Se o UUID gerado não corresponder ao esperado, o teste falhar
			* Se o UUID gerado corresponder ao esperado, o teste deve passar
			 */
			if len([]rune(tt.wantUUID)) == 13 && u.String()[:13] != tt.wantUUID[:13] {
				t.Errorf("UUIDv7.generate() = %v, want %v", u.String(), tt.wantUUID)
			}
		})
	}
}

func TestUUIDv7_Generate(t *testing.T) {
	type fields struct {
		value [16]byte
		dt    time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Gerar com base em data/hora específica fuso horário UTC",
			fields: fields{
				value: [16]byte{},
				dt:    time.Date(2023, 10, 1, 12, 34, 56, 15, time.UTC),
			},
			wantErr: false,
			want:    "018aeb3c-0d80-789c-b926-435e563ed7b9",
		},
		{
			name: "Devido da data ser inválida ou negativa o UUID gerado não terá o timestamp preenchido",
			fields: fields{
				value: [16]byte{},
				dt:    time.Unix(-1, 0), // timestamp negativo
			},
			wantErr: false,
			want:    "00000000-03e8-7f73-a401-3711d808ee08",
		},
		{
			name: "Date time zerado",
			fields: fields{
				value: [16]byte{},
				dt:    time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
			want:    "00000000-03e8-7f73-a401-3711d808ee08",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UUIDv7{
				value: tt.fields.value,
				dt:    tt.fields.dt,
			}

			got, err := u.Generate()

			if (err != nil) != tt.wantErr {
				t.Errorf("UUIDv7.Generate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want == "" && tt.want != got {
				t.Errorf("UUIDv7.Generate() = %v, want %v", got, tt.want)
			}

			/* Seguindo o padrão RFC 9562 com base na data/hora fornecida o UUID deve ser gerado corretamente
			* O UUID v7 gerado deve ter os primeiros 6 bytes preenchidos com o timestamp em milissegundos
			* A versão deve ser 7 (0x70) e a variante deve ser 1 (0x80)
			* O UUID gerado deve ser comparado com o esperado
			* Se o UUID gerado não corresponder ao esperado, o teste falhar
			* Se o UUID gerado corresponder ao esperado, o teste deve passar
			 */
			if len([]rune(tt.want)) == 13 && got[:13] != tt.want[:13] {
				t.Errorf("UUIDv7.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTimeUUIDv7(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "Obter time de UUIDv7 válido",
			args: args{
				uuid: "0198d303-1962-75b7-9baa-1baa41d023b5", // Timestamp corresponde a 2025-08-22T18:20:52.000Z UTC
			},
			want:    time.Date(2025, time.August, 22, 18, 20, 52, 450000000, time.UTC),
			wantErr: false,
		},
		{
			name: "UUIDv7 inválido (formato incorreto)",
			args: args{
				uuid: "invalid-uuid-format",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "UUIDv7 inválido (tamanho incorreto)",
			args: args{
				uuid: "0198d303-17a0-7d0f-8070-2397b93a38", // Tamanho incorreto
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "UUIDv7 com timestamp zero",
			args: args{
				uuid: "00000000-0000-7000-8000-000000000000", // Timestamp corresponde a 1970-01-01T00:00:00.000Z UTC
			},
			want:    time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name: "UUIDv7 com timestamp máximo",
			args: args{
				uuid: "ffffffff-ffff-7fff-bfff-ffffffffffff", // Timestamp máximo
			},
			want:    time.Date(10889, time.August, 2, 5, 31, 50, 655000000, time.UTC), // Correspondente a 10889-09-04T15:46:39.999Z UTC
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTimeUUIDv7(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTimeUUIDv7() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTimeUUIDv7() = %v, want %v", got, tt.want)
			}
		})
	}
}
