package main

import (
	"fmt"

	"github.com/GuNanHai/toolkit"
)

func main() {
	// url := "https://hq.sinajs.cn/list=sh603683"
	url := "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData?page=1&num=40&sort=changepercent&asc=0&node=hs_a&symbol=&_s_r_a=init"

	//从网页获取JSON，fetch()传回字符串形式的JSON
	chFetch := make(chan string)
	go fetch(url, chFetch)

	//将JSON 转变为  struct数组类，存入allStock
	var allStock stockInfoList
	toolkit.JSONToStruct(<-chFetch, &allStock)
	fmt.Println(allStock)

}
