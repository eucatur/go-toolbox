package main

import (
	"github.com/eucatur/go-toolbox/database"
)

func main() {
	db, err := database.ConfigFromEnvFile("../sqlite3-example.json")

	if err != nil {
		panic(err)
	}

	var users map[string]string

	err = db.Select(&users, `select * from users`)

	if err != nil {
		panic(err)
	}
}
