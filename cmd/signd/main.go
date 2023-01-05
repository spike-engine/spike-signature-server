package main

import (
	"_spike-signature-server/cmd/signd/cmd"
	logger "github.com/ipfs/go-log"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	cmd.Execute()

}
