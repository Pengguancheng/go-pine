package pine_binance

import (
	"github.com/stretchr/testify/assert"
	"go-pine/core/ta"
	"testing"
)

func TestNewTaFutureSource(t *testing.T) {
	source, err := NewTaFutureSource(API_KEY, SECRET_KEY, "BTCUSDT", 99, ta.Interval1h)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, source)
	s := source.(*FutureSource)
	for i := 1; i < len(s.KLArr); i++ {
		assert.GreaterOrEqual(t, s.KLArr[i-1].OpenTime, s.KLArr[i].OpenTime,
			"Klines should be sorted from max to min based on timestamp")
	}
}

func TestFutureSource_Close(t *testing.T) {
	source, err := NewTaFutureSource(API_KEY, SECRET_KEY, "BTCUSDT", 99, ta.Interval1h)
	if err != nil {
		t.Fatal(err)
	}
	closeArr := source.Close()
	assert.NotEmpty(t, closeArr)
}
