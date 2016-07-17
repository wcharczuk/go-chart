package chart

import (
	"testing"
	"time"

	"github.com/blendlabs/go-assert"
)

func TestTimeValueFormatterWithFormat(t *testing.T) {
	assert := assert.New(t)

	d := time.Now()
	di := TimeToFloat64(d)
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

func TestFloatValueFormatterWithFormat(t *testing.T) {
	assert := assert.New(t)

	v := 123.456
	sv := FloatValueFormatterWithFormat(v, "%.3f")
	assert.Equal("123.456", sv)
	assert.Equal("", FloatValueFormatterWithFormat(123, "%.3f"))
}
