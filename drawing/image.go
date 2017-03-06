package drawing

import "image"

// Image is a helper wraper that allows (sane) access to pixel info.
type Image struct {
	Inner *image.RGBA
}

// Width returns the image's width in pixels.
func (i Image) Width() int {
	return i.Inner.Rect.Size().X
}

// Height returns the image's height in pixels.
func (i Image) Height() int {
	return i.Inner.Rect.Size().Y
}

// At returns a pixel color at a given x/y.
func (i Image) At(x, y int) Color {
	return ColorFromAlphaMixedRGBA(i.Inner.At(x, y).RGBA())
}
