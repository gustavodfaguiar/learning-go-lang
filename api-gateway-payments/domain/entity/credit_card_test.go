package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("90000000000000", "Gustavo Aguiar", 12, 2023, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5273970210446806", "Gustavo Aguiar", 14, 2023, 645)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5273970210446806", "Gustavo Aguiar", 11, 2023, 645)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5273970210446806", "Gustavo Aguiar", 0, 2023, 645)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5273970210446806", "Gustavo Aguiar", 14, 2023, 645)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard("5273970210446806", "Gustavo Aguiar", 14, lastYear.Year(), 645)
	assert.Equal(t, "invalid expiration year", err.Error())
}
