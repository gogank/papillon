package handler

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/gogank/papillon/configuration"
	"github.com/gogank/papillon/utils"
)

//NewPost New a post by post name and config path
func NewPost(postName string, confPath string) {
	cnf := config.NewConfig(confPath)
	sourceDir := cnf.GetString(utils.DIR_POSTS)
	author := cnf.GetString(utils.COMMON_AUTHOR)
	filePath := path.Join(sourceDir, parsePostName(postName)+".md")
	date := time.Now().Format("2006/01/02")
	contentTpl := `---
title: %s
date: %s
author: %s
---
`
	contentBuf := bytes.NewBufferString("")
	fmt.Fprintf(contentBuf, contentTpl, postName, date, author)
	if !utils.Mkfile(filePath, contentBuf.Bytes()) {
		fmt.Println("创建文章失败!")
	} else {
		fmt.Println("创建文章成功：", filePath)
	}

}

func parsePostName(postName string) string {
	if strings.Contains(postName, " ") {
		postName = strings.Replace(postName, " ", "_", -1)
	}
	// TODO 更多的文件名检查
	return postName
}
