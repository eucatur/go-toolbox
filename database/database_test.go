package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func TestGetByFile(t *testing.T) {
	_, err := GetByFile("mysql-example.json")
	if err != nil {
		t.Error(err)
	}

	_, err = GetByFile("postgres-example.json")
	if err != nil {
		t.Error(err)
	}

	_, err = GetByFile("sqlite3-example.json")
	if err != nil {
		t.Error(err)
	}

	_, err = GetByFile("no-file-example.json")
	if err == nil {
		t.Error("The GetByFile function should not find the file.")
	}
}

func TestSetConnectionByFile(t *testing.T) {
	type args struct {
		filePath string
		db       *sqlx.DB
	}
	tests := []struct {
		name        string
		args        args
		prepareMock func()
	}{
		{
			name: "Parâmetros insuficientes para definir uma nova conexão",
			args: args{},
		},
		{
			name: "Conexão por arquivo já definida",
			args: args{
				filePath: "path/foo",
				db:       &sqlx.DB{},
			},
			prepareMock: func() {

				connections["path/foo"] = &sqlx.DB{}
			},
		},
		{
			name: "Definição de conexão de forma externa através de arquivo",
			args: args{
				filePath: "path/foo-foo",
				db:       &sqlx.DB{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.prepareMock != nil {
				tt.prepareMock()
			}

			SetConnectionByFile(tt.args.filePath, tt.args.db)
		})
	}
}

func TestSetConnectionByConfig(t *testing.T) {
	type args struct {
		config DBConfig
		db     *sqlx.DB
	}
	tests := []struct {
		name        string
		args        args
		prepareMock func(p *args)
	}{
		{
			name: "Parâmetros insuficientes para definir uma nova conexão",
			args: args{},
		},
		{
			name: "Conexão por configuração já definida",
			args: args{
				config: DBConfig{
					Type:     "postgres",
					DataBase: "foo-db",
					Host:     "foo-host",
					Port:     1234,
				},
				db: &sqlx.DB{},
			},
			prepareMock: func(p *args) {

				connections[getDataConnection(p.config)] = &sqlx.DB{}
			},
		},
		{
			name: "Definição de conexão de forma externa através da configuração",
			args: args{
				config: DBConfig{
					Type:     "mysql",
					DataBase: "foo-db",
					Host:     "foo-host",
					Port:     1234,
				},
				db: &sqlx.DB{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.prepareMock != nil {
				tt.prepareMock(&tt.args)
			}

			SetConnectionByConfig(tt.args.config, tt.args.db)
		})
	}
}

func TestClearConnectionByFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name        string
		args        args
		prepareMock func(p *args, sqlMock sqlmock.Sqlmock, dbx *sqlx.DB)
	}{
		{
			name:        "Arquivo não informado não fazer nada",
			args:        args{},
			prepareMock: nil,
		},
		{
			name: "Conexão não localizada, não fazer nada",
			args: args{
				filePath: "sdfjsdlkf",
			},
			prepareMock: nil,
		},
		{
			name: "Remover a conexão com base no caminho informado",
			args: args{
				filePath: "foo/foo.fo",
			},
			prepareMock: func(p *args, sqlMock sqlmock.Sqlmock, dbx *sqlx.DB) {

				sqlMock.ExpectPing()

				SetConnectionByFile(p.filePath, dbx)

				SetConnectionByFile("foofoofoo/fofoofoo/foo.fo", &sqlx.DB{})
			},
		},
		{
			name: "Conexão já fechada não chamar o método Close()",
			args: args{
				filePath: "foo/foo.fo",
			},
			prepareMock: func(p *args, sqlMock sqlmock.Sqlmock, dbx *sqlx.DB) {

				SetConnectionByFile(p.filePath, dbx)

				dbx.Close()

				SetConnectionByFile("foofoofoo/fofoofoo/foo.fo", &sqlx.DB{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.prepareMock != nil {

				db, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))

				if err != nil {
					t.Fatal(err)
				}

				tt.prepareMock(&tt.args, sqlMock, sqlx.NewDb(db, "mysql"))
			}

			ClearConnectionByFile(tt.args.filePath)

			db, found := connections[tt.args.filePath]

			if found {
				t.Error("Dados de conexão Existente dados não removidos")
			}

			if db != nil {
				t.Errorf("Objeto conexão não removido")
			}

		})

	}
}

func TestClearConnectionByConfig(t *testing.T) {
	type args struct {
		config DBConfig
	}
	tests := []struct {
		name        string
		args        args
		prepareMock func(p *args, sqlMock sqlmock.Sqlmock, dbx *sqlx.DB)
	}{
		{
			name: "Conexão não localizada, não fazer nada",
			args: args{
				config: DBConfig{
					Type:     "mysql",
					DataBase: "foo",
					Host:     "foo",
					Port:     1234,
				},
			},
			prepareMock: nil,
		},
		{
			name: "Remover a conexão com base nas configurações de conexão informado",
			args: args{
				config: DBConfig{
					Type:     "mysql",
					DataBase: "foo",
					Host:     "foo",
					Port:     1234,
				},
			},
			prepareMock: func(p *args, sqlMock sqlmock.Sqlmock, dbx *sqlx.DB) {

				sqlMock.ExpectPing()

				SetConnectionByConfig(p.config, dbx)

				SetConnectionByConfig(DBConfig{
					Type:     "mysql",
					DataBase: "foofoofoo",
					Host:     "foofoofoo",
					Port:     4321,
				}, &sqlx.DB{})
			},
		},
		{
			name: "Conexão já fechada não chamar o método Close()",
			args: args{
				config: DBConfig{
					Type:     "mysql",
					DataBase: "foo",
					Host:     "foo",
					Port:     1234,
				},
			},
			prepareMock: func(p *args, sqlMock sqlmock.Sqlmock, dbx *sqlx.DB) {

				SetConnectionByConfig(p.config, dbx)

				dbx.Close()

				SetConnectionByConfig(DBConfig{
					Type:     "mysql",
					DataBase: "foofoofoo",
					Host:     "foofoofoo",
					Port:     4321,
				}, &sqlx.DB{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareMock != nil {

				db, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))

				if err != nil {
					t.Fatal(err)
				}

				tt.prepareMock(&tt.args, sqlMock, sqlx.NewDb(db, "mysql"))
			}

			ClearConnectionByConfig(tt.args.config)

			db, found := connections[getDataConnection(tt.args.config)]

			if found {
				t.Error("Dados de conexão Existente dados não removidos")
			}

			if db != nil {
				t.Errorf("Objeto conexão não removido")
			}
		})
	}
}
