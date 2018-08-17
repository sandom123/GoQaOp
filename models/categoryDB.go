package models

import (
	"fmt"
)

//添加分类
func CategoryAdd(name string, pid, sort, status string) int{
	db := Connect()

	sql := fmt.Sprintf("INSERT %s SET name=?,pid=?,sort=?,status=?", "qa_category")
	stmt, err := db.Prepare(sql)
	if err != nil{
		return 0
	}
	res, err := stmt.Exec(name, pid, sort, status)
	if err != nil{
		return 0
	}
	id, err := res.LastInsertId()
	if err != nil{
		return 0
	}

	return int(id)
}

//修改分类
func CategoryEdit(id int, name string, pid, sort, status string) int{
	db := Connect()

	sql := fmt.Sprintf("UPDATE %s SET name=?,pid=?,sort=?,status=? WHERE id=?", "qa_category")
	stmt, err := db.Prepare(sql)
	if err != nil{
		return 0
	}
	res, err := stmt.Exec(name, pid, sort, status, id)
	if err != nil{
		return 0
	}
	affectedRowNum, err :=res.RowsAffected()
	if err != nil{
		return 0
	}

	return int(affectedRowNum)
}

//删除分类
func CategoryDel(id int) bool {
	db := Connect()
	sql :=fmt.Sprintf("DELETE FROM %s WHERE id=?", "qa_category")
	stmt, err :=db.Prepare(sql)
	if err != nil{
		return false
	}
	res, err :=stmt.Exec(id) //绑定参数
	if err != nil{
		return false
	}
	affectedRowNum, err := res.RowsAffected()
	if err != nil{
		return false
	}
	if(affectedRowNum > 0){
		return true
	}else{
		return false
	}
}