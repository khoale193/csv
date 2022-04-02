package model

import (
	"fmt"
	"github.com/khoale193/csv/models"
	"github.com/khoale193/csv/models/dbcon"
	"github.com/khoale193/csv/pkg/s"
)

type OhlcResponse struct {
	UnixTime int64   `db:"unix"`
	Symbol   string  `db:"symbol"`
	Open     float64 `db:"open"`
	High     float64 `db:"high"`
	Low      float64 `db:"low"`
	Close    float64 `db:"close"`
}

func SelectOhlc(symbol string, limit, offset int) ([]*OhlcResponse, int64, error) {
	selectSQLInit := `select %[1]s.symbol, %[1]s.unix, %[1]s.open, %[1]s.high, %[1]s.low, %[1]s.close
	from %[1]s
	where %[1]s.symbol = ?
	order by %[1]s.unix %[2]s
	` + models.LimitCondition(limit) + `
	` + models.OffsetCondition(offset) + `;`
	selectSQL := fmt.Sprintf(selectSQLInit, (models.Ohlc{}).TableName(), s.OrderAsc)
	var data []*OhlcResponse
	if err := dbcon.GetSqlXDB().Select(&data, selectSQL, symbol); err != nil {
		return nil, 0, err
	}
	countSQLInit := `select count(*)
	from %[1]s
	where %[1]s.symbol = ?;`
	countSQL := fmt.Sprintf(countSQLInit, (models.Ohlc{}).TableName())
	var count int64
	if err := dbcon.GetSqlXDB().Get(&count, countSQL, symbol); err != nil {
		return nil, 0, err
	}
	return data, count, nil
}
