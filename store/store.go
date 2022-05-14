package store

import "errors"

type Record struct {
	Title  string
	Artist string
	Copies uint
	Price  float32
}

func (r *Record) Buy(quantity uint) error {
	if r.Copies == 0 {
		return errors.New("error: not exists any copies to sell")
	}

	r.Copies -= quantity

	return nil
}

func (r *Record) Discount(percentage float32) float32 {
	return (percentage * r.Price) / 100
}
