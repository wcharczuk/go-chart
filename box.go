package chart

import "fmt"

// Box represents the main 4 dimensions of a box.
type Box struct {
	Top    int
	Left   int
	Right  int
	Bottom int
}

// IsZero returns if the box is set or not.
func (b Box) IsZero() bool {
	return b.Top == 0 && b.Left == 0 && b.Right == 0 && b.Bottom == 0
}

// String returns a string representation of the box.
func (b Box) String() string {
	return fmt.Sprintf("box(%d,%d,%d,%d)", b.Top, b.Left, b.Right, b.Bottom)
}

// GetTop returns a coalesced value with a default.
func (b Box) GetTop(defaults ...int) int {
	if b.Top == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return 0
	}
	return b.Top
}

// GetLeft returns a coalesced value with a default.
func (b Box) GetLeft(defaults ...int) int {
	if b.Left == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return 0
	}
	return b.Left
}

// GetRight returns a coalesced value with a default.
func (b Box) GetRight(defaults ...int) int {
	if b.Right == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return 0
	}
	return b.Right
}

// GetBottom returns a coalesced value with a default.
func (b Box) GetBottom(defaults ...int) int {
	if b.Bottom == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return 0
	}
	return b.Bottom
}

// Width returns the width
func (b Box) Width() int {
	return Math.AbsInt(b.Right - b.Left)
}

// Height returns the height
func (b Box) Height() int {
	return Math.AbsInt(b.Bottom - b.Top)
}

// Center returns the center of the box
func (b Box) Center() (x, y int) {
	w, h := b.Width(), b.Height()
	return b.Left + w>>1, b.Top + h>>1
}

// Aspect returns the aspect ratio of the box.
func (b Box) Aspect() float64 {
	return float64(b.Width()) / float64(b.Height())
}

// Clone returns a new copy of the box.
func (b Box) Clone() Box {
	return Box{
		Top:    b.Top,
		Left:   b.Left,
		Right:  b.Right,
		Bottom: b.Bottom,
	}
}

// IsBiggerThan returns if a box is bigger than another box.
func (b Box) IsBiggerThan(other Box) bool {
	return b.Top < other.Top ||
		b.Bottom > other.Bottom ||
		b.Left < other.Left ||
		b.Right > other.Right
}

// IsSmallerThan returns if a box is smaller than another box.
func (b Box) IsSmallerThan(other Box) bool {
	return b.Top > other.Top &&
		b.Bottom < other.Bottom &&
		b.Left > other.Left &&
		b.Right < other.Right
}

// Equals returns if the box equals another box.
func (b Box) Equals(other Box) bool {
	return b.Top == other.Top &&
		b.Left == other.Left &&
		b.Right == other.Right &&
		b.Bottom == other.Bottom
}

// Grow grows a box based on another box.
func (b Box) Grow(other Box) Box {
	return Box{
		Top:    Math.MinInt(b.Top, other.Top),
		Left:   Math.MinInt(b.Left, other.Left),
		Right:  Math.MaxInt(b.Right, other.Right),
		Bottom: Math.MaxInt(b.Bottom, other.Bottom),
	}
}

// Shift pushes a box by x,y.
func (b Box) Shift(x, y int) Box {
	return Box{
		Top:    b.Top + y,
		Left:   b.Left + x,
		Right:  b.Right + x,
		Bottom: b.Bottom + y,
	}
}

// Fit is functionally the inverse of grow.
// Fit maintains the original aspect ratio of the `other` box,
// but constrains it to the bounds of the target box.
func (b Box) Fit(other Box) Box {
	ba := b.Aspect()
	oa := other.Aspect()

	if oa == ba {
		return b.Clone()
	}

	bw, bh := float64(b.Width()), float64(b.Height())
	bw2 := int(bw) >> 1
	bh2 := int(bh) >> 1
	if oa > ba { // ex. 16:9 vs. 4:3
		var noh2 int
		if oa > 1.0 {
			noh2 = int(bw/oa) >> 1
		} else {
			noh2 = int(bh*oa) >> 1
		}
		return Box{
			Top:    (b.Top + bh2) - noh2,
			Left:   b.Left,
			Right:  b.Right,
			Bottom: (b.Top + bh2) + noh2,
		}
	}
	var now2 int
	if oa > 1.0 {
		now2 = int(bh/oa) >> 1
	} else {
		now2 = int(bw*oa) >> 1
	}
	return Box{
		Top:    b.Top,
		Left:   (b.Left + bw2) - now2,
		Right:  (b.Left + bw2) + now2,
		Bottom: b.Bottom,
	}
}

// Constrain is similar to `Fit` except that it will work
// more literally like the opposite of grow.
func (b Box) Constrain(other Box) Box {
	newBox := b.Clone()

	newBox.Top = Math.MaxInt(newBox.Top, other.Top)
	newBox.Left = Math.MaxInt(newBox.Left, other.Left)
	newBox.Right = Math.MinInt(newBox.Right, other.Right)
	newBox.Bottom = Math.MinInt(newBox.Bottom, other.Bottom)

	return newBox
}

// OuterConstrain is similar to `Constraint` with the difference
// that it applies corrections
func (b Box) OuterConstrain(bounds, other Box) Box {
	newBox := b.Clone()
	if other.Top < bounds.Top {
		delta := bounds.Top - other.Top
		newBox.Top = b.Top + delta
	}

	if other.Left < bounds.Left {
		delta := bounds.Left - other.Left
		newBox.Left = b.Left + delta
	}

	if other.Right > bounds.Right {
		delta := other.Right - bounds.Right
		newBox.Right = b.Right - delta
	}

	if other.Bottom > bounds.Bottom {
		delta := other.Bottom - bounds.Bottom
		newBox.Bottom = b.Bottom - delta
	}
	return newBox
}
