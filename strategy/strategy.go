package strategy

import (
	"fmt"
)

type PaymentContxt struct {
	Name, CardId string
	Money        int
	payment      PaymentStrategy
}

func NewPaymentContext(name, cardId string, money int, payment PaymentStrategy) *PaymentContxt {
	return &PaymentContxt{
		Name:    name,
		CardId:  cardId,
		Money:   money,
		payment: payment,
	}
}

func (p *PaymentContxt) Pay() {
	p.payment.Pay(p)
}

type PaymentStrategy interface {
	Pay(contxt *PaymentContxt)
}

type Cash struct {
}

func (*Cash) Pay(ctx *PaymentContxt) {
	fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

type Bank struct {
}

func (*Bank) Pay(ctx *PaymentContxt) {
	fmt.Printf("Pay $%d to %s by bank\n", ctx.Money, ctx.Name)
}
