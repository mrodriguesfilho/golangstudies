package web

import (
	"encoding/json"
	"gitbook/application2/internal/usecases"
	"net/http"
)

type PriceCalculatorHandler struct {
	PriceCalculatorUseCase *usecases.CalculateFinalPrice
}

func NewPriceCalculatorHandler(priceCalculatorUseCase *usecases.CalculateFinalPrice) *PriceCalculatorHandler {
	return &PriceCalculatorHandler{
		PriceCalculatorUseCase: priceCalculatorUseCase,
	}
}

func (p *PriceCalculatorHandler) CalculateFinalPriceHandler(w http.ResponseWriter, r *http.Request) {
	var input usecases.OrderInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.PriceCalculatorUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set(contentType, applicationJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
