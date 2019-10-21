package redis

import (
	"encoding/json"
	"testing"
	"time"
)

func TestFull(t *testing.T) {
	key := "KEY"
	expirationSeconds := 1

	client := Client{
		Host: "localhost",
		Port: 6379,
	}

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
		t.Error(err)
		return
	}

	err = client.Set(key, string(vJSON), expirationSeconds)
	if err != nil {
		t.Error(err)
		return
	}

	personOut := Person{}

	data, ok := client.MustGet(key)
	if !ok {
		t.Error("Não foi possivel obter o cache.")
		return
	}

	err = json.Unmarshal([]byte(data), &personOut)
	if err != nil {
		t.Error(err)
		return
	}

	if personIn != personOut {
		t.Error("O valor obtido é diferente do informado.")
		return
	}

	time.Sleep(time.Duration(expirationSeconds) * time.Second)

	data, ok = client.MustGet(key)
	if ok {
		t.Error("O cache não expirou.")
		return
	}

}
