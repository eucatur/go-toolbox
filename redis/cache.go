package redis

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	commandredis "github.com/eucatur/go-toolbox/redis/command_redis"
	"github.com/eucatur/go-toolbox/text"
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
	MaxIdle         int
	MaxActive       int
	IdleTimeout     int
	pool            *redigo.Pool
	ConnStatemented redigo.Conn

	_exitRecursive bool
}

// DefaultClient is a default client to connect to local redis
var DefaultClient = &Client{
	Host:   "localhost",
	DB:     0,
	Port:   6379,
	Prefix: "",
}

func (c *Client) setDefaultParamsConnection() {

	const (
		defatulMaxIdle     = 30
		defaultMaxActive   = 30
		defaultIdleTimeout = 40
	)

	if c.MaxIdle <= 0 {
		c.MaxIdle = defatulMaxIdle
	}

	if c.MaxActive <= 0 {
		c.MaxActive = defaultMaxActive
	}

	if c.IdleTimeout <= 0 {
		c.IdleTimeout = defaultIdleTimeout
	}

}

func NewConnection(param Client) redigo.Conn {

	client := &param

	return client.GetConnectionFromPool()

}

func (c *Client) GetConnectionFromPool() (conn redigo.Conn) {

	checkRecreatePool := func(connection redigo.Conn) bool {

		if connection == nil {
			return false
		}

		err := connection.Err()

		if err == nil {
			return false
		}

		const (
			msgRedigoPoolOverload = "redigo: connection pool exhausted"
			msgRedigoPoolClosed   = "redigo: get on closed pool"
		)

		isErrorReasonToRecreatePool := []string{
			msgRedigoPoolClosed,
			msgRedigoPoolOverload,
		}

		errMsg := strings.ToLower(connection.Err().Error())

		switch {
		case strings.EqualFold(errMsg, msgRedigoPoolOverload):
			log.Default().Printf("Recriado o pool de conexão devido estar sobrecarregado. Para evitar considere rever os parâmetros definido para MaxIdle, MaxActive e IdleTimeout\n")
		}

		return slices.Contains(isErrorReasonToRecreatePool, errMsg)
	}

	c.setDefaultParamsConnection()

	if c.pool == nil {

		c.pool = &redigo.Pool{
			DialContext: func(ctx context.Context) (redigo.Conn, error) {

				if c.ConnStatemented != nil {

					return c.ConnStatemented, nil
				}

				conn, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), redigo.DialDatabase(c.DB))

				if err != nil {

					msgErr := fmt.Sprintf("Falha ao estabelecer conexão com servidor de cache %s:%d. Detalhes: %s\n", c.Host, c.Port, err.Error())

					invalidDataConnection := []bool{
						text.StringIsEmptyOrWhiteSpace(c.Host),
						c.Port <= 0,
					}

					if slices.Contains(invalidDataConnection, true) {
						msgErr += ". Verifique se as informações de Host e/ou Port foram informadas\n"
					}

					log.Default().Printf(msgErr)

					return nil, err
				}

				return conn, nil
			},
			MaxIdle:     c.MaxIdle,
			MaxActive:   c.MaxActive,
			IdleTimeout: time.Duration(c.IdleTimeout * int(time.Second)),
		}

	}

	conn = c.pool.Get()

	if checkRecreatePool(conn) {

		if c._exitRecursive {
			return
		}

		c.pool = nil
		conn = c.GetConnectionFromPool()

		if checkRecreatePool(conn) {
			c._exitRecursive = true
		}

	}

	return conn

}

// Conn returns a redis connection to execute commands
func (c *Client) Conn() (conn redigo.Conn) {
	return c.GetConnectionFromPool()
}

func (c *Client) Ping() (err error) {

	conn := c.Conn()

	defer conn.Close()

	_, err = conn.Do(commandredis.Ping.String())

	if err != nil {
		return
	}

	return

}

// Set the string value of a key
func (c Client) Set(key, value string, expirationSeconds int) (err error) {

	key = c.Prefix + key

	conn := c.Conn()

	defer conn.Close()

	_, err = conn.Do(commandredis.Set.String(), key, value)

	if expirationSeconds > 0 {
		_, err = conn.Do(commandredis.Expire.String(), key, expirationSeconds)
	}

	return
}

// Get the value of a key
func (c Client) Get(key string) (value string, err error) {
	conn := c.Conn()

	defer conn.Close()

	return redigo.String(conn.Do(commandredis.Get.String(), c.Prefix+key))
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
	conn := c.Conn()

	defer conn.Close()

	_, err = conn.Do(commandredis.Delete.String(), c.Prefix+key)
	return
}

// Delete todas as chaves onde contém o pattern localizado
func (c Client) DeleteLike(pattern string) (err error) {

	conn := c.Conn()

	defer conn.Close()

	iter := 0
	for {
		arr, err := redigo.Values(conn.Do(commandredis.Scan.String(), iter, commandredis.Match.String(), "*"+pattern+"*"))
		if err != nil {
			return fmt.Errorf("error retrieving '%s' keys", c.Prefix+pattern)
		}

		iter, _ = redigo.Int(arr[0], nil)
		keys, _ := redigo.Strings(arr[1], nil)

		for _, key := range keys {
			_, err = conn.Do(commandredis.Delete.String(), key)

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
	conn := c.Conn()

	defer conn.Close()

	return conn.Do(comando, args...)
}
