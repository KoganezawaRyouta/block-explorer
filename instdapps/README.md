# instdapps

## install the dependencies
```
npm install && cd ./client/; npm install
```

## create migration file for deploy ( get contract address )
```
$(npm bin)/zos init Instdapps
$(npm bin)/zos add InstagramPosting
$(npm bin)/zos push --network development
```

### run web server
```
$ cd ./client/; npm run dev
```