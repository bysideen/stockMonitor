package main

import (
	"fmt"
)

type stockInfo struct {
	Symbol        string //股票标记
	Code          string //股票代码
	Name          string //股票名称
	Trade         string //买一 （/元）
	Pricechange   string //涨跌额（/元）
	Changepercent string //涨跌幅（百分比）
	Buy           string //最新价（/元）
	Sell          string //卖一	（/元）
	Settlement    string //昨天收盘价（/元）
	Open          string //开盘价（/元）
	High          string //今日最高价（/元）
	Low           string //今日最低价（/元）
	Volume        int    //成交量（/股）
	Amount        int    //成交额（/元）
	Ticktime      string //数据记录时间
	Per           float32
	Pb            float32 //市净率
	Mktcap        float32 //总市值（/万元）
	Nmc           float32 //流通值（/万元）
	Turnoverratio float32 //换手率（百分号）
}
type stockInfoList []stockInfo

func (stks stockInfoList) Print() {
	fmt.Printf("%+v", stks)
}

type stockInfoDetail struct {
	Name       string //股票名称
	Open       string //开盘价（/元）
	Settlement string //昨天收盘价（/元）
	Buy        string //最新价（/元）
	High       string //今日最高价（/元）
	Low        string //今日最低价（/元）
	BuyIn      string //竞买价，即“买一”报价
	SellOut    string //竞卖价，即“卖一”报价
	Volume     string //成交量（/股）
	Amount     string //成交额（/元）
	BuyVol1    string //买一股数
	Buy1       string //买一价格
	BuyVol2    string //买二股数
	Buy2       string //买二价格
	BuyVol3    string //买三股数
	Buy3       string //买三价格
	BuyVol4    string //买四股数
	Buy4       string //买四价格
	BuyVol5    string //买五股数
	Buy5       string //买五价格

	SellVol1 string //卖一股数
	Sell1    string //卖一价格
	SellVol2 string //卖二股数
	Sell2    string //卖二价格
	SellVol3 string //卖三股数
	Sell3    string //卖三价格
	SellVol4 string //卖四股数
	Sell4    string //卖四价格
	SellVol5 string //卖五股数
	Sell5    string //卖五价格
	Date     string //数据记录日期
	Ticktime string //数据记录时间
}
