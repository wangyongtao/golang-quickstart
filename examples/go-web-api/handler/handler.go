package handler

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var err error

// Handler : 结构体
type Handler struct {
	MyDB *sql.DB
}

// IndexHandler 首页
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("----> IndexHandler ")
	// 解析指定文件生成模板对象
	tpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Println("读取模板文件失败, err: ", err)
		return
	}
	// 渲染模板
	data := map[string]string{
		"Title": "首页 login",
	}
	w.Header().Set("content-type", "text/html")
	tpl.Execute(w, data)
}

// HelloHandler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HelloHandler Hit!")

	// 返回json
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "200",
		"message": "ok",
		"data":    "hello",
	})
}

// HealthCheckHandler
func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("----> HealthCheckHandler: health-check ")
	// log.Println(c.ParamNames())
	// log.Println(c.ParamValues())

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "200",
		"message": "ok",
	})
}
