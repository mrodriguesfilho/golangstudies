package entity

type OrderRepsitoryInterface interface {
	Save(order *Order) error
	GetTotalTransactions() (int, error)
	GetOrderById(id string) (Order, error)
}
