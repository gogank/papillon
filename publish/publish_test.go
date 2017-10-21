package publish

import "testing"

const (
	shellUrl     = "localhost:5001"
)

func TestPublishImpl_AddFile(t *testing.T) {
	publish := NewPublishImpl(shellUrl)
	hash,err := publish.AddFile("./test/gogank.jpg")
	if err != nil {
		t.Error(err)
	}
	t.Log(hash)
}