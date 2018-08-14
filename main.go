package main

import (
	"net/http"
	"log"
	"github.com/sandom123/GoQaOp/users"
	"fmt"
	"html/template"
	_"strconv"
	"github.com/sandom123/GoQaOp/question"
)

//后台首页
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Println("首页")
	//// read cookie
	cookie, err := r.Cookie("uid")
	if err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
	}
	uid := cookie.Value
	if uid == ""{ //未登录
		http.Redirect(w, r, "/login/", http.StatusFound)
	}
	t , err:= template.ParseFiles("views/index.html")
	if err != nil{
		log.Fatal("parse file:", err)
	}

	t.Execute(w, nil)
}

//程序入口文件
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/login/", users.Login) //设置访问的路由
	http.HandleFunc("/clist/", question.CategoryList)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 启动静态文件服务
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
