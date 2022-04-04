package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// encoding/json
// Go语言对json的处理可以用系统自带的 encoding/json 包
// 执行代码: go run json.go

//定义一个简单的结构体 Person
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Status     int    `json:"status"`
	CreateTime int    `json:"create_time"`
}

func main() {
	fmt.Println("---- json -----")

	var user User

	user.ID = 1001
	user.Name = "兔兔"
	user.Email = "tu@a.com"
	user.Status = 1
	user.CreateTime = int(time.Now().Unix())

	fmt.Println("---- user: ", user)

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("---- 编码： 结构体生成 json 出错了： ", err.Error())
		return
	}
	// json.Marshal 返回的是 byte 数组，使用string()转成字符串
	fmt.Println("---- 编码： 结构体生成 json ok => ", string(data))
}
