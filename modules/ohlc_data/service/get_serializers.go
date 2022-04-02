package service

import (
	"time"

	"github.com/khoale193/csv/modules/ohlc_data/model"
	"github.com/khoale193/csv/pkg/e"
)

type OhlcResponse struct {
	Time   time.Time `json:"time"`
	Symbol string    `json:"symbol"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
}

type OhlcSerializer model.OhlcResponse

func (s OhlcSerializer) Response() *OhlcResponse {
	timezone, _ := time.LoadLocation(e.DefaultTimezone)
	response := OhlcResponse{
		Time:   time.UnixMilli(s.UnixTime).In(timezone),
		Symbol: s.Symbol,
		Open:   s.Open,
		High:   s.High,
		Low:    s.Low,
		Close:  s.Close,
	}
	return &response
}

type OhlcListSerializer []*model.OhlcResponse

func (s OhlcListSerializer) Response() []*OhlcResponse {
	var response []*OhlcResponse
	for _, item := range s {
		response = append(response, OhlcSerializer(*item).Response())
	}
	return response
}
