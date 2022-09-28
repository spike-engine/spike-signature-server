package config

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
	"strings"
)

func InitWallet() {
	if !strings.HasSuffix(Cfg.Wallet.WalletFolder, "/") {
		Cfg.Wallet.WalletFolder = Cfg.Wallet.WalletFolder + "/"
	}

	_, err := os.Open(Cfg.Wallet.WalletFolder + Cfg.Wallet.WalletFile)
	if err != nil {
		ks := keystore.NewKeyStore(Cfg.Wallet.WalletFolder, keystore.StandardScryptN, keystore.StandardScryptP)
		account, err := ks.NewAccount(Cfg.Wallet.PassPhrase)
		if err != nil {
			panic(err)
		}
		err = os.Rename(account.URL.Path, Cfg.Wallet.WalletFolder+Cfg.Wallet.WalletFile)
		if err != nil {
			panic(err)
		}

	}

}
