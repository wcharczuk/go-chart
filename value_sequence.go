package chart

// ValueSequence returns a sequence for a given values set.
func ValueSequence(values ...float64) Seq[float64] {
	return Seq[float64]{
		Array(values...),
	}
}
