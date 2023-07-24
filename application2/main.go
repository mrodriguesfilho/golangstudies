package main

import (
	"gitbook/application2/internal/infra/database"
	"gitbook/application2/internal/usecases"

	_ "github.com/lib/pq"
)

func main() {
	dbConnection := Init()
	defer dbConnection.Close()

	orderRepository := database.NewOrderRepository(dbConnection)
	usecase := usecases.NewCalculateFinalPrice(orderRepository)

	input := usecases.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		panic(err)
	}

	println(output)
}
