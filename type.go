package main

import (
	"fmt"
)

type stockInfoList []stockInfo

func (stks stockInfoList) Print() {
	fmt.Printf("%+v", stks)
}
