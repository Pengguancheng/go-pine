package pine_binance

import (
	"context"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/shopspring/decimal"
	"go-pine/core/ta"
)

func NewTaFutureSource(apiKey, secretKey string, len int, interval ta.Interval) (ta.Source, error) {
	svc := &FutureSource{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		Len:       len,
		Interval:  interval,
		Client:    futures.NewClient(apiKey, secretKey),
	}
	res, err := svc.Client.NewKlinesService().Limit(svc.Len).Interval(string(svc.Interval)).Do(context.Background())
	if err != nil {
		return nil, err
	}
	svc.KLArr = res
	return svc, nil
}

type FutureSource struct {
	ApiKey    string
	SecretKey string
	Len       int
	Interval  ta.Interval
	Client    *futures.Client
	KLArr     []*futures.Kline
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
