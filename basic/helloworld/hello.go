package main

import "fmt"

func main() {
    // 我的hello world程序
    fmt.Println("Hello, 世界!");

    // fmt.Println 支持多个参数 (每个参数会自动多个空格)
    fmt.Println("Hello", ",", "多个参数", "!");

    // 字符串拼接 
    fmt.Println("Hello" + "," + "字符串拼接" + "!");
}

// 输出示例：
// $ go run helloworld/hello.go 
// Hello, 世界!
// Hello , 多个参数 !
// Hello,字符串拼接!
