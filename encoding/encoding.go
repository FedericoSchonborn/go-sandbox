package encoding

type UnmarshalFunc func([]byte, any) error

func Unmarshal[T any](f UnmarshalFunc, data []byte) (T, error) {
	var value T
	err := f(data, &value)
	return value, err
}

func UnmarshalSlice[T any, S ~[]T](f UnmarshalFunc, data []byte) (S, error) {
	var value S
	err := f(data, &value)
	return value, err
}

func UnmarshalMap[K comparable, V any, M ~map[K]V](f UnmarshalFunc, data []byte) (M, error) {
	var value M
	err := f(data, &value)
	return value, err
}

type MarshalFunc func(any) ([]byte, error)

func Marshal[T any](f MarshalFunc, value T) ([]byte, error) {
	data, err := f(value)
	return data, err
}

func MarshalSlice[T any, S ~[]T](f MarshalFunc, value S) ([]byte, error) {
	data, err := f(value)
	return data, err
}

func MarshalMap[K comparable, V any, M ~map[K]V](f MarshalFunc, value M) ([]byte, error) {
	data, err := f(value)
	return data, err
}
