package main

import (
	"greeter"
	"log"
	"os"
)

const name = "Gophers"

func main() {
	log.Println("# Greeter")
	greeter.Greeter(os.Stdout, name, greeter.English)
}
