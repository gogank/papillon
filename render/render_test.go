package render

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRenderImpl_RenderSingle(t *testing.T) {
	r := &renderer{}
	_,o,e  := r.Single("../test/single.md")
	assert.Nil(t,e)
	t.Log(string(o))
}

func TestReadPostConfig(t *testing.T) {
	r := &renderer{}
	_,o,e  := r.Single("../test/single.md")
	assert.Nil(t,e)
	ReadPostConfig(o)
}