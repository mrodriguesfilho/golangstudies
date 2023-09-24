package database

import (
	"database/sql"
	"gitbook/application3/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	insertQuery := "insert into mowdb.tb_orders_created (OrderId, BatchId, Status) values ($1, $2, $3)"
	_, err := r.Db.Exec(insertQuery, order.OrderId, order.BatchId, order.Status)

	if err != nil {
		return err
	}

	return nil
}
