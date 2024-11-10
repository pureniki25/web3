package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func account() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/_gI0dP6yhERf2Rlx9zcasrkwRIkTEQ6m")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xd79911b5E62A7CDb1Bc820Cd8667e6d973607A09")
	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Fatal(err)
	}
	pendBalance, err2 := client.PendingBalanceAt(context.Background(), account)

	if err2 != nil {
		log.Fatal(err2)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)    // 25.729324269165216041
	fmt.Println(balance)     // 25893180161173005034
	fmt.Println(pendBalance) // 25893180161173005034
}
