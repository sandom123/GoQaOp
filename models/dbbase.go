package models

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/go-sql-driver/mysql"
	"github.com/sandom123/GoQaOp/util"
)

//数据库连接基类
type Dbbase struct{
	host string
	username string
	password string
	port int
	dbname string
	charset string
}

//数据库连接
func Connect() *sql.DB{
	ip :=util.GetIntranetIp()
	d := Dbbase{}
	if ip == "192.168.0.103"{//家里的ip
		d = Dbbase{
			"127.0.0.1","root", "123456", 3306, "GoQaOp", "utf8",
		}
	}else{
		d = Dbbase{
			"192.168.99.100","root", "xy123456", 3306, "GoQaOp", "utf8",
		}
	}

	dirver := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		d.username, d.password, d.host, d.port, d.dbname, d.charset,
	)
	db, err := sql.Open("mysql", dirver)
	if err != nil{
		log.Fatal("db connect:", err)
	}

	return db
}
type CateCloum struct{
	id int
	name string
	pid,sort,status int
}
// 获取表数据
func GetList(db *sql.DB, s string) (map[int]map[string]string, error){

	result := make(map[int]map[string]string)
	rows, err := db.Query(s)
	if err != nil {
		return map[int]map[string]string{}, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return map[int]map[string]string{}, err
	}
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}
	ii := 0
	for rows.Next() {
		err := rows.Scan(dest...)
		if err != nil{
			log.Fatal("Scan error:", err)
		}
		cls := make(map[string]string)
		 for i, raw := range rawResult {
			cls[cols[i]] = string(raw)
		}
		result[ii] = cls
		 ii ++
	}

	return result,nil
}




