package main

import (
	"log"
	"time"

	"github.com/eucatur/go-toolbox/cache"
)

func main() {
	key := "foo"
	value := "bar"

	// The default duration in cache is 1 minute
	cache.Set(key, value)

	// Add on cache with 5 minutes
	cache.Set(key, value, 5*time.Minute)

	v, found := cache.Get(key)

	if v != nil {
		value = v.(string)
	}

	if found {
		log.Println("Found: ", value)
	}
}
