package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMkdir(t *testing.T) {
	res := Mkdir("test")
	assert.Equal(t, true, res)
}

func TestMkfile(t *testing.T) {
	res := Mkfile("./test/test.html", []byte("papillon"))
	assert.Equal(t, true, res)
}

func TestExistDir(t *testing.T) {
	res := ExistDir("./test")
	assert.Equal(t, true, res)
}

func TestListDir(t *testing.T) {
	file, _ := ListDir("./test", "html")
	assert.Equal(t, "test.html", file[0])
}
