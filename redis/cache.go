package redis

import (
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
)

// Args is a helper for constructing command arguments from structured values.
type Args []interface{}

// Add returns the result of appending value to args.
func (args Args) Add(value ...interface{}) Args {
	return append(args, value...)
}

// Client is the structure used to create a redis connection client
type Client struct {
	Host            string
	Port            int
	DB              int
	Prefix          string
	ConnectionRedis *redigo.Conn
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
	if c.ConnectionRedis != nil {
		return *c.ConnectionRedis
	}

	conn, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), redigo.DialDatabase(c.DB))
	if err != nil {
		panic(err.Error())
	}

	c.ConnectionRedis = &conn
	return *c.ConnectionRedis
}

func (c *Client) Ping() {

	_, err := c.Conn().Do("PING")

	if err != nil {
		panic(fmt.Errorf("não foi possível estabelecer conexão com redis. Detalhes: %s", err.Error()))
	}

}

// Set the string value of a key
func (c Client) Set(key, value string, expirationSeconds int) (err error) {
	key = c.Prefix + key

	_, err = c.Conn().Do("SET", key, value)
	if err != nil {
		return
	}

	if expirationSeconds > 0 {
		_, err = c.Conn().Do("EXPIRE", key, expirationSeconds)
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

// Send Abre uma conexão com o Redis, executa o comando e depois a fecha
func (c Client) Do(comando string, args ...interface{}) (interface{}, error) {
	/*	value := c.Conn().(comando, args...)
		if value != nil {
			return nil, value
		}
		return value, nil*/

	return c.Conn().Do(comando, args...)
}
