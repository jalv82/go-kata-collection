package main

import (
	"go-kata-collection/fizzbuzz"
	"go-kata-collection/greeter"
	"log"
)

func main() {
	log.Println(greeter.Greeter("Gophers", greeter.English))

	log.Println(fizzbuzz.FizzBuzz(3))
}
