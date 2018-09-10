package util

import (
	"testing"
	"time"

	"github.com/blend/go-sdk/assert"
)

func TestTimeDiffDays(t *testing.T) {
	assert := assert.New(t)

	t1 := time.Date(2017, 02, 27, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2017, 01, 10, 3, 0, 0, 0, time.UTC)
	t3 := time.Date(2017, 02, 24, 16, 0, 0, 0, time.UTC)

	assert.Equal(48, Time.DiffDays(t2, t1))
	assert.Equal(2, Time.DiffDays(t3, t1)) // technically we should round down.
}

func TestTimeDiffHours(t *testing.T) {
	assert := assert.New(t)

	t1 := time.Date(2017, 02, 27, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2017, 02, 24, 16, 0, 0, 0, time.UTC)
	t3 := time.Date(2017, 02, 28, 12, 0, 0, 0, time.UTC)

	assert.Equal(68, Time.DiffHours(t2, t1))
	assert.Equal(24, Time.DiffHours(t1, t3))
}
