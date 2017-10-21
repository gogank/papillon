package publish

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
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
	objs,_ := pub.shell.FileList(fmt.Sprintf("/ipfs/%s", "QmXqBwdxJHZPS8he9LZYpx1APyh4Kx8jEHccnmuVKyseag"))
	lists := objs.Links
	t.Log(len(lists))
}