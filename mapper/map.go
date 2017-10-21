package mapper

import (
	"github.com/gogank/papillon/publish"
	"fmt"
	"github.com/pkg/errors"
	"github.com/gogank/papillon/utils"
	"encoding/hex"
)

var linkMap map[string]string
var publisher *publish.PublishImpl

func init(){
	linkMap = make(map[string]string)
	publisher = publish.NewPublishImpl("localhost:5001")
}

func Get(key string) string {
	key = hex.EncodeToString(utils.ByteHash([]byte(key)))
	if hash,ok := linkMap[key];ok {
		return hash
	}
	return ""
}

func Put(key string) (string,error) {
	hash,err := publisher.AddFile(key)
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
