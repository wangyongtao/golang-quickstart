*golang-quick-start*

## 关于 go 语言

Go 语言(即 Golang ），是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。

Go的语法接近C语言，但对于变量的声明有所不同。 罗伯特·格瑞史莫（Robert Griesemer），罗勃·派克（Rob Pike）及肯·汤普逊（Ken Thompson）于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。

在 2016 年，Go 被软件评价公司 TIOBE 选为“TIOBE 2016 年最佳语言”。

## Go 语言的版本发布历史

```sh
go1.16 (released 2021/02/16)
go1.15 (released 2020/08/11)
go1.14 (released 2020/02/25)
go1.13 (released 2019/09/03)
go1.12 (released 2019/02/25)  
go1.11 (released 2018/08/24)  
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
go1.16.windows-amd64.msi (119MB)

(2) Apple macOS
macOS 10.12 or later, Intel 64-bit processor
go1.16.darwin-amd64.pkg (124MB)

(3) Linux
Linux 2.6.23 or later, Intel 64-bit processor
go1.16.linux-amd64.tar.gz (123MB)

(4) Source 
go1.16.src.tar.gz (20MB)
```

查看go命令的路径:

```sh
$ which go
/usr/local/go/bin/go
```

如果安装好 go 语言后，可以使用 `go version` 查看当前版本

```sh
$ go version  
go version go1.16 darwin/amd64
```

可以看到，目前安装的是 `go1.16` 版本。


如果是 macOS 系统， 且已经安装了 `homebrew` 可以使用以下命令安装： 

```sh
$ brew install go
```

通过 `brew` 命令，go 会被安装在 `/usr/local/Cellar/go/` 目录下: 

```sh
$ /usr/local/Cellar/go/1.15.8/bin/go version
go version go1.15.8 darwin/amd64
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

## 查看 go 系统变量

使用 `go env` 可以查看全部的 go 环境变量：

```sh
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/wang/Library/Caches/go-build"
GOENV="/Users/wang/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/wang/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/wang/go"
GOPRIVATE=""
GOPROXY="https://goproxy.cn,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GOVCS=""
GOVERSION="go1.16"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch x86_64 -m64 ..."
```

使用 `go env 变量名` 可以查看单个的环境变量，比如查看 `GOPATH`： 

```sh
$ go env GOPATH
/Users/wangtom/go
```
 
*设置代理*

[Goproxy China 代理](https://github.com/goproxy/goproxy.cn)

在终端执行：（Go 1.13 及以上）

```sh
$ go env -w GOPROXY=https://goproxy.cn,direct
```

关于 GOPATH： 

The `GOPATH` environment variable specifies the location of your workspace. 

It defaults to a directory named `go`inside your home directory, so `$HOME/go` on Unix, `$home/go` on Plan 9, and `%USERPROFILE%\go` (usually `C:\Users\YourName\go`) on Windows.


```sh
$ tree -L 3 ~/go 
/Users/wangtom/go
├── bin
│   ├── go-outline
│   ├── goimports
│   └── gopls
├── pkg
│   ├── darwin_amd64
│   │   └── github.com
│   ├── mod
│   │   ├── cache
│   │   ├── github.com
│   │   ├── golang.org
│   │   ├── google.golang.org
│   │   ├── gopkg.in
│   │   └── honnef.co
│   └── sumdb
│       └── sum.golang.org
└── src
    └── github.com
        └── go-sql-driver
```


## 创建 go modules 模块: 

我们知道 `go get` 获取的代码会放在 `$GOPATH/src` 下面，而`go build`会在`$GOROOT/src`和`$GOPATH/src`下面按照import path去搜索`package`。

每个 go.mod 文件定义了一个 module ，而放置 go.mod 文件的目录被称为 module root 目录（通常对应一个 repo 的 root 目录，但不是必须的）。
module root 目录以及其子目录下的所有 Go package 均归属于该 module，除了那些自身包含 go.mod 文件的子目录。


大多数编程语言都会有包管理工具，像 Node 有`npm`，PHP 有`composer`，Java 有`Maven`和`Gradle`。 
可是，Go 语言一直缺乏一个官方的包管理(曾有一个`Dep`被称为官方试验品`official experiment`)。
终于，在`go1.11` 版本中，新增了`module`管理模块功能，用来管理依赖包。需要将`GO111MODULE`设置为`on`。自 go1.16 起默认 `GO111MODULE=on`。

```sh
$ go mod init github.com/cnwyt/go-quick-start
go: creating new go.mod: module github.com/cnwyt/go-quick-start
go: to add module requirements and sums:
    go mod tidy
```

会生成一个 go.mod 文件，指明了模块名和 go 的版本: 

```sh
module github.com/cnwyt/go-quick-start
go 1.16
```

## 引入一个第三方包 uniplaces/carbon: 

可以使用 go get 直接下载包: 

```sh
$ go get github.com/uniplaces/carbon
```

或者使用 go mod edit 命令编辑mod文件，使用 go mod download 来下载包: 

```sh
$ go mod edit -require github.com/uniplaces/carbon@latest
$ go mod download
```

注意使用 go mod edit 必须指定版本 `path@version`，如 `xxx@latest` 或 `xxx@v0.1.6`， 否则会提示如下错误: 

```sh
$ go mod edit -require github.com/uniplaces/carbon
$ go mod: -require=github.com/uniplaces/carbon: need path@version
```

查看 go.mod 文件: 

```sh
module github.com/cnwyt/go-quick-start

go 1.16

require github.com/uniplaces/carbon latest // indirect
```

可以看到 多了一行 require 语句，指定第三方包的 url路径 和 版本号。

在 main.go 中使用 import 来引入包: 

```java
package main

import "fmt"
import "github.com/uniplaces/carbon"

func main() {
    fmt.Println("Hello, World!");
    fmt.Printf("时间戳 Unix timestamp:  %d \n", time.Now().Unix())
    fmt.Printf("当前时间 Right now is:  %s \n", carbon.Now().DateTimeString())
}
```

使用 go run 命令编译并执行 main.go 文件: 

```sh
$ go run main.go 
Hello, World!
时间戳 Unix timestamp:  1614132935 
当前时间 Right now is:  2021-02-24 10:15:35
```

## 使用第三方包操作 MySQL 数据库 

再引入一个第三方包 go-sql-driver/mysql:

```sh
$ go mod edit -require github.com/go-sql-driver/mysql@latest
$ go mod download  
```

查看 go.mod 文件，require 部分已经指定了 go-sql-driver/mysql 包，多个包默认是用圆括号包起来的。

```java
module github.com/cnwyt/go-quick-start

go 1.16

require (
    github.com/go-sql-driver/mysql v1.5.0
    github.com/uniplaces/carbon v0.1.6 // indirect
)
```

修改 main.go 代码:

```java
package main

import (
	"fmt"
	"time"
	"log"
	"github.com/uniplaces/carbon"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Hello, World!");

    fmt.Printf("时间戳 Unix timestamp:  %d \n", time.Now().Unix())
    fmt.Printf("当前时间 Right now is:  %s \n", carbon.Now().DateTimeString())

    // 连接数据库
	db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/test")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(3 * time.Second)
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Fatal("----> 数据库连接失败.")
	    panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("----> 数据库连接成功.");
	}
	defer db.Close() 
}
```

运行代码，结果如下: 

```sh
$ go run main.go
Hello, World!
时间戳 Unix timestamp:  1614157214 
当前时间 Right now is:  2021-02-24 17:00:14 
----> 数据库连接成功.
```


[END]
