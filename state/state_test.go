package state

import "testing"

func TestExampleWeek(t *testing.T) {
	ctx := NewDayContext()

	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := 0; i < 15; i++ {
		todayAndNext()
	}
}
