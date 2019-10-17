package state

import "fmt"

type Week interface {
	Today()
	Next(*DayContext)
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

func (d *DayContext) Today() {
	d.today.Today()
}
func (d *DayContext) Next() {
	d.today.Next(d)
}

type Sunday struct {
}

func (*Sunday) Today() {
	fmt.Println("Sunday")
}

func (*Sunday) Next(cxt *DayContext) {
	cxt.today = &Monday{}
}

type Monday struct {
}

func (*Monday) Today() {
	fmt.Println("Monday")
}

func (*Monday) Next(cxt *DayContext) {
	cxt.today = &Tuesday{}
}

type Tuesday struct {
}

func (*Tuesday) Today() {
	fmt.Println("Tuesday")
}

func (*Tuesday) Next(cxt *DayContext) {
	cxt.today = &Wednesday{}
}

type Wednesday struct{}

func (*Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct{}

func (*Thursday) Today() {
	fmt.Printf("Thursday\n")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct{}

func (*Friday) Today() {
	fmt.Printf("Friday\n")
}

func (*Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct{}

func (*Saturday) Today() {
	fmt.Printf("Saturday\n")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}
