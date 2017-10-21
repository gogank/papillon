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

func Mkdir(dir string) bool {
	var path string
	if os.IsPathSeparator('\\') {  //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	fmt.Println(path)
	pwd, _ := os.Getwd()  //当前的目录
	err := os.Mkdir(pwd+path+dir, os.ModePerm)  //在当前目录下生成md目录
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("创建目录" + path + dir)
	return true
}
