package messaging

import (
	"log"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

type RabbitMqConnectionManager struct {
	connectionURL string
	conn          *ampq.Connection
}

func NewRabbitMqConnectionManager(ampqURI string) *RabbitMqConnectionManager {
	return &RabbitMqConnectionManager{
		connectionURL: ampqURI,
	}
}

func (m *RabbitMqConnectionManager) Connect() error {
	conn, err := ampq.DialConfig(m.connectionURL, ampq.Config{
		Heartbeat: 10 * time.Second,
	})

	if err != nil {
		return err
	}

	m.conn = conn

	return nil
}

func (m *RabbitMqConnectionManager) Reconnect() error {
	err := m.Connect()

	if err == nil {
		log.Println("Reconnected to RabbitMQ successfully.")
		return nil
	}

	return err
}

func (m *RabbitMqConnectionManager) StartConnectionWatch() {
	err := <-m.conn.NotifyClose(make(chan *ampq.Error))

	if err != nil {
		for {
			log.Printf("RabbitMQ connection problem detected, retrying! %v", err)
			reconnErr := m.Reconnect()

			if reconnErr == nil {
				log.Printf("Still can't connect. Retrying again!")
				continue
			}

			log.Printf("Connection restablished!")
		}
	}
}
