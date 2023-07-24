package database

import (
	"database/sql"
	"gitbook/application2/internal/entity"
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
	insertQuery := "insert into mowdb.tb_orders (id, price, tax, final_price) values ($1, $2, $3, $4)"
	_, err := r.Db.Exec(insertQuery, order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {
	var total int

	err := r.Db.QueryRow("select count(id) from mowdb.tb_orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
