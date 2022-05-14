package generics

type MyNumber interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Sum[T MyNumber](numbers ...T) T {
	var result T

	for _, number := range numbers {
		result += number
	}

	return result
}
