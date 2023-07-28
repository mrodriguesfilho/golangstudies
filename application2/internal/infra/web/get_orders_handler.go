package web

import (
	"encoding/json"
	"gitbook/application2/internal/usecases"
	"net/http"
)

type GetOrdersHandler struct {
	GetOrdersUseCase *usecases.GetOrders
}

func NewGetOrdersHandler(getOrdersUseCase *usecases.GetOrders) *GetOrdersHandler {
	return &GetOrdersHandler{
		GetOrdersUseCase: getOrdersUseCase,
	}
}

func (g *GetOrdersHandler) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var input usecases.GetOrdersInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := g.GetOrdersUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set(contentType, applicationJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
