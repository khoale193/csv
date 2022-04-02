package ohlc_data

import (
	"github.com/gin-gonic/gin"

	"github.com/khoale193/csv/middleware/authorization"
	"github.com/khoale193/csv/modules/ohlc_data/handler"
)

func OhlcRouter(router *gin.RouterGroup) {
	// Member Site Router
	router.Use(authorization.JWT())
	{
		router.GET("/ohlc", handler.GetOhlcData)
	}
}
