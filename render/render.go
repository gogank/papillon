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
	"net/url"
)

type renderer struct {
}

func New() *renderer {
	return &renderer{}
}

// DoRender 进行渲染，将markdown 原始内容 和模板传入，并进行渲染
// 如果 raw = nil, 则仅进行 tpl + user_ctx 渲染
// 如果 raw != nil 则先对raw 进行markdown解析，然后把解析结果 user_ctx["content"] = parsedHtml
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

//readPostConfig 读取文章，过滤文章元信息，并返回需要进行markdown parse的内容
func readPostConfig(raw []byte) (map[string]string, []byte, error) {
	sr := strings.NewReader(string(raw))
	buf := bufio.NewReaderSize(sr, 4096)

	content := make([]byte, 0)
	postConf := make(map[string]string)

	str_flag := false
	end_flag := false

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) >= 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		// 如果 isPrefix 是true的话，则需要再读一行，然后继续处理
		for isPrefix {
			var newLine []byte
			newLine, isPrefix,err = buf.ReadLine()
			if err != nil && err != io.EOF{
				return nil,nil,err
			}
			line = append(line,newLine...)
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


//GetMeta 取得文章元信息
//返回文章元信息
func GetMeta(raw []byte) (map[string]string, error) {
	meta, _, err := readPostConfig(raw)
	return meta, err
}

// 将Html内容中的所有内部链接进行转换
// 0. 如果是内部链接则继续
// 1. 如果能够在ipfs网络中找到，则替换
// 2. 如果不能在ipfs网络中找到，则不处理
func (render *renderer)ConvertLink(raw []byte) ([]byte, error) {
	sr := strings.NewReader(string(raw))
	doc, err := goquery.NewDocumentFromReader(sr)
	if err != nil {
		return nil, err
	}
	tagList := map[string]func(int, *goquery.Selection){}
	tagList["link"] = changeHref
	tagList["a"] = changeHref
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

// changeSrc 将img/script中的 src属性 进行替换
func changeSrc(i int, s *goquery.Selection) {
	if src, ok := s.Attr("src"); ok && isInternal(src) {
		fmt.Println("%%%%%%%%%%%%%%%%%%%%%", parseLink(src))
		if ipfs_link, ok := mapper.Get(parseLink(src)); ok {
				fmt.Println("convert ", parseLink(src), " to ", ipfs_link)
				s.SetAttr("src", addIPFSPrefix(ipfs_link))
		}
	}
}

// changeHref 将 link/a 中的 link href 属性进行替换
func changeHref(i int, s *goquery.Selection) {
	if src, ok := s.Attr("link"); ok && isInternal(src) {
		// 如果是内部链接，进行处理
		if ipfs_link, ok := mapper.Get(parseLink(src)); ok {
				s.SetAttr("link", addIPFSPrefix(ipfs_link))
		}
	}
}

//addIPFSPrefix 添加 IPFS 网络前缀
func addIPFSPrefix(hash string) string{
	return "https://ipfs.io/ipfs/" + hash
}

//解析link 返回能够查询的静态资源key
func parseLink(link string) string{
	if isInternal(link) {
			if isSlashEnd(link) {
				link = link + "index.html"
			}
			if isRelative(link){
				fmt.Println("relative " + link)
				if link[0] == '.' {
					link = link[1:]
				}
				link = string(append([]byte("/"),link...))
			}
	}
	return link
}

//isInternal 判断是否是内部链接
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

//isSlashEnd 判断是否由/结尾
func isSlashEnd(link string) bool {
	return strings.HasSuffix(link, "/")
}

//isRelative 判断是否是相对url
func isRelative(link string) bool{
	url,err := url.Parse(link)
	if err != nil{
		return false
	}
	return !url.IsAbs() && !isSlashStart(link)
}

//isSlashStart 是否由`/`开始
func isSlashStart(link string) bool {
	return strings.HasPrefix(link, "/")
}
