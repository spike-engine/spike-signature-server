package main

import (
	"_spike-signature-server/config"
	"_spike-signature-server/constant"
	"_spike-signature-server/initialize"
	logger "github.com/ipfs/go-log"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	constant.Viper = config.InitViper()
	config.InitWallet()
	initialize.RunServer()
}
