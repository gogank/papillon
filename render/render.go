package render

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aymerick/raymond"
	"github.com/gogank/papillon/mapper"
	"gopkg.in/russross/blackfriday.v2"
)

//Renderer main renderer function supplier
type Renderer struct {
}

//New return a Renderer instance
func New() *Renderer {
	return &Renderer{}
}

// DoRender 进行渲染，将markdown 原始内容 和模板传入，并进行渲染
// 如果 raw = nil, 则仅进行 tpl + user_ctx 渲染
// 如果 raw != nil 则先对raw 进行markdown解析，然后把解析结果 user_ctx["content"] = parsedHtml
func (render *Renderer) DoRender(raw []byte, tpl []byte, userCtx map[string]interface{}) (map[string]interface{}, []byte, error) {
	innerCtx := make(map[string]interface{})
	if raw != nil {
		postCtx, rawContent, err := readPostConfig(raw)
		if err != nil {
			return nil, nil, err
		}
		HTMLContent := blackfriday.Run(rawContent)
		//TODO 临时方案
		innerCtx["content"] = string(HTMLContent)
		for key, value := range postCtx {
			innerCtx[key] = value
		}
	}

	if userCtx != nil {
		for key, value := range userCtx {
			innerCtx[key] = value
		}
	}

	resultS, err := raymond.Render(string(tpl), innerCtx)
	if err != nil {
		return nil, nil, err
	}
	result := []byte(resultS)
	return innerCtx, result, nil
}

