// Example main file for a native dapp, replace with application code
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/KoganezawaRyouta/block-explorer/perigord/bindings"
	_ "github.com/KoganezawaRyouta/block-explorer/perigord/migrations"
	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"
)

var greeter *bindings.Greeter

func main() {
	// 初期化
	err := network.InitNetworks()
	if err != nil {
		log.Fatalln("Could not connect to dev network: ", err)
	}

	// perigord.yamlから情報を取得し、接続を行う
	nw, err := network.Dial("dev")
	if err != nil {
		log.Fatalln("Could not connect to dev network: ", err)
	}

	// マイグレーション実行(デプロイ)
	if err := migration.RunMigrations(context.Background(), nw, false); err != nil {
		log.Fatalln("Error running migrations: ", err)
	}

	// コントラクト接続
	address := contract.AddressOf("Greeter")
	greeter, err = bindings.NewGreeter(address, nw.Client())

	// イベント監視用チャンネル
	eventChan := make(chan *bindings.GreeterResult)
	sub, err := greeter.WatchResult(nil, eventChan)
	if err != nil {
		log.Println("error listening for incoming events:", err)
		return
	}

	// データ取得
	ret, err := greeter.Greet(nil)
	if err != nil {
		log.Println("error listening for incoming events:", err)
		return
	}
	fmt.Printf("更新前: %s\n", ret)

	// データ更新
	tran, err := greeter.SetGreeting(nil, "Hello, World!:D")
	if err != nil {
		log.Println("error listening for incoming events:", err)
		return
	}
	fmt.Printf("Event結果: %v\n", tran)

	for {
		event := <-eventChan
		fmt.Println("Event結果\n")
		fmt.Printf("from: %v\n", event.Raw)
		fmt.Printf("stored: %v\n", event.Stored)
		// 先ほど更新した内容("Hello, World!:D")に更新されているか確認
		m, err := greeter.Greet(nil)
		if err != nil {
			log.Println("error listening for incoming events:", err)
			return
		}
		fmt.Printf("更新後: %s\n", m)
		sub.Unsubscribe()
	}

}
