package game

import (
	"fmt"
	"log"
	"time"
	"yuanshen/csvs"
)

type ShowRole struct {
	RoleId    int
	RoleLevel int
}

type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int
	WorldLevelNow  int
	WorldLevelCool int64
	Birth          int
	ShowTeam       []*ShowRole
	HideShowTeam   int //展示隐藏的开关
	ShowCard       []int

	//看不见的
	Prohibit int
	IsGM     int
}

func (m *ModPlayer) SetIcon(iconId int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconId) {
		//通知客户端非法
		fmt.Println("没有头像")
		return
	}
	// player.ModPlayer.Icon = iconId
	m.Icon = iconId
	fmt.Println("当前头像:", m.Icon)
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

func (m *ModPlayer) ReduceWorldLevel() {
	if m.WorldLevel < csvs.REDUCE_WORLD_LEVEL_START {
		fmt.Println("操作失败:, ---当前世界等级：", m.WorldLevel)
		return
	}

	if m.WorldLevel-m.WorldLevelNow >= csvs.REDUCE_WORLD_LEVEL_MAX {
		fmt.Println("操作失败:, ---当前世界等级：", m.WorldLevel, "---真实世界等级：", m.WorldLevelNow)
		return
	}

	if time.Now().Unix() < m.WorldLevelCool {
		fmt.Println("操作失败:, ---冷却中")
		return
	}

	m.WorldLevelNow -= 1
	m.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL_TIME
	fmt.Println("操作成功:, ---当前世界等级：", m.WorldLevel, "---真实世界等级：", m.WorldLevelNow)
}

func (m *ModPlayer) ReturnWorldLevel() {
	if m.WorldLevelNow == m.WorldLevel {
		fmt.Println("操作失败:, ---当前世界等级：", m.WorldLevel, "---真实世界等级：", m.WorldLevelNow)
		return
	}

	if time.Now().Unix() < m.WorldLevelCool {
		fmt.Println("操作失败:, ---冷却中")
		return
	}

	m.WorldLevelNow += 1
	m.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL_TIME
	fmt.Println("操作成功:, ---当前世界等级：", m.WorldLevel, "---真实世界等级：", m.WorldLevelNow)
}

func (m *ModPlayer) SetBirth(birth int) {
	if m.Birth > 0 {
		fmt.Println("已设置过生日!")
		return
	}
	month := birth / 100
	day := birth % 100

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	case 2:
		if day <= 0 || day > 29 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	default:
		fmt.Println("没有", month, "月！")
		return
	}
	m.Birth = birth
	fmt.Println("设置成功，生日为:", month, "月", day, "日")

	if m.IsBirthDay() {
		fmt.Println("今天是你的生日，生日快乐！")
	} else {
		fmt.Println("期待你生日的到来!")
	}

}

func (m *ModPlayer) IsBirthDay() bool {
	month := time.Now().Month()
	day := time.Now().Day()
	if m.Birth/100 == int(month) && m.Birth%100 == day {
		return true
	}
	return false
}

func (m *ModPlayer) SetShowCard(showCard []int, player *Player) {
	if len(showCard) > csvs.SHOW_SIZE {
		return
	}

	cardExist := make(map[int]int)
	newList := make([]int, 0)
	for _, cardId := range showCard {
		if _, ok := cardExist[cardId]; ok {
			continue
		}
		if !player.ModCard.IsHasCard(cardId) {
			continue
		}
		newList = append(newList, cardId)
		cardExist[cardId] = 1
	}
	m.ShowCard = newList
	fmt.Println("新的卡池:", m.ShowCard)
}

func (m *ModPlayer) SetShowTeam(showRole []int, player *Player) {
	if len(showRole) > csvs.SHOW_SIZE {
		return
	}
	roleExist := make(map[int]int)
	newList := make([]*ShowRole, 0)
	for _, roleId := range showRole {
		if _, ok := roleExist[roleId]; ok {
			continue
		}
		if !player.ModRole.IsHasRole(roleId) {
			continue
		}
		newList = append(newList, &ShowRole{
			RoleId:    roleId,
			RoleLevel: player.ModRole.GetRoleLevel(roleId),
		})
		fmt.Println("新的ROLEID:", roleId)

		roleExist[roleId] = 1
	}
	m.ShowTeam = newList
	fmt.Println("新的队伍:", m.ShowTeam)
}

func (p *ModPlayer) SetHideShowTeam(isHide int) {
	if isHide != csvs.LOGIC_FALSE && isHide != csvs.LOGIC_TRUE {
		return
	}
	p.HideShowTeam = isHide
}

func (p *ModPlayer) SetProhibit(prohibit int) {
	p.Prohibit = prohibit
}

func (p *ModPlayer) SetIsGM(isGM int) {
	p.IsGM = isGM
}

func (p *ModPlayer) IsCanEnter() bool {
	return int64(p.Prohibit) < time.Now().Unix()
}
