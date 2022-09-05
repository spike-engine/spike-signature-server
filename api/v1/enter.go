package v1

import (
	"_spike-signature-server/api/v1/txApi"
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("api")

type RouterGroup struct {
	TxGroup txApi.TxGroup
}

func NewRouterGroup() (RouterGroup, error) {
	tx, err := txApi.NewTxGroup()
	if err != nil {
		log.Error("===Spike log:", err)
		return RouterGroup{}, err
	}
	return RouterGroup{
		TxGroup: tx,
	}, nil
}
