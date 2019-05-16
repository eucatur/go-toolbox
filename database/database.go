package database

import (
	"errors"
	"fmt"

	"github.com/eucatur/go-toolbox/json"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type dbConfig struct {
	FilePath           string
	PathToDBFile       string `json:"path_to_db_file"`
	Type               string `json:"type"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	User               string `json:"user"`
	Password           string `json:"password"`
	DataBase           string `json:"database"`
	MaxOpenConnections int    `json:"max_open_connections"`
	SSLMode            string `json:"ssl_mode"`
}

var connections = map[string]*sqlx.DB{}

func get(filePath string) (*sqlx.DB, error) {
	if db, found := connections[filePath]; found {
		return db, nil
	}
	return nil, fmt.Errorf(`Error database not found. File path: %s`, filePath)
}

func connect(config dbConfig) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
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
		db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DataBase,
		))

	case "sqlite3":
		db, err = sqlx.Connect("sqlite3", config.PathToDBFile)

	default:
		return nil, errors.New("Error database type is not supported")
	}

	if err != nil {
		return nil, fmt.Errorf(`Error connecting to database of type "%s" because of: %s`, config.Type, err.Error())
	}

	if config.MaxOpenConnections > 0 {
		db.SetMaxOpenConns(config.MaxOpenConnections)
	}

	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(-1)

	connections[config.FilePath] = db
	return db, nil
}

// GetByFile Create a database connection through
// the path of a file
func GetByFile(filePath string) (*sqlx.DB, error) {

	if db, err := get(filePath); err == nil {
		return db, nil
	}

	var (
		config dbConfig
		err    error
	)

	if err = json.UnmarshalFile(filePath, &config); err != nil {
		return nil, err
	}

	config.FilePath = filePath

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
