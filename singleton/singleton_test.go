package singleton

import "testing"

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()

	if ins1 != ins2 {
		t.Fatalf("instance is noot equal")
	}
}
