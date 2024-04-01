package rabbit_amqp

import (
	"ContactsService/internal/models/dbmodels"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type Broker struct {
	PublisherChannel *amqp.Channel
	ContactsQ        amqp.Queue
	ErrorQ           amqp.Queue
}

func NewBroker(conn *amqp.Connection) *Broker {
	pubCh, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	contactsQ, err := pubCh.QueueDeclare(
		"contacts",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	errorQ, err := pubCh.QueueDeclare(
		"errors",
		false,
		false,
		false,
		false,
		nil,
	)
	return &Broker{
		PublisherChannel: pubCh,
		ContactsQ:        contactsQ,
		ErrorQ:           errorQ,
	}
}

func (b *Broker) ErrorMessage(err error) {
	body := []byte(err.Error())
	err = b.PublisherChannel.Publish(
		"",
		b.ErrorQ.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		log.Println(err)
	}
}

func (b *Broker) ContactMessage(contact dbmodels.Contact) {
	body, err := json.Marshal(contact)
	if err != nil {
		log.Println(err)
		return
	}
	err = b.PublisherChannel.Publish(
		"",
		b.ContactsQ.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Println(err)
	}
}

func (b *Broker) Close() {
	b.PublisherChannel.Close()
}
