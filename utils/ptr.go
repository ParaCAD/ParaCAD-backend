package utils

func ValueOrDefault[T any](value *T) T {
	if value == nil {
		return *new(T)
	}
	return *value
}
