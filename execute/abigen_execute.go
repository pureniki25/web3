package execute

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"Test1/store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SepoliaUrl        = "https://eth-sepolia.g.alchemy.com/v2/_gI0dP6yhERf2Rlx9zcasrkwRIkTEQ6m"
	WssSepoliaUrl     = "wss://eth-sepolia.g.alchemy.com/v2/_gI0dP6yhERf2Rlx9zcasrkwRIkTEQ6m"
	ContractAddr      = "0xB1d757C224893925728270925Cb30085E15c35c1"
	AccountPrivateKey = "27581412d290be11702d251cecdf811e1a735cb26dff2f9046420958957915b0"
)

func execute() {
	client, err := ethclient.Dial(SepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store.NewStore(common.HexToAddress(ContractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(AccountPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	(key[:], []byte("demo_save_key"))
	(value[:], []byte("demo_save_value11111"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
}
