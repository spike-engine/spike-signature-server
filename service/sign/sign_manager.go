package sign

import (
	"_spike-signature-server/config"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	logger "github.com/ipfs/go-log"
	"math/big"
	"os"
)

var log = logger.Logger("sign")

type SignatureManager struct {
	transactOpts *bind.TransactOpts
}

func New() (*SignatureManager, error) {

	walletPath, err := os.Open(config.Cfg.Wallet.WalletFolder + config.Cfg.Wallet.WalletFile)
	if err != nil {
		panic(err)
	}

	transactOpts, err := bind.NewTransactorWithChainID(walletPath, config.Cfg.Wallet.PassPhrase, big.NewInt(config.Cfg.Wallet.ChainId))
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
