package mapper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	key := "./test/gogank.jpg"
	hash1,err := Put("./test/gogank.jpg")
	if err != nil{
		t.Error(err)
	}
	t.Log(hash1)
	hash2 := Get(key)
	assert.Equal(t,hash1,hash2)
}


func TestWalkDir(t *testing.T) {
	file,_ := WalkDir("./test")
	t.Log(len(file))
	for i:=0;i<len(file);i++ {
		t.Log(file[i])
	}
}
