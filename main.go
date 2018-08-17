package main

import (
	"net/http"
	"log"
	"github.com/sandom123/GoQaOp/users"
	"html/template"
	_"strconv"
	"github.com/sandom123/GoQaOp/question"
	"runtime"
	"os"
	"strconv"
	"github.com/sandom123/GoQaOp/util"
	"time"
	"github.com/sandom123/GoQaOp/category"
)

type Data struct {
	Action string
}

func Index(w http.ResponseWriter, r *http.Request){
	_, username := util.GetLoginUid(w, r)
	t , err:= template.ParseFiles(
		"views/welcome.html")
	sysinfo := make(map[string]string)
	sysinfo["goVersion"] = runtime.Version() //go语言版本
	sysinfo["ip"] = util.GetIntranetIp()
	sysinfo["goos"] = runtime.GOOS //操作系统名称
	sysinfo["goHost"] = runtime.GOARCH //操作系统架构
	sysinfo["hostName"] = ""
	name, _ := os.Hostname()
	sysinfo["hostName"] = name //获取本机名称
	sysinfo["cpuCount"] = strconv.Itoa(runtime.GOMAXPROCS(0)) // 获取本机CPU个数
	sysinfo["username"] = username
	sysinfo["currentTime"] = time.Now().Format("2006-01-02 15:04:05")
	if err != nil{
		log.Fatal("parse file:", err)
	}
	t.Execute(w, sysinfo)
}
//后台首页
func Main(w http.ResponseWriter, r *http.Request){
	_, username := util.GetLoginUid(w, r)
	t , err:= template.ParseFiles(
		"views/public/layout.html",
		"views/public/header.html",
		"views/public/lefter.html",
		"views/public/footer.html",
		)
	if err != nil{
		log.Fatal("parse file:", err)
	}
	data := make(map[string]string)
	data["username"] = username
	t.Execute(w, data)
}

//程序入口文件
func main() {
	http.HandleFunc("/", Main)
	http.HandleFunc("/index/", Index)
	http.HandleFunc("/login/", users.Login) //设置访问的路由
	http.HandleFunc("/clist/", category.List)
	http.HandleFunc("/cadd/" , category.Add)

	http.HandleFunc("/qalist/", question.QaList)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 启动静态文件服务
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
