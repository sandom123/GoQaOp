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

type listTree struct{
	info map[string]string
	child map[int]map[string]string
}

type listData struct{
	list map[int]listTree
	count int
}

//获取列表树
func getListTree()  map[int]listTree{
	pwhere := "pid = 0"
	pdata :=models.GetCategoryList(pwhere)
	data := make(map[int]listTree)
	for i,val := range pdata{
		if val["pid"] != "0"{
			childWhere := "pid ="+val["pid"]
			cdata := models.GetCategoryList(childWhere)
			listS := listTree{info:val, child:cdata} //fmt.Println(listS)
			data[i] = listS
		}else{
			data[i] = listTree{info:val, child:map[int]map[string]string{}}
		}
	}

	return data
}
//分类列表
func List(w http.ResponseWriter, r *http.Request){
	util.GetLoginUid(w, r)
	lidata := getListTree()
	count := 10
	//data := util.Struct2Map(listData{lidata, count})
	data := listData{lidata, count}
	fmt.Println(data)
	//fmt.Fprint(w, "分类列表")
	t, err:= template.ParseFiles("views/clist.html", "views/public/common.html")
	if err != nil{
		log.Fatal("template parse:", err)
	}
	t.Execute(w, data)
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