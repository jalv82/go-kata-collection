package greeter

import "fmt"

func Greeter(name string) string {
	return fmt.Sprintf("Hello, %v", name)
}
