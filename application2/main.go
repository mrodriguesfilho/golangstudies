package main

import (
	"gitbook/application2/internal/infra/database"
	"gitbook/application2/internal/infra/web"
	"gitbook/application2/internal/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	dbConnection := Init()
	defer dbConnection.Close()

	repository := database.NewOrderRepository(dbConnection)
	priceCalculatorUseCase := usecases.NewCalculateFinalPrice(repository)
	priceCalculatorHandler := web.NewPriceCalculatorHandler(priceCalculatorUseCase)

	r := chi.NewRouter()
	r.Post("/CalculatePrice", priceCalculatorHandler.CalculateFinalPriceHandler)

	http.ListenAndServe(":5050", r)
}
