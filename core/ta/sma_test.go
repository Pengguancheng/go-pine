package ta

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSMA(t *testing.T) {
	tests := []struct {
		name     string
		data     []decimal.Decimal
		period   int
		expected decimal.Decimal
	}{
		{
			name:     "Basic SMA",
			data:     decimalSlice(1, 2, 3, 4, 5),
			period:   3,
			expected: decimal.NewFromFloat((1 + 2 + 3) / 3),
		},
		{
			name:     "SMA with longer period",
			data:     decimalSlice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			period:   5,
			expected: decimal.NewFromFloat((1 + 2 + 3 + 4 + 5) / 5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sma, err := Sma(tt.data, tt.period)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected.String(), sma.String())
		})
	}
}

// Helper function to create a slice of decimal.Decimal from int values
func decimalSlice(values ...int) []decimal.Decimal {
	decimals := make([]decimal.Decimal, len(values))
	for i, v := range values {
		decimals[i] = decimal.NewFromInt(int64(v))
	}
	return decimals
}
