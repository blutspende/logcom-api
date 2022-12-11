package logcom

func toPtr[T any](value T) *T {
	return &value
}
