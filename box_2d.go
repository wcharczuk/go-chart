package chart

import (
	"fmt"
	"math"

	util "github.com/wcharczuk/go-chart/util"
)

// Box2d is a box with (4) independent corners.
// It is used when dealing with ~rotated~ boxes.
type Box2d struct {
	TopLeft, TopRight, BottomRight, BottomLeft Point
}

// Points returns the constituent points of the box.
func (bc Box2d) Points() []Point {
	return []Point{
		bc.TopRight,
		bc.BottomRight,
		bc.BottomLeft,
		bc.TopLeft,
	}
}

// Box return the Box2d as a regular box.
func (bc Box2d) Box() Box {
	return Box{
		Top:    int(math.Min(bc.TopLeft.Y, bc.TopRight.Y)),
		Left:   int(math.Min(bc.TopLeft.X, bc.BottomLeft.X)),
		Right:  int(math.Max(bc.TopRight.X, bc.BottomRight.X)),
		Bottom: int(math.Max(bc.BottomLeft.Y, bc.BottomRight.Y)),
	}
}

// Width returns the width
func (bc Box2d) Width() float64 {
	minLeft := math.Min(bc.TopLeft.X, bc.BottomLeft.X)
	maxRight := math.Max(bc.TopRight.X, bc.BottomRight.X)
	return maxRight - minLeft
}

// Height returns the height
func (bc Box2d) Height() float64 {
	minTop := math.Min(bc.TopLeft.Y, bc.TopRight.Y)
	maxBottom := math.Max(bc.BottomLeft.Y, bc.BottomRight.Y)
	return maxBottom - minTop
}

// Center returns the center of the box
func (bc Box2d) Center() (x, y float64) {
	left := util.Math.Mean(bc.TopLeft.X, bc.BottomLeft.X)
	right := util.Math.Mean(bc.TopRight.X, bc.BottomRight.X)
	x = ((right - left) / 2.0) + left

	top := util.Math.Mean(bc.TopLeft.Y, bc.TopRight.Y)
	bottom := util.Math.Mean(bc.BottomLeft.Y, bc.BottomRight.Y)
	y = ((bottom - top) / 2.0) + top

	return
}

// Rotate rotates the box.
func (bc Box2d) Rotate(thetaDegrees float64) Box2d {
	cx, cy := bc.Center()

	thetaRadians := util.Math.DegreesToRadians(thetaDegrees)

	tlx, tly := util.Math.RotateCoordinate(int(cx), int(cy), int(bc.TopLeft.X), int(bc.TopLeft.Y), thetaRadians)
	trx, try := util.Math.RotateCoordinate(int(cx), int(cy), int(bc.TopRight.X), int(bc.TopRight.Y), thetaRadians)
	brx, bry := util.Math.RotateCoordinate(int(cx), int(cy), int(bc.BottomRight.X), int(bc.BottomRight.Y), thetaRadians)
	blx, bly := util.Math.RotateCoordinate(int(cx), int(cy), int(bc.BottomLeft.X), int(bc.BottomLeft.Y), thetaRadians)

	return Box2d{
		TopLeft:     Point{float64(tlx), float64(tly)},
		TopRight:    Point{float64(trx), float64(try)},
		BottomRight: Point{float64(brx), float64(bry)},
		BottomLeft:  Point{float64(blx), float64(bly)},
	}
}

// Shift shifts a box by a given x and y value.
func (bc Box2d) Shift(x, y float64) Box2d {
	return Box2d{
		TopLeft:     bc.TopLeft.Shift(x, y),
		TopRight:    bc.TopRight.Shift(x, y),
		BottomRight: bc.BottomRight.Shift(x, y),
		BottomLeft:  bc.BottomLeft.Shift(x, y),
	}
}

// Equals returns if the box equals another box.
func (bc Box2d) Equals(other Box2d) bool {
	return bc.TopLeft.Equals(other.TopLeft) &&
		bc.TopRight.Equals(other.TopRight) &&
		bc.BottomRight.Equals(other.BottomRight) &&
		bc.BottomLeft.Equals(other.BottomLeft)
}

// Overlaps returns if two boxes overlap.
func (bc Box2d) Overlaps(other Box2d) bool {
	pa := bc.Points()
	pb := other.Points()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			pa0 := pa[i]
			pa1 := pa[(i+1)%4]

			pb0 := pb[j]
			pb1 := pb[(j+1)%4]

			if util.Math.LinesIntersect(pa0.X, pa0.Y, pa1.X, pa1.Y, pb0.X, pb0.Y, pb1.X, pb1.Y) {
				return true
			}
		}
	}
	return false
}

func (bc Box2d) String() string {
	return fmt.Sprintf("Box2d{%s,%s,%s,%s}", bc.TopLeft.String(), bc.TopRight.String(), bc.BottomRight.String(), bc.BottomLeft.String())
}

// Point is an X,Y pair
type Point struct {
	X, Y float64
}

// Shift shifts a point.
func (p Point) Shift(x, y float64) Point {
	return Point{
		X: p.X + x,
		Y: p.Y + y,
	}
}

// DistanceTo calculates the distance to another point.
func (p Point) DistanceTo(other Point) float64 {
	dx := math.Pow(float64(p.X-other.X), 2)
	dy := math.Pow(float64(p.Y-other.Y), 2)
	return math.Pow(dx+dy, 0.5)
}

// Equals returns if a point equals another point.
func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

// String returns a string representation of the point.
func (p Point) String() string {
	return fmt.Sprintf("P{%.2f,%.2f}", p.X, p.Y)
}
