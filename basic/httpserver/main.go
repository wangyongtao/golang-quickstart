package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 指定htpp端口号
const httpPort string = ":8001"

// User 用户结构体
type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CreateTime int    `json:"create_time"`
}

// JSONResult 返回结果结构体
type JSONResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []any  `json:"data"`
}

// IndexHandler 首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("----> IndexHandler ")
	// 解析指定文件生成模板对象
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("读取模板文件失败, err: ", err)
		return
	}
	// 渲染模板
	data := map[string]string{
		"Title": "首页 home",
	}
	w.Header().Set("content-type", "text/html")
	tpl.Execute(w, data)
}

// AboutHanlder : 输出字符串
func AboutHanlder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "text/html")

	io.WriteString(w, "<meta charset='utf-8'/>") // 指定html编码
	io.WriteString(w, "<p>AboutHanlder! <p/>")
	io.WriteString(w, "<p>hello, world! <p/>")
	io.WriteString(w, "<a href='/'>返回首页</a>")
}

// UserListHanlder : handler "/api/users"
func UserListHanlder(w http.ResponseWriter, req *http.Request) {
	fmt.Println("---> UserListHanlder hit")

	// 获取参数
	limit := req.FormValue("limit")
	fmt.Println("---> parame limit : ", limit)

	ret := new(JSONResult)
	ret.Code = 0
	ret.Message = "success"
	// 追加一条数据
	ret.Data = append(ret.Data, &User{
		Id:         10001,
		Name:       "zhang3",
		CreateTime: int(time.Now().Unix()),
	})
	// 追加一条数据
	ret.Data = append(ret.Data, &User{
		Id:         10002,
		Name:       "li4",
		CreateTime: int(time.Now().Unix()),
	})
	retJSON, _ := json.Marshal(ret)

	w.Header().Set("content-type", "text/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(retJSON))
	return
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
	id := r.FormValue("id")
	if len(id) <= 0 {
		log.Println("param error: id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uid, _ := strconv.Atoi(id)
	// 根据用户 id 去数据查询用户数据
	u := &User{
		Id:         uid,
		Name:       "用户名" + id,
		CreateTime: int(time.Now().Unix()),
	}
	w.Header().Set("content-type", "application/json")
	// io.WriteString(w, string(retJSON))
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "ok",
		"data":    u,
	})
	return
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

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHanlder)
	http.HandleFunc("/api/users", UserListHanlder)
	http.HandleFunc("/api/users/detail", UserDetailHanlder)
	http.HandleFunc("/api/users/:id", UserDetailHanlder) // <---- 默认不支持path参数

	// 启动服务
	log.Println("ListenAndServe ok. 访问地址 http://localhost" + httpPort)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
