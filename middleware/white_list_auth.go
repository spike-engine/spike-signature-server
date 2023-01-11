package middleware

import (
	"github.com/gin-gonic/gin"

	"_spike-signature-server/config"
	"_spike-signature-server/response"
)

func WhiteListAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		log.Infof("clientIp : %s", clientIp)
		for _, ip := range config.Cfg.TxApiWhiteList.IpList {
			if ip == clientIp {
				c.Next()
				return
			}
		}
		response.FailWithMessage("Yor are not in whiteList", c)
		c.Abort()
	}
}
