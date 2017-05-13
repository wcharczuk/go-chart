package util

import (
	"math"
	"time"
)

const (
	_pi   = math.Pi
	_2pi  = 2 * math.Pi
	_3pi4 = (3 * math.Pi) / 4.0
	_4pi3 = (4 * math.Pi) / 3.0
	_3pi2 = (3 * math.Pi) / 2.0
	_5pi4 = (5 * math.Pi) / 4.0
	_7pi4 = (7 * math.Pi) / 4.0
	_pi2  = math.Pi / 2.0
	_pi4  = math.Pi / 4.0
	_d2r  = (math.Pi / 180.0)
	_r2d  = (180.0 / math.Pi)
)

var (
	// Math contains helper methods for common math operations.
	Math = &mathUtil{}
)

type mathUtil struct{}

// Max returns the maximum value of a group of floats.
func (m mathUtil) Max(values ...float64) float64 {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, v := range values {
		if max < v {
			max = v
		}
	}
	return max
}

// MinAndMax returns both the min and max in one pass.
func (m mathUtil) MinAndMax(values ...float64) (min float64, max float64) {
	if len(values) == 0 {
		return
	}
	min = values[0]
	max = values[0]
	for _, v := range values[1:] {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return
}

// MinAndMaxOfTime returns the min and max of a given set of times
// in one pass.
func (m mathUtil) MinAndMaxOfTime(values ...time.Time) (min time.Time, max time.Time) {
	if len(values) == 0 {
		return
	}

	min = values[0]
	max = values[0]

	for _, v := range values[1:] {
		if max.Before(v) {
			max = v
		}
		if min.After(v) {
			min = v
		}
	}
	return
}

// GetRoundToForDelta returns a `roundTo` value for a given delta.
func (m mathUtil) GetRoundToForDelta(delta float64) float64 {
	startingDeltaBound := math.Pow(10.0, 10.0)
	for cursor := startingDeltaBound; cursor > 0; cursor /= 10.0 {
		if delta > cursor {
			return cursor / 10.0
		}
	}

	return 0.0
}

// RoundUp rounds up to a given roundTo value.
func (m mathUtil) RoundUp(value, roundTo float64) float64 {
	d1 := math.Ceil(value / roundTo)
	return d1 * roundTo
}

// RoundDown rounds down to a given roundTo value.
func (m mathUtil) RoundDown(value, roundTo float64) float64 {
	d1 := math.Floor(value / roundTo)
	return d1 * roundTo
}

// Normalize returns a set of numbers on the interval [0,1] for a given set of inputs.
// An example: 4,3,2,1 => 0.4, 0.3, 0.2, 0.1
// Caveat; the total may be < 1.0; there are going to be issues with irrational numbers etc.
func (m mathUtil) Normalize(values ...float64) []float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	output := make([]float64, len(values))
	for x, v := range values {
		output[x] = m.RoundDown(v/total, 0.0001)
	}
	return output
}

// MinInt returns the minimum of a set of integers.
func (m mathUtil) MinInt(values ...int) int {
	min := math.MaxInt32
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// MaxInt returns the maximum of a set of integers.
func (m mathUtil) MaxInt(values ...int) int {
	max := math.MinInt32
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// AbsInt returns the absolute value of an integer.
func (m mathUtil) AbsInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// AbsInt64 returns the absolute value of a long.
func (m mathUtil) AbsInt64(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

// Mean returns the mean of a set of values
func (m mathUtil) Mean(values ...float64) float64 {
	return m.Sum(values...) / float64(len(values))
}

// MeanInt returns the mean of a set of integer values.
func (m mathUtil) MeanInt(values ...int) int {
	return m.SumInt(values...) / len(values)
}

// Sum sums a set of values.
func (m mathUtil) Sum(values ...float64) float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	return total
}

// SumInt sums a set of values.
func (m mathUtil) SumInt(values ...int) int {
	var total int
	for _, v := range values {
		total += v
	}
	return total
}

// PercentDifference computes the percentage difference between two values.
// The formula is (v2-v1)/v1.
func (m mathUtil) PercentDifference(v1, v2 float64) float64 {
	if v1 == 0 {
		return 0
	}
	return (v2 - v1) / v1
}

// DegreesToRadians returns degrees as radians.
func (m mathUtil) DegreesToRadians(degrees float64) float64 {
	return degrees * _d2r
}

// RadiansToDegrees translates a radian value to a degree value.
func (m mathUtil) RadiansToDegrees(value float64) float64 {
	return math.Mod(value, _2pi) * _r2d
}

// PercentToRadians converts a normalized value (0,1) to radians.
func (m mathUtil) PercentToRadians(pct float64) float64 {
	return m.DegreesToRadians(360.0 * pct)
}

// RadianAdd adds a delta to a base in radians.
func (m mathUtil) RadianAdd(base, delta float64) float64 {
	value := base + delta
	if value > _2pi {
		return math.Mod(value, _2pi)
	} else if value < 0 {
		return math.Mod(_2pi+value, _2pi)
	}
	return value
}

// DegreesAdd adds a delta to a base in radians.
func (m mathUtil) DegreesAdd(baseDegrees, deltaDegrees float64) float64 {
	value := baseDegrees + deltaDegrees
	if value > _2pi {
		return math.Mod(value, 360.0)
	} else if value < 0 {
		return math.Mod(360.0+value, 360.0)
	}
	return value
}

// DegreesToCompass returns the degree value in compass / clock orientation.
func (m mathUtil) DegreesToCompass(deg float64) float64 {
	return m.DegreesAdd(deg, -90.0)
}

// CirclePoint returns the absolute position of a circle diameter point given
// by the radius and the theta.
func (m mathUtil) CirclePoint(cx, cy int, radius, thetaRadians float64) (x, y int) {
	x = cx + int(radius*math.Sin(thetaRadians))
	y = cy - int(radius*math.Cos(thetaRadians))
	return
}

func (m mathUtil) RotateCoordinate(cx, cy, x, y int, thetaRadians float64) (rx, ry int) {
	tempX, tempY := float64(x-cx), float64(y-cy)
	rotatedX := tempX*math.Cos(thetaRadians) - tempY*math.Sin(thetaRadians)
	rotatedY := tempX*math.Sin(thetaRadians) + tempY*math.Cos(thetaRadians)
	rx = int(rotatedX) + cx
	ry = int(rotatedY) + cy
	return
}
