package chart

import (
	"fmt"
	"time"

	"github.com/wcharczuk/go-chart/date"
)

// MarketHoursRange is a special type of range that compresses a time range into just the
// market (i.e. NYSE operating hours and days) range.
type MarketHoursRange struct {
	Min time.Time
	Max time.Time

	MarketOpen  time.Time
	MarketClose time.Time

	HolidayProvider date.HolidayProvider

	ValueFormatter ValueFormatter

	Domain int
}

// IsZero returns if the range is setup or not.
func (mhr MarketHoursRange) IsZero() bool {
	return mhr.Min.IsZero() && mhr.Max.IsZero()
}

// GetMin returns the min value.
func (mhr MarketHoursRange) GetMin() float64 {
	return TimeToFloat64(mhr.Min)
}

// GetMax returns the max value.
func (mhr MarketHoursRange) GetMax() float64 {
	return TimeToFloat64(mhr.GetEffectiveMax())
}

// GetEffectiveMax gets either the close on the max, or the max itself.
func (mhr MarketHoursRange) GetEffectiveMax() time.Time {
	maxClose := date.On(mhr.MarketClose, mhr.Max)
	if maxClose.After(mhr.Max) {
		return maxClose
	}
	return mhr.Max
}

// SetMin sets the min value.
func (mhr *MarketHoursRange) SetMin(min float64) {
	mhr.Min = Float64ToTime(min)
}

// SetMax sets the max value.
func (mhr *MarketHoursRange) SetMax(max float64) {
	mhr.Max = Float64ToTime(max)
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
func (mhr MarketHoursRange) GetHolidayProvider() date.HolidayProvider {
	if mhr.HolidayProvider == nil {
		return date.DefaultHolidayProvider
	}
	return mhr.HolidayProvider
}

// GetTicks returns the ticks for the range.
// This is to override the default continous ticks that would be generated for the range.
func (mhr *MarketHoursRange) GetTicks(vf ValueFormatter) []Tick {
	var ticks []Tick

	cursor := date.On(mhr.MarketClose, mhr.Min)
	maxClose := date.On(mhr.MarketClose, mhr.Max)

	for date.BeforeDate(cursor, maxClose) {
		if date.IsWeekDay(cursor.Weekday()) && !mhr.GetHolidayProvider()(cursor) {
			ticks = append(ticks, Tick{
				Value: TimeToFloat64(cursor),
				Label: vf(cursor),
			})
		}

		cursor = cursor.AddDate(0, 0, 1)
	}

	endMarketClose := date.On(mhr.MarketClose, cursor)
	if date.IsWeekDay(endMarketClose.Weekday()) && !mhr.GetHolidayProvider()(endMarketClose) {
		ticks = append(ticks, Tick{
			Value: TimeToFloat64(endMarketClose),
			Label: vf(endMarketClose),
		})
	}

	return ticks
}

func (mhr MarketHoursRange) String() string {
	return fmt.Sprintf("MarketHoursRange [%s, %s] => %d", mhr.Min.Format(DefaultDateMinuteFormat), mhr.Max.Format(DefaultDateMinuteFormat), mhr.Domain)
}

// Translate maps a given value into the ContinuousRange space.
func (mhr MarketHoursRange) Translate(value float64) int {
	valueTime := Float64ToTime(value)
	valueTimeEastern := valueTime.In(date.Eastern())
	totalSeconds := date.CalculateMarketSecondsBetween(mhr.Min, mhr.GetEffectiveMax(), mhr.MarketOpen, mhr.MarketClose, mhr.HolidayProvider)
	valueDelta := date.CalculateMarketSecondsBetween(mhr.Min, valueTimeEastern, mhr.MarketOpen, mhr.MarketClose, mhr.HolidayProvider)
	translated := int((float64(valueDelta) / float64(totalSeconds)) * float64(mhr.Domain))
	return translated
}
