package pointer

func New[T any](value T) *T {
	return &value
}

func Unwrap[T any](value *T, initial T) T {
	if value == nil {
		return initial
	}
	return *value
}
