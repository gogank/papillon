package render

import (
	"github.com/russross/blackfriday"
	"strings"
	"bufio"
	"fmt"
	"io"
	"github.com/aymerick/raymond"
)

type renderer struct {
}

func New() *renderer {
	return &renderer{}
}
// Single parse the single post content, return the config map and html content bytes
func (render *renderer) DoRender(raw []byte, tpl []byte) (ctx map[string]string,result []byte, err error) {
	ctx, raw_content,err := ReadPostConfig(raw)
	html_content := blackfriday.MarkdownCommon(raw_content)
	//TODO 临时方案
	ctx["content"] = string(html_content)
	result_s, err := raymond.Render(string(tpl), ctx)
	if err != nil {
		return nil,nil,err
	}
	result = []byte(result_s)
	return
}


func ReadPostConfig(raw []byte) (map[string]string, []byte, error) {
	sr := strings.NewReader(string(raw))
	buf := bufio.NewReaderSize(sr, 4096)

	content := make([]byte,0)
	postConf := make(map[string]string)

	str_flag := false
	end_flag := false

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) >= 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		if isPrefix {
			// TODO: 这里可能出现单行缓冲区溢出
			panic("buffer overflow")
		}
		if err != io.EOF && err != nil  {
			return nil,nil,err
		}
		// 开始读取第一行
		if !str_flag && ! end_flag {
			line_str :=  strings.TrimSpace(string(line))
			dash_line := string(line_str[0:3])
			if dash_line == "---" {
				str_flag = true
			}
			// 判断是否是配置区域
		} else if str_flag && !end_flag {
			line_str :=  strings.TrimSpace(string(line))
			dash_line := string(line_str[0:3])
			if dash_line != "---" && strings.Contains(line_str,":"){
				tmp := strings.Split(string(line), ":")
				key := strings.TrimSpace(tmp[0])
				fmt.Println("key:", key)
				value := strings.TrimSpace(tmp[1])
				fmt.Println("value:", value)
				postConf[key] = value
			} else if dash_line == "---"{
				end_flag = true
			}
		} else if str_flag && end_flag {
			content = append(content,line...)
		}
	}

	return postConf,content,nil
}

func (render *renderer) Directory(dir string, dst string) {

}
