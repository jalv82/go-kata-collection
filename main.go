package main

import (
	"go-kata-collection/fizzbuzz"
	"go-kata-collection/generics"
	"go-kata-collection/greeter"
	"go-kata-collection/store"
	"log"
)

func main() {
	log.Println("# Greeter")
	log.Println(greeter.Greeter("Gophers", greeter.English))

	log.Println("# FizzBuzz")
	log.Println(fizzbuzz.FizzBuzz(3))

	log.Println("# Store Record")
	playStoreRecord()

	log.Println("# Sum with Generics")
	log.Println("Integers: ", generics.Sum(3, 4, 5, 6))
	log.Println("Floats: ", generics.Sum(3.6, 4.5, 5.4, 6.3))
}

func playStoreRecord() {
	record := store.Record{
		Title:        "The Jar Dance",
		Artist:       "Chiquito de la calzada",
		Copies:       10,
		Price:        12.99,
		IsDiscounted: true,
	}

	_ = record.Buy(1)

	discount, _ := record.Discount(10)
	log.Println(discount)

	json, _ := record.Details()
	log.Println(string(json))

}
