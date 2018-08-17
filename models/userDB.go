package models

import (
	"fmt"
	_"log"
)
var tableConfig = make(map[string]string)
var (
	UserpassSalt = "GoQaOp"
)

type UserInfo map[string]string

//通过用户名获取用户信息
func GetInfoByName(name string) (UserInfo, error){
	db := Connect()
	sql := fmt.Sprintf("Select id,username,password FROM %s WHERE username = '%s' limit 1", "qa_admin", name)
	var id,username, password string
	err := db.QueryRow(sql).Scan(&id, &username, &password)
	if err != nil{
		return UserInfo{},err
	}

	return UserInfo{"id":id,"username":username,"password":password}, nil
}

