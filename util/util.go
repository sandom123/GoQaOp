package util

import (
	"path/filepath"
	"os"
	"net"
	"crypto/md5"
	"encoding/hex"
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
