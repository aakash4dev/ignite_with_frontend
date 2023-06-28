package main

import (
	"context"
	"fmt"
	"log"
    "os"
	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	// "blog/x/blog/types"
	"blogclient/types"


	// "path/filepath"
	// "github.com/cosmos/cosmos-sdk/crypto/hd"
	// "github.com/cosmos/cosmos-sdk/crypto/keyring"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/tendermint/tendermint/libs/cli"
)

func main() {

	// rootDir := os.ExpandEnv("/ignitecli")
	// keyringDir := filepath.Join(rootDir, "keyring")
	// kb,_ := keyring.New(sdk.KeyringServiceName(), keyring.BackendFile, keyringDir, nil)
	// accountName := "myaccount"
	// mnemonic := "your mnemonic phrase here"
	// info, err := kb.NewAccount(accountName, mnemonic, "", hd.Secp256k1)
	// if err != nil {
	// 	fmt.Println("Error creating/recovering key:", err)
	// 	return
	// }

	// // Get the account
	// addr := sdk.AccAddress(info.GetPubKey().Address().Bytes())
	// account, err := kb.GetAccount(addr)
	// if err != nil {
	// 	fmt.Println("Error retrieving account:", err)
	// 	return
	// }

	// // Print the account information
	// fmt.Println("Account address:", account.GetAddress())

	// // os.Exit(0)
	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress("http://localhost:26657"))
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Name)
	fmt.Println(account.Address)

	os.Exit(0)

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	// Define a message to create a post
	msg := &types.MsgCreatePost{
		Creator: addr,
		Title:   "Hello!",
		Body:    "This is the first post",
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
	}

	// Print response from broadcasting a transaction
	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)

	// Instantiate a query client for your `blog` blockchain
	queryClient := types.NewQueryClient(client.Context())

	// Query the blockchain using the client's `PostAll` method
	// to get all posts store all posts in queryResp
	queryResp, err := queryClient.PostAll(ctx, &types.QueryAllPostRequest{})
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)
}
