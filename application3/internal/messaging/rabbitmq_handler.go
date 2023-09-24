package messaging

import (
	"encoding/json"
	"gitbook/application3/internal/application"
	"log"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

type RabbitMQHandler struct {
	connManager *RabbitMqConnectionManager
	channel     *ampq.Channel
	queue       *ampq.Queue
}

func NewRabbitMQHandler(rabbitMqConnection *RabbitMqConnectionManager, queueName string) (*RabbitMQHandler, error) {

	channel, err := rabbitMqConnection.conn.Channel()

	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &RabbitMQHandler{
		channel:     channel,
		queue:       &queue,
		connManager: rabbitMqConnection,
	}, nil
}

func (h *RabbitMQHandler) ConsumeMessages(saveOrder application.SaveOrder, consumerId int) {
	for {

		if h.channel.IsClosed() {
			log.Print("Channel closed, recreating")

			newChannel, err := h.connManager.conn.Channel()

			if err != nil {
				log.Printf("Failed to recreate channel: %v", err)
				time.Sleep(time.Second * 10)
				continue
			}

			h.channel = newChannel
			log.Print("Channel recreated")
		}

		msgs, err := h.channel.Consume(
			h.queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			log.Printf("Failed to start consuming %v", err)
			time.Sleep(time.Second * 10)
			continue
		}

		for msg := range msgs {

			var orderCreated application.OrderCreatedInput

			if err := json.Unmarshal(msg.Body, &orderCreated); err != nil {
				log.Printf("Error unmarshaling RabbitMQ message: %v", err)
				msg.Nack(false, false)
				continue
			}

			log.Printf("Consumer %d Order data received: [Id:%s] [OrderId:%s] [Status:%v]", consumerId, orderCreated.OrderId, orderCreated.BatchId, orderCreated.Status)

			err = saveOrder.Execute(orderCreated)

			if err != nil {
				log.Printf("Error saving order: %v", err)
				msg.Nack(false, true)
				continue
			}

			msg.Ack(true)

			log.Printf("Consumer %d Order data saved: [Id:%s] [OrderId:%s] [Status:%v]", consumerId, orderCreated.OrderId, orderCreated.BatchId, orderCreated.Status)
		}
	}
}
