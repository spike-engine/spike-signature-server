package config

import "github.com/ethereum/go-ethereum/accounts/keystore"

func InitWallet() {

	ks := keystore.NewKeyStore(Cfg.Wallet.WalletPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if len(ks.Accounts()) == 0 {
		_, err := ks.NewAccount(Cfg.Wallet.PassPhrase)
		if err != nil {
			panic(err)
		}
	}
}
