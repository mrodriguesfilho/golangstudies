package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Init() *sql.DB {
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

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected sucessfully to the database")

	return db
}
