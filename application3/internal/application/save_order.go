package application

import (
	"gitbook/application3/internal/entity"
	"log"
	"time"
)

type OrderCreatedInput struct {
	OrderId string
	BatchId string
	Status  entity.Status
}

type SaveOrder struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewSaveOrder(orderRepository entity.OrderRepositoryInterface) *SaveOrder {
	return &SaveOrder{
		OrderRepository: orderRepository,
	}
}

func (s *SaveOrder) Execute(orderCreated OrderCreatedInput) error {
	order := entity.NewOrder(orderCreated.OrderId, orderCreated.BatchId, orderCreated.Status)
	err := s.OrderRepository.Save(order)

	if err != nil {
		log.Printf("Failed to save [OrderId:%s] [BatchId:%s] [Status:%v] ", orderCreated.OrderId, orderCreated.BatchId, orderCreated.Status)
		return err
	}

	time.Sleep(time.Second * 1)
	return nil
}
