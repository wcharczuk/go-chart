package chart

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Float is an alias for float64 that provides a better .String() method.
type Float float64

// String returns the string representation of a float.
func (f Float) String() string {
	return fmt.Sprintf("%.2f", f)
}

// TimeToFloat64 returns a float64 representation of a time.
func TimeToFloat64(t time.Time) float64 {
	return float64(t.UnixNano())
}

// Float64ToTime returns a time from a float64.
func Float64ToTime(tf float64) time.Time {
	return time.Unix(0, int64(tf))
}

// MinAndMax returns both the min and max in one pass.
func MinAndMax(values ...float64) (min float64, max float64) {
	if len(values) == 0 {
		return
	}
	min = values[0]
	max = values[0]
	for _, v := range values {
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
func MinAndMaxOfTime(values ...time.Time) (min time.Time, max time.Time) {
	if len(values) == 0 {
		return
	}

	min = values[0]
	max = values[0]

	for _, v := range values {
		if max.Before(v) {
			max = v
		}
		if min.After(v) {
			min = v
		}
	}
	return
}

// Slices generates N slices that span the total.
// The resulting array will be intermediate indexes until total.
func Slices(count int, total float64) []float64 {
	var values []float64
	sliceWidth := float64(total) / float64(count)
	for cursor := 0.0; cursor < total; cursor += sliceWidth {
		values = append(values, cursor)
	}
	return values
}

// GetRoundToForDelta returns a `roundTo` value for a given delta.
func GetRoundToForDelta(delta float64) float64 {
	startingDeltaBound := math.Pow(10.0, 10.0)
	for cursor := startingDeltaBound; cursor > 0; cursor /= 10.0 {
		if delta > cursor {
			return cursor / 10.0
		}
	}

	return 0.0
}

// RoundUp rounds up to a given roundTo value.
func RoundUp(value, roundTo float64) float64 {
	d1 := math.Ceil(value / roundTo)
	return d1 * roundTo
}

// RoundDown rounds down to a given roundTo value.
func RoundDown(value, roundTo float64) float64 {
	d1 := math.Floor(value / roundTo)
	return d1 * roundTo
}

// Normalize returns a set of numbers on the interval [0,1] for a given set of inputs.
// An example: 4,3,2,1 => 0.4, 0.3, 0.2, 0.1
// Caveat; the total may be < 1.0; there are going to be issues with irrational numbers etc.
func Normalize(values ...float64) []float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	output := make([]float64, len(values))
	for x, v := range values {
		output[x] = RoundDown(v/total, 0.001)
	}
	return output
}

// MinInt returns the minimum of a set of integers.
func MinInt(values ...int) int {
	min := math.MaxInt32
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// MaxInt returns the maximum of a set of integers.
func MaxInt(values ...int) int {
	max := math.MinInt32
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// AbsInt returns the absolute value of an integer.
func AbsInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// Seq produces an array of floats from [start,end] by optional steps.
func Seq(start, end float64, steps ...float64) []float64 {
	var values []float64
	step := 1.0
	if len(steps) > 0 {
		step = steps[0]
	}

	if start < end {
		for x := start; x <= end; x += step {
			values = append(values, x)
		}
	} else {
		for x := start; x >= end; x = x - step {
			values = append(values, x)
		}
	}
	return values
}

// SeqRand generates a random sequence.
func SeqRand(samples int, scale float64) []float64 {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	values := make([]float64, samples)

	for x := 0; x < samples; x++ {
		values[x] = rnd.Float64() * scale
	}

	return values
}

// Sum sums a set of values.
func Sum(values ...float64) float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	return total
}

// SumInt sums a set of values.
func SumInt(values ...int) int {
	var total int
	for _, v := range values {
		total += v
	}
	return total
}

// SeqDays generates a sequence of timestamps by day, from -days to today.
func SeqDays(days int) []time.Time {
	var values []time.Time
	for day := days; day >= 0; day-- {
		values = append(values, time.Now().AddDate(0, 0, -day))
	}
	return values
}

// PercentDifference computes the percentage difference between two values.
// The formula is (v2-v1)/v1.
func PercentDifference(v1, v2 float64) float64 {
	return (v2 - v1) / v1
}

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

// DegreesToRadians returns degrees as radians.
func DegreesToRadians(degrees float64) float64 {
	return degrees * _d2r
}

// RadiansToDegrees translates a radian value to a degree value.
func RadiansToDegrees(value float64) float64 {
	return math.Mod(value, _2pi) * _r2d
}

// PercentToRadians converts a normalized value (0,1) to radians.
func PercentToRadians(pct float64) float64 {
	return DegreesToRadians(360.0 * pct)
}

// RadianAdd adds a delta to a base in radians.
func RadianAdd(base, delta float64) float64 {
	value := base + delta
	if value > _2pi {
		return math.Mod(value, _2pi)
	} else if value < 0 {
		return math.Mod(_2pi+value, _2pi)
	}
	return value
}

// DegreesAdd adds a delta to a base in radians.
func DegreesAdd(baseDegrees, deltaDegrees float64) float64 {
	value := baseDegrees + deltaDegrees
	if value > _2pi {
		return math.Mod(value, 360.0)
	} else if value < 0 {
		return math.Mod(360.0+value, 360.0)
	}
	return value
}

// DegreesToCompass returns the degree value in compass / clock orientation.
func DegreesToCompass(deg float64) float64 {
	return DegreesAdd(deg, -90.0)
}

// CirclePoint returns the absolute position of a circle diameter point given
// by the radius and the angle.
func CirclePoint(cx, cy int, radius, angleRadians float64) (x, y int) {
	x = cx + int(radius*math.Sin(angleRadians))
	y = cy - int(radius*math.Cos(angleRadians))
	return
}
