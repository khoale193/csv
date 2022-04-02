package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/khoale193/csv/modules/ohlc_data"
	"github.com/khoale193/csv/modules/sync"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Router Sync Open-High-Low-Close Data
	sync.SyncRouter(r.Group("/api"))

	// OHLC Data
	ohlc_data.OhlcRouter(r.Group("/api"))

	return r
}
