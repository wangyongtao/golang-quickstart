package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

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
func UserDetailHanlder(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world1!\n")
}

// IndexHanlder : 用户详情
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

func main() {
	// http.HandleFunc("/", IndexHanlder)
	// http.HandleFunc("/hello", HelloHanlder)
	// http.HandleFunc("/user", UserListHanlder)
	// http.HandleFunc("/user/:id", UserDetailHanlder)

	http.Handle("/", hook(loggingHandler(http.HandlerFunc(IndexHanlder))))
	http.HandleFunc("/hello", HelloHanlder)
	http.HandleFunc("/user", UserListHanlder)
	http.HandleFunc("/user/:id", UserDetailHanlder)

	port := ":8082" // 监听端口号

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// log.Println("ListenAndServe: 访问地址: http://localhost: " + port)
}
