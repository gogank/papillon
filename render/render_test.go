package render

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
)

var TestPostPath = "../test/single.md"
var TestTempl = "../test/post.hbs"

func TestRenderer_DoRender(t *testing.T) {
	r := &renderer{}
	b,e := ioutil.ReadFile(TestPostPath)
	assert.Nil(t,e)

	tb,e := ioutil.ReadFile(TestTempl)
	assert.Nil(t,e)

	_,o,e  := r.DoRender(b,tb,nil)
	assert.Nil(t,e)
	t.Log(string(o))

}

func TestFilterLink(t *testing.T) {
	r := &renderer{}
	b,e := ioutil.ReadFile(TestPostPath)
	assert.Nil(t,e)

	tb,e := ioutil.ReadFile(TestTempl)
	assert.Nil(t,e)

	_,o,e  := r.DoRender(b,tb,nil)
	assert.Nil(t,e)
	//t.Log(string(o))

	ConvertLink(o)
}

func TestIsInternal(t *testing.T) {
	assert.False(t,IsInternal("http://www.papillon.io"))
	assert.True(t,IsInternal("./internal.css"))
}

func TestIsInternal2(t *testing.T) {
	link := "http://www.baidu.com"
	b,e := regexp.MatchString("https?://\\S+",link)
	assert.Nil(t,e)
	assert.True(t,b)

	link = "./css/style.css"
	b,e = regexp.MatchString("^[http|https]://*.?$",link)
	assert.Nil(t,e)
	assert.True(t,b)
}
