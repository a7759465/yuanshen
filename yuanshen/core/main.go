package main

import (
	"fmt"
	"runtime"
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
	playerGM.ModPlayer.AddExp(50000000, playerGM)
	// ticker := time.NewTicker(time.Second * 1)
	// fmt.Print("test")
	// for range ticker.C {
	// 	if time.Now().Unix()%2 == 0 {
	// 		playerGM.ModPlayer.AddExp(50000000, playerGM)
	// 	} else {
	// 		// playerGM.ModUniqueTask.MyTaskInfo[10001] = &game.TaskInfo{
	// 		// 	TaskId: 10001,
	// 		// 	State:  1,
	// 		// }
	// 		playerGM.ModUniqueTask.MyTaskInfo[10001] = &game.TaskInfo{}

	// 	}
	// }

	fmt.Println("======================数据测试======================")
	runtime.GOMAXPROCS(1)
	go playerSet(playerGM)
	go playerGet(playerGM)

	for {

	}
}

func playerSet(player *game.Player) {
	s1 := time.Now().Nanosecond()
	for i := 0; i < 1000000; i++ {
		player.ModUniqueTask.Locker.Lock()
		player.ModUniqueTask.MyTaskInfo[10001] = &game.TaskInfo{}
		player.ModUniqueTask.Locker.Unlock()
	}
	e1 := time.Now().Nanosecond() - s1
	fmt.Println("lock time", e1/1000000)

	s3 := time.Now().Nanosecond()
	for i := 0; i < 1000000; i++ {
		player.ModUniqueTask.MyTaskInfo[10001] = &game.TaskInfo{}
	}
	e3 := time.Now().Nanosecond() - s3
	fmt.Println("no lock time", e3/1000000)
}

func playerGet(player *game.Player) {
	s2 := time.Now().Nanosecond()
	for i := 0; i < 1000000; i++ {
		player.ModUniqueTask.Locker.RLock()
		_, ok := player.ModUniqueTask.MyTaskInfo[10001]
		player.ModUniqueTask.Locker.RUnlock()
		if !ok {
			// panic("no ok")
		}
	}
	e2 := time.Now().Nanosecond() - s2
	fmt.Println("rlock time", e2/1000000)

	s4 := time.Now().Nanosecond()
	for i := 0; i < 1000000; i++ {
		_, ok := player.ModUniqueTask.MyTaskInfo[10001]
		if !ok {
			panic("no ok")
		}
	}
	e4 := time.Now().Nanosecond() - s4
	fmt.Println("no rlock time", e4/1000000)
}
