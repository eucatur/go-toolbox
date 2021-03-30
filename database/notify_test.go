package database

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewNotifier(t *testing.T) {

	// Carregar as configurações
	MustGetByFile("postgres-example.json").DB.Close()

	notify, err := NewNotifier("events", 10*time.Second, 20*time.Second)

	if reflect.DeepEqual(err, nil) {
		t.Failed()
	}

	if reflect.DeepEqual(notify, nil) {
		t.Fatal("Notify não iniciado")
	}

	err = notify.Close()

	if !reflect.DeepEqual(err, nil) {
		t.Failed()
	}

}

func TestFetch(t *testing.T) {

	MustGetByFile("postgres-example.json").DB.Close()

	notify, err := NewNotifier("events", 2*time.Second, 5*time.Second)

	if reflect.DeepEqual(err, nil) {
		t.Failed()
	}

	if reflect.DeepEqual(notify, nil) {
		t.Fatal("Notify não iniciado")
	}

	data := make(chan []byte)

	go func(notfy *Notifier) {

		for {

			time.Sleep(20 * time.Second)

			err = notfy.Close()

			if !reflect.DeepEqual(err, nil) {
				fmt.Println("Deu esse problema: ", err)
			}

			if notfy == nil {
				break
			}

			fmt.Println("Conexão fechada devido ter demorado demais pra notificação retornar algo")

		}

	}(notify)

	go func(d chan []byte) {

		err = notify.Fetch(d)

		if !reflect.DeepEqual(err, nil) {
			t.Failed()
		}

	}(data)

	fmt.Println("Data return from notifier: ", string(<-data))

	notify.Close()
}
