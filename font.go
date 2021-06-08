package chart

import (
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2/roboto"
)

var _defaultFont defaultFont

// GetDefaultFont returns the default font (Roboto-Medium).
func GetDefaultFont() (*truetype.Font, error) {
	return _defaultFont.Font()
}

type defaultFont struct {
	font *truetype.Font
	err  error
	once sync.Once
}

var _testingHook = func() {}

func (df *defaultFont) Font() (*truetype.Font, error) {
	df.once.Do(func() {
		df.font, df.err = truetype.Parse(roboto.Roboto)
		_testingHook()
	})
	return df.font, df.err
}
