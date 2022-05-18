package zero

func Zero[T any]() T {
	var zero T
	return zero
}

// TODO: DeepIsZero

func IsZero[T comparable](value T) bool {
	var zero T
	return value == zero
}
