package config

import (
	"os"
	"strings"
)

func InitWallet() {
	if !strings.HasSuffix(Cfg.Wallet.WalletFolder, "/") {
		Cfg.Wallet.WalletFolder = Cfg.Wallet.WalletFolder + "/"
	}

	_, err := os.Open(Cfg.Wallet.WalletFolder + Cfg.Wallet.WalletFile)
	if err != nil {
		panic(err)
	}

}
