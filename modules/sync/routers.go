package sync

import (
	"github.com/gin-gonic/gin"

	"github.com/khoale193/csv/middleware/authorization"
	"github.com/khoale193/csv/modules/sync/handler"
)

func SyncRouter(router *gin.RouterGroup) {
	// Member Site Router
	router.Use(authorization.JWT())
	{
		router.POST("/sync", handler.SyncOhlcData)
	}
}
