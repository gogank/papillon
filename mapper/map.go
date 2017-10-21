package mapper

import (
	"github.com/gogank/papillon/publish"
	"fmt"
	"github.com/pkg/errors"
	"github.com/gogank/papillon/utils"
	"encoding/hex"
	"path/filepath"
	"os"
	"strings"
)

var linkMap map[string]string
var publisher *publish.PublishImpl

func init(){
	linkMap = make(map[string]string)
	publisher = publish.NewPublishImpl("localhost:5001")
}

func Get(key string) (string,bool) {
	key = hex.EncodeToString(utils.ByteHash([]byte(key)))
	if hash,ok := linkMap[key];ok {
		return hash,true
	}
	return "",false
}

func Put(key string,dir string) (string,error) {
	hash,err := publisher.AddFile(key)
	dirPthByte := []rune(dir)
	lenDir := len(dirPthByte)
	filenameByte := []rune(key)
	key = string(filenameByte[lenDir:])
	key = hex.EncodeToString(utils.ByteHash([]byte(key)))
	if err!= nil {
		return "",err
	}
	if _,ok := linkMap[key];ok {
		fmt.Println("This file has alreadly upload.")
		return "",errors.New("This file has alreadly upload.")
	}
	linkMap[key] = hash
	return hash,nil
}

func WalkDir(dirPth string) (hashs []string, err error) {
	files := make([]string, 0, 30)
	hashs = make([]string, 0, 30)
	dirPthByte := []rune(dirPth)
	bol := strings.EqualFold("./",string(dirPthByte[:len([]rune("./"))]))
	if bol {
		dirPth = string(dirPthByte[len([]rune("./")):])
	}
	//fmt.Println(dirPth)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		files = append(files, filename)
		hash,err := Put(filename,dirPth)
		hashs = append(hashs,hash)
		if err != nil{
			fmt.Println(err)
			return err
		}
		return nil
	})
	return hashs, err
}
