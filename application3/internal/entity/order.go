package entity

type Order struct {
	OrderId string
	BatchId string
	Status  Status
}

type Status int

const (
	Exporting Status = iota
	Exported
	Settled
)

func NewOrder(orderId string, batchId string, status Status) *Order {
	order := &Order{
		OrderId: orderId,
		BatchId: batchId,
		Status:  status,
	}

	return order
}
