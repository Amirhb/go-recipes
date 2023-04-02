package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"fmt"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/")
	if err != nil {
		fmt.Println(err)
	}
	client.Close()

	clientWithContext, err := ethclient.DialContext(context.Background(), "https://mainnet.infura.io/v3")
	if err != nil {
		fmt.Println(err)
	}
	clientWithContext.Close()

	clientRpc, err := rpc.Dial("https://mainnet.infura.io/v3/")
	if err != nil {
		fmt.Println(err)
	}
	clientWithRpc := ethclient.NewClient(clientRpc)
	clientWithRpc.Close()
}
