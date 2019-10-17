package main

import (
	"fmt"
	"time"

	"github.com/GuNanHai/toolkit"
)

func main() {
	// url := "https://hq.sinajs.cn/list=sh603683"

	//从网页获取JSON，fetch()传回字符串形式的JSON

	var allStock = map[string]stockInfo{}

	for {
		fetchAllStock(&allStock)

		fmt.Println("--------")
		fmt.Println("股票名： ", allStock["sz300202"].Name)
		fmt.Println("个股当前价：", allStock["sz300202"].Buy)
		fmt.Println(len(allStock))

		time.Sleep(2 * time.Second)
	}

}
func stkLookup(symbol string, stkList stockInfoList) stockInfo {
	for _, stk := range stkList {
		if stk.Symbol == symbol {
			return stk
		}
	}
	return stockInfo{}
}
func fetchAllStock(stks *map[string]stockInfo) {

	chFetch := make(chan string)

	for i := 1; i < 40; i++ {
		url := genStockInfoURL(i)
		go fetch(url, chFetch)
	}

	for i := 1; i < 40; i++ {
		page := <-chFetch

		if len(page) < 10 {
			continue
		}

		// fmt.Println(page)
		var stkInfoListTemp = stockInfoList{}
		toolkit.JSONToStruct(page, &stkInfoListTemp)
		for _, stkInfo := range stkInfoListTemp {
			(*stks)[stkInfo.Symbol] = stkInfo
		}
	}

}
