package game

import (
	"fmt"
	"regexp"
	"yuanshen/csvs"
)

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string
	BanWordExtra []string
}

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = &ManageBanWord{}
		// manageBanWord.BanWordBase = []string{"外挂", "工具"}
		// manageBanWord.BanWordExtra = []string{"原神"}
	}
	return manageBanWord
}

func (m *ManageBanWord) Run() {

	m.BanWordBase = csvs.GetBanWordBase()

	// ticker := time.NewTicker(time.Second * 1)
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		fmt.Println("time now:", time.Now(), ", time unix:", time.Now().Unix())
	// 	}
	// }
}

func (m *ManageBanWord) IsBanWOrd(txt string) bool {
	for _, v := range m.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		if match {
			fmt.Println(v, "是敏感词")
			return match
		}
	}
	for _, v := range m.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		if match {
			fmt.Println(v, "是敏感词")
			return match
		}
	}
	return false
}
