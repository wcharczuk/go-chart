package matrix

import (
	"math"
)

func minInt(values ...int) int {
	min := math.MaxInt32

	for x := 0; x < len(values); x++ {
		if values[x] < min {
			min = values[x]
		}
	}
	return min
}

func maxInt(values ...int) int {
	max := math.MinInt32

	for x := 0; x < len(values); x++ {
		if values[x] > max {
			max = values[x]
		}
	}
	return max
}
