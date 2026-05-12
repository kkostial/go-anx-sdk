package ptr

// To is a helper function that returns a pointer to the passed in value.
// It is used in places where it is not possible to take a pointer directly, like from a literal or function return value.
func To[T any](v T) *T {
	return &v
}
