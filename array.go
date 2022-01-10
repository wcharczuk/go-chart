package chart

var (
	_ Sequence[int] = (*array[int])(nil)
)

// NewArray returns a new array from a given set of values.
// Array implements Sequence, which allows it to be used with the sequence helpers.
func Array[A any](values ...A) Sequence[A] {
	return array[A](values)
}

// Array is a wrapper for an array of floats that implements `ValuesProvider`.
type array[A any] []A

// Len returns the value provider length.
func (a array[A]) Len() int {
	return len(a)
}

// GetValue returns the value at a given index.
func (a array[A]) GetValue(index int) A {
	return a[index]
}
