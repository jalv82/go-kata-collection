package store

import "errors"

type Record struct {
	Title        string
	Artist       string
	Copies       uint
	Price        float32
	IsDiscounted bool
}

func (r *Record) Buy(quantity uint) error {
	if r.Copies == 0 {
		return errors.New("error: not exists any copies to sell")
	}

	r.Copies -= quantity

	return nil
}

func (r *Record) Discount(percentage float32) (float32, error) {
	if !r.IsDiscounted {
		return 0, errors.New("error: not exists discount")
	}

	return (percentage * r.Price) / 100, nil
}
