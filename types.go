package fda

func Ptr[T any](v T) *T {
	return &v
}
