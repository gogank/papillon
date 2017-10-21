package utils

import (
	"os"
	"path/filepath"
	"fmt"
	"path"
	"io/ioutil"
	"strings"
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

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, fi.Name())
		}
	}
	return files, nil
}

func ReadFile(filename string) ([]byte,error) {
	contents,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}
	return contents,nil
}

func ExistDir(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(err)
	return false
}

func RemoveDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}