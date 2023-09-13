package redis

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock/v3"
	"github.com/stretchr/testify/require"
)

func TestClient_Ping(t *testing.T) {
	type fields struct {
		Host            string
		Port            int
		DB              int
		Prefix          string
		ConnectionRedis redigo.Conn
	}
	tests := []struct {
		name         string
		prepareMock  func(f *fields, connMock *redigomock.Conn) *redigomock.Cmd
		expectedTest func(test *testing.T, connRedigoMock *redigomock.Conn, cmdRedis *redigomock.Cmd, clientForTest *Client)
	}{
		{
			name: "Teste conexao com falha",
			prepareMock: func(f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("PING").ExpectError(errors.New("mockError"))

			},
			expectedTest: func(test *testing.T, connRedigoMock *redigomock.Conn, cmdRedis *redigomock.Cmd, clientForTest *Client) {

				require.Panics(t, func() {
					clientForTest.Ping()
				})

				if connRedigoMock.Stats(cmdRedis) != 1 {
					t.Error("comando não encontrado")
				}
			},
		},
		{
			name: "Teste conexao com successo",
			prepareMock: func(f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("PING").Expect("PONG")

			},
			expectedTest: func(test *testing.T, connRedigoMock *redigomock.Conn, cmdRedis *redigomock.Cmd, clientForTest *Client) {

				require.NotPanics(test, func() {
					clientForTest.Ping()
				})

				if connRedigoMock.Stats(cmdRedis) != 1 {
					t.Error("commando não encontrado")
				}

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn := redigomock.NewConn()

			cmd := &redigomock.Cmd{}

			defer conn.Close()

			f := fields{
				ConnectionRedis: conn,
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&f, conn)
			}

			c := &Client{
				Host:            f.Host,
				Port:            f.Port,
				DB:              f.DB,
				Prefix:          f.Prefix,
				ConnectionRedis: &f.ConnectionRedis,
			}

			if tt.expectedTest != nil {
				tt.expectedTest(t, conn, cmd, c)
			}
		})
	}
}

func TestClient_Set(t *testing.T) {
	type fields struct {
		Host            string
		Port            int
		DB              int
		Prefix          string
		ConnectionRedis redigo.Conn
	}
	type args struct {
		key               string
		value             string
		expirationSeconds int
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		prepareMock func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd
	}{
		{
			name: "Falha ao persistir em cache",
			args: args{
				key:               "FOO-KEY",
				value:             "FOO-VALUE",
				expirationSeconds: 0,
			},
			wantErr: true,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("SET", p.key, p.value).ExpectError(errors.New("mockError"))
			},
		},
		{
			name: "Definir cache",
			args: args{
				key:               "FOO-KEY",
				value:             "FOO-VALUE",
				expirationSeconds: 0,
			},
			wantErr: false,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("SET", p.key, p.value).Expect("success")

			},
		},
		{
			name: "Definir cache com tempo de expiração",
			args: args{
				key:               "FOO-KEY",
				value:             "FOO-VALUE",
				expirationSeconds: 120,
			},
			wantErr: false,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				connMock.Command("SET", p.key, p.value).Expect("success")

				return connMock.Command("EXPIRE", p.key, p.expirationSeconds)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn := redigomock.NewConn()

			cmd := &redigomock.Cmd{}

			defer conn.Close()

			f := fields{
				ConnectionRedis: conn,
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			c := Client{
				Host:            f.Host,
				Port:            f.Port,
				DB:              f.DB,
				Prefix:          f.Prefix,
				ConnectionRedis: &f.ConnectionRedis,
			}
			if err := c.Set(tt.args.key, tt.args.value, tt.args.expirationSeconds); (err != nil) != tt.wantErr {
				t.Errorf("Client.Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			if conn.Stats(cmd) != 1 {
				t.Error("comando não encontrado")
			}

		})
	}
}

