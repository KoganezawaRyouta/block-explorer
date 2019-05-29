FROM golang:1.11.5

RUN mkdir -p /go/src/github.com/KoganezawaRyouta/block-explorer/go-dapps/contracts
RUN go get -u github.com/ethereum/go-ethereum
WORKDIR /go/src/github.com/ethereum/go-ethereum
RUN make
RUN make devtools
RUN curl -o /usr/bin/solc -fL https://github.com/ethereum/solidity/releases/download/v0.5.0/solc-static-linux
RUN chmod u+x /usr/bin/solc
WORKDIR /go/src/github.com/KoganezawaRyouta/block-explorer/go-dapps