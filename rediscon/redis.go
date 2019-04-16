package rediscon

import (
	"fmt"
	"time"

	"github.com/eucatur/go-toolbox/json"
	"github.com/garyburd/redigo/redis"
)

// ConnectByFile ...
func ConnectByFile(env_json_file_path string) (c redis.Conn, err error) {
	var m map[string]interface{}

	err = json.UnmarshalFile(env_json_file_path, &m)
	if err != nil {
		return
	}

	address := fmt.Sprintf("%v:%v", m["REDIS_IP"].(string), m["REDIS_PORT"].(string))

	c, err = redis.Dial("tcp", address, redis.DialReadTimeout(10*time.Second), redis.DialWriteTimeout(10*time.Second))
	if err != nil {
		return
	}

	return
}

// ConnectByFile ...
func MustConnectByFile(env_json_file_path string) (c redis.Conn) {
	var m map[string]interface{}

	err := json.UnmarshalFile(env_json_file_path, &m)
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%v:%v", m["REDIS_IP"].(string), m["REDIS_PORT"].(string))

	c, err = redis.Dial("tcp", address, redis.DialReadTimeout(10*time.Second), redis.DialWriteTimeout(10*time.Second))
	if err != nil {
		panic(err)
	}

	return
}
