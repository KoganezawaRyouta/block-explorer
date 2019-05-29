package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func main() {
	//approveFnSignature()

	//transferFnSignature()

	test()
}

func approveFnSignature() {
	fmt.Println("approveFnSignature ===============")

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("ed19d0e3fc1e8d3bb92389bf993943949c6c96f17f4bf506bb0b5c5194ee780b") // 0xed19d0e3fc1e8d3bb92389bf993943949c6c96f17f4bf506bb0b5c5194ee780b
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

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xf69fc8a0aa6b2d0f3bafa7d40ee501a788b0d65e")
	tokenAddress := common.HexToAddress("0x213C11560828125125844A1942B26A29A98cfDBf")

	approveFnSignature := []byte("approve(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(approveFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("10", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256
	fmt.Println(nonce)
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func test() {
	fmt.Println("approveFnSignature ===============")

	client, err := ethclient.Dial("http://localhost:8545")
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

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0x213C11560828125125844A1942B26A29A98cfDBf")

	approveFnSignature := []byte("totalSupply()")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(approveFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	var data []byte
	data = append(data, methodID...)
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256
	fmt.Println(nonce)
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

//
//func transferFnSignature() {
//	fmt.Println("transferFnSignature ===============")
//
//	client, err := ethclient.Dial("http://localhost:8545")
//	if err != nil {
//		fmt.Println("error: client.Dial(http://localhost:8545)")
//		log.Fatal(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA("ed19d0e3fc1e8d3bb92389bf993943949c6c96f17f4bf506bb0b5c5194ee780b")
//	if err != nil {
//		fmt.Println("error: crypto.HexToECDSA")
//		log.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		fmt.Println("error: publicKey.(*ecdsa.PublicKey)")
//		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		fmt.Println("error: client.PendingNonceAt(context.Background(), fromAddress)")
//		log.Fatal(err)
//	}
//
//	value := big.NewInt(0)
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		fmt.Println("error: client.SuggestGasPrice(context.Background())")
//		log.Fatal(err)
//	}
//
//	//fromAddress = common.HexToAddress("0xf69fc8a0aa6b2d0f3bafa7d40ee501a788b0d65e")
//	toAddress := common.HexToAddress("0x9e26730bf972d6ae8840b8f6003a61f57e25ff92")
//	tokenAddress := common.HexToAddress("0x213C11560828125125844A1942B26A29A98cfDBf")
//
//	transferFromFnSignature := []byte("transferFrom(address,address,uint256)")
//	hash := sha3.NewLegacyKeccak256()
//	hash.Write(transferFromFnSignature)
//	methodID := hash.Sum(nil)[:4]
//	fmt.Println(hexutil.Encode(methodID))
//
//	paddedFromAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
//	fmt.Println(hexutil.Encode(paddedFromAddress))
//
//	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
//	fmt.Println(hexutil.Encode(paddedAddress))
//
//	amount := new(big.Int)
//	amount.SetString("10", 10)
//
//	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
//	fmt.Println(hexutil.Encode(paddedAmount))
//
//	var data []byte
//	data = append(data, methodID...)
//	data = append(data, paddedFromAddress...)
//	data = append(data, paddedAddress...)
//	data = append(data, paddedAmount...)
//
//	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
//		To:   &tokenAddress,
//		Data: data,
//	})
//	if err != nil {
//		fmt.Println("error: client.EstimateGas(context.Background(), ethereum.CallMsg{")
//		log.Fatal(err)
//	}
//	fmt.Println(gasLimit) // 23256
//	fmt.Println(nonce)
//
//	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		fmt.Println("error: client.NetworkID(context.Background())")
//		log.Fatal(err)
//	}
//
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//	if err != nil {
//		fmt.Println("error: types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)")
//		log.Fatal(err)
//	}
//
//	err = client.SendTransaction(context.Background(), signedTx)
//	if err != nil {
//		fmt.Println("error: client.SendTransaction(context.Background(), signedTx)")
//		log.Fatal(err)
//	}
//
//	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
//}
