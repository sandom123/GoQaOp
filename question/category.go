package question

import (
	"net/http"
	"fmt"
)

func CategoryList(w http.ResponseWriter, r *http.Request){

	fmt.Fprint(w, "分类列表")
}
