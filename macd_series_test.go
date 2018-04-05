package chart

import (
	"fmt"
	"testing"

	"github.com/blend/go-sdk/assert"
)

var (
	macdExpected = []float64{
		0,
		0.06381766382,
		0.1641441222,
		0.2817201894,
		0.4033023481,
		0.3924673744,
		0.2983093823,
		0.1561821464,
		-0.008916708129,
		-0.05210332292,
		-0.01649503993,
		0.06667130899,
		0.1751344574,
		0.1657328378,
		0.08257097469,
		-0.04265109369,
		-0.1875741257,
		-0.2091853882,
		-0.1518975486,
		-0.04781419838,
		0.08025242841,
		0.08881960494,
		0.02183529775,
		-0.08904155476,
		-0.2214141128,
		-0.2321805992,
		-0.1656331722,
		-0.05373789678,
		0.08083727586,
		0.09475354363,
		0.03209767112,
		-0.07534076818,
		-0.2050442354,
		-0.2138010557,
		-0.1458045181,
		-0.03293263556,
		0.1022243734,
		0.1163957964,
		0.05372761902,
		-0.05393941791,
		-0.1840438454,
		-0.1933365048,
		-0.1259788988,
		-0.01382225715,
		0.1205656194,
		0.1339326478,
		0.07044017167,
		-0.03805851969,
		-0.1689918111,
		-0.1791024416,
	}
)

func TestMACDSeries(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValuesProvider{
		emaXValues,
		emaYValues,
	}
	assert.Equal(50, mockSeries.Len())

	mas := &MACDSeries{
		InnerSeries: mockSeries,
	}

	var yvalues []float64
	for x := 0; x < mas.Len(); x++ {
		_, y := mas.GetValues(x)
		yvalues = append(yvalues, y)
	}

	assert.NotEmpty(yvalues)
	for index, vy := range yvalues {
		assert.InDelta(vy, macdExpected[index], emaDelta, fmt.Sprintf("delta @ %d actual: %0.9f expected: %0.9f", index, vy, macdExpected[index]))
	}
}
