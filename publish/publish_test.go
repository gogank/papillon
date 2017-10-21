package publish

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	shellUrl     = "localhost:5001"
)

var pub *PublishImpl

func init()  {
	pub = NewPublishImpl(shellUrl)
}

func TestPublishImpl_AddFile(t *testing.T) {
	hash,err := pub.AddFile("./test/gogank.jpg")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t,err)
	t.Log(hash)
}

func TestPublishImpl_AddDir(t *testing.T) {
	hash,err := pub.AddDir("./test")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t,err)
	t.Log(hash)
}

func TestPublishImpl_AddFile2(t *testing.T) {
	res,err := pub.AddDirCmd("./test")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t,err)
	t.Log(res)
}


