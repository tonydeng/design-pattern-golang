package command

import "testing"

func TestExampleCommand(t *testing.T) {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButtion1()
	box1.PressButtion2()

	box2 := NewBox(rebootCommand, startCommand)
	box2.PressButtion1()
	box2.PressButtion2()
}
