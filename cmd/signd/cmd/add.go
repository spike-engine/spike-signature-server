package cmd

import (
	"_spike-signature-server/constant"
	"bufio"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

func AddWalletCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add local wallet",
		Long: `Derive a new private key and encrypt to disk.
Use the tools provided by go-ethereum to generate ethereum ecologically compatible wallets.

You can specify the file path through --wallet-dir,and the file name through --file-name.
`,
		RunE: runAddCmd,
	}

	return cmd
}

func runAddCmd(cmd *cobra.Command, args []string) error {
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

	buf := bufio.NewReader(cmd.InOrStdin())
	firstInputPassword, err := input.GetPassword("Enter wallet passphrase:\n", buf)
	if err != nil {
		return err
	}

	secondInputPassword, err := input.GetPassword("Re-enter keyring passphrase:\n", buf)
	if err != nil {
		return err
	}

	if !strings.EqualFold(firstInputPassword, secondInputPassword) {
		return errors.New("two inputs are inconsistent. Please re-try")
	}

	walletFilePath := path.Join(walletDir, fileName)

	_, err = os.Open(walletFilePath)
	if err == nil {
		return errors.New("this file already exists and overwriting files is prohibited")
	}

	ks := keystore.NewKeyStore(walletDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(firstInputPassword)
	if err != nil {
		return err
	}
	err = os.Rename(account.URL.Path, walletFilePath)
	if err != nil {
		return err
	}

	fmt.Println("please check your wallet,", walletFilePath)
	return nil
}
