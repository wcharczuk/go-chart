package util

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
)

func TestTimeFromFloat64(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()

	assert.InTimeDelta(now, Time.FromFloat64(Time.ToFloat64(now)), time.Microsecond)
}
