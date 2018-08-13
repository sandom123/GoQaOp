package models

import (
	"fmt"
	_"log"
)
var tableConfig = make(map[string]string)
var (
	tablename = "qa_admin"
	userpassSalt = "GoQaOp"
)

func GetInfoByName(name string){
	db := Connect()
	sql := fmt.Sprintf("Select * FROM %s WHERE username = '%s' limit 1", tablename, name)
	info, _ := Get(db, sql)
	fmt.Println(info)


}

