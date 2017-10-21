package utils

import (
	"os"
	"path/filepath"
	"fmt"
	"path"
)

func Exist(filename string) bool {
	path := Abs(filename)
	_,err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 返回绝对路径
func Abs(filename string) string {
	path,err := filepath.Abs(filename)
	if err != nil{
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
	err := os.Mkdir(pwd+path+dir, os.ModePerm)  //在当前目录下生成目录
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Mkfile(filename string,file []byte) bool {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_,err = f.Write(file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	f.Close();
	return true
}

// 取得文件的后缀
func Ext(filepath string) string{
	return path.Ext(filepath)
}
