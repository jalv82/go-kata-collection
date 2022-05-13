package greeter

import "testing"

func TestGreeter(t *testing.T) {
	got := Greeter("Coudurance")

	want := "Hello, Coudurance"

	if got != want {
		t.Errorf("Got: %v != Want: %v", got, want)
	}
}
