package dyntuple

import "fmt"

type DynTuple struct {
	s []any
	l int
}

func New(values ...any) *DynTuple {
	return &DynTuple{
		s: values,
		l: len(values),
	}
}

func Get[T any](dt *DynTuple, i int) (T, error) {
	var zero T

	if i < 0 || i > dt.l {
		return zero, fmt.Errorf("index %d is outside of range 0 to %d", i, dt.l-1)
	}

	value := dt.s[i]
	result, ok := value.(T)
	if !ok {
		return zero, fmt.Errorf("failed to convert value of type %T to type %T", value, zero)
	}

	return result, nil
}
