package chart

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/drawing"
)

func TestVectorRendererPath(t *testing.T) {
	assert := assert.New(t)

	vr, err := SVG(100, 100)
	assert.Nil(err)

	typed, isTyped := vr.(*vectorRenderer)
	assert.True(isTyped)

	typed.MoveTo(0, 0)
	typed.LineTo(100, 100)
	typed.LineTo(0, 100)
	typed.Close()
	typed.FillStroke()

	buffer := bytes.NewBuffer([]byte{})
	err = typed.Save(buffer)
	assert.Nil(err)

	raw := string(buffer.Bytes())

	assert.True(strings.HasPrefix(raw, "<svg"))
	assert.True(strings.HasSuffix(raw, "</svg>"))
}

func TestVectorRendererMeasureText(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	vr, err := SVG(100, 100)
	assert.Nil(err)

	vr.SetDPI(DefaultDPI)
	vr.SetFont(f)
	vr.SetFontSize(12.0)

	tb := vr.MeasureText("Ljp")
	assert.Equal(21, tb.Width())
	assert.Equal(15, tb.Height())
}

func TestCanvasStyleSVG(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Font:        f,
		Padding:     DefaultBackgroundPadding,
	}

	canvas := &canvas{dpi: DefaultDPI}

	svgString := canvas.styleAsSVG(set)
	assert.NotEmpty(svgString)
	assert.True(strings.HasPrefix(svgString, "style=\""))
	assert.True(strings.Contains(svgString, "stroke:rgba(255,255,255,1.0)"))
	assert.True(strings.Contains(svgString, "stroke-width:5"))
	assert.True(strings.Contains(svgString, "fill:rgba(255,255,255,1.0)"))
	assert.True(strings.HasSuffix(svgString, "\""))
}

func TestCanvasClassSVG(t *testing.T) {
	as := assert.New(t)

	set := Style{
		ClassName: "test-class",
	}

	canvas := &canvas{dpi: DefaultDPI}

	as.Equal("class=\"test-class\"", canvas.styleAsSVG(set))
}

func TestCanvasCustomInlineStylesheet(t *testing.T) {
	b := strings.Builder{}

	canvas := &canvas{
		w:     &b,
		css:   ".background { fill: red }",
	}

	canvas.Start(200, 200)

	assert.New(t).Contains(b.String(), fmt.Sprintf(`<style type="text/css"><![CDATA[%s]]></style>`, canvas.css))
}

func TestCanvasCustomInlineStylesheetWithNonce(t *testing.T) {
	b := strings.Builder{}

	canvas := &canvas{
		w:     &b,
		css:   ".background { fill: red }",
		nonce: "RAND0MSTRING",
	}

	canvas.Start(200, 200)

	assert.New(t).Contains(b.String(), fmt.Sprintf(`<style type="text/css" nonce="%s"><![CDATA[%s]]></style>`, canvas.nonce, canvas.css))
}