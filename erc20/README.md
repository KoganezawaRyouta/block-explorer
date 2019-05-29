# erc20

$(npm bin)/truffle init
$(npm bin)/zos init erc-20
$(npm bin)/zos link openzeppelin-eth

## install the dependencies
```
npm install && cd ./client/; npm install
```

## run ethereum simulator
Set Metamask network to localhost:8545
```
$(npm bin)/ganache-cli -p 8545 -h 127.0.0.1 -d
```

## create migration file for deploy ( get contract address )
```
$(npm bin)/zos add InstagramPosting
$(npm bin)/zos push --network local
```

### run web server
```
$ cd ./client/; npm run dev
```
