package flyweight

import "testing"

func TestExampleFlyweight(t *testing.T) {
	viewer := NewImageViewer("image1.png")
	viewer.Display()
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fail()
	}
}
