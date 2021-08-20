package redis

import (
	"fmt"

	redigo "github.com/garyburd/redigo/redis"
)

// Client is the structure used to create a redis connection client
type Client struct {
	Host   string
	Port   int
	DB     int
	Prefix string
	conn   *redigo.Conn
}

// DefaultClient is a default client to connect to local redis
var DefaultClient = Client{
	Host:   "localhost",
	DB:     0,
	Port:   6379,
	Prefix: "",
}

// Conn returns a redis connection to execute commands
func (c *Client) Conn() (conn redigo.Conn) {
	if c.conn != nil {
		return *c.conn
	}

	conn, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), redigo.DialDatabase(c.DB))
	if err != nil {
		panic(err.Error())
	}

	c.conn = &conn
	return *c.conn
}

// Set the string value of a key
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

// Get the value of a key
func (c Client) Get(key string) (value string, err error) {
	return redigo.String(c.Conn().Do("GET", c.Prefix+key))
}

// MustGet the value of a key and you can check for a boolean returned
func (c Client) MustGet(key string) (value string, ok bool) {
	var err error
	value, err = c.Get(key)
	if err != nil || value == "" {
		return "", false
	}

	return value, true
}

// Delete a key
func (c Client) Delete(key string) (err error) {
	_, err = c.Conn().Do("DEL", c.Prefix+key)
	return
}

// Delete todas as chaves onde contém o pattern localizado
func (c Client) DeleteLike(pattern string) (err error) {
	iter := 0
	for {
		arr, err := redigo.Values(c.Conn().Do("SCAN", iter, "MATCH", "*"+pattern+"*"))
		if err != nil {
			return fmt.Errorf("error retrieving '%s' keys", c.Prefix+pattern)
		}

		iter, _ = redigo.Int(arr[0], nil)
		keys, _ := redigo.Strings(arr[1], nil)

		for _, key := range keys {
			_, err = c.Conn().Do("DEL", key)

			if err != nil {
				return err
			}
		}

		if iter == 0 {
			break
		}
	}

	return nil
}
