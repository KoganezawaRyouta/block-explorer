package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"

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
	totalSupply()
}

func totalSupply() {
	fmt.Println("totalSupply ===============")

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/e768b31d63f34e939640c06edaaf2dc1")
	if err != nil {
		log.Fatal(err)
	}
	tokenAddress := common.HexToAddress("0xaBfceE43417680B51c823c4A0219C9D43ACfd489")
	approveFnSignature := []byte("totalSupply()")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(approveFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	var data []byte
	data = append(data, methodID...)
	totalSupply, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// hexadecimal []byte -> hexadecimal string -> hexadecimal -> decimal number
	v, _ := strconv.ParseUint(hex.EncodeToString(totalSupply), 16, 0)
	fmt.Println(v)
}

func approveFnSignature() {
	fmt.Println("approveFnSignature ===============")

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/e768b31d63f34e939640c06edaaf2dc1")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("79191cad68a10e8170d8240313bbff1083a874b79d0a22b9481c4ab971008572") // 0xed19d0e3fc1e8d3bb92389bf993943949c6c96f17f4bf506bb0b5c5194ee780b
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

	toAddress := common.HexToAddress("0xf45299bf372194caa1a9ea384b9147febd22b95f")
	tokenAddress := common.HexToAddress("0xaBfceE43417680B51c823c4A0219C9D43ACfd489")

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

func transferFnSignature() {
	fmt.Println("transferFnSignature ===============")

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/e768b31d63f34e939640c06edaaf2dc1")
	if err != nil {
		fmt.Println("error: client.Dial")
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("79191cad68a10e8170d8240313bbff1083a874b79d0a22b9481c4ab971008572")
	if err != nil {
		fmt.Println("error: crypto.HexToECDSA")
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error: publicKey.(*ecdsa.PublicKey)")
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("error: client.PendingNonceAt(context.Background(), fromAddress)")
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("error: client.SuggestGasPrice(context.Background())")
		log.Fatal(err)
	}

	//fromAddress = common.HexToAddress("0x21422af071f0368eff2b52ba1ffbacddaacfc4e2")
	toAddress := common.HexToAddress("0xf45299bf372194caa1a9ea384b9147febd22b95f")
	tokenAddress := common.HexToAddress("0xaBfceE43417680B51c823c4A0219C9D43ACfd489")

	transferFromFnSignature := []byte("transferFrom(address,address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFromFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedFromAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedFromAddress))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("10", 10)

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedFromAddress...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//gas: 4465030,
	//gasPrice: 10000000000,
	fmt.Println(gasPrice) // 23256
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		//To:   &tokenAddress,
		////Value: value,
		////Gas:      4465030,
		////GasPrice: gasPrice,
		//Data: data,
		// 上記パラメータを設定すると、 `gas required exceeds allowance or always failing transaction` というエラーが発生する
	})
	if err != nil {
		fmt.Println("error: client.EstimateGas")
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256
	fmt.Println(nonce)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("error: client.NetworkID(context.Background())")
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("error: types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)")
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("error: client.SendTransaction(context.Background(), signedTx)")
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
