package initialize

import (
	v1 "_spike-signature-server/api/v1"
	"_spike-signature-server/middleware"
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("initialize")

func initRouter() (*gin.Engine, error) {
	var r = gin.Default()
	r.Use(middleware.Cors())
	publicGroup := r.Group("")
	{
		// health
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	routerGroupApp, err := v1.NewRouterGroup()
	if err != nil {
		return nil, err
	}

	txGroup := routerGroupApp.TxGroup
	txApiGroup := r.Group("/tx-api/v1")
	txGroup.InitTxGroup(txApiGroup)
	return r, nil
}
