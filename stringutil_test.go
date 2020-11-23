package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestSplitCSV(t *testing.T) {
	// replaced new assertions helper

	testutil.AssertEmpty(t, SplitCSV(""))
	testutil.AssertEqual(t, []string{"foo"}, SplitCSV("foo"))
	testutil.AssertEqual(t, []string{"foo", "bar"}, SplitCSV("foo,bar"))
	testutil.AssertEqual(t, []string{"foo", "bar"}, SplitCSV("foo, bar"))
	testutil.AssertEqual(t, []string{"foo", "bar"}, SplitCSV(" foo , bar "))
	testutil.AssertEqual(t, []string{"foo", "bar", "baz"}, SplitCSV("foo,bar,baz"))
	testutil.AssertEqual(t, []string{"foo", "bar", "baz,buzz"}, SplitCSV("foo,bar,\"baz,buzz\""))
	testutil.AssertEqual(t, []string{"foo", "bar", "baz,'buzz'"}, SplitCSV("foo,bar,\"baz,'buzz'\""))
	testutil.AssertEqual(t, []string{"foo", "bar", "baz,'buzz"}, SplitCSV("foo,bar,\"baz,'buzz\""))
	testutil.AssertEqual(t, []string{"foo", "bar", "baz,\"buzz\""}, SplitCSV("foo,bar,'baz,\"buzz\"'"))
}
