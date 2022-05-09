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
	// player := game.NewTestPlayer()
	playerGM := game.NewTestPlayer()

	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		if time.Now().Unix()%2 == 0 {
			playerGM.ModPlayer.SetShowTeam([]int{1001, 1001, 1001, 1002, 1001, 1005, 1006}, playerGM)
			playerGM.ModPlayer.SetShowTeam([]int{}, playerGM)
			playerGM.ModPlayer.SetShowTeam([]int{1009}, playerGM)
		}
	}

}
