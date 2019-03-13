package cache

import (
	"time"

	GoCache "github.com/patrickmn/go-cache"
)

var (
	DefaultExpiration = 1 * time.Minute
	cache             = GoCache.New(5*time.Minute, DefaultExpiration)
)

func Set(key string, value interface{}, d ...time.Duration) {
	duration := DefaultExpiration

	if len(d) > 0 {
		duration = d[0]
	}

	cache.Set(key, value, duration)
}

func Get(key string) (interface{}, bool) {
	return cache.Get(key)
}
