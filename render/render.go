package render

import (
	"github.com/russross/blackfriday"
	"strings"
	"bufio"
	"fmt"
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

	sr := strings.NewReader(string(raw))
	buf := bufio.NewReaderSize(sr, 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFGHIJKLMNOP"   true   <nil>
	// "QRSTUVWXYZ"   false   <nil>
	// "1234567890"   false   <nil>
	// ""   false   EOF


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
