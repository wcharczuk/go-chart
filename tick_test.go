package chart

import (
	"math"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestGenerateContinuousTicks(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:    0.0,
		Max:    10.0,
		Domain: 256,
	}

	vf := FloatValueFormatter

	ticks := GenerateContinuousTicks(r, ra, false, Style{}, vf)
	assert.NotEmpty(ticks)
	assert.Len(ticks, 11)
	assert.Equal(0.0, ticks[0].Value)
	assert.Equal(10, ticks[len(ticks)-1].Value)
}

func TestGenerateContinuousTicksDescending(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:        0.0,
		Max:        10.0,
		Domain:     256,
		Descending: true,
	}

	vf := FloatValueFormatter

	ticks := GenerateContinuousTicks(r, ra, false, Style{}, vf)
	assert.NotEmpty(ticks)
	assert.Len(ticks, 11)
	assert.Equal(10.0, ticks[0].Value)
	assert.Equal(9.0, ticks[1].Value)
	assert.Equal(1.0, ticks[len(ticks)-2].Value)
	assert.Equal(0.0, ticks[len(ticks)-1].Value)
}

func TestGenerateContinuousPrettyTicks(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:    37.5,
		Max:    60.1,
		Domain: 256,
	}

	vf := FloatValueFormatter

	ticks := GeneratePrettyContinuousTicks(r, ra, false, Style{}, vf)
	assert.NotEmpty(ticks)
	assert.Equal(ticks, []Tick{
		{Label: "38.00", Value: 38},
		{Label: "40.00", Value: 40},
		{Label: "42.00", Value: 42},
		{Label: "44.00", Value: 44},
		{Label: "46.00", Value: 46},
		{Label: "48.00", Value: 48},
		{Label: "50.00", Value: 50},
		{Label: "52.00", Value: 52},
		{Label: "54.00", Value: 54},
		{Label: "56.00", Value: 56},
		{Label: "58.00", Value: 58},
		{Label: "60.00", Value: 60}})
}

func TestGeneratePrettyTicksForVerySmallRange(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:    1e-100,
		Max:    1e-99,
		Domain: 256,
	}

	vf := FloatValueFormatter

	ticks := GeneratePrettyContinuousTicks(r, ra, false, Style{}, vf)
	assert.NotEmpty(ticks)
	assert.Len(ticks, 9)
}

func TestGeneratePrettyTicksForVeryLargeRange(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:    1e-100,
		Max:    1e+100,
		Domain: 256,
	}

	vf := FloatValueFormatter

	ticks := GeneratePrettyContinuousTicks(r, ra, false, Style{}, vf)
	assert.NotEmpty(ticks)
	assert.Len(ticks, 10)
}

func TestGeneratePrettyTicksForVerySmallDomain(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:    0.0,
		Max:    10.0,
		Domain: 1,
	}

	vf := FloatValueFormatter

	ticks := GeneratePrettyContinuousTicks(r, ra, false, Style{}, vf)
	assert.Empty(ticks)
}

func TestGeneratePrettyTicksForVeryLargeDomain(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:    0.0,
		Max:    10.0,
		Domain: math.MaxInt32,
	}

	vf := FloatValueFormatter

	ticks := GeneratePrettyContinuousTicks(r, ra, false, Style{}, vf)
	assert.NotEmpty(ticks)
	assert.Len(ticks, 1001)
}
