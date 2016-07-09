package drawing

import (
	"fmt"
	"strconv"
)

var (
	// ColorTransparent is a fully transparent color.
	ColorTransparent = Color{}

	// ColorWhite is white.
	ColorWhite = Color{R: 255, G: 255, B: 255, A: 255}

	// ColorBlack is black.
	ColorBlack = Color{R: 0, G: 0, B: 0, A: 255}

	// ColorRed is red.
	ColorRed = Color{R: 255, G: 0, B: 0, A: 255}

	// ColorGreen is green.
	ColorGreen = Color{R: 0, G: 255, B: 0, A: 255}

	// ColorBlue is blue.
	ColorBlue = Color{R: 0, G: 0, B: 255, A: 255}
)

func parseHex(hex string) uint8 {
	v, _ := strconv.ParseInt(hex, 16, 16)
	return uint8(v)
}

// ColorFromHex returns a color from a css hex code.
func ColorFromHex(hex string) Color {
	var c Color
	if len(hex) == 3 {
		c.R = parseHex(string(hex[0])) * 0x11
		c.G = parseHex(string(hex[1])) * 0x11
		c.B = parseHex(string(hex[2])) * 0x11
	} else {
		c.R = parseHex(string(hex[0:2]))
		c.G = parseHex(string(hex[2:4]))
		c.B = parseHex(string(hex[4:6]))
	}
	c.A = 255
	return c
}

// Color is our internal color type because color.Color is bullshit.
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// RGBA returns the color as a pre-alpha mixed color set.
func (c Color) RGBA() (r, g, b, a uint32) {
	fa := float64(c.A) / 255.0
	r = uint32(float64(uint32(c.R)) * fa)
	r |= r << 8
	g = uint32(float64(uint32(c.G)) * fa)
	g |= g << 8
	b = uint32(float64(uint32(c.B)) * fa)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

// IsZero returns if the color has been set or not.
func (c Color) IsZero() bool {
	return c.R == 0 && c.G == 0 && c.B == 0 && c.A == 0
}

// IsTransparent returns if the colors alpha channel is zero.
func (c Color) IsTransparent() bool {
	return c.A == 0
}

// WithAlpha returns a copy of the color with a given alpha.
func (c Color) WithAlpha(a uint8) Color {
	return Color{
		R: c.R,
		G: c.G,
		B: c.B,
		A: a,
	}
}

// String returns a css string representation of the color.
func (c Color) String() string {
	fa := float64(c.A) / float64(255)
	return fmt.Sprintf("rgba(%v,%v,%v,%.1f)", c.R, c.G, c.B, fa)
}
