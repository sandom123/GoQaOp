package question

import (
	"net/http"
	"html/template"
	"log"
)

func CategoryList(w http.ResponseWriter, r *http.Request){

	//fmt.Fprint(w, "分类列表")
	t, err:= template.ParseFiles("views/clist.html")
	if err != nil{
		log.Fatal("template parse:", err)
	}
	t.Execute(w, nil)
}
