package chart

import (
	"testing"
	"time"

	"github.com/blendlabs/go-assert"
)

func TestTimeValueFormatterWithFormat(t *testing.T) {
	assert := assert.New(t)

	d := time.Now()
	di := Time.ToFloat64(d)
	df := float64(di)

	s := TimeValueFormatterWithFormat(d, DefaultDateFormat)
	si := TimeValueFormatterWithFormat(di, DefaultDateFormat)
	sf := TimeValueFormatterWithFormat(df, DefaultDateFormat)
	assert.Equal(s, si)
	assert.Equal(s, sf)

	sd := TimeValueFormatter(d)
	sdi := TimeValueFormatter(di)
	sdf := TimeValueFormatter(df)
	assert.Equal(s, sd)
	assert.Equal(s, sdi)
	assert.Equal(s, sdf)
}

func TestFloatValueFormatter(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1234.00", FloatValueFormatter(1234.00))
}

func TestFloatValueFormatterWithFloat32Input(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1234.00", FloatValueFormatter(float32(1234.00)))
}

func TestFloatValueFormatterWithIntegerInput(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1234.00", FloatValueFormatter(1234))
}

func TestFloatValueFormatterWithInt64Input(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1234.00", FloatValueFormatter(int64(1234)))
}

func TestFloatValueFormatterWithFormat(t *testing.T) {
	assert := assert.New(t)

	v := 123.456
	sv := FloatValueFormatterWithFormat(v, "%.3f")
	assert.Equal("123.456", sv)
	assert.Equal("123.000", FloatValueFormatterWithFormat(123, "%.3f"))
}
