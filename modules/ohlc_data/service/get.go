package service

import (
	"github.com/khoale193/csv/modules/ohlc_data/model"
	"github.com/khoale193/csv/pkg/pagination"
)

type GetOhlcDataForm struct {
	Symbol     string
	Pagination *pagination.Paginator
}

func (a *GetOhlcDataForm) GetOhlcData() ([]*OhlcResponse, error) {
	if data, count, err := model.SelectOhlc(a.Symbol, a.Pagination.PerPageNums(), a.Pagination.Offset()); err != nil {
		return nil, err
	} else {
		a.Pagination.SetNums(count)
		return OhlcListSerializer(data).Response(), nil
	}

}
