package service

import "github.com/khoale193/csv/modules/sync/model"

type SyncOhlcService struct {
	Data model.OhlcList
}

func (a *SyncOhlcService) SyncData() error {
	a.Data.CreateInsertUpdateOhlcData()
	return nil
}
