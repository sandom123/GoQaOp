package users

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"errors"
	_"crypto/md5"
	"github.com/sandom123/GoQaOp/models"
	"github.com/sandom123/GoQaOp/util"
	_"time"
)

type Data struct {
	Action string
	Data string
}

type JsonOut struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("登录")
	err := r.ParseForm()//解析url传递的参数，对于post则解析响应包的主体

	if err != nil {
		fmt.Println("解析表单数据失败!")
	}

	if r.Method == "GET"{
		t, err := template.ParseFiles("views/login.html")
		if err != nil{
			log.Fatal("parseFile:", err)
		}
		d :=Data{Action:"login"}
		t.Execute(w, d)
	}else{
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		token := r.Form.Get("token")
		id, err := checkLogin(username, password, token)
		if(err != nil){
			util.JsonOutPut(w,"0", 201, err.Error())
			return
		}
		cookie := http.Cookie{Name: "username", Value: string(username), Path: "/", MaxAge: 86400}
		http.SetCookie(w, &cookie)
		cookie2 := http.Cookie{Name: "uid", Value: string(id), Path: "/", MaxAge: 86400}
		http.SetCookie(w, &cookie2)

		util.JsonOutPut(w,"1", 200, "登录成功")
		return
	}

}

//检测用户是否登录
func checkLogin(username, password , token string) (string, error) {
	if len(username) == 0{
		return "",errors.New("用户名不能为空")
	}
	if len(password) == 0{
		return "",errors.New("密码不能为空")
	}
	userInfo,err:=models.GetInfoByName(username)
	if err != nil{
		return "",errors.New("用户信息查询错误")
	}
	pas:= userInfo["password"]
	repas := fmt.Sprintf("%s%s", util.Md5hash(password), models.UserpassSalt)
	if  repas!= pas{
		return "",errors.New("用户名或者密码输入错误")
	}
	id:= string(userInfo["id"])
	return id,nil
}
