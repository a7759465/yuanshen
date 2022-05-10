package game

import (
	"fmt"
	"yuanshen/csvs"
)

type Icon struct {
	IconId int
}

type ModIcon struct {
	IconInfo map[int]*Icon
}

func (m *ModIcon) IsHasIcon(iconId int) bool {
	_, ok := m.IconInfo[iconId]
	return ok
}

func (m *ModIcon) AddItem(itemId int) {
	if _, ok := m.IconInfo[itemId]; ok {
		fmt.Println("已存在头像")
		return
	}
	itemConfig := csvs.GetIconConfig(itemId)
	if itemConfig == nil {
		fmt.Println("非法头像:", itemId)
		return
	}
	m.IconInfo[itemId] = &Icon{IconId: itemId}
	fmt.Println("获得头像:", itemId)
}
