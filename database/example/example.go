package main

import (
	"github.com/eucatur/go-toolbox/database"
)

func main() {
	db, err := database.ConfigFromEnvFile("../sqlite3-example.json")

	if err != nil {
		panic(err)
	}

	var now string

	err = db.Get(&now, `select now()`)

	if err != nil {
		panic(err)
	}
}
