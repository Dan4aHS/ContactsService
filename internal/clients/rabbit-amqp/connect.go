package rabbit_amqp

import (
	"ContactsService/internal/config"
	"fmt"
	"github.com/streadway/amqp"
)

func ConnectAMQP(cfg *config.Config) (*amqp.Connection, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.Broker.Username,
		cfg.Broker.Password,
		cfg.Broker.Host,
		cfg.Broker.Port,
	)
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
