package category

import (
	"net/http"
	"html/template"
	"log"
	"github.com/sandom123/GoQaOp/util"
	"fmt"
	"github.com/sandom123/GoQaOp/models"
	"strconv"
)

//分类列表
func List(w http.ResponseWriter, r *http.Request){
	util.GetLoginUid(w, r)
	//fmt.Fprint(w, "分类列表")
	t, err:= template.ParseFiles("views/clist.html", "views/public/common.html")
	if err != nil{
		log.Fatal("template parse:", err)
	}
	t.Execute(w, nil)
}

//添加分类
func Add(w http.ResponseWriter, r *http.Request)  {
	util.GetLoginUid(w, r)
	err := r.ParseForm()//解析url传递的参数，对于post则解析响应包的主体

	if err != nil {
		fmt.Println("解析表单数据失败!")
		return
	}
	if r.Method == "GET"{
		//fmt.Fprint(w, "分类列表")
		t, err:= template.ParseFiles("views/clist_add.html", "views/public/common.html")
		if err != nil{
			log.Fatal("template parse:", err)
		}
		t.Execute(w, nil)
	}else{
		name := r.Form.Get("cname")
		pid := r.Form.Get("pid")
		sort := r.Form.Get("sort")
		status := r.Form.Get("status")

		insertId := models.CategoryAdd(name, pid, sort, status)
		if insertId > 0{ //插入成功
			util.JsonOutPut(w,strconv.Itoa(insertId), 200, "添加成功")
			return
		}else{
			util.JsonOutPut(w,"0", 201, "添加失败")
			return
		}
	}
}