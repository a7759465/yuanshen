package game

import (
	"fmt"
)

const (
	TASK_STATE_INIT   = 0
	TASK_STATE_DOING  = 1
	TASK_STATE_FINISH = 2
)

type Player struct {
	ModPlayer     *ModPlayer
	ModIcon       *ModIcon
	ModCard       *ModCard
	ModUniqueTask *ModUniqueTask
	ModRole       *ModRole
}

func (p *Player) RecvSetIcon(iconId int) {
	p.ModPlayer.SetIcon(iconId, p)
	fmt.Println("当前图标:", p.ModPlayer.Icon)
}

func (p *Player) RecvSetCard(cardId int) {
	p.ModPlayer.SetCard(cardId, p)
	fmt.Println("当前名片:", p.ModPlayer.Card)
}

func (p *Player) RecvSetName(name string) {
	p.ModPlayer.SetName(name, p)
	fmt.Println("当前名字:", p.ModPlayer.Name)
}

func (p *Player) RecvSetSign(sign string) {
	p.ModPlayer.SetSign(sign, p)
	fmt.Println("当前签名:", p.ModPlayer.Sign)
}

func (p *Player) ReduceWorldLevel(sign string) {
	p.ModPlayer.ReduceWorldLevel()
	fmt.Println("当前签名:", p.ModPlayer.Sign)
}

func (p *Player) SetBirth(birth int) {
	p.ModPlayer.SetBirth(birth)
}

func (p *Player) SetShowCard(card []int) {
	p.ModPlayer.SetShowCard(card, p)
}

func (p *Player) SetShowTeam(showRole []int) {
	p.ModPlayer.SetShowTeam(showRole, p)
}

func (p *Player) SetHideShowTeam(isHide int) {
	p.ModPlayer.SetHideShowTeam(isHide)
}

func NewTestPlayer() *Player {
	player := &Player{}
	player.ModPlayer = &ModPlayer{}
	player.ModIcon = &ModIcon{}
	player.ModCard = &ModCard{}
	player.ModUniqueTask = &ModUniqueTask{}
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	// player.ModUniqueTask.Locker = &sync.RWMutex{}

	player.ModRole = &ModRole{}

	player.ModPlayer.PlayerLevel = 1
	player.ModPlayer.WorldLevel = 5
	player.ModPlayer.WorldLevelNow = 4
	return player
}
