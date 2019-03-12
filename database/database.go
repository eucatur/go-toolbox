package database

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/eucatur/go-toolbox/json2env"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var conections = make(map[string]*sqlx.DB)

// Because of Heroku, your default max conections is 20
var max_open_conns = 20

const ERROR_CONNECT = "Error in connect with database"

func ConfigFromEnvFile(env_file_path string) (db *sqlx.DB, err error) {
	db, exist := conections[env_file_path]
	if !exist {
		if err = json2env.LoadFile(env_file_path); err != nil {
			return db, err
		}

		if os.Getenv("DB_TYPE") == "postgres" {
			db, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s port=%s password=%s host=%s dbname=%s sslmode=%s",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_DATABASE"),
				os.Getenv("DB_SSLMODE")),
			)
		}

		if os.Getenv("DB_TYPE") == "mysql" {
			db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_DATABASE")),
			)
		}

		if os.Getenv("DB_TYPE") == "sqlite3" {
			db, err = sqlx.Connect("sqlite3", os.Getenv("PATH_TO_DB_FILE"))
		}

		if err != nil {
			return db, errors.New(fmt.Sprintf(`Error connecting to database of type "%s" because of: %s`, os.Getenv("DB_TYPE"), err.Error()))
		}

		if os.Getenv("MAX_OPEN_CONNS") != "" {
			max_open_conns, err = strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))

			if err != nil {
				return db, errors.New("ENV MAX_OPEN_CONNS is not int valid: " + env_file_path)
			}
		}

		db.SetMaxOpenConns(max_open_conns)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(2 * time.Minute)
		conections[env_file_path] = db
	}

	return db, err
}
