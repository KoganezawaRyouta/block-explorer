package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	store "github.com/KoganezawaRyouta/block-explorer/go-dapps/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	deployMode  = flag.Bool("d", false, "")
	loadMode    = flag.Bool("l", false, "")
	setItemMode = flag.Bool("s", false, "")
)

func init() {
	flag.Parse()
}

func main() {
	if *deployMode {
		deploy()
	}
	if *loadMode {
		load()
	}
	if *setItemMode {
		setItem()
	}
}

func load() {
	client, err := ethclient.Dial("http://ganache:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x808253c9562a2919a2d4cd362f11a1ffa7998531")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("instance version: %s\n", version)
	_ = instance
}

func deploy() {
	client, err := ethclient.Dial("http://ganache:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 0xed19d0e3fc1e8d3bb92389bf993943949c6c96f17f4bf506bb0b5c5194ee780b
	privateKey, err := crypto.HexToECDSA("ed19d0e3fc1e8d3bb92389bf993943949c6c96f17f4bf506bb0b5c5194ee780b")
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

	auth := bind.NewKeyedTransactor(privateKey)
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
	fmt.Println(address.String())

	fmt.Println(tx.Hash().Hex())
	fmt.Println(tx.Hash().String())

	_ = instance
}

func setItem() {
	client, err := ethclient.Dial("http://ganache:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("ed19d0e3fc1e8d3bb92389bf993943949c6c96f17f4bf506bb0b5c5194ee780b")
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x808253c9562a2919a2d4cd362f11a1ffa7998531")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "bar"
}
