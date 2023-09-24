package main

import (
	"fmt"
	"gitbook/application3/internal/application"
	"gitbook/application3/internal/infra/database"
	"gitbook/application3/internal/messaging"
	"log"
	"runtime"

	_ "github.com/lib/pq"
)

func main() {
	numberOfCores := runtime.NumCPU()
	fmt.Printf("Number of cores on Apple M1: %d\n", numberOfCores)

	dbConnection := Init()
	defer dbConnection.Close()

	repository := database.NewOrderRepository(dbConnection)
	saveOrder := application.NewSaveOrder(repository)
	rabbitMqConnectionManager := messaging.NewRabbitMqConnectionManager("amqp://guest:guest@localhost:5672/")

	err := rabbitMqConnectionManager.Connect()

	if err != nil {
		log.Fatalf("Cannot connect to rabbit %v", err)
	}

	for i := 0; i < 1; i++ {
		rabbitMqHandler, err := messaging.NewRabbitMQHandler(rabbitMqConnectionManager, "OrderCreatedQueue")

		if err != nil {
			log.Fatalf("Cannot create consummer %v", err)
		}

		go rabbitMqConnectionManager.StartConnectionWatch()
		go rabbitMqHandler.ConsumeMessages(*saveOrder, i)
	}

	select {}
}
