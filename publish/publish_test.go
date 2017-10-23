package publish

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var pub *PublishImpl

func init() {
	pub = NewPublishImpl()
}

func TestPublishImpl_AddFile(t *testing.T) {
	hash, err := pub.AddFile("./test/gogank.jpg")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
	t.Log(hash)
}

func TestPublishImpl_AddDir(t *testing.T) {
	hash, err := pub.AddDir("./test")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
	t.Log(hash)
}

func TestPublishImpl_AddFile2(t *testing.T) {
	res, err := pub.AddDirCmd("./test")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
	t.Log(res)
}

func TestPublishImpl_PublishCmd(t *testing.T) {
	peer, err := pub.PublishCmd()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
	t.Log(peer)
}

func TestPublishImpl_NamePublish(t *testing.T) {
	id, err := pub.LocalID()
	assert.Nil(t, err)
	t.Log(id)
	//err = pub.NamePublish(id,"QmPgHm5A9vzb1xETRFC9jzSbye14yDkVwozkRyZ6zokp83")
	//assert.Nil(t,err)
}
