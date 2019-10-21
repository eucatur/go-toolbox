# redis #

cache regis é um wrapper do [redigo](https://github.com/gomodule/redigo) uma lib de cache utilizando o REDIS, basicamente tem somente o método SET e GET 

## Exemplo ##

```code
package main

import (
	"encoding/json"
	"log"

	"github.com/eucatur/go-toolbox/redis"
)

func main() {
	client := redis.Client{
		Host: "localhost",
		Port: 6379,
	}

	key := "KEY"
	expirationSeconds := 1

	type Person struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	personIn := Person{
		Name:  "Gael Félix Bertani",
		Phone: "(99) 99999-9999",
	}

	vJSON, err := json.Marshal(personIn)
	if err != nil {
		log.Println(err)
		return
	}

	err = client.Set(key, string(vJSON), expirationSeconds)
	if err != nil {
		log.Println(err)
		return
	}

	personOut := Person{}

	data, ok := client.MustGet(key)
	if !ok {
		log.Println(err)
		return
	}

	err = json.Unmarshal([]byte(data), &personOut)
	if err != nil {
		log.Println(err)
		return
	}
}

```