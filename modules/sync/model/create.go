package model

import (
	"fmt"

	"github.com/khoale193/csv/models"
	"github.com/khoale193/csv/models/dbcon"
)

type OhlcList []models.Ohlc

func (a OhlcList) CreateInsertUpdateOhlcData() error {
	bar := []string{"Unix", "Symbol", "Open", "High", "Low", "Close"}
	columnName := models.ColumnsName((models.Ohlc{}).TableName(), &models.Ohlc{}, bar)
	columnNameValues := models.ColumnsNameValueList(&models.Ohlc{}, bar)
	for _, item := range a {
		if _, err := dbcon.GetSqlXDB().NamedExec(fmt.Sprintf("insert into %[1]s (%[2]s) values (%[3]s) on duplicate key update open = :open, high = :high, low = :low, close = :close",
			(models.Ohlc{}).TableName(), columnName, columnNameValues), item); err != nil {
			return err
		}
	}
	return nil
}
