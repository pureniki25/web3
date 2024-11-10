package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func connect() {
	// 连接到以太坊节点
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// 获取最新块号
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 打印最新块号
	fmt.Printf("Latest block number: %d\n", blockNumber)
}
