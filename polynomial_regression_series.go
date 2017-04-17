package chart

import "fmt"

// PolynomialRegressionSeries implements a polynomial regression over a given
// inner series.
type PolynomialRegressionSeries struct {
	Name  string
	Style Style
	YAxis YAxisType

	Limit       int
	Offset      int
	Order       int
	InnerSeries ValueProvider

	coeffs []float64
}

// GetName returns the name of the time series.
func (prs PolynomialRegressionSeries) GetName() string {
	return prs.Name
}

// GetStyle returns the line style.
func (prs PolynomialRegressionSeries) GetStyle() Style {
	return prs.Style
}

// GetYAxis returns which YAxis the series draws on.
func (prs PolynomialRegressionSeries) GetYAxis() YAxisType {
	return prs.YAxis
}

// Len returns the number of elements in the series.
func (prs PolynomialRegressionSeries) Len() int {
	return Math.MinInt(prs.GetLimit(), prs.InnerSeries.Len()-prs.GetOffset())
}

// GetLimit returns the window size.
func (prs PolynomialRegressionSeries) GetLimit() int {
	if prs.Limit == 0 {
		return prs.InnerSeries.Len()
	}
	return prs.Limit
}

// GetEndIndex returns the effective limit end.
func (prs PolynomialRegressionSeries) GetEndIndex() int {
	offset := prs.GetOffset() + prs.Len()
	innerSeriesLastIndex := prs.InnerSeries.Len() - 1
	return Math.MinInt(offset, innerSeriesLastIndex)
}

// GetOffset returns the data offset.
func (prs PolynomialRegressionSeries) GetOffset() int {
	if prs.Offset == 0 {
		return 0
	}
	return prs.Offset
}

// Validate validates the series.
func (prs *PolynomialRegressionSeries) Validate() error {
	if prs.InnerSeries == nil {
		return fmt.Errorf("linear regression series requires InnerSeries to be set")
	}
	return nil
}

// GetValue returns the series value for a given index.
func (prs *PolynomialRegressionSeries) GetValue(index int) (x, y float64) {
	if prs.InnerSeries == nil || prs.InnerSeries.Len() == 0 {
		return
	}
	return
}

func (prs *PolynomialRegressionSeries) computeCoefficients() {
	vandMatrix := make([][]float64, prs.Len(), prs.Order+1)
	var xvalue float64
	for i := 0; i < prs.Len(); i++ {
		_, xvalue = prs.InnerSeries.GetValue(i)
		var mult float64 = 1.0
		for j := 0; j < prs.Order+1; j++ {
			vandMatrix[i][j] = mult
			mult = mult * xvalue
		}
	}
}
