package main

type TabelogResult struct {
	Address string
	URL     string
	Name    string
}

func GetAddressFromTabelogURL(url string) TabelogResult {
	return TabelogResult{
		Address: "東京都渋谷区恵比寿南1-16-12 ＡＢＣ・ＭＡＭＩＥＳ　３Ｆ",
		URL:     url,
		Name:    "レストラン名",
	}
}
