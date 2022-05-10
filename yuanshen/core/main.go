package main

import (
	"fmt"
	"time"

	"yuanshen/game"

	"yuanshen/csvs"
)

func main() {
	fmt.Println("======================原神服务器======================")

	//加载配置
	csvs.CheckLoadCsv()
	go game.GetManageBanWord().Run()
	fmt.Println("======================数据测试======================")

	ticker := time.NewTicker(time.Second * 10)
	for range ticker.C {
		player := game.NewTestPlayer()
		go player.Run()

	}

}
