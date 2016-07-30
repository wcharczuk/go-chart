package chart

// ConcatSeries is a special type of series that concatenates its `InnerSeries`.
type ConcatSeries []Series

// Len returns the length of the concatenated set of series.
func (cs ConcatSeries) Len() int {
	total := 0
	for _, s := range cs {
		if typed, isValueProvider := s.(ValueProvider); isValueProvider {
			total += typed.Len()
		}
	}

	return total
}

// GetValue returns the value at the (meta) index (i.e 0 => totalLen-1)
func (cs ConcatSeries) GetValue(index int) (x, y float64) {
	cursor := 0
	for _, s := range cs {
		if typed, isValueProvider := s.(ValueProvider); isValueProvider {
			len := typed.Len()
			if index < cursor+len {
				x, y = typed.GetValue(index - cursor) //FENCEPOSTS.
				return
			}
			cursor += typed.Len()
		}
	}
	return
}
