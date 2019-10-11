package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	// gbk编码类型字符bytes的解析器
	chDecoder = simplifiedchinese.GB18030.NewDecoder()
)

// request url
//成功  通过通道传出网页内容	字符串
//失败	通过通道传出url 	 字符串
func fetch(url string, ch chan string) {
	//连接网址,    如果失败则传出失败的url
	resp, err := http.Get(url)
	if err != nil {
		ch <- url
		fmt.Println(url, " 网页内容获取失败: ", err)
		return
	}

	//读取网页源代码
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	//对网页源代码进行转码，避免中文乱码
	bodyDecoded, err := chDecoder.Bytes(body)
	if err != nil {
		fmt.Println(err)
	}

	//将网页源代码从[]byte转换为string,然后传出去
	ch <- string(bodyDecoded)
	fmt.Println(url, " 网页内容获取成功。")
}
