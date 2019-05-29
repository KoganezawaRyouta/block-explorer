// Invokes the perigord driver application

package main

import (
	_ "github.com/KoganezawaRyouta/block-explorer/perigord/migrations"
	"github.com/polyswarm/perigord/stub"
)

func main() {
	stub.StubMain()
}
