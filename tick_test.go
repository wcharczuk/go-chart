package chart

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
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
	assert.Len(11, ticks)
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
	assert.Len(11, ticks)
	assert.Equal(10.0, ticks[0].Value)
	assert.Equal(9.0, ticks[1].Value)
	assert.Equal(1.0, ticks[len(ticks)-2].Value)
	assert.Equal(0.0, ticks[len(ticks)-1].Value)
}
