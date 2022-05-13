package greeter

import "fmt"

type Language string

const (
	English Language = "english"
	Spanish Language = "spanish"
)

func Greeter(name string, language Language) string {
	greet := selector(language)

	return fmt.Sprintf("%v, %v", greet, name)
}

func selector(language Language) string {
	var greet string

	switch language {
	case English:
		greet = "Hello"
	case Spanish:
		greet = "Hola"
	}

	return greet
}
