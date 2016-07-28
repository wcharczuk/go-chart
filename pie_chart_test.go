package chart

import (
	"bytes"
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestPieChart(t *testing.T) {
	assert := assert.New(t)

	pie := PieChart{
		Canvas: Style{
			FillColor: ColorLightGray,
		},
		Values: []Value{
			{Value: 10, Label: "Blue"},
			{Value: 9, Label: "Green"},
			{Value: 8, Label: "Gray"},
			{Value: 7, Label: "Orange"},
			{Value: 6, Label: "HEANG"},
			{Value: 5, Label: "??"},
			{Value: 2, Label: "!!"},
		},
	}

	b := bytes.NewBuffer([]byte{})
	pie.Render(PNG, b)
	assert.NotZero(b.Len())
}
