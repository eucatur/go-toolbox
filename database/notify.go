package database

import (
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

// Notifier encapsulates the state of the listener connection.
type Notifier struct {
	listener *pq.Listener
	failed   chan error
}

// NewNotifier - creates a new notifier for given PostgreSQL credentials.
func NewNotifier(channelName string, minIntervalConnection, maxIntervalConnection time.Duration) (*Notifier, error) {
	n := &Notifier{failed: make(chan error, 2)}

	validateConfigNotify()

	switch _config.Type {
	case "postgres":

		dsn := fmt.Sprintf("user=%s port=%d password=%s host=%s dbname=%s sslmode=%s",
			_config.User,
			_config.Port,
			_config.Password,
			_config.Host,
			_config.DataBase,
			_config.SSLMode,
		)

		listener := pq.NewListener(
			dsn,
			minIntervalConnection, maxIntervalConnection,
			n.LogListener)

		if err := listener.Listen(channelName); err != nil {
			listener.Close()
			log.Println("ERROR!: ", err)
			return nil, err
		}

		n.listener = listener

	default:

		panic(fmt.Sprintf("Banco de dados %s não definido para notificação", _config.Type))

	}

	return n, nil
}

// LogListener - is the state change callback for the listener.
func (n *Notifier) LogListener(event pq.ListenerEventType, err error) {
	if err != nil {
		log.Printf("listener error: %s\n", err)
	}
	if event == pq.ListenerEventConnectionAttemptFailed {
		n.failed <- err
	}
}

// Fetch - main to fetch data in DB through notifier
// Please use the method into goroutines
func (n *Notifier) Fetch(data chan []byte) error {
	var fetchCounter uint64
	for {
		select {
		case e := <-n.listener.Notify:
			if e == nil {
				continue
			}
			fetchCounter++
			data <- []byte(e.Extra)
		case err := <-n.failed:
			return err
		case <-time.After(time.Minute):
			go n.listener.Ping()
		}
	}
}

// Close - Close the connection with listener
func (n *Notifier) Close() error {
	close(n.failed)
	return n.listener.Close()
}

// validateConfigNotify validation of settings
func validateConfigNotify() {

	if _config.User == "" {
		panic("User of DB not setted")
	}
	if _config.Port <= 0 {
		panic("Port of DB not setted")
	}
	if _config.Password == "" {
		panic("Password of DB not setted")
	}
	if _config.Host == "" {
		panic("Host of DB not setted")
	}
	if _config.DataBase == "" {
		panic("Name of DB not setted")
	}
	if _config.SSLMode == "" {
		panic("SSL Mode of DB not setted")
	}
	if _config.Type == "" {
		panic("Type of DB not setted")
	}

}
