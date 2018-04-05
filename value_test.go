package chart

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

func TestValuesValues(t *testing.T) {
	assert := assert.New(t)

	vs := []Value{
		{Value: 10, Label: "Blue"},
		{Value: 9, Label: "Green"},
		{Value: 8, Label: "Gray"},
		{Value: 7, Label: "Orange"},
		{Value: 6, Label: "HEANG"},
		{Value: 5, Label: "??"},
		{Value: 2, Label: "!!"},
	}

	values := Values(vs).Values()
	assert.Len(7, values)
	assert.Equal(10, values[0])
	assert.Equal(9, values[1])
	assert.Equal(8, values[2])
	assert.Equal(7, values[3])
	assert.Equal(6, values[4])
	assert.Equal(5, values[5])
	assert.Equal(2, values[6])
}

func TestValuesValuesNormalized(t *testing.T) {
	assert := assert.New(t)

	vs := []Value{
		{Value: 10, Label: "Blue"},
		{Value: 9, Label: "Green"},
		{Value: 8, Label: "Gray"},
		{Value: 7, Label: "Orange"},
		{Value: 6, Label: "HEANG"},
		{Value: 5, Label: "??"},
		{Value: 2, Label: "!!"},
	}

	values := Values(vs).ValuesNormalized()
	assert.Len(7, values)
	assert.Equal(0.2127, values[0])
	assert.Equal(0.0425, values[6])
}

func TestValuesNormalize(t *testing.T) {
	assert := assert.New(t)

	vs := []Value{
		{Value: 10, Label: "Blue"},
		{Value: 9, Label: "Green"},
		{Value: 8, Label: "Gray"},
		{Value: 7, Label: "Orange"},
		{Value: 6, Label: "HEANG"},
		{Value: 5, Label: "??"},
		{Value: 2, Label: "!!"},
	}

	values := Values(vs).Normalize()
	assert.Len(7, values)
	assert.Equal(0.2127, values[0].Value)
	assert.Equal(0.0425, values[6].Value)
}
