package greeter

import "fmt"

func Greeter(name string, language string) string {
	greet := selector(language)

	return fmt.Sprintf("%v, %v", greet, name)
}

func selector(language string) string {
	var greet string

	switch language {
	case "english":
		greet = "Hello"
	case "spanish":
		greet = "Hola"
	}

	return greet
}
