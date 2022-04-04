package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	handler "go-web-api/handler"
)

// 指定htpp端口号
const httpPort string = ":8001"

// User 用户结构体
type User struct {
	Name string
	Age  int
}
type result struct {
	Code  int
	Param string
	Msg   string
	Data  []User
}

// HelloHanlder : handler "/hello"
func HelloHanlder(w http.ResponseWriter, req *http.Request) {
	data := User{Name: "why", Age: 18}

	ret := new(result)
	id := req.FormValue("id")
	//id := req.PostFormValue('id')

	ret.Code = 0
	ret.Param = id
	ret.Msg = "success"
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	retJSON, _ := json.Marshal(ret)

	io.WriteString(w, string(retJSON))
}

// UserListHanlder : 用户列表
func UserListHanlder(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world1!\n")
}

// UserDetailHanlder : 用户详情
func UserDetailHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---> UserDetailHanlder hit")
	// 解析form表单
	err := r.ParseForm()
	if err != nil {
		log.Println("could not parse query: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 获取表单参数
	name := r.FormValue("name")

	io.WriteString(w, "UserDetailHanlder\n")
	io.WriteString(w, name)
}

// IndexHanlder : 首页
func IndexHanlder(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `
	<a href='/'>index</a> 
	<a href='/hello'>hello</a> 
	<a href='/user'>users</a> 
	<a href='/user/111'>user detail (id=111)</a> 
	`)
}

// http中间件: 实现一个函数签名为func(http.Handler) http.Handler的函数即可
// func middlewareHandler(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// 执行handler之前的逻辑
// 		next.ServeHTTP(w, r)
// 		// 执行完毕handler后的逻辑
// 	})
// }
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Comleted %s in %v", r.URL.Path, time.Since(start))
	})
}
func hook(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("---------------------")
		log.Println("before hook")
		next.ServeHTTP(w, r)
		log.Println("after hook")

	})
}
func init() {
	log.Println("----- init -----")
}

func main() {
	log.Println("----- main -----")

	h := &handler.Handler{}
	// http.HandleFunc("/", IndexHanlder)
	// http.HandleFunc("/hello", HelloHanlder)
	// http.HandleFunc("/user", UserListHanlder)
	// http.HandleFunc("/user/:id", UserDetailHanlder)

	// http.Handle("/", hook(loggingHandler(http.HandlerFunc(IndexHanlder))))
	http.HandleFunc("/", h.IndexHandler)
	http.HandleFunc("/hello", HelloHanlder)
	http.HandleFunc("/user", UserListHanlder)
	http.HandleFunc("/user/:id", UserDetailHanlder)

	err := http.ListenAndServe(httpPort, nil)
	log.Println("ListenAndServe ok: 访问地址: http://localhost: " + httpPort)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
