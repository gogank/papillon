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
	hash2 ,_ := Get(key)
	assert.Equal(t,hash1,hash2)
}


func TestWalkDir(t *testing.T) {
	file,e := WalkDir("./test")
	assert.Nil(t,e)
	t.Log(len(file))
	for i:=0;i<len(file);i++ {
		t.Log(file[i])
	}
}
