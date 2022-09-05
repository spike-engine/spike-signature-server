package sign

import (
	"_spike-signature-server/config"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	logger "github.com/ipfs/go-log"
	"math/big"
)

var log = logger.Logger("sign")

type SignatureManager struct {
	transactOpts *bind.TransactOpts
}

func New() (*SignatureManager, error) {

	ks := keystore.NewKeyStore(config.Cfg.Wallet.WalletPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if len(ks.Accounts()) > 1 {
		return nil, errors.New("===Spike log: walletPath exist more than one json")
	}

	err := ks.Unlock(ks.Accounts()[0], config.Cfg.Wallet.PassPhrase)
	if err != nil {
		log.Error("===Spike log:", err)
		return nil, err
	}

	transactOpts, err := bind.NewKeyStoreTransactorWithChainID(ks, ks.Accounts()[0], big.NewInt(config.Cfg.Wallet.ChainId))
	if err != nil {
		log.Error("===Spike log:", err)
		return nil, err
	}

	return &SignatureManager{
		transactOpts: transactOpts,
	}, nil
}

func (s *SignatureManager) SignatureTransaction(transaction *types.Transaction) ([]byte, error) {
	signedTx, err := s.transactOpts.Signer(s.transactOpts.From, transaction)
	if err != nil {
		log.Error("===Spike log:", err)
		return nil, err
	}

	marshal, err := json.Marshal(signedTx)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
