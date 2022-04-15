package maps

func Get[M ~map[K]V, K comparable, V any](m M, key K) V {
	return m[key]
}

func Lookup[M ~map[K]V, K comparable, V any](m M, key K) (value V, ok bool) {
	value, ok = m[key]
	return
}

func GetOr[M ~map[K]V, K comparable, V any](m M, key K, or V) V {
	value, ok := m[key]
	if !ok {
		return or
	}

	return value
}

func GetOrElse[M ~map[K]V, K comparable, V any](m M, key K, fn func() V) V {
	value, ok := m[key]
	if !ok {
		return fn()
	}

	return value
}
