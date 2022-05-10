package main

import (
	"fmt"

	"yuanshen/game"

	"yuanshen/csvs"
)

func main() {
	fmt.Println("======================原神服务器======================")

	//加载配置
	csvs.CheckLoadCsv()
	go game.GetManageBanWord().Run()
	fmt.Println("======================数据测试======================")

	player := game.NewTestPlayer()
	// player.ModBag.AddItem(1000003)
	// player.ModBag.AddItem(1000005)
	// player.ModBag.AddItem(1000006)
	// player.ModBag.AddItem(4000004, player)
	// player.ModBag.AddItem(2000001)
	// player.ModBag.AddItem(2000002)

	// player.ModIcon.AddItem(3000001)
	// player.ModIcon.AddItem(3000002)
	// player.ModIcon.AddItem(3000003)
	// player.ModIcon.AddItem(3000004)
	// player.ModIcon.AddItem(4000004)

	// player.ModPlayer.SetIcon(3000004, player)

	// player.ModBag.AddItem()

	player.ModBag.RemoveItem(1000004, 3)
	player.ModBag.RemoveItem(1000005, 3)
	player.ModBag.RemoveItem(1000004, 3)
	// ticker := time.NewTicker(time.Second * 10)
	// for range ticker.C {
	// 	player := game.NewTestPlayer()
	// 	go player.Run()

	// }

}
