package chart

import (
	"math"
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestBoxClone(t *testing.T) {
	// replaced new assertions helper
	a := Box{Top: 5, Left: 5, Right: 15, Bottom: 15}
	b := a.Clone()
	testutil.AssertTrue(t, a.Equals(b))
	testutil.AssertTrue(t, b.Equals(a))
}

func TestBoxEquals(t *testing.T) {
	// replaced new assertions helper

	a := Box{Top: 5, Left: 5, Right: 15, Bottom: 15}
	b := Box{Top: 10, Left: 10, Right: 30, Bottom: 30}
	c := Box{Top: 5, Left: 5, Right: 15, Bottom: 15}
	testutil.AssertTrue(t, a.Equals(a))
	testutil.AssertTrue(t, a.Equals(c))
	testutil.AssertTrue(t, c.Equals(a))
	testutil.AssertFalse(t, a.Equals(b))
	testutil.AssertFalse(t, c.Equals(b))
	testutil.AssertFalse(t, b.Equals(a))
	testutil.AssertFalse(t, b.Equals(c))
}

func TestBoxIsBiggerThan(t *testing.T) {
	// replaced new assertions helper

	a := Box{Top: 5, Left: 5, Right: 25, Bottom: 25}
	b := Box{Top: 10, Left: 10, Right: 20, Bottom: 20} // only half bigger
	c := Box{Top: 1, Left: 1, Right: 30, Bottom: 30}   //bigger
	testutil.AssertTrue(t, a.IsBiggerThan(b))
	testutil.AssertFalse(t, a.IsBiggerThan(c))
	testutil.AssertTrue(t, c.IsBiggerThan(a))
}

func TestBoxIsSmallerThan(t *testing.T) {
	// replaced new assertions helper

	a := Box{Top: 5, Left: 5, Right: 25, Bottom: 25}
	b := Box{Top: 10, Left: 10, Right: 20, Bottom: 20} // only half bigger
	c := Box{Top: 1, Left: 1, Right: 30, Bottom: 30}   //bigger
	testutil.AssertFalse(t, a.IsSmallerThan(b))
	testutil.AssertTrue(t, a.IsSmallerThan(c))
	testutil.AssertFalse(t, c.IsSmallerThan(a))
}

func TestBoxGrow(t *testing.T) {
	// replaced new assertions helper

	a := Box{Top: 1, Left: 2, Right: 15, Bottom: 15}
	b := Box{Top: 4, Left: 5, Right: 30, Bottom: 35}
	c := a.Grow(b)
	testutil.AssertFalse(t, c.Equals(b))
	testutil.AssertFalse(t, c.Equals(a))
	testutil.AssertEqual(t, 1, c.Top)
	testutil.AssertEqual(t, 2, c.Left)
	testutil.AssertEqual(t, 30, c.Right)
	testutil.AssertEqual(t, 35, c.Bottom)
}

func TestBoxFit(t *testing.T) {
	// replaced new assertions helper

	a := Box{Top: 64, Left: 64, Right: 192, Bottom: 192}
	b := Box{Top: 16, Left: 16, Right: 256, Bottom: 170}
	c := Box{Top: 16, Left: 16, Right: 170, Bottom: 256}

	fab := a.Fit(b)
	testutil.AssertEqual(t, a.Left, fab.Left)
	testutil.AssertEqual(t, a.Right, fab.Right)
	testutil.AssertTrue(t, fab.Top < fab.Bottom)
	testutil.AssertTrue(t, fab.Left < fab.Right)
	testutil.AssertTrue(t, math.Abs(b.Aspect()-fab.Aspect()) < 0.02)

	fac := a.Fit(c)
	testutil.AssertEqual(t, a.Top, fac.Top)
	testutil.AssertEqual(t, a.Bottom, fac.Bottom)
	testutil.AssertTrue(t, math.Abs(c.Aspect()-fac.Aspect()) < 0.02)
}

func TestBoxConstrain(t *testing.T) {
	// replaced new assertions helper

	a := Box{Top: 64, Left: 64, Right: 192, Bottom: 192}
	b := Box{Top: 16, Left: 16, Right: 256, Bottom: 170}
	c := Box{Top: 16, Left: 16, Right: 170, Bottom: 256}

	cab := a.Constrain(b)
	testutil.AssertEqual(t, 64, cab.Top)
	testutil.AssertEqual(t, 64, cab.Left)
	testutil.AssertEqual(t, 192, cab.Right)
	testutil.AssertEqual(t, 170, cab.Bottom)

	cac := a.Constrain(c)
	testutil.AssertEqual(t, 64, cac.Top)
	testutil.AssertEqual(t, 64, cac.Left)
	testutil.AssertEqual(t, 170, cac.Right)
	testutil.AssertEqual(t, 192, cac.Bottom)
}

func TestBoxOuterConstrain(t *testing.T) {
	// replaced new assertions helper

	box := NewBox(0, 0, 100, 100)
	canvas := NewBox(5, 5, 95, 95)
	taller := NewBox(-10, 5, 50, 50)

	c := canvas.OuterConstrain(box, taller)
	testutil.AssertEqual(t, 15, c.Top, c.String())
	testutil.AssertEqual(t, 5, c.Left, c.String())
	testutil.AssertEqual(t, 95, c.Right, c.String())
	testutil.AssertEqual(t, 95, c.Bottom, c.String())

	wider := NewBox(5, 5, 110, 50)
	d := canvas.OuterConstrain(box, wider)
	testutil.AssertEqual(t, 5, d.Top, d.String())
	testutil.AssertEqual(t, 5, d.Left, d.String())
	testutil.AssertEqual(t, 85, d.Right, d.String())
	testutil.AssertEqual(t, 95, d.Bottom, d.String())
}

func TestBoxShift(t *testing.T) {
	// replaced new assertions helper

	b := Box{
		Top:    5,
		Left:   5,
		Right:  10,
		Bottom: 10,
	}

	shifted := b.Shift(1, 2)
	testutil.AssertEqual(t, 7, shifted.Top)
	testutil.AssertEqual(t, 6, shifted.Left)
	testutil.AssertEqual(t, 11, shifted.Right)
	testutil.AssertEqual(t, 12, shifted.Bottom)
}

func TestBoxCenter(t *testing.T) {
	// replaced new assertions helper

	b := Box{
		Top:    10,
		Left:   10,
		Right:  20,
		Bottom: 30,
	}
	cx, cy := b.Center()
	testutil.AssertEqual(t, 15, cx)
	testutil.AssertEqual(t, 20, cy)
}

func TestBoxCornersCenter(t *testing.T) {
	// replaced new assertions helper

	bc := BoxCorners{
		TopLeft:     Point{5, 5},
		TopRight:    Point{15, 5},
		BottomRight: Point{15, 15},
		BottomLeft:  Point{5, 15},
	}

	cx, cy := bc.Center()
	testutil.AssertEqual(t, 10, cx)
	testutil.AssertEqual(t, 10, cy)
}

func TestBoxCornersRotate(t *testing.T) {
	// replaced new assertions helper

	bc := BoxCorners{
		TopLeft:     Point{5, 5},
		TopRight:    Point{15, 5},
		BottomRight: Point{15, 15},
		BottomLeft:  Point{5, 15},
	}

	rotated := bc.Rotate(45)
	testutil.AssertTrue(t, rotated.TopLeft.Equals(Point{10, 3}), rotated.String())
}
