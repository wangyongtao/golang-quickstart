package main

import (
	"fmt"
	"go-mod-demo/util"
)

func main() {
	fmt.Println("hello, golang dev docs!")

	// 测试调用 util 包
	curr := util.UnixTime()

	fmt.Println("util time => ", curr)
	fmt.Println("util datetime => ", util.UnixTime2DateTime(curr))
}
