README.md


```
$ mkdir go-mod-demo 
$ cd go-mod-demo

$ ll
total 16
-rw-r--r--  1 wangtom  staff     9B 12  2 11:09 README.md
-rw-r--r--  1 wangtom  staff   192B 12  2 11:08 main.go 
```


$ go mod init go-mod-demo
go: creating new go.mod: module go-mod-demo

```
$ cat go.mod 
module go-mod-demo

go 1.13
```


go list -m输出的信息被称为build list，也就是构建当前module所要构建的所有相关package（及版本）的列表。

```
$ go list -m -json all
{
	"Path": "go-mod-demo",
	"Main": true,
	"Dir": "/Users/wangtom/Code/golang-quick-start/examples/go-mod-demo",
	"GoMod": "/Users/wangtom/Code/golang-quick-start/examples/go-mod-demo/go.mod",
	"GoVersion": "1.13"
}
```


```sh
$ go get -u github.com/joho/godotenv
$ ll
total 32
-rw-r--r--  1 wangtom  staff   695B 12  2 11:19 README.md
-rw-r--r--  1 wangtom  staff    81B 12  2 11:21 go.mod
-rw-r--r--  1 wangtom  staff   167B 12  2 11:21 go.sum
-rw-r--r--  1 wangtom  staff   192B 12  2 11:08 main.go
```


```sh
$ cat go.mod 
module go-mod-demo

go 1.13

require github.com/joho/godotenv v1.3.0 // indirect
```
➜  go-mod-demo git:(master) ✗ 


使用 go mod 的依赖包会被安装在 `~/go/pkg/mod/` 命令下：

```sh
$ ll ~/go/pkg/mod/github.com/joho/godotenv@v1.3.0/
total 80
-r--r--r--  1 wangtom  staff   1.0K 12  2 09:40 LICENCE
-r--r--r--  1 wangtom  staff   4.8K 12  2 09:40 README.md
dr-xr-xr-x  3 wangtom  staff    96B 12  2 09:40 autoload
dr-xr-xr-x  3 wangtom  staff    96B 12  2 09:40 cmd
dr-xr-xr-x  8 wangtom  staff   256B 12  2 09:40 fixtures
-r--r--r--  1 wangtom  staff   8.6K 12  2 09:40 godotenv.go
-r--r--r--  1 wangtom  staff    12K 12  2 09:40 godotenv_test.go
```



vi .env

APP_NAME=Go-Mod-DEMO
APP_DEBUG=true
APP_VERSION=0.0.1


