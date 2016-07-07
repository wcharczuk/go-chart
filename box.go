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
	return fmt.Sprintf("Box(%d,%d,%d,%d)", b.Top, b.Left, b.Right, b.Bottom)
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
