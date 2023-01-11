package txApi

import (
	"_spike-signature-server/middleware"
	"_spike-signature-server/request"
	"_spike-signature-server/response"
	"_spike-signature-server/service/sign"
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("txApi")

type TxGroup struct {
	manager *sign.SignatureManager
}

func NewTxGroup() (TxGroup, error) {

	manager, err := sign.New()
	if err != nil {
		log.Error("===Spike log:", err)
		return TxGroup{}, err
	}
	return TxGroup{
		manager: manager,
	}, nil
}

func (api *TxGroup) InitTxGroup(g *gin.RouterGroup) {
	g.Use(middleware.WhiteListAuth())
	tx := g.Group("tx")
	{
		tx.POST("/signature", api.SignatureTransaction)
	}
}

func (api *TxGroup) SignatureTransaction(c *gin.Context) {
	var service request.SignTXService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
	}

	transaction, err := api.manager.SignatureTransaction(service.Tx)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(string(transaction), c)
}
