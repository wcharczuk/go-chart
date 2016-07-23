package chart

import (
	"fmt"
	"time"

	"github.com/wcharczuk/go-chart/date"
)

// NYSEMarketHoursRange is a special type of range that compresses a time range into just the
// market (i.e. NYSE operating hours and days) range.
type NYSEMarketHoursRange struct {
	Min    time.Time
	Max    time.Time
	Domain int
}

// IsZero returns if the range is setup or not.
func (mhr NYSEMarketHoursRange) IsZero() bool {
	return mhr.Min.IsZero() && mhr.Max.IsZero()
}

// GetMin returns the min value.
func (mhr NYSEMarketHoursRange) GetMin() float64 {
	return TimeToFloat64(mhr.Min)
}

// GetMax returns the max value.
func (mhr NYSEMarketHoursRange) GetMax() float64 {
	return TimeToFloat64(mhr.Max)
}

// SetMin sets the min value.
func (mhr *NYSEMarketHoursRange) SetMin(min float64) {
	mhr.Min = Float64ToTime(min)
}

// SetMax sets the max value.
func (mhr *NYSEMarketHoursRange) SetMax(max float64) {
	mhr.Max = Float64ToTime(max)
}

// GetDelta gets the delta.
func (mhr NYSEMarketHoursRange) GetDelta() float64 {
	min := TimeToFloat64(mhr.Min)
	max := TimeToFloat64(mhr.Max)
	return max - min
}

// GetDomain gets the domain.
func (mhr NYSEMarketHoursRange) GetDomain() int {
	return mhr.Domain
}

// SetDomain sets the domain.
func (mhr *NYSEMarketHoursRange) SetDomain(domain int) {
	mhr.Domain = domain
}

func (mhr NYSEMarketHoursRange) String() string {
	return fmt.Sprintf("MarketHoursRange [%s, %s] => %d", mhr.Min.Format(DefaultDateFormat), mhr.Max.Format(DefaultDateFormat), mhr.Domain)
}

// Translate maps a given value into the ContinuousRange space.
func (mhr NYSEMarketHoursRange) Translate(value float64) int {
	valueTime := Float64ToTime(value)
	deltaSeconds := date.CalculateMarketSecondsBetween(mhr.Min, mhr.Max)
	valueDelta := date.CalculateMarketSecondsBetween(mhr.Min, valueTime)

	translated := int((float64(valueDelta) / float64(deltaSeconds)) * float64(mhr.Domain))
	fmt.Printf("nyse translating: %s to %d ~= %d", valueTime.Format(time.RFC3339), deltaSeconds, valueDelta)
	return translated
}
