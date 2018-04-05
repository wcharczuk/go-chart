package drawing

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

type point struct {
	X, Y float64
}

type mockLine struct {
	inner []point
}

func (ml *mockLine) LineTo(x, y float64) {
	ml.inner = append(ml.inner, point{x, y})
}

func (ml mockLine) Len() int {
	return len(ml.inner)
}

func TestTraceQuad(t *testing.T) {
	assert := assert.New(t)

	// Quad
	// x1, y1, cpx1, cpy2, x2, y2 float64
	// do the 9->12 circle segment
	quad := []float64{10, 20, 20, 20, 20, 10}
	liner := &mockLine{}
	TraceQuad(liner, quad, 0.5)
	assert.NotZero(liner.Len())
}
