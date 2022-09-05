package config

import (
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("config")

var Cfg Config

type Config struct {
	Wallet Wallet `toml:"wallet"`
	System System `toml:"system"`
}

type Wallet struct {
	WalletPath string `toml:"walletPath"`
	PassPhrase string `toml:"passPhrase"`
	ChainId    int64  `toml:"chainId"`
}

type System struct {
	Port      string `toml:"port"`
	MachineId string `toml:"machineId"`
}
