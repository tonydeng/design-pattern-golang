package simple_factory

import "testing"

func TestType1(t *testing.T)  {
	api := NewAPI(1)
	s := api.Say("Tony")

	if s !="Hi, Tony"{
		t.Fatalf("Type1 test fail")
	}
}

func TestType2(t *testing.T)  {
	api := NewAPI(2)
	s := api.Say("Tony")

	if s !="Hello, Tony"{
		t.Fatalf("Type1 test fail")
	}
}