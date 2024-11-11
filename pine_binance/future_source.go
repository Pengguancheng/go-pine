package pine_binance

import (
	"context"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/shopspring/decimal"
	"go-pine/core/ta"
)

func NewTaFutureSource(apiKey, secretKey, symbol string, dataLen int, interval ta.Interval) (ta.Source, error) {
	svc := &FutureSource{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		Len:       dataLen,
		Interval:  interval,
		Client:    futures.NewClient(apiKey, secretKey),
		KLArr:     nil,
		Symbol:    symbol,
	}
	res, err := svc.Client.NewKlinesService().
		Symbol(svc.Symbol).
		Limit(svc.Len).
		Interval(string(svc.Interval)).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	svc.KLArr = make([]*futures.Kline, len(res))
	// reverse arr sort
	for i := 0; i < len(res); i++ {
		reverseIndex := len(res) - 1 - i
		svc.KLArr[reverseIndex] = res[i]
	}
	return svc, nil
}

type FutureSource struct {
	ApiKey    string
	SecretKey string
	Len       int
	Interval  ta.Interval
	Client    *futures.Client
	KLArr     []*futures.Kline
	Symbol    string
}

func (f *FutureSource) Close() []decimal.Decimal {
	rs := make([]decimal.Decimal, len(f.KLArr))
	for i, v := range f.KLArr {
		rs[i], _ = decimal.NewFromString(v.Close)
	}
	return rs
}

func (f *FutureSource) Open() []decimal.Decimal {
	rs := make([]decimal.Decimal, len(f.KLArr))
	for i, v := range f.KLArr {
		rs[i], _ = decimal.NewFromString(v.Open)
	}
	return rs
}

func (f *FutureSource) High() []decimal.Decimal {
	rs := make([]decimal.Decimal, len(f.KLArr))
	for i, v := range f.KLArr {
		rs[i], _ = decimal.NewFromString(v.High)
	}
	return rs
}

func (f *FutureSource) Low() []decimal.Decimal {
	rs := make([]decimal.Decimal, len(f.KLArr))
	for i, v := range f.KLArr {
		rs[i], _ = decimal.NewFromString(v.Low)
	}
	return rs
}
