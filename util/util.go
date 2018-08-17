package util

import (
	"path/filepath"
	"os"
	"net"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
)

//公共方法
 func GetCurrentPath() string{
 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
 	if err != nil{
 		return ""
	}

	return dir
 }

//获取客户端IP
func GetIntranetIp() string{
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}

//md5加密字符串
func Md5hash(s string)  string{
	m := md5.New()
	m.Write([]byte(s))
	cipherStr := m.Sum(nil)
	str :=hex.EncodeToString(cipherStr)

	return str
}

type JsonOut struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

//json输出
func JsonOutPut(w http.ResponseWriter, data string, code int, message string) {
	var d JsonOut
	d.Code = code
	d.Message = message
	d.Data = data
	result, _ := json.Marshal(d)
	os.Stdout.Write(result)
	w.Write(result)
	return
}

//检测登录,获取登录的用户ID
func GetLoginUid(w http.ResponseWriter, r *http.Request) (int, string){
	//// read cookie
	cookie, err := r.Cookie("uid")
	if err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return 0, ""
	}
	cuid := cookie.Value
	if cuid == ""{ //未登录
		http.Redirect(w, r, "/login/", http.StatusFound)
		return 0, ""
	}
	uid, err := strconv.Atoi(cuid)
	if err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return 0, ""
	}
	ucookie, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return 0, ""
	}
	uusername := ucookie.Value
	if uusername == ""{ //未登录
		http.Redirect(w, r, "/login/", http.StatusFound)
		return 0, ""
	}

	return uid, uusername
}