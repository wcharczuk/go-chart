package chart

// Points are an array of points.
import (
	"fmt"
	"strings"
)

// Point represents a x,y coordinate.
type Point struct {
	X float64
	Y float64
}

// Points represents a group of points.
type Points []Point

// String returns a string representation of the points.
func (p Points) String() string {
	var values []string
	for _, v := range p {
		values = append(values, fmt.Sprintf("%d,%d", int(v.X), int(v.Y)))
	}
	return strings.Join(values, "\n")
}

// Len returns the length of the points set.
func (p Points) Len() int {
	return len(p)
}

// Swap swaps two elments.
func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less returns if the X value of one element is less than another.
// This is the default sort for charts where you plot by x values in order.
func (p Points) Less(i, j int) bool {
	return p[i].X < p[j].X
}
