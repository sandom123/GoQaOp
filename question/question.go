package question

import (
	"net/http"
	"html/template"
	"log"
	"github.com/sandom123/GoQaOp/util"
)

func QaList(w http.ResponseWriter, r *http.Request){
	//fmt.Fprint(w, "问答列表")
	util.GetLoginUid(w, r)
	t, err:= template.ParseFiles("views/qlist.html")
	if err != nil{
		log.Fatal("template parse:", err)
	}
	t.Execute(w, nil)
}
