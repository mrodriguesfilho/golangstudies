package database

import (
	"database/sql"
	"fmt"
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

func (r *OrderRepository) GetOrderById(id string) (entity.Order, error) {
	var orderFound entity.Order

	row := r.Db.QueryRow("select id, price, tax, final_price from mowdb.tb_orders where id =$1", id)
	err := row.Scan(&orderFound.ID, &orderFound.Price, &orderFound.Tax, &orderFound.FinalPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Order{}, fmt.Errorf("Order with ID %s was not found", id)
		}
	}

	return orderFound, nil
}
