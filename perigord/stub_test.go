package main

import (
	"testing"

	_ "github.com/KoganezawaRyouta/block-explorer/perigord/migrations"
	_ "github.com/KoganezawaRyouta/block-explorer/perigord/tests"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
