package redis

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/eucatur/go-toolbox/text"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock/v3"
	"github.com/stretchr/testify/require"
)

func TestClient_Ping(t *testing.T) {
	type fields struct {
		Host           string
		Port           int
		DB             int
		Prefix         string
		PoolConnection redigo.Pool
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

				err := clientForTest.Ping()

				require.Error(t, err)

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

			f := fields{
				PoolConnection: redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return conn, nil
					},
				},
			}

			defer f.PoolConnection.Close()

			connFromPool := f.PoolConnection.Get()

			defer connFromPool.Close()

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&f, conn)
			}

			c := &Client{
				Host:   f.Host,
				Port:   f.Port,
				DB:     f.DB,
				Prefix: f.Prefix,
				pool:   &f.PoolConnection,
			}

			if tt.expectedTest != nil {
				tt.expectedTest(t, conn, cmd, c)
			}
		})
	}
}

func TestClient_Set(t *testing.T) {
	type fields struct {
		Host           string
		Port           int
		DB             int
		Prefix         string
		PoolConnection redigo.Pool
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
				PoolConnection: redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return conn, nil
					},
				},
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			c := Client{
				Host:   f.Host,
				Port:   f.Port,
				DB:     f.DB,
				Prefix: f.Prefix,
				pool:   &f.PoolConnection,
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
		Host           string
		Port           int
		DB             int
		Prefix         string
		PoolConnection redigo.Pool
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
				PoolConnection: redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return conn, nil
					},
				},
			}

			c := Client{
				Host:   f.Host,
				Port:   f.Port,
				DB:     f.DB,
				Prefix: f.Prefix,
				pool:   &f.PoolConnection,
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
		Host           string
		Port           int
		DB             int
		Prefix         string
		PoolConnection redigo.Pool
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
				PoolConnection: redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return conn, nil
					},
				},
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)

			}

			c := Client{
				Host:   f.Host,
				Port:   f.Port,
				DB:     f.DB,
				Prefix: f.Prefix,
				pool:   &f.PoolConnection,
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
		Host           string
		Port           int
		DB             int
		Prefix         string
		PoolConnection redigo.Pool
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
				PoolConnection: redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return conn, nil
					},
				},
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			c := Client{
				Host:   f.Host,
				Port:   f.Port,
				DB:     f.DB,
				Prefix: f.Prefix,
				pool:   &f.PoolConnection,
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
		Host           string
		Port           int
		DB             int
		Prefix         string
		PoolConnection redigo.Pool
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
				PoolConnection: redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return conn, nil
					},
				},
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			c := Client{
				Host:   f.Host,
				Port:   f.Port,
				DB:     f.DB,
				Prefix: f.Prefix,
				pool:   &f.PoolConnection,
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
		Host           string
		Port           int
		DB             int
		Prefix         string
		PoolConnection redigo.Pool
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
				PoolConnection: redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return conn, nil
					},
				},
			}

			if tt.prepareMock != nil {
				cmd = tt.prepareMock(&tt.args, &f, conn)
			}

			c := Client{
				Host:   f.Host,
				Port:   f.Port,
				DB:     f.DB,
				Prefix: f.Prefix,
				pool:   &f.PoolConnection,
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

func TestClient_getConnectionFromPool(t *testing.T) {
	type fields struct {
		Host        string
		Port        int
		DB          int
		Prefix      string
		MaxIdle     int
		MaxActive   int
		IdleTimeout int
		pool        *redigo.Pool

		_connectionStatemented redigo.Conn
	}
	tests := []struct {
		name        string
		wantConn    func(tt *testing.T, f fields, conn redigo.Conn)
		prepareMock func(f *fields, connMock *redigomock.Conn)
	}{
		{
			name: "Falha ao estabelecer conexão irá printar a infromação",
			wantConn: func(tt *testing.T, f fields, conn redigo.Conn) {

				require.Error(tt, conn.Err())

			},
		},
		{
			name: "Conexão com servidor estabelecida com sucesso",
			wantConn: func(tt *testing.T, f fields, conn redigo.Conn) {

				require.NoError(tt, conn.Err())
			},
			prepareMock: func(f *fields, connMock *redigomock.Conn) {

				f.pool = &redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return connMock, nil
					},
				}
			},
		},
		{
			name: "Conexão fechada no pool - deverá retornar uma nova conexão",
			wantConn: func(tt *testing.T, f fields, conn redigo.Conn) {

				require.NoError(tt, conn.Err())
			},
			prepareMock: func(f *fields, connMock *redigomock.Conn) {

				f.pool = &redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return connMock, nil
					},
				}

				defer connMock.Close()

			},
		},
		{
			/*
				ATENÇÃO: ESSE TESTE FAZ-SE NECESSÁRIO RODAR EM INFRA DEVIDO NA RETENTATIVA DE OBTER UM NOVO
				POOL DE CONEXÃO SER DESCONSIDERADO O MOCK PARAMETRIZA E APLICAR UMA CONEXÃO REAL.

				COMO SUGESTÃO PARA RODAR INFRA BASTA LEVANTAR UM CONTAINER DOCKER DO REDIS E PARAMETRIZAR
				AS INFORMAÇÕES DE CONEXÃO NESTE TESTE NA SEÇÃO prepareMock nas linhas:
				f.Host
				f.Port
				f.DB

			*/
			name: "Pool de conexões fechada, iremos instanciar novamente e deverá retornar a conexão",
			wantConn: func(tt *testing.T, f fields, conn redigo.Conn) {

				isCheckResultTest := !text.StringIsEmptyOrWhiteSpace(f.Host) &&
					f.Port > 0 &&
					f.DB > 0

				if isCheckResultTest {

					require.NoError(tt, conn.Err())

				}

			},
			prepareMock: func(f *fields, connMock *redigomock.Conn) {

				f.Host = "localhost"
				f.Port = 6380
				f.DB = 8

				f.pool = &redigo.Pool{
					DialContext: func(ctx context.Context) (redigo.Conn, error) {
						return connMock, nil
					},
				}

				defer f.pool.Close()
			},
		},
		{
			name: "Usando pool de conexão com conexão declarada",
			wantConn: func(tt *testing.T, f fields, conn redigo.Conn) {

				_, err := conn.Do("PING")

				require.NoError(tt, err)

			},
			prepareMock: func(f *fields, connMock *redigomock.Conn) {

				connMock.Command("PING").Expect("PONG")

				f._connectionStatemented = connMock

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn := redigomock.NewConn()

			defer conn.Close()

			f := fields{}

			if tt.prepareMock != nil {
				tt.prepareMock(&f, conn)
			}

			c := &Client{
				Host:            f.Host,
				Port:            f.Port,
				DB:              f.DB,
				Prefix:          f.Prefix,
				MaxIdle:         f.MaxIdle,
				MaxActive:       f.MaxActive,
				IdleTimeout:     f.IdleTimeout,
				pool:            f.pool,
				ConnStatemented: f._connectionStatemented,
			}

			gotConn := c.GetConnectionFromPool()

			tt.wantConn(t, f, gotConn)

		})
	}
}
