package store

import (
	"encoding/json"
	"errors"
	"fmt"
)

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

	if percentage <= 0 {
		return 0, errors.New(fmt.Sprintf("error: not is possible apply this discount (%.2f %%)", percentage))
	}

	return (percentage * r.Price) / 100, nil
}

func (r *Record) Details() ([]byte, error) {
	return json.Marshal(r)
}
