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

	// player := game.NewTestPlayer()
	playerGM := game.NewTestPlayer()

	ticker := time.NewTicker(time.Second * 1)
	fmt.Print("test")
	for range ticker.C {
		if time.Now().Unix()%2 == 0 {
			playerGM.ModPlayer.AddExp(50000000, playerGM)
		} else {
			// playerGM.ModUniqueTask.MyTaskInfo[10001] = &game.TaskInfo{
			// 	TaskId: 10001,
			// 	State:  1,
			// }
		}
	}

}
