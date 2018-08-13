package users

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"errors"
	_"crypto/md5"
	"github.com/sandom123/GoQaOp/models"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("登录")
	r.ParseForm() //解析url传递的参数，对于post则解析响应包的主体
	if r.Method == "GET"{
		t, err := template.ParseFiles("views/login.html")
		if err != nil{
			log.Fatal("parseFile:", err)
		}
		//h := md5.New()
		//token := fmt.Sprintf("%x", h.Sum(nil))
		t.Execute(w, nil)
	}else{
		//接收表单数据，并进行处理
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		token := r.Form.Get("token")
		checkLogin(username, password, token)
		
	}

}

//检测用户是否登录
func checkLogin(username, password , token string) (error) {
	//fmt.Println(username, password, token)
	if len(username) == 0{
		return errors.New("用户名不能为空")
	}
	if len(password) == 0{
		return errors.New("密码不能为空")
	}
	models.GetInfoByName(username)

	return nil
}
