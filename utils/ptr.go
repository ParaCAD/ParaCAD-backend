package utils

func GetPtr[T any](value T) *T {
	return &value
}

func ValueOrEmpty[T any](value *T) T {
	if value == nil {
		return *new(T)
	}
	return *value
}

func ValueOrDefault[T any](value *T, def T) T {
	if value == nil {
		return def
	}
	return *value
}
