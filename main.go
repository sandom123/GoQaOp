package main

import (
	"net/http"
	"log"
	"github.com/sandom123/GoQaOp/users"
)

//程序入口文件

func main() {

	http.HandleFunc("/login", users.Login) //设置访问的路由
	//http.Handle("/static/", http.FileServer(http.Dir("/static/")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 启动静态文件服务
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
