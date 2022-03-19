Golang 学习笔记

## 关于 go 语言

Go 语言(即 Golang ），是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。

Go的语法接近C语言，但对于变量的声明有所不同。 罗伯特·格瑞史莫（Robert Griesemer），罗勃·派克（Rob Pike）及肯·汤普逊（Ken Thompson）于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。

在 2016 年，Go 被软件评价公司 TIOBE 选为“TIOBE 2016 年最佳语言”。

## Go 语言的版本发布历史

```sh
go1.18 (released 2022-03-15) 增加泛型 
go1.17 (released 2021-08-16) 
go1.16 (released 2021-02-16) 
go1.15 (released 2020-08-11)
go1.14 (released 2020-02-25)
go1.13 (released 2019/09/03)
go1.12 (released 2019/02/25)  
go1.11 (released 2018/08/24) 新增 go mod 模块管理（Modules）  
go1.10 (released 2018/02/16)  
go1.9  (released 2017/08/24)  
go1.8  (released 2017/02/16)  
go1.7  (released 2016/08/15)  
go1.6  (released 2016/02/17)  
go1.5  (released 2015/08/19)  
go1.4  (released 2014/12/10)  
go1.3  (released 2014/06/18)  
go1.2  (released 2013/12/01)  
go1.1  (released 2013/05/13)  
go1    (released 2012/03/28)  
```

更多详细的发布历史详见[Golang Release History](https://golang.google.cn/doc/devel/release.html)。

## 安装 go 环境

最简单的方式是，直接去官网下载安装包: 

安装包下载地址为：`https://golang.org/dl`  

如果打不开可以使用这个国内镜像地址：`https://golang.google.cn/dl` 

下载对应系统的安装包即可：

```sh
(1) Microsoft Windows
Windows 7 or later, Intel 64-bit processor
go1.18.windows-amd64.msi (132MB)

(2) Apple macOS
macOS 10.11 or later, Intel 64-bit processor
go1.18.darwin-amd64.pkg (138MB)

(3) Linux
Linux 2.6.23 or later, Intel 64-bit processor
go1.18.linux-amd64.tar.gz (135MB)

(4) Source 
go1.18.src.tar.gz (22MB)
```

如果安装好 go 语言后，可以使用 `go version` 查看当前版本

```sh
$ go version  
go version go1.18 darwin/amd64
```

可以看到，目前安装的是 go1.18 版本。

查看go命令的路径:

```sh
$ which go
/usr/local/go/bin/go
```

## 第一个 HelloWord 程序

创建一个名为 golang-quick-start 的目录, 在目录中创建一个 `main.go` 文件: 

```sh
$ mkdir golang-quick-start
$ cd 
$ vi main.go
```

代码文件 `main.go` 内容如下: 

```java
// golang-quick-start/main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!");
}
```

## 编译并运行代码

(1) 直接使用 go run 命令运行 main.go 程序:

```sh
$ go run main.go 
Hello, World!
```

(2) 使用 go build 构建代码，并执行编译后的可执行文件:

```sh
$ go build main.go
$ ./main 
Hello, World!
```

[END]
