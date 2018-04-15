package chart

import (
	"math"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestBoxClone(t *testing.T) {
	assert := assert.New(t)
	a := Box{Top: 5, Left: 5, Right: 15, Bottom: 15}
	b := a.Clone()
	assert.True(a.Equals(b))
	assert.True(b.Equals(a))
}

func TestBoxEquals(t *testing.T) {
	assert := assert.New(t)

	a := Box{Top: 5, Left: 5, Right: 15, Bottom: 15}
	b := Box{Top: 10, Left: 10, Right: 30, Bottom: 30}
	c := Box{Top: 5, Left: 5, Right: 15, Bottom: 15}
	assert.True(a.Equals(a))
	assert.True(a.Equals(c))
	assert.True(c.Equals(a))
	assert.False(a.Equals(b))
	assert.False(c.Equals(b))
	assert.False(b.Equals(a))
	assert.False(b.Equals(c))
}

func TestBoxIsBiggerThan(t *testing.T) {
	assert := assert.New(t)

	a := Box{Top: 5, Left: 5, Right: 25, Bottom: 25}
	b := Box{Top: 10, Left: 10, Right: 20, Bottom: 20} // only half bigger
	c := Box{Top: 1, Left: 1, Right: 30, Bottom: 30}   //bigger
	assert.True(a.IsBiggerThan(b))
	assert.False(a.IsBiggerThan(c))
	assert.True(c.IsBiggerThan(a))
}

func TestBoxIsSmallerThan(t *testing.T) {
	assert := assert.New(t)

	a := Box{Top: 5, Left: 5, Right: 25, Bottom: 25}
	b := Box{Top: 10, Left: 10, Right: 20, Bottom: 20} // only half bigger
	c := Box{Top: 1, Left: 1, Right: 30, Bottom: 30}   //bigger
	assert.False(a.IsSmallerThan(b))
	assert.True(a.IsSmallerThan(c))
	assert.False(c.IsSmallerThan(a))
}

func TestBoxGrow(t *testing.T) {
	assert := assert.New(t)

	a := Box{Top: 1, Left: 2, Right: 15, Bottom: 15}
	b := Box{Top: 4, Left: 5, Right: 30, Bottom: 35}
	c := a.Grow(b)
	assert.False(c.Equals(b))
	assert.False(c.Equals(a))
	assert.Equal(1, c.Top)
	assert.Equal(2, c.Left)
	assert.Equal(30, c.Right)
	assert.Equal(35, c.Bottom)
}

func TestBoxFit(t *testing.T) {
	assert := assert.New(t)

	a := Box{Top: 64, Left: 64, Right: 192, Bottom: 192}
	b := Box{Top: 16, Left: 16, Right: 256, Bottom: 170}
	c := Box{Top: 16, Left: 16, Right: 170, Bottom: 256}

	fab := a.Fit(b)
	assert.Equal(a.Left, fab.Left)
	assert.Equal(a.Right, fab.Right)
	assert.True(fab.Top < fab.Bottom)
	assert.True(fab.Left < fab.Right)
	assert.True(math.Abs(b.Aspect()-fab.Aspect()) < 0.02)

	fac := a.Fit(c)
	assert.Equal(a.Top, fac.Top)
	assert.Equal(a.Bottom, fac.Bottom)
	assert.True(math.Abs(c.Aspect()-fac.Aspect()) < 0.02)
}

func TestBoxConstrain(t *testing.T) {
	assert := assert.New(t)

	a := Box{Top: 64, Left: 64, Right: 192, Bottom: 192}
	b := Box{Top: 16, Left: 16, Right: 256, Bottom: 170}
	c := Box{Top: 16, Left: 16, Right: 170, Bottom: 256}

	cab := a.Constrain(b)
	assert.Equal(64, cab.Top)
	assert.Equal(64, cab.Left)
	assert.Equal(192, cab.Right)
	assert.Equal(170, cab.Bottom)

	cac := a.Constrain(c)
	assert.Equal(64, cac.Top)
	assert.Equal(64, cac.Left)
	assert.Equal(170, cac.Right)
	assert.Equal(192, cac.Bottom)
}

func TestBoxOuterConstrain(t *testing.T) {
	assert := assert.New(t)

	box := NewBox(0, 0, 100, 100)
	canvas := NewBox(5, 5, 95, 95)
	taller := NewBox(-10, 5, 50, 50)

	c := canvas.OuterConstrain(box, taller)
	assert.Equal(15, c.Top, c.String())
	assert.Equal(5, c.Left, c.String())
	assert.Equal(95, c.Right, c.String())
	assert.Equal(95, c.Bottom, c.String())

	wider := NewBox(5, 5, 110, 50)
	d := canvas.OuterConstrain(box, wider)
	assert.Equal(5, d.Top, d.String())
	assert.Equal(5, d.Left, d.String())
	assert.Equal(85, d.Right, d.String())
	assert.Equal(95, d.Bottom, d.String())
}

func TestBoxShift(t *testing.T) {
	assert := assert.New(t)

	b := Box{
		Top:    5,
		Left:   5,
		Right:  10,
		Bottom: 10,
	}

	shifted := b.Shift(1, 2)
	assert.Equal(7, shifted.Top)
	assert.Equal(6, shifted.Left)
	assert.Equal(11, shifted.Right)
	assert.Equal(12, shifted.Bottom)
}

func TestBoxCenter(t *testing.T) {
	assert := assert.New(t)

	b := Box{
		Top:    10,
		Left:   10,
		Right:  20,
		Bottom: 30,
	}
	cx, cy := b.Center()
	assert.Equal(15, cx)
	assert.Equal(20, cy)
}

func TestBoxCornersCenter(t *testing.T) {
	assert := assert.New(t)

	bc := BoxCorners{
		TopLeft:     Point{5, 5},
		TopRight:    Point{15, 5},
		BottomRight: Point{15, 15},
		BottomLeft:  Point{5, 15},
	}

	cx, cy := bc.Center()
	assert.Equal(10, cx)
	assert.Equal(10, cy)
}

func TestBoxCornersRotate(t *testing.T) {
	assert := assert.New(t)

	bc := BoxCorners{
		TopLeft:     Point{5, 5},
		TopRight:    Point{15, 5},
		BottomRight: Point{15, 15},
		BottomLeft:  Point{5, 15},
	}

	rotated := bc.Rotate(45)
	assert.True(rotated.TopLeft.Equals(Point{10, 3}), rotated.String())
}
