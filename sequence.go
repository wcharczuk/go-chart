package chart

// Sequence is a provider for values for a seq.
type Sequence[A any] interface {
	Len() int
	GetValue(int) A
}
