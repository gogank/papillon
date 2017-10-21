package render

import (
	"github.com/russross/blackfriday"
	"strings"
	"bufio"
	"fmt"
	"io"
	"github.com/Joker/jade"
	"github.com/gogank/papillon/utils"
)

type renderer struct {
}

func New() *renderer {
	return &renderer{}
}
// Single parse the single post content, return the config map and html content bytes
func (render *renderer) Single(raw []byte) (map[string]string, []byte, error) {
	post_conf, raw_content,err := ReadPostConfig(raw)
	output := blackfriday.MarkdownCommon(raw_content)
	return post_conf, output, err
}

// DoRender apply the jade template and return rendered html content
// html - html bytes content, generally from Single
// template - jade template file path
func (render *renderer) DoRender(html string,template string)(rendered_html []byte, err error){
	tmpl_b , err := utils.ReadFile(template)
	if err != nil{
		return
	}
	tpl, err := jade.Parse("default", string(tmpl_b))

	if err != nil {
	fmt.Printf("Parse error: %v", err)
	return
	}

	fmt.Printf( "Output:\n\n%s", tpl  )
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
		fmt.Println(string(line))
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
