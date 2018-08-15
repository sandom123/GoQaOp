package question

import (
	"net/http"
	"fmt"
	_"html/template"
	_"log"
)

type Data struct {
	Action string
}


func CategoryList(w http.ResponseWriter, r *http.Request){

	fmt.Fprint(w, "分类列表")
	//t, err:= template.ParseFiles("views/clist.html")
	//if err != nil{
	//	log.Fatal("template parse:", err)
	//}
	//d :=Data{Action:"clist"}
	//t.Execute(w, d)
}
