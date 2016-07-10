package chart

import (
	"fmt"
	"time"
)

// ValueFormatter is a function that takes a value and produces a string.
type ValueFormatter func(v interface{}) string

// TimeValueFormatter is a ValueFormatter for timestamps.
func TimeValueFormatter(v interface{}) string {
	return TimeValueFormatterWithFormat(v, DefaultDateFormat)
}

// TimeValueFormatterWithFormat is a ValueFormatter for timestamps with a given format.
func TimeValueFormatterWithFormat(v interface{}, dateFormat string) string {
	if typed, isTyped := v.(time.Time); isTyped {
		return typed.Format(dateFormat)
	}
	if typed, isTyped := v.(int64); isTyped {
		return time.Unix(typed, 0).Format(dateFormat)
	}
	if typed, isTyped := v.(float64); isTyped {
		return time.Unix(int64(typed), 0).Format(dateFormat)
	}
	return ""
}

// FloatValueFormatter is a ValueFormatter for float64.
func FloatValueFormatter(v interface{}) string {
	return FloatValueFormatterWithFormat(v, "%.2f")
}

// FloatValueFormatterWithFormat is a ValueFormatter for float64 with a given format.
func FloatValueFormatterWithFormat(v interface{}, floatFormat string) string {
	if typed, isTyped := v.(float64); isTyped {
		return fmt.Sprintf(floatFormat, typed)
	}
	return ""
}
