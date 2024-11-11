package ta

import "github.com/shopspring/decimal"

type Source interface {
	Close() []decimal.Decimal
	Open() []decimal.Decimal
	High() []decimal.Decimal
	Low() []decimal.Decimal
}
