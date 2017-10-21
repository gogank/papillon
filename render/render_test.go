package render

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

var TestPostPath = "../test/single.md"
var TestTempl = "../test/post.hbs"



func TestReadPostConfig(t *testing.T) {
	b,e := ioutil.ReadFile(TestTempl)
	assert.Nil(t,e)
	ReadPostConfig(b)
}

func TestRenderer_DoRender(t *testing.T) {
	r := &renderer{}
	b,e := ioutil.ReadFile(TestPostPath)
	assert.Nil(t,e)

	tb,e := ioutil.ReadFile(TestTempl)
	assert.Nil(t,e)

	_,o,e  := r.DoRender(b,tb)
	assert.Nil(t,e)
	t.Log(string(o))

}