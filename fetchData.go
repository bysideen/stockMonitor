package main

import (
	"fmt"
	"strconv"

	"github.com/GuNanHai/godom"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	// gbk编码类型字符bytes的解析器
	chDecoder = simplifiedchinese.GB18030.NewDecoder()
)

// genStockInfoURL : 生成不同页面的Url
func genStockInfoURL(i int) string {
	result := "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData?page=" + strconv.Itoa(i) + "&num=100&sort=changepercent&asc=0&node=hs_a&symbol=&_s_r_a=init"
	return result
}

// request url
//成功  通过通道传出网页内容	字符串
//失败	通过通道传出url 	 字符串
func fetch(url string, ch chan string) {
	//连接网址,    如果失败则传出失败的url
	resp := godom.Fetch(url, 20)

	//对网页源代码进行转码，避免中文乱码
	bodyDecoded, err := chDecoder.Bytes([]byte(resp.Raw))
	if err != nil {
		fmt.Println(err)
	}

	//将网页源代码从[]byte转换为string,然后传出去
	ch <- string(bodyDecoded)
}
