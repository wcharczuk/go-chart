package drawing

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Basic Colors from:
// https://www.w3.org/wiki/CSS/Properties/color/keywords
var (
	// ColorTransparent is a fully transparent color.
	ColorTransparent = Color{R: 255, G: 255, B: 255, A: 0}
	// ColorWhite is white.
	ColorWhite = Color{R: 255, G: 255, B: 255, A: 255}
	// ColorBlack is black.
	ColorBlack = Color{R: 0, G: 0, B: 0, A: 255}
	// ColorRed is red.
	ColorRed = Color{R: 255, G: 0, B: 0, A: 255}
	// ColorGreen is green.
	ColorGreen = Color{R: 0, G: 128, B: 0, A: 255}
	// ColorBlue is blue.
	ColorBlue = Color{R: 0, G: 0, B: 255, A: 255}
	// ColorSilver is a known color.
	ColorSilver = Color{R: 192, G: 192, B: 192, A: 255}
	// ColorMaroon is a known color.
	ColorMaroon = Color{R: 128, G: 0, B: 0, A: 255}
	// ColorPurple is a known color.
	ColorPurple = Color{R: 128, G: 0, B: 128, A: 255}
	// ColorFuchsia is a known color.
	ColorFuchsia = Color{R: 255, G: 0, B: 255, A: 255}
	// ColorLime is a known color.
	ColorLime = Color{R: 0, G: 255, B: 0, A: 255}
	// ColorOlive is a known color.
	ColorOlive = Color{R: 128, G: 128, B: 0, A: 255}
	// ColorYellow is a known color.
	ColorYellow = Color{R: 255, G: 255, B: 0, A: 255}
	// ColorNavy is a known color.
	ColorNavy = Color{R: 0, G: 0, B: 128, A: 255}
	// ColorTeal is a known color.
	ColorTeal = Color{R: 0, G: 128, B: 128, A: 255}
	// ColorAqua is a known color.
	ColorAqua = Color{R: 0, G: 255, B: 255, A: 255}
)

func parseHex(hex string) uint8 {
	v, _ := strconv.ParseInt(hex, 16, 16)
	return uint8(v)
}

// ParseColor parses a color from a string.
func ParseColor(rawColor string) Color {
	if strings.HasPrefix(rawColor, "rgba") {
		return ColorFromRGBA(rawColor)
	}
	if strings.HasPrefix(rawColor, "rgb") {
		return ColorFromRGB(rawColor)
	}
	if strings.HasPrefix(rawColor, "#") {
		return ColorFromHex(rawColor)
	}
	return ColorFromKnown(rawColor)
}

var rgbaexpr = regexp.MustCompile(`rgba\((?P<R>.+),(?P<G>.+),(?P<B>.+),(?P<A>.+)\)`)

// ColorFromRGBA returns a color from an `rgba()` css function.
func ColorFromRGBA(rgba string) (output Color) {
	values := rgbaexpr.FindStringSubmatch(rgba)
	for i, name := range rgbaexpr.SubexpNames() {
		if i == 0 {
			continue
		}
		if i >= len(values) {
			break
		}
		switch name {
		case "R":
			value := strings.TrimSpace(values[i])
			parsed, _ := strconv.ParseInt(value, 10, 16)
			output.R = uint8(parsed)
		case "G":
			value := strings.TrimSpace(values[i])
			parsed, _ := strconv.ParseInt(value, 10, 16)
			output.G = uint8(parsed)
		case "B":
			value := strings.TrimSpace(values[i])
			parsed, _ := strconv.ParseInt(value, 10, 16)
			output.B = uint8(parsed)
		case "A":
			value := strings.TrimSpace(values[i])
			parsed, _ := strconv.ParseFloat(value, 32)
			if parsed > 1 {
				parsed = 1
			} else if parsed < 0 {
				parsed = 0
			}
			output.A = uint8(parsed * 255)
		}
	}
	return
}

var rgbexpr = regexp.MustCompile(`rgb\((?P<R>.+),(?P<G>.+),(?P<B>.+)\)`)

// ColorFromRGB returns a color from an `rgb()` css function.
func ColorFromRGB(rgb string) (output Color) {
	output.A = 255
	values := rgbexpr.FindStringSubmatch(rgb)
	for i, name := range rgbaexpr.SubexpNames() {
		if i == 0 {
			continue
		}
		if i >= len(values) {
			break
		}
		switch name {
		case "R":
			value := strings.TrimSpace(values[i])
			parsed, _ := strconv.ParseInt(value, 10, 16)
			output.R = uint8(parsed)
		case "G":
			value := strings.TrimSpace(values[i])
			parsed, _ := strconv.ParseInt(value, 10, 16)
			output.G = uint8(parsed)
		case "B":
			value := strings.TrimSpace(values[i])
			parsed, _ := strconv.ParseInt(value, 10, 16)
			output.B = uint8(parsed)
		}
	}
	return
}

// ColorFromHex returns a color from a css hex code.
//
// NOTE: it will trim a leading '#' character if present.
func ColorFromHex(hex string) Color {
	if strings.HasPrefix(hex, "#") {
		hex = strings.TrimPrefix(hex, "#")
	}
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

// ColorFromKnown returns an internal color from a known (basic) color name.
func ColorFromKnown(known string) Color {
	switch strings.ToLower(known) {
	case "transparent":
		return ColorTransparent
	case "white":
		return ColorWhite
	case "black":
		return ColorBlack
	case "red":
		return ColorRed
	case "blue":
		return ColorBlue
	case "green":
		return ColorGreen
	case "silver":
		return ColorSilver
	case "maroon":
		return ColorMaroon
	case "purple":
		return ColorPurple
	case "fuchsia":
		return ColorFuchsia
	case "lime":
		return ColorLime
	case "olive":
		return ColorOlive
	case "yellow":
		return ColorYellow
	case "navy":
		return ColorNavy
	case "teal":
		return ColorTeal
	case "aqua":
		return ColorAqua
	default:
		return Color{}
	}
}

// ColorFromAlphaMixedRGBA returns the system alpha mixed rgba values.
func ColorFromAlphaMixedRGBA(r, g, b, a uint32) Color {
	fa := float64(a) / 255.0
	var c Color
	c.R = uint8(float64(r) / fa)
	c.G = uint8(float64(g) / fa)
	c.B = uint8(float64(b) / fa)
	c.A = uint8(a | (a >> 8))
	return c
}

// ColorChannelFromFloat returns a normalized byte from a given float value.
func ColorChannelFromFloat(v float64) uint8 {
	return uint8(v * 255)
}

// Color is our internal color type because color.Color is bullshit.
type Color struct {
	R, G, B, A uint8
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

// Equals returns true if the color equals another.
func (c Color) Equals(other Color) bool {
	return c.R == other.R &&
		c.G == other.G &&
		c.B == other.B &&
		c.A == other.A
}

// AverageWith averages two colors.
func (c Color) AverageWith(other Color) Color {
	return Color{
		R: (c.R + other.R) >> 1,
		G: (c.G + other.G) >> 1,
		B: (c.B + other.B) >> 1,
		A: c.A,
	}
}

// String returns a css string representation of the color.
func (c Color) String() string {
	fa := float64(c.A) / float64(255)
	return fmt.Sprintf("rgba(%v,%v,%v,%.1f)", c.R, c.G, c.B, fa)
}
