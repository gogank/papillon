package mapper

import (
	"github.com/gogank/papillon/publish"
	"fmt"
	"github.com/pkg/errors"
)

var linkMap map[string]string
var publisher *publish.PublishImpl

func init(){
	linkMap = make(map[string]string)
	publisher = publish.NewPublishImpl("localhost:5001")
}

func Get(key string) string {
	if hash,ok := linkMap[key];ok {
		return hash
	}
	return ""
}

func Put(key string) (string,error) {
	hash,err := publisher.AddFile(key)
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
