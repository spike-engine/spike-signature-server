package cmd

import (
	"_spike-signature-server/constant"
	"bufio"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

func ReSetWalletPassWordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "reset wallet password",
		Long: `Reset the passphrase of the wallet you specified.
You can specify the file path through --wallet-dir,and the file name through --file-name.
`,
		RunE: runResetCmd,
	}
	return cmd
}

func runResetCmd(cmd *cobra.Command, args []string) error {
	var err error

	walletDir, err := cmd.Flags().GetString(constant.FlagWalletDir)
	if err != nil {
		return err
	}

	if !path.IsAbs(walletDir) {
		return errors.New("can't input wallet-dir or input path is not an absolute path")
	}

	fileName, err := cmd.Flags().GetString(constant.FlagWalletFileName)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(fileName, ".json") {
		return errors.New("can't input wallet-dir or fileName is Illegal")
	}

	walletFilePath := path.Join(walletDir, fileName)
	jsonFile, err := os.Open(walletFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	fileinfo, err := jsonFile.Stat()
	if err != nil {
		return err
	}
	jsonFileBuffer := make([]byte, fileinfo.Size())

	_, err = jsonFile.Read(jsonFileBuffer)
	if err != nil {
		return err
	}

	inputBuf := bufio.NewReader(cmd.InOrStdin())
	oldPassword, err := input.GetPassword("Enter old wallet passphrase:\n", inputBuf)
	if err != nil {
		return err
	}

	newPassword, err := input.GetPassword("Enter new wallet passphrase:\n", inputBuf)
	if err != nil {
		return err
	}

	ks := keystore.NewKeyStore(walletDir, keystore.StandardScryptN, keystore.StandardScryptP)

	key, err := keystore.DecryptKey(jsonFileBuffer, oldPassword)
	if key != nil && key.PrivateKey != nil {
		defer func(key *keystore.Key) {
			for i := range key.PrivateKey.D.Bits() {
				key.PrivateKey.D.Bits()[i] = 0
			}
		}(key)
	}

	if err != nil {
		return err
	}

	account, err := ks.Find(accounts.Account{Address: key.Address})
	if err != nil {
		return err
	}
	err = ks.Update(account, oldPassword, newPassword)
	if err != nil {
		return err
	}

	fmt.Println("passphrase modify success")
	return nil
}
