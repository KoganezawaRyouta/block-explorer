# instdapps

## install the dependencies
```
npm install
```

## contracts compile
```
$ $(npm bin)/truffle compile
```

## run ethereum simulator
```
$ $(npm bin)/truffle develop
Truffle Develop started at http://127.0.0.1:9545/ // Set this to RPC on MetaMask 
```

## create migration file for deploy ( get contract address )
```
$ $(npm bin)/truffle create migration DeploySafeMath
$ $(npm bin)/truffle develop
$ truffle(develop)> migrate
```

### run web server
```
$ cd ./client/; npm run dev
```
