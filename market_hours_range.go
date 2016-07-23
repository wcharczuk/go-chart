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
	return TimeToFloat64(mhr.Max)
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
	min := TimeToFloat64(mhr.Min)
	max := TimeToFloat64(mhr.Max)
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

func (mhr MarketHoursRange) String() string {
	return fmt.Sprintf("MarketHoursRange [%s, %s] => %d", mhr.Min.Format(DefaultDateFormat), mhr.Max.Format(DefaultDateFormat), mhr.Domain)
}

// Translate maps a given value into the ContinuousRange space.
func (mhr MarketHoursRange) Translate(value float64) int {
	valueTime := Float64ToTime(value)
	deltaSeconds := date.CalculateMarketSecondsBetween(mhr.Min, mhr.Max, mhr.MarketOpen, mhr.MarketClose, mhr.HolidayProvider)
	valueDelta := date.CalculateMarketSecondsBetween(mhr.Min, valueTime, mhr.MarketOpen, mhr.MarketClose, mhr.HolidayProvider)

	translated := int((float64(valueDelta) / float64(deltaSeconds)) * float64(mhr.Domain))
	fmt.Printf("nyse translating: %s to %d ~= %d", valueTime.Format(time.RFC3339), deltaSeconds, valueDelta)
	return translated
}
