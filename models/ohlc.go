package models

import (
	"github.com/khoale193/csv/pkg/e"
)

type Ohlc struct {
	Unix   int64   `gorm:"column:unix;not null;index" db:"unix"`
	Symbol string  `gorm:"column:symbol;size:128;not null;index" db:"symbol"`
	Open   float64 `gorm:"column:open;not null" db:"open"`
	High   float64 `gorm:"column:high;not null" db:"high"`
	Low    float64 `gorm:"column:low;not null" db:"low"`
	Close  float64 `gorm:"column:close;not null" db:"close"`
}

func (Ohlc) TableName() string {
	return e.OhlcTable
}
