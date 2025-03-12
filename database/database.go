package database

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/XSAM/otelsql"
	"github.com/eucatur/go-toolbox/json"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	FilePath     string
	PathToDBFile string `json:"path_to_db_file"`
	Type         string `json:"type"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DataBase     string `json:"database"`
	SSLMode      string `json:"ssl_mode"`

	MaxLifeTime       int `json:"max_life_time"`
	MaxOpenConnection int `json:"max_open_connection"`
	MaxIdleConnection int `json:"max_idle_connection"`

	Context *context.Context
}

var connections = map[string]*sqlx.DB{}

var _config = DBConfig{}

func get(filePath string) (*sqlx.DB, error) {
	if db, found := connections[filePath]; found {
		return db, nil
	}
	return nil, fmt.Errorf(`error database not found. File path: %s`, filePath)
}

func SetConnectionByFile(filePath string, db *sqlx.DB) {
	
	if strings.EqualFold(filePath, "") || db == nil {
		return
	}

	if _, found := connections[filePath]; found {
		return
	}

	connections[filePath] = db
	
}

func getDataConnection(config DBConfig) string {
	return fmt.Sprintf("%s-%s-%s-%d", config.Type, config.DataBase, config.Host, config.Port)
}

func getConnection(config DBConfig) (*sqlx.DB, error) {

	dataConnection := getDataConnection(config)

	if db, found := connections[dataConnection]; found {
		return db, nil
	}

	return nil, fmt.Errorf("error database not found by data connection. Datas of connection: %s", dataConnection)
}

func SetConnectionByConfig(config DBConfig, db *sqlx.DB) {

	if db == nil {
		return
	}

	dataConnection := getDataConnection(config)

	if _, found := connections[dataConnection]; found {
		return
	}

	connections[dataConnection] = db

}

func connect(config DBConfig) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
		driverName string
	)

	switch config.Type {
	case "postgres":
		db, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s port=%d password=%s host=%s dbname=%s sslmode=%s",
			config.User,
			config.Port,
			config.Password,
			config.Host,
			config.DataBase,
			config.SSLMode,
		))

	case "mysql":

		driverName, err = otelsql.Register("mysql", otelsql.WithAttributes())

		if err != nil {
			return nil, fmt.Errorf(`erro ao registrar o driver mysql para otel. Detalhes: %s`, err.Error())
		}

		db, err = sqlx.Connect(driverName, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DataBase,
		))

	case "sqlite3":
		db, err = sqlx.Connect("sqlite3", config.PathToDBFile)

	default:
		return nil, errors.New("error database type is not supported")
	}

	if err != nil {
		return nil, fmt.Errorf(`error connecting to database of type "%s" because of: %s`, config.Type, err.Error())
	}

	maxLifeTime := time.Duration(config.MaxLifeTime)

	db.SetMaxOpenConns(config.MaxOpenConnection)
	db.SetMaxIdleConns(config.MaxIdleConnection)
	db.SetConnMaxLifetime(maxLifeTime * time.Minute)

	if !strings.EqualFold(config.FilePath, "") {

		connections[config.FilePath] = db

	} else {

		connections[getDataConnection(config)] = db

	}

	return db, nil
}

// GetByFile Create a database connection through
// the path of a file
func GetByFile(filePath string) (*sqlx.DB, error) {

	if db, err := get(filePath); err == nil {
		return db, nil
	}

	var (
		config DBConfig
		err    error
	)

	if err = json.UnmarshalFile(filePath, &config); err != nil {
		return nil, err
	}

	config.FilePath = filePath
	_config = config

	return connect(config)
}

// MustGetByFile Create a database connection through
// the path of a file and generates a panic in case of error
func MustGetByFile(filePath string) *sqlx.DB {
	db, err := GetByFile(filePath)
	if err != nil {
		panic(err)
	}
	return db
}

func GetConnection(config DBConfig) (*sqlx.DB, error) {

	if db, err := getConnection(config); err == nil {
		return db, nil
	}

	_config = config

	return connect(config)

}

func ClearConnectionByFile(filePath string) {

	if strings.EqualFold(filePath, "") {
		return
	}

	db, found := connections[filePath]

	if !found {
		return
	}

	if db != nil {
		if err := db.Ping(); err == nil {
			db.Close()
		}
	}

	connections[filePath] = nil

	delete(connections, filePath)

}

func ClearConnectionByConfig(config DBConfig) {

	dataConnection := getDataConnection(config)

	db, found := connections[dataConnection]

	if !found {
		return
	}

	if db != nil {
		if err := db.Ping(); err == nil {
			db.Close()
		}
	}

	connections[dataConnection] = nil

	delete(connections, dataConnection)

}
