package render

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"github.com/gogank/papillon/mapper"
	"fmt"
)

var TestPostPath = "../test/single.md"
var TestTempl = "../test/post.hbs"

func TestRenderer_DoRender(t *testing.T) {
	r := &renderer{}
	b, e := ioutil.ReadFile(TestPostPath)
	assert.Nil(t, e)

	tb, e := ioutil.ReadFile(TestTempl)
	assert.Nil(t, e)

	_, o, e := r.DoRender(b, tb, nil)
	assert.Nil(t, e)
	t.Log(string(o))

}

func TestConvertLink(t *testing.T) {
	r := &renderer{}
	b, e := ioutil.ReadFile(TestPostPath)
	assert.Nil(t, e)

	tb, e := ioutil.ReadFile(TestTempl)
	assert.Nil(t, e)

	_, o, e := r.DoRender(b, tb, nil)
	assert.Nil(t, e)

	mapper.WalkDir("../test/asserts")
	// "style/style.css"
	html, e := r.ConvertLink(o, "../test/asserts")
	assert.Nil(t, e)

	t.Log(string(html))

}

func TestIsInternal(t *testing.T) {
	assert.False(t, isInternal("http://www.papillon.io"))
	assert.True(t, isInternal("./internal.css"))
}

func TestIsInternal2(t *testing.T) {
	links := make(map[string]bool)
	links["http://www.papillon.io"] = false
	links["style.css"] = true

	for k, v := range links {
		t.Log(k)
		assert.Equal(t, isInternal(k), v)
	}
}

func TestIsSlashEnd(t *testing.T) {
	link := "www.papillon.com/"
	assert.True(t, isSlashEnd(link))
}

func TestParseHtree(t *testing.T) {
	line1 := "# line1"
	line2 := "## line1"
	line3 := "### line3"
	st := ParseHtree(line1, nil)
	st = ParseHtree(line2, st)
	st = ParseHtree(line3, st)
	for !st.empty() {
		node := st.pop()
		fmt.Println(node.level)
		fmt.Println(node.content)
		subtree := node.subTree

		fmt.Println(subtree[0].level)
		fmt.Println(subtree[0].content)
		fmt.Println(subtree[0].subTree)
	}
}