func TestClient_Get(t *testing.T) {
	type fields struct {
		Host            string
		Port            int
		DB              int
		Prefix          string
		ConnectionRedis redigo.Conn
	}
	type args struct {
		key string
	}
	tests := []struct {
		name        string
		args        args
		wantValue   string
		wantErr     bool
		prepareMock func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd
	}{
		{
			name: "Obter valor do cache",
			args: args{
				key: "GET-FROM-CACHE",
			},
			wantValue: "FOO FOO",
			wantErr:   false,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("GET", p.key).Expect("FOO FOO")

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn := redigomock.NewConn()

			cmd := &redigomock.Cmd{}

			defer conn.Close()

			f := fields{
				ConnectionRedis: conn,
			}

			c := Client{
				Host:            f.Host,
				Port:            f.Port,
				DB:              f.DB,
				Prefix:          f.Prefix,
				ConnectionRedis: &f.ConnectionRedis,
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			gotValue, err := c.Get(tt.args.key)

			if conn.Stats(cmd) != 1 {
				t.Error("comando não encontrado")
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotValue != tt.wantValue {
				t.Errorf("Client.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestClient_MustGet(t *testing.T) {
	type fields struct {
		Host            string
		Port            int
		DB              int
		Prefix          string
		ConnectionRedis redigo.Conn
	}
	type args struct {
		key string
	}
	tests := []struct {
		name        string
		args        args
		wantValue   string
		wantOk      bool
		prepareMock func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd
	}{
		{
			name: "Nenhuma valor retornado",
			args: args{
				key: "FOO-KEY",
			},
			wantValue: "",
			wantOk:    false,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("GET", p.key).Expect("")
			},
		},
		{
			name: "valor retornado",
			args: args{
				key: "FOO-KEY",
			},
			wantValue: "FOO-VALUE",
			wantOk:    true,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("GET", p.key).Expect("FOO-VALUE")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn := redigomock.NewConn()

			cmd := &redigomock.Cmd{}

			defer conn.Close()

			f := fields{
				ConnectionRedis: conn,
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)

			}

			c := Client{
				Host:            f.Host,
				Port:            f.Port,
				DB:              f.DB,
				Prefix:          f.Prefix,
				ConnectionRedis: &f.ConnectionRedis,
			}

			gotValue, gotOk := c.MustGet(tt.args.key)

			if gotValue != tt.wantValue {
				t.Errorf("Client.MustGet() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Client.MustGet() gotOk = %v, want %v", gotOk, tt.wantOk)
			}

			if conn.Stats(cmd) != 1 {
				t.Error("commando não encontrado")
			}
		})
	}
}

func TestClient_Delete(t *testing.T) {
	type fields struct {
		Host            string
		Port            int
		DB              int
		Prefix          string
		ConnectionRedis redigo.Conn
	}
	type args struct {
		key string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		prepareMock func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd
	}{
		{
			name: "falha Remover valor do cache",
			args: args{
				key: "FOO-KEY",
			},
			wantErr: true,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("DEL", p.key).ExpectError(errors.New("mockError"))

			},
		},
		{
			name: "Remover valor do cache",
			args: args{
				key: "FOO-KEY",
			},
			wantErr: false,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("DEL", p.key).Expect("success")

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn := redigomock.NewConn()

			cmd := &redigomock.Cmd{}

			defer conn.Close()

			f := fields{
				ConnectionRedis: conn,
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			c := Client{
				Host:            f.Host,
				Port:            f.Port,
				DB:              f.DB,
				Prefix:          f.Prefix,
				ConnectionRedis: &f.ConnectionRedis,
			}
			if err := c.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Client.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			if conn.Stats(cmd) != 1 {
				t.Error("comando não encontrado")
			}
		})
	}
}

func TestClient_DeleteLike(t *testing.T) {
	type fields struct {
		Host            string
		Port            int
		DB              int
		Prefix          string
		ConnectionRedis redigo.Conn
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		prepareMock func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd
	}{
		{
			name: "Falha ao obter o valor do cache para remoção por combinação",
			args: args{
				pattern: "FOO",
			},
			wantErr: true,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("SCAN", 0, "MATCH", fmt.Sprintf("*%s*", p.pattern)).ExpectError(errors.New("mockError"))

			},
		},
		{
			name: "Cache encontrado falha ao remover",
			args: args{
				pattern: "FOO",
			},
			wantErr: true,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				value := []interface{}{
					p.pattern,
				}

				connMock.Command("SCAN", 0, "MATCH", fmt.Sprintf("*%s*", p.pattern)).ExpectSlice("FOO-KEY", value).Expect(p.pattern)

				return connMock.Command("DEL", p.pattern).ExpectError(errors.New("mockError"))

			},
		},
		{
			name: "Cache encontrado e removido com sucesso",
			args: args{
				pattern: "FOO",
			},
			wantErr: false,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				value := []interface{}{
					p.pattern,
				}

				connMock.Command("SCAN", 0, "MATCH", fmt.Sprintf("*%s*", p.pattern)).ExpectSlice("FOO-KEY", value).Expect(p.pattern)

				return connMock.Command("DEL", p.pattern).Expect("success")

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn := redigomock.NewConn()

			cmd := &redigomock.Cmd{}

			defer conn.Close()

			f := fields{
				ConnectionRedis: conn,
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			c := Client{
				Host:            f.Host,
				Port:            f.Port,
				DB:              f.DB,
				Prefix:          f.Prefix,
				ConnectionRedis: &f.ConnectionRedis,
			}
			if err := c.DeleteLike(tt.args.pattern); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteLike() error = %v, wantErr %v", err, tt.wantErr)
			}

			if conn.Stats(cmd) != 1 {
				t.Error("comando não encontrado")
			}
		})
	}
}

func TestClient_Do(t *testing.T) {
	type fields struct {
		Host            string
		Port            int
		DB              int
		Prefix          string
		ConnectionRedis redigo.Conn
	}
	type args struct {
		comando string
		args    []interface{}
	}
	tests := []struct {
		name        string
		args        args
		want        interface{}
		wantErr     bool
		prepareMock func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd
	}{
		{
			name: "Falha ao executar o comando",
			args: args{
				comando: "GET",
				args:    []interface{}{},
			},
			want:    nil,
			wantErr: true,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("GET", p.args...).ExpectError(errors.New("mockError"))

			},
		},
		{
			name: "Comando executado com sucesso",
			args: args{
				comando: "GET",
				args: []interface{}{
					"FOO-KEY",
				},
			},
			want:    "FOO-VALUE",
			wantErr: false,
			prepareMock: func(p *args, f *fields, connMock *redigomock.Conn) *redigomock.Cmd {

				return connMock.Command("GET", p.args...).Expect("FOO-VALUE")

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn := redigomock.NewConn()

			cmd := &redigomock.Cmd{}

			defer conn.Close()

			f := fields{
				ConnectionRedis: conn,
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			c := Client{
				Host:            f.Host,
				Port:            f.Port,
				DB:              f.DB,
				Prefix:          f.Prefix,
				ConnectionRedis: &f.ConnectionRedis,
			}
			got, err := c.Do(tt.args.comando, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Do() = %v, want %v", got, tt.want)
			}

			if conn.Stats(cmd) != 1 {
				t.Error("comando não localizado")
			}
		})
	}
}
