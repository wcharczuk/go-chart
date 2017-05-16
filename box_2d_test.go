package chart

import (
	"fmt"
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestBox2dCenter(t *testing.T) {
	assert := assert.New(t)

	bc := Box2d{
		TopLeft:     Point{5, 5},
		TopRight:    Point{15, 5},
		BottomRight: Point{15, 15},
		BottomLeft:  Point{5, 15},
	}

	cx, cy := bc.Center()
	assert.Equal(10, cx)
	assert.Equal(10, cy)
}

func TestBox2dRotate(t *testing.T) {
	assert := assert.New(t)

	bc := Box2d{
		TopLeft:     Point{5, 5},
		TopRight:    Point{15, 5},
		BottomRight: Point{15, 15},
		BottomLeft:  Point{5, 15},
	}

	rotated := bc.Rotate(45)
	assert.True(rotated.TopLeft.Equals(Point{10, 3}), rotated.String())
}

func TestBox2dOverlaps(t *testing.T) {
	assert := assert.New(t)

	bc := Box2d{
		TopLeft:     Point{5, 5},
		TopRight:    Point{15, 5},
		BottomRight: Point{15, 15},
		BottomLeft:  Point{5, 15},
	}

	// shift meaningfully the full width of bc right.
	bc2 := bc.Shift(bc.Width()+1, 0)
	assert.False(bc.Overlaps(bc2), fmt.Sprintf("%v\n\t\tshould not overlap\n\t%v", bc, bc2))

	// shift meaningfully the full height of bc down.
	bc3 := bc.Shift(0, bc.Height()+1)
	assert.False(bc.Overlaps(bc3), fmt.Sprintf("%v\n\t\tshould not overlap\n\t%v", bc, bc3))

	bc4 := bc.Shift(5, 0)
	assert.True(bc.Overlaps(bc4))

	bc5 := bc.Shift(0, 5)
	assert.True(bc.Overlaps(bc5))

	bcr := bc.Rotate(45)
	bcr2 := bc.Rotate(45).Shift(bc.Height(), 0)
	assert.False(bcr.Overlaps(bcr2), fmt.Sprintf("%v\n\t\tshould not overlap\n\t%v", bcr, bcr2))
}
