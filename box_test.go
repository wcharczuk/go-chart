package chart

import (
	"math"
	"testing"

	"github.com/blendlabs/go-assert"
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
