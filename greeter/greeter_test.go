package greeter

import "testing"

func TestGreeter(t *testing.T) {
	tests := []struct {
		description string
		name        string
		language    string
		want        string
	}{
		{description: "English greeting", name: "Codurance", language: "english", want: "Hello, Codurance"},
		{description: "Spanish greeting", name: "Codurance", language: "spanish", want: "Hola, Codurance"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := Greeter("Codurance", test.language)

			if got != test.want {
				t.Errorf("Got: %v != Want: %v", got, test.want)
			}
		})
	}
}
