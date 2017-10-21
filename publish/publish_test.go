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
	hash,err := pub.AddDir("/Users/DeepSea/Documents/workspace/workspace_go/src/github.com/gogank/papillon/publish/test/.")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t,err)
	t.Log(hash)
}

func TestPublishImpl_AddFile2(t *testing.T) {
	err := pub.shell.Publish("QmRHfsJ9vR44vnr3W1Eq4Ef1J4dEgaQvsXziuEU7bjpEDT","QmcPdt8s9AJRzS8Yg6aLEwwqEzzDxmitaco44Hmke2tzLJ")
	if err != nil {
		t.Error(err)
	}

}

