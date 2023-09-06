package rediscon

import (
	"testing"
)

func TestLoadFile(t *testing.T) {

	conn := MustConnectByFile("redis.json")

	if conn == nil {
		t.Error("conexão com redis não retornada")
	}

}
