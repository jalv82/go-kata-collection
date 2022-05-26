package greeter

import "testing"

const name string = "Codurance"

func TestGreeter(t *testing.T) {
	tests := []struct {
		description string
		name        string
		language    Language
		want        string
	}{
		{description: "English greeting", name: name, language: English, want: "Hello, " + name},
		{description: "Spanish greeting", name: name, language: Spanish, want: "Hola, " + name},
		{description: "Italian greeting", name: name, language: Italian, want: "Ciao, " + name},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := Greeter(name, test.language)

			if got != test.want {
				t.Errorf("Got: %v != Want: %v", got, test.want)
			}
		})
	}
}
