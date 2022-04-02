package migration

import (
	"fmt"

	"github.com/khoale193/csv/models"
	"github.com/khoale193/csv/models/dbcon"
)

func (a *Migration) Migrate1Up() {
	alterSQL := fmt.Sprintf("alter table %s add constraint idx_ohlc_unix_symbol unique (unix, symbol)", (models.Ohlc{}).TableName())
	dbcon.GetSqlXDB().Exec(alterSQL)
}
