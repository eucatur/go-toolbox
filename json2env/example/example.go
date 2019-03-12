package main

import (
	"log"
	"os"

	"github.com/eucatur/go-toolbox/json2env"
)

func main() {
	if err := json2env.LoadFile("../test.json"); err != nil {
		panic(err)
	}

	value := os.Getenv("json")

	log.Println(value)
}
