package csvs

import "fmt"

type ConfigBanWord struct {
	Id  int
	Txt string
}

var ConfigBanWordSlice []*ConfigBanWord

func init() {

	// Loadcsv(ConfigBanWordSlice, 'banword csv')

	ConfigBanWordSlice = append(ConfigBanWordSlice,
		&ConfigBanWord{Id: 1, Txt: "外挂"},
		&ConfigBanWord{Id: 1, Txt: "原神"},
		&ConfigBanWord{Id: 1, Txt: "工具"},
		&ConfigBanWord{Id: 1, Txt: "辅助"},
		&ConfigBanWord{Id: 1, Txt: "卧槽"},
	)

	fmt.Println("csc_banword init...")
}

func GetBanWordBase() []string {
	relString := make([]string, 0)
	for _, v := range ConfigBanWordSlice {
		relString = append(relString, v.Txt)
	}
	return relString
}

func GetBanWordExtra() []string {
	relString := make([]string, 0)
	for _, v := range ConfigBanWordSlice {
		relString = append(relString, v.Txt)
	}
	return relString
}
