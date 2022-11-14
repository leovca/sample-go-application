package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenEmptyID_WhenCreateANewORder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenEmptyPrice_WhenCreateANewORder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenEmptyTax_WhenCreateANewORder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenValidParams_WhenANewOrder_ThenShouldReceiveCrateOrderWithAllParams(t *testing.T) {
	order := Order{
		ID:    "123",
		Price: 10,
		Tax:   2,
	}

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)

	assert.Nil(t, order.IsValid())
}

func TestGivenValidParams_WhenCallNewNewOrderFunc_ThenShouldReceiveCrateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder(
		"123",
		10,
		2.0,
	)

	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)

	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
