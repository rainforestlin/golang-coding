package adapter

import "testing"

func TestAdapter_Request(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != adaptee.SpecificRequest() {
		t.Fatal()
	}
}
