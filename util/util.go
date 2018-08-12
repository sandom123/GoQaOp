package util

import (
	"path/filepath"
	"os"
)

//公共方法
 func GetCurrentPath() string{
 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
 	if err != nil{
 		return ""
	}

	return dir
 }
