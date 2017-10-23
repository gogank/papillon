package mapper

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogank/papillon/publish"
	"github.com/gogank/papillon/utils"
)

var linkMap map[string]string
var publisher *publish.Impl

func init() {
	linkMap = make(map[string]string)
	publisher = publish.NewImpl()
}

//Get a value from the mapper
func Get(key string) (string, bool) {
	key = hex.EncodeToString(utils.ByteHash([]byte(key)))
	if hash, ok := linkMap[key]; ok {
		return hash, true
	}
	return "", false
}

//Put a value into the mapper
func Put(key string, dir string) (string, error) {
	hash, err := publisher.AddFile(key)
	dirPthByte := []rune(dir)
	lenDir := len(dirPthByte)
	filenameByte := []rune(key)
	key = string(filenameByte[lenDir:])
	fmt.Println("put: ", key)
	key = hex.EncodeToString(utils.ByteHash([]byte(key)))
	if err != nil {
		return "", err
	}
	linkMap[key] = hash
	return hash, nil
}

//WalkDirCmd walk the spific dir and put the relative dir path into mapper
func WalkDirCmd(dirPth string) ([]string, error) {
	files := make([]string, 0, 30)
	dirPthByte := []rune(dirPth)
	bol := strings.EqualFold("./", string(dirPthByte[:len([]rune("./"))]))
	if bol {
		dirPth = string(dirPthByte[len([]rune("./")):])
	}
	rootHash, err := publisher.AddDir(dirPth)
	if err != nil {
		return nil, err
	}
	rootkey := hex.EncodeToString(utils.ByteHash([]byte("/")))
	linkMap[rootkey] = rootHash

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		files = append(files, filename)
		dirPthByte := []rune(dirPth)
		filenameByte := []rune(filename)

		key := string(filenameByte[len(dirPthByte):])
		value := rootHash + string(filenameByte[len(dirPthByte):])

		key = hex.EncodeToString(utils.ByteHash([]byte(key)))
		if err != nil {
			return err
		}
		linkMap[key] = value
		return nil
	})
	return files, err
}
