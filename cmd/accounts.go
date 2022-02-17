package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"strings"
)
var errType error
var flowConnectionAddr = "127.0.0.1:3569"
var connectionError = errors.New("connection error")
var invalidAccountError = errors.New("invalid account error")

func init() {
	rootCmd.AddCommand(accountsCmd)
	accountsCmd.AddCommand(getAccountsCmd)
}

var accountsCmd = &cobra.Command{
	Use:     "accounts",
	Short:   "Manage Flow accounts",
	Long:    `This command can be used to manage Flow accounts`,
}

var getAccountsCmd = &cobra.Command{
	Use:   "get",
	Short: "Get flow account information",
	Long:  `This command can be used to retrieve Flow account and its properties`,
	Example: "flowcli accounts get 179b6b1cb6755e32",
	Args:    cobra.ExactArgs(1),
	Run: getAccount,
}

func getAccount(cmd *cobra.Command, args []string)  {
	fmt.Println("retrieving Flow account info")

	ctx := context.Background()
	//initialize flow client
	flowClient, err := client.New(flowConnectionAddr, grpc.WithInsecure())
	if err != nil {
		handleError(err)
		return
	}
	defer func() {
		err := flowClient.Close()
		if err != nil {
			handleError(err)
		}
	}()

	address := flow.HexToAddress(args[0])
	//get account info through Flow Go SDK using given address
	account, err := flowClient.GetAccount(ctx, address)
	if err != nil {
		handleError(err)
		return
	}
	printAccount(account)
}

func getErrorType(err error) error {
	es := err.Error()
	if strings.Contains(es, "connection refused") {
		return connectionError
	} else if strings.Contains(es, "invalid for chain") {
		return invalidAccountError
	}
	return err
}

func handleError(err error) {
	fmt.Println("err:", err.Error())
	errType = getErrorType(err)
	switch {
	case errors.Is(errType, connectionError):
		fmt.Println("Please make sure Flow Emulator is running and try again")
	case errors.Is(errType, invalidAccountError):
		fmt.Println("Please make sure valid account address is entered and try again")
	default:
	}
}

func printAccount(account *flow.Account) {
	fmt.Printf("\nAddress: %s", account.Address.String())
	fmt.Printf("\nBalance: %d", account.Balance)
	fmt.Printf("\nContracts: %d", len(account.Contracts))
	fmt.Printf("\nKeys: %d\n", len(account.Keys))
}
