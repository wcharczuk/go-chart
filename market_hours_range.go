package chart

import (
	"fmt"
	"time"

	"github.com/wcharczuk/go-chart/seq"
	"github.com/wcharczuk/go-chart/util"
)

// MarketHoursRange is a special type of range that compresses a time range into just the
// market (i.e. NYSE operating hours and days) range.
type MarketHoursRange struct {
	Min time.Time
	Max time.Time

	MarketOpen  time.Time
	MarketClose time.Time

	HolidayProvider util.HolidayProvider

	ValueFormatter ValueFormatter

	Descending bool
	Domain     int
}

// IsDescending returns if the range is descending.
func (mhr MarketHoursRange) IsDescending() bool {
	return mhr.Descending
}

// GetTimezone returns the timezone for the market hours range.
func (mhr MarketHoursRange) GetTimezone() *time.Location {
	return mhr.GetMarketOpen().Location()
}

// IsZero returns if the range is setup or not.
func (mhr MarketHoursRange) IsZero() bool {
	return mhr.Min.IsZero() && mhr.Max.IsZero()
}

// GetMin returns the min value.
func (mhr MarketHoursRange) GetMin() float64 {
	return util.Time.ToFloat64(mhr.Min)
}

// GetMax returns the max value.
func (mhr MarketHoursRange) GetMax() float64 {
	return util.Time.ToFloat64(mhr.GetEffectiveMax())
}

// GetEffectiveMax gets either the close on the max, or the max itself.
func (mhr MarketHoursRange) GetEffectiveMax() time.Time {
	maxClose := util.Date.On(mhr.MarketClose, mhr.Max)
	if maxClose.After(mhr.Max) {
		return maxClose
	}
	return mhr.Max
}

// SetMin sets the min value.
func (mhr *MarketHoursRange) SetMin(min float64) {
	mhr.Min = util.Time.FromFloat64(min)
	mhr.Min = mhr.Min.In(mhr.GetTimezone())
}

// SetMax sets the max value.
func (mhr *MarketHoursRange) SetMax(max float64) {
	mhr.Max = util.Time.FromFloat64(max)
	mhr.Max = mhr.Max.In(mhr.GetTimezone())
}

// GetDelta gets the delta.
func (mhr MarketHoursRange) GetDelta() float64 {
	min := mhr.GetMin()
	max := mhr.GetMax()
	return max - min
}

// GetDomain gets the domain.
func (mhr MarketHoursRange) GetDomain() int {
	return mhr.Domain
}

// SetDomain sets the domain.
func (mhr *MarketHoursRange) SetDomain(domain int) {
	mhr.Domain = domain
}

// GetHolidayProvider coalesces a userprovided holiday provider and the date.DefaultHolidayProvider.
func (mhr MarketHoursRange) GetHolidayProvider() util.HolidayProvider {
	if mhr.HolidayProvider == nil {
		return func(_ time.Time) bool { return false }
	}
	return mhr.HolidayProvider
}

// GetMarketOpen returns the market open time.
func (mhr MarketHoursRange) GetMarketOpen() time.Time {
	if mhr.MarketOpen.IsZero() {
		return util.NYSEOpen()
	}
	return mhr.MarketOpen
}

// GetMarketClose returns the market close time.
func (mhr MarketHoursRange) GetMarketClose() time.Time {
	if mhr.MarketClose.IsZero() {
		return util.NYSEClose()
	}
	return mhr.MarketClose
}

// GetTicks returns the ticks for the range.
// This is to override the default continous ticks that would be generated for the range.
func (mhr *MarketHoursRange) GetTicks(r Renderer, defaults Style, vf ValueFormatter) []Tick {
	times := seq.Time.MarketHours(mhr.Min, mhr.Max, mhr.GetMarketOpen(), mhr.GetMarketClose(), mhr.GetHolidayProvider())
	timesWidth := mhr.measureTimes(r, defaults, vf, times)
	if timesWidth <= mhr.Domain {
		return mhr.makeTicks(vf, times)
	}

	times = seq.Time.MarketHourQuarters(mhr.Min, mhr.Max, mhr.GetMarketOpen(), mhr.GetMarketClose(), mhr.GetHolidayProvider())
	timesWidth = mhr.measureTimes(r, defaults, vf, times)
	if timesWidth <= mhr.Domain {
		return mhr.makeTicks(vf, times)
	}

	times = seq.Time.MarketDayCloses(mhr.Min, mhr.Max, mhr.GetMarketOpen(), mhr.GetMarketClose(), mhr.GetHolidayProvider())
	timesWidth = mhr.measureTimes(r, defaults, vf, times)
	if timesWidth <= mhr.Domain {
		return mhr.makeTicks(vf, times)
	}

	times = seq.Time.MarketDayAlternateCloses(mhr.Min, mhr.Max, mhr.GetMarketOpen(), mhr.GetMarketClose(), mhr.GetHolidayProvider())
	timesWidth = mhr.measureTimes(r, defaults, vf, times)
	if timesWidth <= mhr.Domain {
		return mhr.makeTicks(vf, times)
	}

	times = seq.Time.MarketDayMondayCloses(mhr.Min, mhr.Max, mhr.GetMarketOpen(), mhr.GetMarketClose(), mhr.GetHolidayProvider())
	timesWidth = mhr.measureTimes(r, defaults, vf, times)
	if timesWidth <= mhr.Domain {
		return mhr.makeTicks(vf, times)
	}

	return GenerateContinuousTicks(r, mhr, false, defaults, vf)

}

func (mhr *MarketHoursRange) measureTimes(r Renderer, defaults Style, vf ValueFormatter, times []time.Time) int {
	defaults.GetTextOptions().WriteToRenderer(r)
	var total int
	for index, t := range times {
		timeLabel := vf(t)

		labelBox := r.MeasureText(timeLabel)
		total += labelBox.Width()
		if index > 0 {
			total += DefaultMinimumTickHorizontalSpacing
		}
	}
	return total
}

func (mhr *MarketHoursRange) makeTicks(vf ValueFormatter, times []time.Time) []Tick {
	ticks := make([]Tick, len(times))
	for index, t := range times {
		ticks[index] = Tick{
			Value: util.Time.ToFloat64(t),
			Label: vf(t),
		}
	}
	return ticks
}

func (mhr MarketHoursRange) String() string {
	return fmt.Sprintf("MarketHoursRange [%s, %s] => %d", mhr.Min.Format(time.RFC3339), mhr.Max.Format(time.RFC3339), mhr.Domain)
}

// Translate maps a given value into the ContinuousRange space.
func (mhr MarketHoursRange) Translate(value float64) int {
	valueTime := util.Time.FromFloat64(value)
	valueTimeEastern := valueTime.In(util.Date.Eastern())
	totalSeconds := util.Date.CalculateMarketSecondsBetween(mhr.Min, mhr.GetEffectiveMax(), mhr.GetMarketOpen(), mhr.GetMarketClose(), mhr.HolidayProvider)
	valueDelta := util.Date.CalculateMarketSecondsBetween(mhr.Min, valueTimeEastern, mhr.GetMarketOpen(), mhr.GetMarketClose(), mhr.HolidayProvider)
	translated := int((float64(valueDelta) / float64(totalSeconds)) * float64(mhr.Domain))

	if mhr.IsDescending() {
		return mhr.Domain - translated
	}

	return translated
}
