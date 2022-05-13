package greeter

import "testing"

func TestGreeter(t *testing.T) {
	got := Greeter("Codurance")

	want := "Hello, Codurance"

	if got != want {
		t.Errorf("Got: %v != Want: %v", got, want)
	}
}
