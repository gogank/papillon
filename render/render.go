package render

import (
	"strings"
	"bufio"
	"io"
	"github.com/aymerick/raymond"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/PuerkitoBio/goquery"
	"errors"
	"fmt"
	"regexp"
	"github.com/gogank/papillon/mapper"
)

type renderer struct {
}

func New() *renderer {
	return &renderer{}
}

// Single parse the single post content, return the config map and html content bytes
func (render *renderer) DoRender(raw []byte, tpl []byte, user_ctx map[string]interface{}) (map[string]interface{}, []byte, error) {
	inner_ctx := make(map[string]interface{})
	if raw != nil {
		post_ctx, raw_content, err := readPostConfig(raw)
		if err != nil {
			return nil, nil, err
		}
		html_content := blackfriday.Run(raw_content)
		//TODO 临时方案
		inner_ctx["content"] = string(html_content)
		for key, value := range post_ctx {
			inner_ctx[key] = value
		}
	}

	if user_ctx != nil {
		for key, value := range user_ctx {
			inner_ctx[key] = value
		}
	}

	result_s, err := raymond.Render(string(tpl), inner_ctx)
	if err != nil {
		return nil, nil, err
	}
	result := []byte(result_s)
	return inner_ctx, result, nil
}

func readPostConfig(raw []byte) (map[string]string, []byte, error) {
	sr := strings.NewReader(string(raw))
	buf := bufio.NewReaderSize(sr, 4096)

	content := make([]byte, 0)
	postConf := make(map[string]string)

	str_flag := false
	end_flag := false

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) >= 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		if isPrefix {
			// TODO: 这里可能出现单行缓冲区溢出
			panic("buffer overflow")
		}
		if err != io.EOF && err != nil {
			return nil, nil, err
		}
		// 开始读取第一行
		if !str_flag && ! end_flag {
			line_str := strings.TrimSpace(string(line))
			// 防止没有 `---`
			if !strings.HasPrefix(line_str, "-") {
				continue
			}
			dash_line := string(line_str[0:3])
			if dash_line == "---" {
				str_flag = true
			}
			// 判断是否是配置区域
		} else if str_flag && !end_flag {
			line_str := strings.TrimSpace(string(line))
			dash_line := string(line_str[0:3])
			if dash_line != "---" && strings.Contains(line_str, ":") {
				tmp := strings.Split(string(line), ":")
				key := strings.TrimSpace(tmp[0])
				value := strings.TrimSpace(tmp[1])
				postConf[key] = value
			} else if dash_line == "---" {
				end_flag = true
			}
		} else if str_flag && end_flag {
			content = append(content, line...)
			// 需要加上换行符，否则markdown会解析错误
			content = append(content, byte('\n'))
		}
	}

	if len(content) == 0 {
		return nil, nil, errors.New("post has no configuration part")
	}

	return postConf, content, nil
}

func GetMeta(raw []byte) (map[string]string, error) {
	meta, _, err := readPostConfig(raw)
	return meta, err
}

// Convert the link as `public` folder  as root
func ConvertLink(raw []byte) ([]byte, error) {
	sr := strings.NewReader(string(raw))
	doc, err := goquery.NewDocumentFromReader(sr)
	if err != nil {
		return nil, err
	}
	tagList := map[string]func(int, *goquery.Selection){}
	tagList["link"] = changeLink
	tagList["a"] = changeLink
	tagList["script"] = changeSrc
	tagList["img"] = changeSrc
	for k,v := range tagList{
		doc.Find(k).Each(v)
	}
	html, err := doc.Html()
	if err != nil {
		return nil, err
	}

	return []byte(html), nil
}

// changeSrc change TAG's `src` attr
func changeSrc(i int, s *goquery.Selection) {
	if src, ok := s.Attr("src"); ok && isInternal(src) {
		if isInternal(src) {
			if isSlashEnd(src) {
				src = src + "index.html"
			}
			fmt.Println("BBBBBBBBB>>>" + src)
			if ipfs_link, ok := mapper.Get(src); ok {
				s.SetAttr("src", ipfs_link +"BBBBBB")
			}
		}

	}
}

// changeSrc change TAG's `link` attr
func changeLink(i int, s *goquery.Selection) {
	if src, ok := s.Attr("link"); ok && isInternal(src) {
		fmt.Println(src)
		// 如果是内部链接，进行处理
		if isInternal(src) {
			if isSlashEnd(src) {
				src = src + "index.html"
			}
			if ipfs_link, ok := mapper.Get(src); ok {
				s.SetAttr("link", ipfs_link + "AAAAAAA")
			}
		}

	}
}

func isInternal(link string) bool {
	reg, err := regexp.Compile("(?i:^http).*")
	if err != nil {
		panic(err)
	}
	if reg.MatchString(link) {
		return false
	}
	return true
}

func isSlashEnd(link string) bool {
	return strings.HasSuffix(link, "/")
}
