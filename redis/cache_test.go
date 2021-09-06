package redis

import (
	"encoding/json"
	"testing"
	"time"
)

func TestFull(t *testing.T) {
	key := "KEY"
	expirationSeconds := 1

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

	err = DefaultClient.Set(key, string(vJSON), expirationSeconds)
	if err != nil {
		t.Error(err)
		return
	}

	personOut := Person{}

	data, ok := DefaultClient.MustGet(key)
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

	data, ok = DefaultClient.MustGet(key)
	if ok {
		t.Error("O cache não expirou.")
		return
	}

	err = DefaultClient.Set(key, string(vJSON), 1)
	if err != nil {
		t.Error(err)
		return
	}

	err = DefaultClient.Delete(key)
	if err != nil {
		t.Error(err)
		return
	}

	_, ok = DefaultClient.MustGet(key)
	if ok {
		t.Error("Não foi possivel deletar o cache.")
		return
	}

}

/*
func TestHM(t *testing.T) {
	key := "KEY"
	//expirationSeconds := 1

	//type Person struct {
	//	Name  string `json:"name"`
	//	Phone string `json:"phone"`
	//}
	//
	//personIn0 := Person{
	//	Name:  "Gael Félix Bertani",
	//	Phone: "(99) 99999-9999",
	//}
	//
	//personIn1 := Person{
	//	Name:  "Fulano da Silva",
	//	Phone: "(99) 99999-9999",
	//}
	//
	//vJSON0, err := json.Marshal(personIn0)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//vJSON1, err := json.Marshal(personIn1)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}

//var args = []string{"chave1", "valor1", "chave2", "valor2"}
	err := DefaultClient.HMSet(key,"chave1", "valor1", "chave2", "valor2")
	if err != nil {
		t.Error(err)
		return
	}

	//var persons []Person


	_, ok := DefaultClient.HMMustGet(key)
	if ok == false {
		t.Error("Não foi possível obter o cache.")
		return
	}

	//buf := &bytes.Buffer{}
	//gob.NewEncoder(buf).Encode(data) //converte []string em bytes

//	err = json.Unmarshal(buf.Bytes(), &persons)
//	if err != nil {
//		t.Error(err)
//		return
//	}

	//if personIn != personOut {
	//	t.Error("O valor obtido é diferente do informado.")
	//	return
	//}
	//
	//time.Sleep(time.Duration(expirationSeconds) * time.Second)
	//
	//data, ok = DefaultClient.MustGet(key)
	//if ok {
	//	t.Error("O cache não expirou.")
	//	return
	//}

	//err = DefaultClient.HMSet(key, expirationSeconds, string(vJSON0),
	//	string(vJSON1))
	//if err != nil {
	//	t.Error(err)
	//	return
	//}

	err = DefaultClient.HMDelete(key)
	if err != nil {
		t.Error(err)
		return
	}

	_, ok = DefaultClient.HMMustGet(key)
	if ok!= true {
		t.Error("Não foi possível deletar o cache.")
		return
	}

}

*/
