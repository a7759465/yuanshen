package game

import (
	"fmt"
	"log"
	"yuanshen/csvs"
)

type ModPlayer struct {
	UserId        int
	Icon          int
	Card          int
	Name          string
	Sign          string
	PlayerLevel   int
	PlayerExp     int
	WordLevel     int
	WordLevelCool int64
	Birth         int
	ShowTeam      []int
	ShowCard      []int

	//看不见的
	IsProhibit int
	IsGM       int
}

func (m *ModPlayer) SetIcon(iconId int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconId) {
		//通知客户端非法
		return
	}
	// player.ModPlayer.Icon = iconId
	m.Icon = iconId
}

func (m *ModPlayer) SetCard(cardId int, player *Player) {
	if !player.ModCard.IsHasCard(cardId) {
		//通知客户端非法
		return
	}
	// player.ModPlayer.Card = cardId
	m.Card = cardId
}

func (m *ModPlayer) SetName(name string, player *Player) {
	//http接口验证 非法名字
	//本地敏感词词库
	if GetManageBanWord().IsBanWOrd(name) {
		return
	}
	// player.ModPlayer.Name = name
	m.Name = name
}

func (m *ModPlayer) SetSign(sign string, player *Player) {
	//http接口验证 非法名字
	//本地敏感词词库
	if GetManageBanWord().IsBanWOrd(sign) {
		return
	}
	player.ModPlayer.Sign = sign
}

func (m *ModPlayer) AddExp(exp int, player *Player) {
	m.PlayerExp += exp

	for {
		config := csvs.GetNowLevelConfig(m.PlayerLevel)
		if config == nil {
			log.Fatal("AddExp panic")
			break
		}
		if config.PlayerExp == 0 { //满级
			break
		}
		if m.PlayerExp < config.PlayerExp { //没够
			break
		}
		//是否完成任务
		if config.ChapterId > 0 && !player.ModUniqueTask.IsTaskFinish(config.ChapterId) {
			break
		}
		m.PlayerLevel += 1
		m.PlayerExp -= config.PlayerExp
	}
	fmt.Println("现在等级:", m.PlayerLevel, ",现在经验:", m.PlayerExp)
}
