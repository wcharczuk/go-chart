package chart

import "fmt"

// LastValueAnnotation returns an annotation series of just the last value of a value provider.
func LastValueAnnotation(innerSeries ValueProvider, vfs ...ValueFormatter) AnnotationSeries {
	var vf ValueFormatter
	if len(vfs) > 0 {
		vf = vfs[0]
	} else if typed, isTyped := innerSeries.(ValueFormatterProvider); isTyped {
		_, vf = typed.GetValueFormatters()
	} else {
		vf = FloatValueFormatter
	}

	var lastValue Value2
	if typed, isTyped := innerSeries.(LastValueProvider); isTyped {
		lastValue.XValue, lastValue.YValue = typed.GetLastValue()
		lastValue.Label = vf(lastValue.YValue)
	} else {
		lastValue.XValue, lastValue.YValue = innerSeries.GetValue(innerSeries.Len() - 1)
		lastValue.Label = vf(lastValue.YValue)
	}

	var seriesName string
	var seriesStyle Style
	if typed, isTyped := innerSeries.(Series); isTyped {
		seriesName = fmt.Sprintf("%s - Last Value", typed.GetName())
		seriesStyle = typed.GetStyle()
	}

	return AnnotationSeries{
		Name:        seriesName,
		Style:       seriesStyle,
		Annotations: []Value2{lastValue},
	}
}
