package encoding

type UnmarshalFunc func([]byte, any) error

func Unmarshal[T any](f UnmarshalFunc, data []byte) (value T, err error) {
	err = f(data, &value)
	return value, err
}

func UnmarshalSlice[T any, S ~[]T](f UnmarshalFunc, data []byte) (value S, err error) {
	err = f(data, &value)
	return value, err
}

func UnmarshalMap[K comparable, V any, M ~map[K]V](f UnmarshalFunc, data []byte) (value M, err error) {
	err = f(data, &value)
	return value, err
}
