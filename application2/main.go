package main

import (
	"database/sql"
	"fmt"
	"gitbook/application2/internal/infra/database"
	"gitbook/application2/internal/usecases"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file:", err)
	}

	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	sslmode := viper.GetString("database.sslmode")

	connectionString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbname, user, password, sslmode)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
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
