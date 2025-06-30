package time

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func mockDateTimeTimezone(t *testing.T, dtValue, iana string) time.Time {

	dt := mockDateTimeNowServer(t, dtValue)

	location := mockLoadLocationTest(t, iana)

	getLocation := dt.In(location).Format("Z07:00")

	dtParsed, err := time.Parse(fmt.Sprintf("%s %sZ07:00", time.DateOnly, time.TimeOnly), fmt.Sprintf("%s%s", dt.Format(time.DateTime), getLocation))

	if err != nil {
		t.Fatal(err)
	}

	return dtParsed

}

func mockDateTimeNowServer(t *testing.T, dateTimeExpected string) time.Time {

	mockDtNow := GetTryParseDate(dateTimeExpected, "")

	return mockDtNow

}

func mockLoadLocationTest(t *testing.T, iana string) *time.Location {

	location, err := time.LoadLocation(iana)

	if err != nil {
		t.Fatal(err)
	}

	return location
}

func TestSet(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type args struct {
		valueTime   time.Time
		valueString string
	}
	tests := []struct {
		debug bool
		name  string
		args  args
		want  DateTime
	}{
		{
			name: "valores totalmente inválidos",
			args: args{
				valueTime:   time.Time{}.In(&time.Location{}),
				valueString: "skdjflskdjflksjdf",
			},
			want: "",
		},
		{
			name: "DataHora mínima",
			args: args{
				valueTime:   mockDateTimeNowServer(t, "0001-01-01-01 00:00:00"),
				valueString: "0001-01-01 00:00:00",
			},
			want: "",
		},
		{
			name: "DataHora mínima com TZ",
			args: args{
				valueTime:   mockDateTimeNowServer(t, "0001-01-01-01T00:00:00Z"),
				valueString: "0001-01-01T00:00:00Z",
			},
			want: "",
		},
		{
			name: "DataHora mínima com timezone definido",
			args: args{
				valueTime:   mockDateTimeNowServer(t, "0001-01-01-01T00:00:00-04:00"),
				valueString: "0001-01-01T00:00:00-04:00",
			},
			want: "",
		},
		{
			name: "DataHora válida com timezone específico",
			args: args{
				valueTime:   mockDateTimeTimezone(t, "2023-11-28 11:16:35", "America/Porto_Velho"),
				valueString: "2023-11-28T11:16:35-04:00",
			},
			want: "2023-11-28T11:16:35-04:00",
		},
		{
			name: "DataHora válida",
			args: args{
				valueTime:   mockDateTimeTimezone(t, "2023-11-28 11:19:14", "America/Sao_Paulo"),
				valueString: "2023-11-28 11:19:14-03:00",
			},
			want: "2023-11-28T11:19:14-03:00",
		},
		{
			debug: true,
			name:  "Data dia 10/06/2024 com fuso horário -03:00",
			args: args{
				valueTime:   mockDateTimeTimezone(t, "2024-06-10 00:00:00", "America/Sao_Paulo"),
				valueString: "10/06/2024",
			},
			want: "2024-06-10T00:00:00-03:00",
		},
	}
	for _, tt := range tests {
		if !tt.debug {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := Set(tt.args.valueTime); got != tt.want {
				t.Errorf("Set() = %#v, want %#v", got, tt.want)
			}
			if got := Set(tt.args.valueString); got != tt.want {
				t.Errorf("Set() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestDateTime_SetLocation(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type fields struct {
		datetime time.Time
		location time.Location
		timezone string
		date     string
		time     string
	}
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name     string
		args     args
		fields   fields
		expected func(oldDateTime, newDateTime DateTime)
	}{
		{
			name: "mesmo timezone",
			args: args{
				location: mockLoadLocationTest(t, "America/Porto_Velho"),
			},
			fields: fields{
				datetime: mockDateTimeTimezone(t, "2023-09-23 17:00:20", "America/Porto_Velho"),
				location: *mockLoadLocationTest(t, "America/Porto_Velho"),
				timezone: "-04:00",
				date:     "2023-09-23",
				time:     "17:00:20",
			},
			expected: func(stroldDateTime, strnewDateTime DateTime) {

				oldDateTime := tryParse(string(stroldDateTime))
				newDateTime := tryParse(string(strnewDateTime))

				require.Equal(t, oldDateTime.datetime, newDateTime.datetime)
				require.Equal(t, oldDateTime.location, newDateTime.location)
				require.Equal(t, oldDateTime.timezone, newDateTime.timezone)
				require.Equal(t, oldDateTime.date, newDateTime.date)
				require.Equal(t, oldDateTime.time, newDateTime.time)

			},
		},
		{
			name: "Mudança de location e timezone",
			args: args{
				location: mockLoadLocationTest(t, "America/Porto_Velho"),
			},
			fields: fields{
				datetime: mockDateTimeTimezone(t, "2023-09-23 17:00:20", "America/Sao_Paulo"),
				location: *mockLoadLocationTest(t, "America/Sao_Paulo"),
				timezone: "-03:00",
				date:     "2023-09-23",
				time:     "17:00:20",
			},
			expected: func(stroldDateTime, strnewDateTime DateTime) {

				oldDateTime := tryParse(string(stroldDateTime))
				newDateTime := tryParse(string(strnewDateTime))

				require.NotEqual(t, oldDateTime.datetime, newDateTime.datetime)
				require.NotEqual(t, oldDateTime.location, newDateTime.location)
				require.NotEqual(t, oldDateTime.timezone, newDateTime.timezone)
				require.Equal(t, oldDateTime.date, newDateTime.date)
				require.Equal(t, oldDateTime.time, newDateTime.time)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := Set(tt.fields.datetime)

			oldDt := dt

			dt.SetLocation(tt.args.location)

			if tt.expected != nil {
				tt.expected(oldDt, dt)
			}
		})
	}
}

func TestDateTime_GetDateTimeWithTimezone(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type fields struct {
		datetime time.Time
		location time.Location
		timezone string
		date     string
		time     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Obter a data e hora completa com fuso horário",
			fields: fields{
				datetime: mockDateTimeTimezone(t, "2023-09-23 17:22:58", "America/Porto_Velho"),
				location: *mockLoadLocationTest(t, "America/Porto_Velho"),
				timezone: "-04:00",
				date:     "2023-09-23",
				time:     "17:22:58",
			},
			want: "2023-09-23T17:22:58-04:00",
		},
		{
			name:   "Valor inicial",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dt := Set(tt.fields.datetime)

			if got := dt.GetDateTimeWithTimezone(); got != tt.want {
				t.Errorf("DateTime.GetDateTimeWithTimezone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_MarshalJSON(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type fields struct {
		datetime time.Time
		location time.Location
		timezone string
		date     string
		time     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Converter parar retorno JSON",
			fields: fields{
				datetime: mockDateTimeTimezone(t, "2023-09-23 17:22:58", "America/Porto_Velho"),
				location: *mockLoadLocationTest(t, "America/Porto_Velho"),
				timezone: "-04:00",
				date:     "2023-09-23",
				time:     "17:22:58",
			},
			want:    []byte(`"2023-09-23T17:22:58-04:00"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dt := Set(tt.fields.datetime)

			got, err := dt.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_UnmarshalJSON(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type fields struct {
		datetime time.Time
		location time.Location
		timezone string
		date     string
		time     string
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Converter informação com data e hora mínimo",
			fields: fields{
				datetime: mockDateTimeTimezone(t, "2023-09-23 17:22:58", "America/Porto_Velho"),
				location: *mockLoadLocationTest(t, "America/Porto_Velho"),
				timezone: "-04:00",
				date:     "2023-09-23",
				time:     "17:22:58",
			},
			args: args{
				bytes: []byte("0001-01-01T00:00:00Z"),
			},
			wantErr: false,
		},
		{
			name: "Informação válida",
			fields: fields{
				datetime: mockDateTimeTimezone(t, "2023-09-23 17:22:58", "America/Porto_Velho"),
				location: *mockLoadLocationTest(t, "America/Porto_Velho"),
				timezone: "-04:00",
				date:     "2023-09-23",
				time:     "17:22:58",
			},
			args: args{
				bytes: []byte(`2023-09-23 17:44:38`),
			},
			wantErr: false,
		},
		{
			name: "Informação válida com aspas",
			fields: fields{
				datetime: mockDateTimeTimezone(t, "2023-09-23 17:22:58", "America/Porto_Velho"),
				location: *mockLoadLocationTest(t, "America/Porto_Velho"),
				timezone: "-04:00",
				date:     "2023-09-23",
				time:     "17:22:58",
			},
			args: args{
				bytes: []byte(`"2023-09-23 17:44:38"`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dt := Set(tt.fields.datetime)

			if err := dt.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDateTime_Value(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type fields struct {
		datetime time.Time
		location time.Location
		timezone string
		date     string
		time     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "Retorno para persistência em DB",
			fields: fields{
				datetime: mockDateTimeTimezone(t, "2023-09-23 17:22:58", "America/Porto_Velho"),
				location: *mockLoadLocationTest(t, "America/Porto_Velho"),
				timezone: "-04:00",
				date:     "2023-09-23",
				time:     "17:22:58",
			},
			want:    "2023-09-23T17:22:58-04:00",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := Set(tt.fields.datetime)

			got, err := dt.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTime.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_Scan(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type args struct {
		value interface{}
	}
	tests := []struct {
		name        string
		dt          DateTime
		args        args
		wantErr     bool
		want        DateTime
		prepareMock func()
	}{
		{
			name: "Informação time.Time",
			dt:   "",
			args: args{
				value: mockDateTimeTimezone(t, "2023-09-23 17:22:58", "America/Porto_Velho"),
			},
			wantErr: false,
			want:    "2023-09-23T17:22:58-04:00",
		},
		{
			name: "Informação string",
			dt:   "",
			args: args{
				value: "2023-09-23 17:22:58",
			},
			wantErr: false,
			want:    DateTime(mockDateTimeTimezone(t, "2023-09-23T17:22:58", "Local").Format(time.RFC3339)),
		},
		{
			name: "Informação []uint8",
			dt:   "",
			args: args{
				value: []uint8("2023-11-28 14:17:20"),
			},
			wantErr: false,
			want:    DateTime(mockDateTimeTimezone(t, "2023-11-28 14:17:20", "Local").Format(time.RFC3339)),
		},
		{
			name: "Informação inválida - valor mínimo",
			dt:   "",
			args: args{
				value: "0001-01-01 00:00:00",
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "Informação nil",
			dt:   "",
			args: args{
				value: nil,
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "Tipo não tratado",
			dt:   "",
			args: args{
				value: math.MaxInt64,
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "Retornar com timezone definido em CurrentTime",
			dt:   "",
			args: args{
				value: "2023-11-28 14:17:20",
			},
			wantErr: false,
			want:    "2023-11-28T14:17:20-03:00",
			prepareMock: func() {

				SetLocationInCurrentTime(mockLoadLocationTest(t, "America/Sao_Paulo"))

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.prepareMock != nil {
				tt.prepareMock()
			}

			if err := tt.dt.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.dt, tt.want) {
				t.Errorf("DateTime.Scan() value = %#v, want %#v", tt.dt, tt.want)
			}
		})
	}
}

func TestDateTime_IsValid(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	tests := []struct {
		name    string
		args    DateTime
		wantErr bool
	}{
		{
			name:    "Date time inválido com timezone",
			args:    "0001-01-01T00:00:00Z",
			wantErr: true,
		},
		{
			name:    "Data e hora impossível interpretar",
			args:    "skldfjsldkfjsldkfjslkdfj",
			wantErr: true,
		},
		{
			name:    "Data e hora inválida",
			args:    "0001-01-01 00:00:00",
			wantErr: true,
		},
		{
			name:    "Data e hora válida sem fuso horário",
			args:    "2023-11-27 22:25:21",
			wantErr: false,
		},
		{
			name:    "Data e hora válida com fuso horário",
			args:    "2023-11-27T22:25:21-04:00",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dt := tt.args

			if err := dt.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_MarshalJSON(t *testing.T) {

	defer func() {
		CurrentTime = time.Now
	}()

	type subfields struct {
		DateTimeFOO DateTime `json:"date_time_foo,omitempty"`
	}

	type fields struct {
		Name         string     `json:"name,omitempty"`
		Dob          Date       `json:"dob,omitempty"`
		CreatedAt    DateTime   `json:"created_at,omitempty"`
		AnotherField string     `json:"another_field,omitempty"`
		SubFields    *subfields `json:"sub_fields,omitempty"`
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Campos declarados vazio",
			fields: fields{
				Name:         "EMPTY FIELDS",
				Dob:          "",
				CreatedAt:    "",
				AnotherField: "",
			},
			want:    []byte(`{"name":"EMPTY FIELDS"}`),
			wantErr: false,
		},
		{
			name: "Declaração dentro de outra struct",
			fields: fields{
				Name:      "TESTE",
				SubFields: &subfields{},
			},
			want:    []byte(`{"name":"TESTE","sub_fields":{}}`),
			wantErr: false,
		},
		{
			name: "Apenas com um campo preenchido",
			fields: fields{
				Name: "TESTE OF TESTES",
			},
			want:    []byte(`{"name":"TESTE OF TESTES"}`),
			wantErr: false,
		},
		{
			name: "Campo dob e o outro campo não deve ser retornado",
			fields: fields{
				Name:      "Teste",
				CreatedAt: Set("2023-11-27 22:19:13"),
			},
			want:    []byte(fmt.Sprintf(`{"name":"Teste","created_at":"%s"}`, mockDateTimeTimezone(t, "2023-11-27 22:19:13", "Local").Format(time.RFC3339))),
			wantErr: false,
		},
		{
			name: "Converter para JSON tudo",
			fields: fields{
				Name:         "Teste",
				Dob:          Date(Set("2001-11-27")),
				CreatedAt:    Set("2023-11-27 19:53:19"),
				AnotherField: "foo-foo",
			},
			want:    []byte(fmt.Sprintf(`{"name":"Teste","dob":"2001-11-27","created_at":"%s","another_field":"foo-foo"}`, mockDateTimeTimezone(t, "2023-11-27 19:53:19", "Local").Format(time.RFC3339))),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			content, err := json.Marshal(tt.fields)

			require.NoError(t, err)

			require.Nil(t, err)

			require.Equal(t, tt.want, content, fmt.Sprintf("want: %s | Content: %s", string(tt.want), content))

		})
	}

}

func Test_tryParse(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type args struct {
		dateTime string
	}
	tests := []struct {
		name        string
		args        args
		want        tpDateTime
		prepareMock func()
	}{
		{
			name: "Data e hora totalmente inválida e zuada",
			args: args{
				dateTime: "l4k32nrdnfkjlq3brkjl43rbwjklef",
			},
			want: tpDateTime{},
		},
		{
			name: "Data e hora inválida",
			args: args{
				dateTime: "0001-01-01 00:00:00",
			},
			want: tpDateTime{},
		},
		{
			name: "Data e hora sem timezone",
			args: args{
				dateTime: "2023-11-27 21:45:31",
			},
			want: tpDateTime{
				datetime: mockDateTimeTimezone(t, "2023-11-27 21:45:31", "Local").In(mockLoadLocationTest(t, "Local")),
				location: mockLoadLocationTest(t, "Local"),
				timezone: mockDateTimeTimezone(t, "2023-11-27 21:45:31", "Local").Format(timezoneformat),
				date:     "2023-11-27",
				time:     "21:45:31",
			},
		},
		{
			name: "Data e hora com timezone",
			args: args{
				dateTime: mockDateTimeTimezone(t, "2023-11-27 21:52:02", "UTC").Format(time.RFC3339),
			},
			want: tpDateTime{
				datetime: mockDateTimeTimezone(t, "2023-11-27 21:52:02", "UTC").In(mockLoadLocationTest(t, "UTC")),
				location: mockLoadLocationTest(t, "UTC"),
				timezone: mockDateTimeTimezone(t, "2023-11-27 21:52:02", "UTC").Format(timezoneformat),
				date:     "2023-11-27",
				time:     "21:52:02",
			},
		},
		{
			name: "Date time com timezone string",
			args: args{
				dateTime: "2023-11-27T21:52:02-03:00",
			},
			want: tpDateTime{
				datetime: mockDateTimeTimezone(t, "2023-11-27 21:52:02", "America/Sao_Paulo").In(mockLoadLocationTest(t, "America/Sao_Paulo")),
				location: mockLoadLocationTest(t, "America/Sao_Paulo"),
				timezone: mockDateTimeTimezone(t, "2023-11-27 21:52:02", "America/Sao_Paulo").Format(timezoneformat),
				date:     "2023-11-27",
				time:     "21:52:02",
			},
			prepareMock: func() {
				_toolParse = nil
				SetLocationInCurrentTime(mockLoadLocationTest(t, "America/Sao_Paulo"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareMock != nil {
				tt.prepareMock()
			}
			if got := tryParse(tt.args.dateTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tryParse() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestDateTime_ChangeTimezone(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name  string
		strdt DateTime
		args  args
		want  DateTime
	}{
		{
			name:  "Alteração de fuso horário",
			strdt: "2023-12-01T22:12:30-03:00",
			args: args{
				location: mockLoadLocationTest(t, "America/Porto_Velho"),
			},
			want: "2023-12-01T21:12:30-04:00",
		},
		{
			name:  "Alteração de fuso horário com muita diferença",
			strdt: "2023-12-02T02:21:40+01:00",
			args: args{
				location: mockLoadLocationTest(t, "America/Porto_Velho"),
			},
			want: "2023-12-01T21:21:40-04:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.strdt.ChangeTimezone(tt.args.location)

			if !reflect.DeepEqual(tt.strdt, tt.want) {
				t.Errorf("DateTime.ChangeTimezone got = %v want = %v", tt.strdt, tt.want)
			}
		})
	}
}

func TestPeriod_PeriodValid(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	tests := []struct {
		name   string
		period Period
		want   bool
	}{
		{
			name:   "Nenhum período informado",
			period: []DateTime{},
			want:   false,
		},
		{
			name: "Período não informados não satisfatórios",
			period: []DateTime{
				"2023-12-04 20:05:33",
			},
			want: false,
		},
		{
			name: "Período com informações inválidas",
			period: []DateTime{
				"2023-12-04 20:06:09",
				"dsfisdjfsdjosdvjfdhfkjs",
			},
			want: false,
		},
		{
			name: "Período válido",
			period: []DateTime{
				"2023-01-04 20:06:37",
				"2023-12-04 20:06:35",
			},
			want: true,
		},
		{
			name: "Período com apenas data válido",
			period: []DateTime{
				"2023-01-04",
				"2023-12-04",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.period.PeriodValid(); got != tt.want {
				t.Errorf("Period.PeriodValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseToFullYear(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	type args struct {
		year int64
	}
	tests := []struct {
		name        string
		args        args
		want        int64
		prepareMock func()
	}{
		{
			name: "23 -> 2023",
			args: args{
				year: 23,
			},
			want:        2023,
			prepareMock: nil,
		},
		{
			name: "2034 -> 2034",
			args: args{
				year: 2034,
			},
			want:        2034,
			prepareMock: nil,
		},
		{
			name: "101 -> 2101",
			args: args{
				year: 101,
			},
			want:        2101,
			prepareMock: nil,
		},
		{
			name: "2201 -> 2201",
			args: args{
				year: 2201,
			},
			want:        2201,
			prepareMock: nil,
		},
		{
			name: "00 -> 2000",
			args: args{
				year: 0,
			},
			want:        2000,
			prepareMock: nil,
		},
		{
			name: "01 -> 2001",
			args: args{
				year: 01,
			},
			want:        2001,
			prepareMock: nil,
		},
		{
			name: "3001 -> 3001",
			args: args{
				year: 3001,
			},
			want: 3001,
			prepareMock: func() {
				CurrentTime = func() time.Time {
					return time.Date(3000, time.December, 5, 9, 37, 59, 0, mockLoadLocationTest(t, "America/Porto_Velho"))
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseToFullYear(tt.args.year); got != tt.want {
				t.Errorf("ParseToFullYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_GetFullYear(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	tests := []struct {
		name string
		dt   DateTime
		want int64
	}{
		{
			name: "Obter o ano completo a partir de ano com 2 dígitos",
			dt:   Set("24-01-03"),
			want: 2024,
		},
		{
			name: `Obter o ano completo a partir de quatro dígitos`,
			dt:   Set("2025-01-03"),
			want: 2025,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.GetFullYear(); got != tt.want {
				t.Errorf("DateTime.GetFullYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_GetShortYear(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	tests := []struct {
		name string
		dt   DateTime
		want int64
	}{
		{
			name: "Obter o ano encurtado a partir de ano com 2 dígitos",
			dt:   Set("24-01-03"),
			want: 24,
		},
		{
			name: `Obter o ano encurtado a partir de quatro dígitos`,
			dt:   Set("2025-01-03"),
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.GetShortYear(); got != tt.want {
				t.Errorf("DateTime.GetShortYear() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestDateTime_GetMonth(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	tests := []struct {
		name string
		dt   DateTime
		want string
	}{
		{
			name: "Obter mês a partir de um dígito",
			dt:   Set("2024-3-13"),
			want: "03",
		},
		{
			name: "Obter mês a partir de dois dígitos",
			dt:   Set("2024-06-13"),
			want: "06",
		},
		{
			name: "Obter mês quando ele for dois dígitos",
			dt:   Set("2024-11-13"),
			want: "11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.GetMonth(); got != tt.want {
				t.Errorf("DateTime.GetMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_String(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	tests := []struct {
		name  string
		strdt DateTime
		want  string
	}{
		{
			name:  "Obter string da data e hora no formato ISO 8601 a partir de uma data informada no formato ISO 8601",
			strdt: Set("2024-03-13 18:36:01"),
			want:  "2024-03-13 18:36:01",
		},
		{
			name:  "Obter string da data e hora no forma ISO 8601 a partir de uma formato esperado pela lib",
			strdt: Set("24-03-13 18:39:39"),
			want:  "2024-03-13 18:39:39",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.strdt.String(); got != tt.want {
				t.Errorf("DateTime.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_GetTimezone(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	tests := []struct {
		name  string
		strdt DateTime
		want  string
	}{
		{
			name:  "Obter o timezone que foi definido",
			strdt: Set("2024-03-13 18:40:28-03:00"),
			want:  "-03:00",
		},
		{
			name:  "Timezone não retornado quando nada é definido",
			strdt: Set(""),
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.strdt.GetTimezone(); got != tt.want {
				t.Errorf("DateTime.GetTimezone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_GetTime(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()
	tests := []struct {
		name  string
		strdt DateTime
		want  string
	}{
		{
			name:  "Obter a hora",
			strdt: Set("2024-03-13 18:42:47"),
			want:  "18:42:47",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.strdt.GetTime(); got != tt.want {
				t.Errorf("DateTime.GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UseDateTimeMethods(t *testing.T) {
	defer func() {
		CurrentTime = time.Now
	}()

	const dateTimeUsedForTest = "2024-03-16 08:45:29"

	dateTimeForTest := Set(dateTimeUsedForTest)

	tests := []struct {
		name    string
		runTest func(test *testing.T, dt *DateTime)
	}{
		{
			name: `O método IsValid() deverá validar se a data definida é válida ou não.
			> Neste teste os valores serão manipulados para teste do método IsValid()
			Após os teste o valor será retornado para não afetar os outros cenários`,
			runTest: func(test *testing.T, dt *DateTime) {

				require.NoError(test, dt.IsValid())

				*dt = "kl12j31kl2j3"

				require.Error(test, dt.IsValid())

				*dt = dateTimeUsedForTest

			},
		},
		{
			name: `O método String() deve retornar a data que foi definida sem quais outras
			informações como timezone ou outro formato`,
			runTest: func(test *testing.T, dt *DateTime) {

				expected := "2024-03-16 08:45:29"

				require.Equal(test, expected, dt.String())

			},
		},
		{
			name: `O método SetLocation() deve retornar a mesma data que foi definida mudando
			apenas o fuso horário
			> Utilizado o método GetDateTimeWithTimezone() para garantir o retorno no padrão 
			RFC339`,
			runTest: func(test *testing.T, dt *DateTime) {

				dt.SetLocation(mockLoadLocationTest(t, "America/Sao_Paulo"))

				expected := "2024-03-16T08:45:29-03:00"

				require.Equal(test, expected, dt.GetDateTimeWithTimezone())

				dt.SetLocation(mockLoadLocationTest(test, "America/Porto_Velho"))

				expected = "2024-03-16T08:45:29-04:00"

				require.Equal(test, expected, dt.GetDateTimeWithTimezone())

				dt.SetLocation(mockLoadLocationTest(test, "America/Rio_Branco"))

				expected = "2024-03-16T08:45:29-05:00"

				require.Equal(test, expected, dt.GetDateTimeWithTimezone())

			},
		},
		{
			name: `O método GetDateTimeWithTimezone() deve retornar a mesma data com a
			informação do fuso horário no padrão RFC3339.
			> Uso do método SetLocation() para compatibilidade em outras regiões ao rodar o
			teste`,
			runTest: func(test *testing.T, dt *DateTime) {

				expected := "2024-03-16T08:45:29-03:00"

				dt.SetLocation(mockLoadLocationTest(test, "America/Sao_Paulo"))

				require.Equal(test, expected, dt.GetDateTimeWithTimezone())

			},
		},
		{
			name: `Alteração de fuso horário com ChangeTimezone() deve se basear a partir do fuso
			horário inicial.
			> Uso do método SetLocation() para compatibilidade em outras regiões ao rodar o
			teste`,
			runTest: func(test *testing.T, dt *DateTime) {

				dt.SetLocation(mockLoadLocationTest(test, "America/Sao_Paulo"))

				expected := "2024-03-16T06:45:29-05:00"

				dt.ChangeTimezone(mockLoadLocationTest(test, "America/Rio_Branco"))

				require.Equal(t, expected, dt.GetDateTimeWithTimezone())

			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if tt.runTest != nil {
				tt.runTest(t, &dateTimeForTest)
			} else {
				t.Fatal("nenhum test implementado")
			}

		})

	}

}

func Test_OrderedWithGetStdTime(t *testing.T) {

	t.Skip("teste em andamento")

	type fakeStruct struct {
		ID       int64
		Name     string
		Document string
		Dob      string
	}

	fakeDatas := []fakeStruct{
		{
			ID:       1,
			Name:     "Fulano",
			Document: "123456789",
			Dob:      "26/01/2024",
		},
		{
			ID:       2,
			Name:     "Ciclano",
			Document: "987654321",
			Dob:      "22/08/2024",
		},
		{
			ID:       3,
			Name:     "Beltrano",
			Document: "456789123",
			Dob:      "27/09/2024",
		},
		{
			ID:       4,
			Name:     "Zé",
			Document: "321654987",
			Dob:      "10/06/2024",
		},
		{
			ID:       5,
			Name:     "Maria",
			Document: "654321789",
			Dob:      "22/08/2024",
		},
		{
			ID:       6,
			Name:     "João",
			Document: "789123456",
			Dob:      "16/04/2024",
		},
	}

	data1006 := Set("10/06/2024")
	data2208 := Set("22/08/2024")

	require.True(t, data1006.GetStdTime().Before(data2208.GetStdTime()), "A data 10/06/2024 não é menor que 22/08/2024")

	return

	expectedsDobsOrderedAsc := []string{
		"26/01/2024",
		"16/04/2024",
		"10/06/2024",
		"22/08/2024",
		"22/08/2024",
		"27/09/2024",
	}

	sort.Slice(fakeDatas, func(i, j int) bool {
		return Set(fakeDatas[i].Dob).GetStdTime().Before(Set(fakeDatas[j].Dob).GetStdTime())
	})

	gotDobsOrdered := []string{}

	for _, fakeData := range fakeDatas {

		gotDobsOrdered = append(gotDobsOrdered, fakeData.Dob)
	}

	require.Equal(t, expectedsDobsOrderedAsc, gotDobsOrdered, "As datas não estão ordenadas corretamente. Esperado: %v, Obtido: %v", expectedsDobsOrderedAsc, gotDobsOrdered)

}

func TestDateTime_LastDayMonth(t *testing.T) {
	tests := []struct {
		name     string
		input    DateTime
		expected int
	}{
		{
			name:     "data inválida",
			input:    Set("0001-01-01 00:00:00"),
			expected: 0,
		},
		{
			name:     "data válida - mês com 31 dias",
			input:    Set("2025-07-15"),
			expected: 31,
		},
		{
			name:     "data válida - mês com 30 dias",
			input:    Set("2025-06-10"),
			expected: 30,
		},
		{
			name:     "data com ano anterior ao atual",
			input:    Set("2020-02-10"),
			expected: 29,
		},
		{
			name:     "dia anterior ao atual",
			input:    Set("2025-06-29"),
			expected: 30,
		},
		{
			name:     "ano bissexto",
			input:    Set("2024-02-10"),
			expected: 29,
		},
		{
			name:     "ano não bissexto",
			input:    Set("2023-02-10"),
			expected: 28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.LastDayMonth()
			if result != tt.expected {
				t.Errorf("LastDayMonth() = %d, want %d", result, tt.expected)
			}
		})
	}
}
