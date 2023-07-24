package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	t.Run("shouldn't allow create an order without id", func(t *testing.T) {
		order := Order{}
		err := order.Validate()
		assert.Error(t, err, "id is required")
		assert.EqualError(t, err, "id is required", "Error message mismatch")
	})

	t.Run("shouldn't allow create an order without price", func(t *testing.T) {
		order := Order{ID: "123"}
		err := order.Validate()
		assert.Error(t, err, "price is requi1red")
		assert.EqualError(t, err, "price must be greater than zero", "Error message mismatch")
	})

	t.Run("shouldn't allow create an order without tax", func(t *testing.T) {
		order := Order{ID: "123", Price: 10.0}
		err := order.Validate()
		assert.Error(t, err, "tax is required")
		assert.EqualError(t, err, "tax must be greater than zero", "Error message mismatch")
	})

	t.Run("should create an order", func(t *testing.T) {
		order, _ := NewOrder("1", 10, 5)
		assert.NoError(t, order.Validate())
	})
}

func TestFinalPrice(t *testing.T) {
	order, _ := NewOrder("123", 10.0, 5.0)
	order.CalculateFinalPrice()
	assert.Equal(t, 15.0, order.FinalPrice)
}
