package main

import (
	"fmt"
	"time"

	"github.com/GuNanHai/toolkit"
)

func main() {
	// url := "https://hq.sinajs.cn/list=sh603683"

	//从网页获取JSON，fetch()传回字符串形式的JSON

	var allStock stockInfoList

	for {
		allStock = []stockInfo{}
		fetchAllStock(&allStock)

		stock := stkLookup("sz300202", allStock)
		fmt.Println("--------")
		fmt.Println("股票名： ", stock.Name)
		fmt.Println("个股当前价：", stock.Buy)
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
func fetchAllStock(stks *stockInfoList) {
	chFetch := make(chan string)

	for i := 1; i < 40; i++ {
		url := genStockInfoURL(i)
		go fetch(url, chFetch)
	}

	var stkInfoListTemp stockInfoList

	for i := 1; i < 40; i++ {
		page := <-chFetch

		if len(page) < 10 {
			continue
		}

		// fmt.Println(page)
		stkInfoListTemp = []stockInfo{}
		toolkit.JSONToStruct(page, &stkInfoListTemp)

		*stks = append(*stks, stkInfoListTemp...)
	}

}
