package chart

import "image"

// Imager interface is implemented by renderers that want to expose the data
// as the image.Image interface.
type Imager interface {
	// ToImage returns a image.Image
	ToImage() image.Image
}
