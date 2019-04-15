package main

import (
	"fmt"

	"github.com/eucatur/go-toolbox/rediscon"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := rediscon.ConnectByFile("../redis.json")

	if err != nil {
		panic(err)
	}

	defer c.Close()

	c.Do("SET", "foo", "bar")

	r, err := redis.String(c.Do("GET", "foo"))

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
