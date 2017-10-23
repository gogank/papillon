package handler

import (
	"github.com/gogank/papillon/configuration"
	"github.com/gogank/papillon/utils"
	"strings"
	"path"
	"fmt"
	"bytes"
	"time"
)

func NewPost(post_name string, conf_path string) {
	cnf := config.NewConfig(conf_path)
	sourceDir := cnf.GetString(utils.DIR_POSTS)
	author := cnf.GetString(utils.COMMON_AUTHOR)
	file_path := path.Join(sourceDir, parsePostName(post_name)+".md")
	date := time.Now().Format("2006/01/02")
	content_tpl := `---
title: %s
date: %s
author: %s
---
`
	content_buf := bytes.NewBufferString("")
	fmt.Fprintf(content_buf, content_tpl, post_name, date, author)
	if !utils.Mkfile(file_path, content_buf.Bytes()) {
		fmt.Println("创建文章失败!")
	} else {
		fmt.Println("创建文章成功：", file_path)
	}

}

func parsePostName(post_name string) string {
	if strings.Contains(post_name, " ") {
		post_name = strings.Replace(post_name, " ", "_", -1)
	}
	// TODO 更多的文件名检查
	return post_name
}
