package greeter

import "fmt"

func Greeter(name string, language string) string {

	var greet string
	switch language {
	case "english":
		greet = "Hello"
	case "spanish":
		greet = "Hola"
	}

	return fmt.Sprintf("%v, %v", greet, name)
}
