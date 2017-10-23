package mapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPut(t *testing.T) {
	key := "./test/gogank.jpg"
	hash1, err := Put("./test/gogank.jpg", "./test")
	if err != nil {
		t.Error(err)
	}
	t.Log(hash1)
	hash2, _ := Get(key)
	assert.Equal(t, hash1, hash2)
}

func TestWalkDir(t *testing.T) {
	file, _ := WalkDir("./test")
	t.Log(len(file))
	for i := 0; i < len(file); i++ {
		t.Log(file[i])
	}
}

func TestWalkDirCmd(t *testing.T) {
	WalkDirCmd("/Users/chenquan/Workspace/go/src/github.com/gogank/papillon/build/blog/public")
}
