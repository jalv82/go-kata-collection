package greeter

import "fmt"

type Language string

func Greeter(name string, language Language) string {
	greet := selector(language)

	return fmt.Sprintf("%v, %v", greet, name)
}

func selector(language Language) string {
	var greet string

	switch language {
	case "english":
		greet = "Hello"
	case "spanish":
		greet = "Hola"
	}

	return greet
}
