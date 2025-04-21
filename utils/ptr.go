package utils

func GetPtr[T any](value T) *T {
	return &value
}

func ValueOrDefault[T any](value *T) T {
	if value == nil {
		return *new(T)
	}
	return *value
}
