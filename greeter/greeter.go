package greeter

import (
	"fmt"
	"io"
)

type Language string

const (
	English Language = "english"
	Italian Language = "italian"
	Spanish Language = "spanish"
)

func Greeter(out io.Writer, name string, language Language) {
	greet := selector(language)

	_, _ = fmt.Fprintf(out, "%v, %v", greet, name)
}

func selector(language Language) string {
	var greet string

	switch language {
	case English:
		greet = "Hello"
	case Italian:
		greet = "Ciao"
	case Spanish:
		greet = "Hola"
	}

	return greet
}
