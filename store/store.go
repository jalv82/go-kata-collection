package store

type Record struct {
	Title  string
	Artist string
	Copies uint
}

func (r *Record) Buy(quantity uint) {
	r.Copies -= quantity
}
