package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMkdir(t *testing.T) {
	res := Mkdir("test")
	assert.Equal(t,true,res)
}
