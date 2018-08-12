package users

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("登录")
	if r.Method == "GET"{
		t, err := template.ParseFiles("github.com/sandom123/GoQaOp/views/login.html")
		if err != nil{
			log.Fatal("parseFile:", err)
		}
		t.Execute(w, nil)
	}else{

	}

}
