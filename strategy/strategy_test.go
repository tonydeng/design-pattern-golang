package strategy

import "testing"

func TestExamplePayByCash(t *testing.T) {
	ctx := NewPaymentContext("Tony", "1234", 100, &Cash{})
	ctx.Pay()
}

func TestExamplePayByBank(t *testing.T) {
	ctx := NewPaymentContext("Tony", "1234", 100, &Bank{})
	ctx.Pay()
}
