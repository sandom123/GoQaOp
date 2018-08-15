package question

import (
	"net/http"
	"html/template"
	"log"
)

func QaList(w http.ResponseWriter, r *http.Request){
	//fmt.Fprint(w, "问答列表")
	t, err:= template.ParseFiles("views/qlist.html")
	if err != nil{
		log.Fatal("template parse:", err)
	}
	t.Execute(w, nil)
}
