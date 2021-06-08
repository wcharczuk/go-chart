package chart

import (
	"testing"

	"github.com/golang/freetype/truetype"
)

func TestDefaultFont(t *testing.T) {
	inits := make(chan struct{}, 2)
	old := _testingHook
	_testingHook = func() { inits <- struct{}{} }
	defer func() { _testingHook = old }()

	type testData struct {
		font *truetype.Font
		err  error
	}
	calls := make(chan testData, 2)
	for i := 0; i < cap(calls); i++ {
		go func() {
			font, err := GetDefaultFont()
			calls <- testData{font, err}
		}()
	}

	var c []testData
	for i := 0; i < cap(calls); i++ {
		c = append(c, <-calls)
	}

	if c[0].err != nil || c[1].err != nil {
		t.Error("GetDefaultFont got unexpected error: ", c)
	}
	if c[0].font != c[1].font {
		t.Error("GetDefaultFont got different fonts: ", c)
	}

	if len(inits) > 1 { // Safe to check since we already have two results.
		t.Error("GetDefaultFont initialized more than once")
	}
}
