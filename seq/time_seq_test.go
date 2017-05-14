package seq

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
)

func TestTimeSeqTimes(t *testing.T) {
	assert := assert.New(t)

	seq := Times(time.Now(), time.Now(), time.Now())
	assert.Equal(3, seq.Len())
}

func parseTime(str string) time.Time {
	tv, _ := time.Parse("2006-01-02 15:04:05", str)
	return tv
}

func TestTimeSeqSort(t *testing.T) {
	assert := assert.New(t)

	seq := Times(
		parseTime("2016-05-14 12:00:00"),
		parseTime("2017-05-14 12:00:00"),
		parseTime("2015-05-14 12:00:00"),
		parseTime("2017-05-13 12:00:00"),
	)

	sorted := seq.Sort()
	assert.Equal(4, sorted.Len())
	min, max := sorted.MinAndMax()
	assert.Equal(parseTime("2015-05-14 12:00:00"), min)
	assert.Equal(parseTime("2017-05-14 12:00:00"), max)

	first, last := sorted.First(), sorted.Last()
	assert.Equal(min, first)
	assert.Equal(max, last)
}

func TestTimeSeqSortDescending(t *testing.T) {
	assert := assert.New(t)

	seq := Times(
		parseTime("2016-05-14 12:00:00"),
		parseTime("2017-05-14 12:00:00"),
		parseTime("2015-05-14 12:00:00"),
		parseTime("2017-05-13 12:00:00"),
	)

	sorted := seq.SortDescending()
	assert.Equal(4, sorted.Len())
	min, max := sorted.MinAndMax()
	assert.Equal(parseTime("2015-05-14 12:00:00"), min)
	assert.Equal(parseTime("2017-05-14 12:00:00"), max)

	first, last := sorted.First(), sorted.Last()
	assert.Equal(max, first)
	assert.Equal(min, last)
}

func TestTimeSeqDays(t *testing.T) {
	assert := assert.New(t)

	seq := Times(
		parseTime("2017-05-10 12:00:00"),
		parseTime("2017-05-10 16:00:00"),
		parseTime("2017-05-11 12:00:00"),
		parseTime("2015-05-12 12:00:00"),
		parseTime("2015-05-12 16:00:00"),
		parseTime("2017-05-13 12:00:00"),
		parseTime("2017-05-14 12:00:00"),
	)

	days := seq.Days()
	assert.Equal(5, days.Len())
	assert.Equal(10, days.First().Day())
	assert.Equal(14, days.Last().Day())
}
