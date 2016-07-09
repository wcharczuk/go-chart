package drawing

import (
	"math"

	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype/raster"
	"github.com/golang/freetype/truetype"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distance(x1, y1, x2, y2 float64) float64 {
	return vectorDistance(x2-x1, y2-y1)
}

func vectorDistance(dx, dy float64) float64 {
	return float64(math.Sqrt(dx*dx + dy*dy))
}

func toFtCap(c LineCap) raster.Capper {
	switch c {
	case RoundCap:
		return raster.RoundCapper
	case ButtCap:
		return raster.ButtCapper
	case SquareCap:
		return raster.SquareCapper
	}
	return raster.RoundCapper
}

func toFtJoin(j LineJoin) raster.Joiner {
	switch j {
	case RoundJoin:
		return raster.RoundJoiner
	case BevelJoin:
		return raster.BevelJoiner
	}
	return raster.RoundJoiner
}

func pointToF64Point(p truetype.Point) (x, y float64) {
	return fUnitsToFloat64(p.X), -fUnitsToFloat64(p.Y)
}

func fUnitsToFloat64(x fixed.Int26_6) float64 {
	scaled := x << 2
	return float64(scaled/256) + float64(scaled%256)/256.0
}