//readPostConfig 读取文章，过滤文章元信息，并返回需要进行markdown parse的内容
func readPostConfig(raw []byte) (map[string]string, []byte, error) {
	sr := strings.NewReader(string(raw))
	buf := bufio.NewReaderSize(sr, 4096)

	content := make([]byte, 0)
	postConf := make(map[string]string)

	strFlag := false
	endFlag := false

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) >= 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		// 如果 isPrefix 是true的话，则需要再读一行，然后继续处理
		for isPrefix {
			var newLine []byte
			newLine, isPrefix, err = buf.ReadLine()
			if err != nil && err != io.EOF {
				return nil, nil, err
			}
			line = append(line, newLine...)
		}
		if err != io.EOF && err != nil {
			return nil, nil, err
		}
		// 开始读取第一行
		if !strFlag && !endFlag {
			lineStr := strings.TrimSpace(string(line))
			// 防止没有 `---`
			if !strings.HasPrefix(lineStr, "-") {
				continue
			}
			dashLine := string(lineStr[0:3])
			if dashLine == "---" {
				strFlag = true
			}
			// 判断是否是配置区域
		} else if strFlag && !endFlag {
			lineStr := strings.TrimSpace(string(line))
			dashLine := string(lineStr[0:3])
			if dashLine != "---" && strings.Contains(lineStr, ":") {
				tmp := strings.Split(string(line), ":")
				key := strings.TrimSpace(tmp[0])
				value := strings.TrimSpace(tmp[1])
				postConf[key] = value
			} else if dashLine == "---" {
				endFlag = true
			}
		} else if strFlag && endFlag {
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

type hTreeNode struct {
	level,
	content string
	subTree []*hTreeNode
}

type stack struct {
	s []*hTreeNode
}

func (st *stack) push(node *hTreeNode) {
	st.s = append(st.s, node)
}

func (st *stack) pop() *hTreeNode {
	if len(st.s) == 0 {
		return nil
	}
	tmpn := st.s[len(st.s)-1]
	st.s = st.s[0 : len(st.s)-1]
	return tmpn
}
func (st *stack) empty() bool {
	return len(st.s) == 0
}

func (st *stack) len() int {
	return len(st.s)
}

func newStack() *stack {
	return &stack{
		s: make([]*hTreeNode, 0),
	}
}

func parseHtree(str string, st *stack) *stack {
	sstr := strings.TrimSpace(str)
	// h1
	if strings.HasPrefix(sstr, "###") {
		//子级别树
	} else if strings.HasPrefix(sstr, "##") {
		// 如果树为空则push root
		if st.empty() {
			root := &hTreeNode{
				level:   "root",
				content: "root",
				subTree: make([]*hTreeNode, 0),
			}
			root.subTree = append(root.subTree,
				&hTreeNode{
					level:   "h1",
					content: "",
					subTree: make([]*hTreeNode, 0),
				})
			st.push(root)
		}
		// 找到根
		node := st.pop()
		for node.level != "h1" && node.level != "root" {
			node = st.pop()
		}
		// 不会是 nil
		if node.level == "h1" {
			node.subTree = append(node.subTree,
				&hTreeNode{
					level:   "h2",
					content: sstr[2:],
					subTree: make([]*hTreeNode, 0),
				})
		} else if node.level == "root" {
			//如果是root
			h1 := &hTreeNode{
				level:   "h1",
				content: "h1",
				subTree: make([]*hTreeNode, 0),
			}
			h1.subTree = append(h1.subTree,
				&hTreeNode{
					level:   "h2",
					content: sstr[2:],
					subTree: make([]*hTreeNode, 0),
				})
			node.subTree = append(node.subTree, h1)
		}
		st.push(node)

	} else if strings.HasPrefix(sstr, "#") {
		if st == nil {
			st = newStack()
		}
		if st.empty() {
			st.push(&hTreeNode{
				level:   "root",
				content: "root",
				subTree: make([]*hTreeNode, 0),
			})
		}
		fmt.Println(st.len())
		// 找到根
		node := st.pop()
		for node.level != "root" {
			node = st.pop()
		}
		// 在根的子树中push
		node.subTree = append(node.subTree, &hTreeNode{
			level:   "h1",
			content: sstr[1:] + "AAA",
			subTree: make([]*hTreeNode, 0),
		})
		st.push(node)
	}

	return st
}

//GetMeta 取得文章元信息
//返回文章元信息
func GetMeta(raw []byte) (map[string]string, error) {
	meta, content, err := readPostConfig(raw)
	if len(content) < 320 {
		meta["abstract"] = string(blackfriday.Run(content))
	} else {
		meta["abstract"] = string(blackfriday.Run(content[:302]))
	}

	return meta, err
}

// ConvertLink 将Html内容中的所有内部链接进行转换
// 0. 如果是内部链接则继续
// 1. 如果能够在ipfs网络中找到，则替换
// 2. 如果不能在ipfs网络中找到，则不处理
func (render *Renderer) ConvertLink(raw []byte, parentDir string) ([]byte, error) {
	// upload files
	_, err := mapper.WalkDirCmd(parentDir)
	if err != nil {
		return nil, err
	}

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
	for k, v := range tagList {
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
		if ipfsLink, ok := mapper.Get(parseLink(src)); ok {
			fmt.Println("Convert: ", src, "to", ipfsLink)
			//fmt.Println("convert ", parseLink(src), " to ", ipfs_link)
			s.SetAttr("src", addIPFSPrefix(ipfsLink))
		}
	}
}

// changeHref 将 link/a 中的 link href 属性进行替换
func changeHref(i int, s *goquery.Selection) {
	if src, ok := s.Attr("href"); ok && isInternal(src) {
		// 如果是内部链接，进行处理
		if ipfsLink, ok := mapper.Get(parseLink(src)); ok {
			//fmt.Println("Convert: ", src, "to",ipfs_link)
			s.SetAttr("href", addIPFSPrefix(ipfsLink))
		}
	}
}

//addIPFSPrefix 添加 IPFS 网络前缀
func addIPFSPrefix(hash string) string {
	return "https://ipfs.io/ipfs/" + hash
}

//解析link 返回能够查询的静态资源key
func parseLink(link string) string {
	if isInternal(link) {
		if isSlashEnd(link) && len(link) > 1 {
			link = link + "index.html"
		}
		if isRelative(link) {
			fmt.Println("relative " + link)
			if link[0] == '.' {
				link = link[1:]
			}
			link = string(append([]byte("/"), link...))
		}
	}
	// 排除 /verndors/script/main.js?v=2.1.3
	if strings.Contains(link, "?") {
		link = strings.Split(link, "?")[0]
	}
	// 转换空格
	link = strings.Replace(link, " ", "_", -1)
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
func isRelative(link string) bool {
	url, err := url.Parse(link)
	if err != nil {
		return false
	}
	return !url.IsAbs() && !isSlashStart(link)
}

//isSlashStart 是否由`/`开始
func isSlashStart(link string) bool {
	return strings.HasPrefix(link, "/")
}
