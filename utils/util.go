package utils

import (
	"os"
	"path/filepath"
	"fmt"
)

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	path := Abs(filename)
	//fmt.Println(path)
	_,err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func Abs(filename string) string {
	path,err := filepath.Abs(filename)
	if err != nil{
		fmt.Errorf("convert abs err %v",err)
		return filename
	}
	return path
}