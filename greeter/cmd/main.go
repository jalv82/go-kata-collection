package main

import (
	"greeter"
	"log"
)

func main() {
	log.Println("# Greeter")
	log.Println(greeter.Greeter("Gophers", greeter.English))
}
