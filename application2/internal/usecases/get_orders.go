package usecases

import "gitbook/application2/internal/entity"

type GetOrdersInput struct {
	Ids []string
}

type GetOrdersOutput struct {
	Orders []OrderOutput
}

type GetOrders struct {
	OrderRepository entity.OrderRepsitoryInterface
}

func NewGetOrders(orderRepository entity.OrderRepsitoryInterface) *GetOrders {
	return &GetOrders{
		OrderRepository: orderRepository,
	}
}

func (g *GetOrders) Execute(input GetOrdersInput) (GetOrdersOutput, error) {
	var getOrderOutput GetOrdersOutput
	for _, id := range input.Ids {
		orderFound, err := g.OrderRepository.GetOrderById(id)

		if err != nil {
			continue
		}

		orderOutput := OrderOutput{
			ID:         orderFound.ID,
			Price:      orderFound.Price,
			Tax:        orderFound.Tax,
			FinalPrice: orderFound.FinalPrice,
		}

		getOrderOutput.Orders = append(getOrderOutput.Orders, orderOutput)
	}

	return getOrderOutput, nil
}
