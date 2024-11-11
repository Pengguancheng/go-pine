package ta

import (
	"errors"
	"github.com/shopspring/decimal"
)

func Sma(data []decimal.Decimal, length int) (decimal.Decimal, error) {
	if len(data) < length {
		return decimal.Zero, errors.New("data length is less than the SMA length")
	}

	sum := decimal.Zero
	for _, d := range data[:length] {
		sum = sum.Add(d)
	}

	return sum.Div(decimal.NewFromInt(int64(length))), nil
}
