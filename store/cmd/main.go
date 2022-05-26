package main

import (
	"log"
	"store"
)

func main() {
	log.Println("# Store Record")
	playStoreRecord()
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
