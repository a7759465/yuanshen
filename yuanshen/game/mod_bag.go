package game

import (
	"fmt"
	"yuanshen/csvs"
)

type ItemInfo struct {
	ItemId  int
	ItemNum int64
}

type ModBag struct {
	BagInfo map[int]*ItemInfo
}

func (m *ModBag) AddItem(itemId int, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println("物品不存在")
		return
	}
	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		m.AddItemToBag(itemId, 1)
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("角色:", itemConfig.ItemName)
	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 9)
	case csvs.ITEMTYPE_WEAPON:
		fmt.Println("武器:", itemConfig.ItemName)
	case csvs.ITEMTYPE_RELICS:
		fmt.Println("圣物:", itemConfig.ItemName)
	case csvs.ITEMTYPE_COOK:
		fmt.Println("厨艺:", itemConfig.ItemName)
	case csvs.ITEMTYPE_HOME_ITEM:
		fmt.Println("家园物品:", itemConfig.ItemName)
	default: //同普通
		m.AddItemToBag(itemId, 1)
	}
}

func (m *ModBag) AddItemToBag(itemId int, num int64) {
	_, ok := m.BagInfo[itemId]
	if ok {
		m.BagInfo[itemId].ItemNum += num
	} else {
		m.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: num}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("获得物品", config.ItemName, "----数量：", num, "----当前数量：", m.BagInfo[itemId].ItemNum)
	}
}

func (m *ModBag) RemoveItem(itemId int, num int64) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		m.RemoveItemToBagGM(itemId, num)
	default: //同普通
		//m.AddItemToBag(itemId, 1)
	}
}

func (m *ModBag) RemoveItemToBagGM(itemId int, num int64) {
	_, ok := m.BagInfo[itemId]
	if ok {
		m.BagInfo[itemId].ItemNum -= num
	} else {
		m.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: 0 - num}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("扣除物品", config.ItemName, "----数量：", num, "----当前数量：", m.BagInfo[itemId].ItemNum)
	}
}
