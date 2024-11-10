package main

import (
	"Test1/store"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func deploy() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/_gI0dP6yhERf2Rlx9zcasrkwRIkTEQ6m")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.GenerateKey()
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// privateKeyHex := hex.EncodeToString(privateKeyBytes)
	// fmt.Println("Private Key:", privateKeyHex)
	privateKey, err := crypto.HexToECDSA("27581412d290be11702d251cecdf811e1a735cb26dff2f9046420958957915b0")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(10000000000))
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
