package render

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

var TestPostPath = "../test/single.md"


func TestRenderImpl_RenderSingle(t *testing.T) {
	r := &renderer{}
	b,e := ioutil.ReadFile(TestPostPath)
	assert.Nil(t,e)
	_,o,e  := r.Single(b)
	assert.Nil(t,e)
	t.Log(string(o))
}

func TestReadPostConfig(t *testing.T) {
	b,e := ioutil.ReadFile(TestPostPath)
	assert.Nil(t,e)
	ReadPostConfig(b)
}