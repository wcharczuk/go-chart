package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestSplitCSV(t *testing.T) {
	assert := assert.New(t)

	assert.Empty(SplitCSV(""))
	assert.Equal([]string{"foo"}, SplitCSV("foo"))
	assert.Equal([]string{"foo", "bar"}, SplitCSV("foo,bar"))
	assert.Equal([]string{"foo", "bar"}, SplitCSV("foo, bar"))
	assert.Equal([]string{"foo", "bar"}, SplitCSV(" foo , bar "))
	assert.Equal([]string{"foo", "bar", "baz"}, SplitCSV("foo,bar,baz"))
	assert.Equal([]string{"foo", "bar", "baz,buzz"}, SplitCSV("foo,bar,\"baz,buzz\""))
	assert.Equal([]string{"foo", "bar", "baz,'buzz'"}, SplitCSV("foo,bar,\"baz,'buzz'\""))
	assert.Equal([]string{"foo", "bar", "baz,'buzz"}, SplitCSV("foo,bar,\"baz,'buzz\""))
	assert.Equal([]string{"foo", "bar", "baz,\"buzz\""}, SplitCSV("foo,bar,'baz,\"buzz\"'"))
}
