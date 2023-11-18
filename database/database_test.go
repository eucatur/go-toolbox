package database

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
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
