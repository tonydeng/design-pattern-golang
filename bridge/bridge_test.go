package bridge

import "testing"

func TestExampleCommonSms(t *testing.T) {
	m := NewCommonMessage(ViaSms())
	m.SendMessage("have a drink?", "Bob")
}

func TestExampleCommonEmail(t *testing.T) {
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via Email
}

func TestExampleUrgencySMS(t *testing.T) {
	m := NewUrgencyMessage(ViaSms())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via SMS
}

func TestExampleUrgencyEmail(t *testing.T) {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}
