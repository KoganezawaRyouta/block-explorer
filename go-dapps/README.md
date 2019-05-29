generate the ABI from a solidity source file
```bash
solc --abi --bin contracts/Store.sol -o ./build
```

convert the ABI to a Go file that we can import
```bash
abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=Store.go
```
