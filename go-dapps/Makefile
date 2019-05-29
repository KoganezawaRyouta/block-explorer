deps:
	go get -t -d ./...

bld: deps
	go build -v -o ${GOPATH}/bin/gd-ship ./cmd/main.go

ship: bld deps
	gd-ship -d

load: bld deps
	gd-ship -l

set: bld deps
	gd-ship -s
