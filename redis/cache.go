package redis

import (
	"fmt"

	redigo "github.com/garyburd/redigo/redis"
)

// DefaultClient ...
var DefaultClient = Client{
	Host:   "localhost",
	Port:   6379,
	Prefix: "",
}

// Client ...
type Client struct {
	Host   string
	Port   int
	Prefix string
	conn   *redigo.Conn
}

// Conn ...
func (c *Client) Conn() (conn redigo.Conn) {
	if c.conn != nil {
		return *c.conn
	}

	conn, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		panic(err.Error())
	}

	c.conn = &conn
	return *c.conn
}

// Set ...
func (c Client) Set(key, value string, expirationSeconds ...int) (err error) {
	key = c.Prefix + key

	_, err = c.Conn().Do("SET", key, value)
	if err != nil {
		return
	}

	if len(expirationSeconds) > 0 {
		_, err = c.Conn().Do("EXPIRE", key, expirationSeconds[0])
	}

	return
}

// Get ...
func (c Client) Get(key string) (value string, err error) {
	return redigo.String(c.Conn().Do("GET", c.Prefix+key))
}

// MustGet ...
func (c Client) MustGet(key string) (value string, ok bool) {
	var err error
	value, err = c.Get(key)
	if err != nil || value == "" {
		return "", false
	}

	return value, true
}
