package chart

import (
	"math"
	"time"
)

// Range is a type that translates values from a range to a domain.
type Range interface {
	GetMin() interface{}
	GetMax() interface{}
	Translate(value interface{}) int
}

// NewRangeOfFloat64 returns a new Range
func NewRangeOfFloat64(domain int, values ...float64) Range {
	min, max := MinAndMax(values...)
	return &RangeOfFloat64{
		MinValue:    min,
		MaxValue:    max,
		MinMaxDelta: max - min,
		Domain:      domain,
	}
}

// RangeOfFloat64 represents a continuous range
// of float64 values mapped to a [0...WindowMaxValue]
// interval.
type RangeOfFloat64 struct {
	MinValue    float64
	MaxValue    float64
	MinMaxDelta float64
	Domain      int
}

// GetMin implements the interface method.
func (r RangeOfFloat64) GetMin() interface{} {
	return r.MinValue
}

// GetMax implements the interface method.
func (r RangeOfFloat64) GetMax() interface{} {
	return r.MaxValue
}

// Translate maps a given value into the range space.
// An example would be a 600 px image, with a min of 10 and a max of 100.
// Translate(50) would yield (50.0/90.0)*600 ~= 333.33
func (r RangeOfFloat64) Translate(value interface{}) int {
	if typedValue, isTyped := value.(float64); isTyped {
		finalValue := ((r.MaxValue - typedValue) / r.MinMaxDelta) * float64(r.Domain)
		return int(math.Floor(finalValue))
	}
	return 0
}

// NewRangeOfTime makes a new range of time with the given time values.
func NewRangeOfTime(domain int, values ...time.Time) Range {
	min, max := MinAndMaxOfTime(values...)
	r := &RangeOfTime{
		MinValue:    min,
		MaxValue:    max,
		MinMaxDelta: max.Unix() - min.Unix(),
		Domain:      domain,
	}
	return r
}

// RangeOfTime represents a timeseries.
type RangeOfTime struct {
	MinValue    time.Time
	MaxValue    time.Time
	MinMaxDelta int64 //unix time difference
	Domain      int
}

// GetMin implements the interface method.
func (r RangeOfTime) GetMin() interface{} {
	return r.MinValue
}

// GetMax implements the interface method.
func (r RangeOfTime) GetMax() interface{} {
	return r.MaxValue
}

// Translate maps a given value into the range space (of time).
// An example would be a 600 px image, with a min of jan-01-2016 and a max of jun-01-2016.
// Translate(may-01-2016) would yield ... something.
func (r RangeOfTime) Translate(value interface{}) int {
	if typed, isTyped := value.(time.Time); isTyped {
		valueDelta := r.MaxValue.Unix() - typed.Unix()
		finalValue := (float64(valueDelta) / float64(r.MinMaxDelta)) * float64(r.Domain)
		return int(math.Floor(finalValue))
	}
	return 0
}
