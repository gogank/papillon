package render

import (
	"github.com/russross/blackfriday"
)

type renderer struct {
}

func New() *renderer {
	return &renderer{}
}

func (render *renderer) Single(raw []byte) (map[string]string, []byte, error) {
	output := blackfriday.MarkdownCommon(raw)
	return nil, output, nil
}

func ReadPostConfig(raw []byte) (map[string]string) {
	str_flag := false
	//end_flag := false
	//buff := ""
	dash1_count := 0
	//dash2_count := 0
	content := string(raw)
	for idx := 0; idx < len(content); {
		// 进行第一行的 `---` 处理
		if !str_flag {
			//过滤掉非`-`字符
			if content[idx] != '-' {
				idx ++
			} else {
				//遇到 `-` 字符,如果下一个字符是`-`则进行计数
				// 遇到两个这样的情况就可以确定，存在3个 `-`
				if content[idx+1] == '-' && dash1_count < 2 {
					dash1_count ++

					// 当前字符是 `-` 但是后一个字符不是 `-`
				} else {
					// 但是已经计数完毕
					if dash1_count >= 2 {
						str_flag = true
						//将指针移动到 '\n'之后 TODO: 需要判断是否是文件末尾
						for content[idx] != '\n' {
							idx++
						}
						//得到 `\n`后的指针之后，进入下一步处理
						idx++

					} else {
						//或者是当前字符是`-`,下个字符不是`-`,但是计数没有结束
						// 重置计数
						dash1_count = 0
						// 该行不符合要求 取到下一行
						for content[idx] != '\n' {
							idx++
						}
						//得到 `\n`后的指针之后，进入下一步处理
						idx++
					}

				}

			}
		}
		// 进行第二行的`---`处理

		// get first line `---`
		if content[idx] == '-' && !str_flag {

			if raw[idx+1] == '-' {
				dash1_count++
			}
			if dash1_count >= 3 {

				str_flag = true
			}
		}

	}
	content := string(raw)
	for _, c := range content {

		if c == '\n' {
			print("true")
		}
	}
	return nil
}

func next(b []byte, idx int) byte {
	if len(b) >= idx+2 {
		return b[idx+1]
	}
	return byte(0)
}

func (render *renderer) Directory(dir string, dst string) {

}
