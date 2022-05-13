package greeter

import "testing"

func TestGreeter(t *testing.T) {
	tests := []struct {
		description string
		name        string
		language    string
		want        string
	}{
		{description: "English greeting", name: "Codurance", language: "english"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := Greeter("Codurance")

			want := "Hello, Codurance"

			if got != want {
				t.Errorf("Got: %v != Want: %v", got, want)
			}
		})
	}
}
