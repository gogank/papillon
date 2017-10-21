package publish

import (
	api"github.com/ipfs/go-ipfs-api"
	"github.com/gogank/papillon/utils"
	"strings"
)

type publish interface {
	AddFile(filename string) (string, error)
	AddDir(dir string) (string, error)
}

type PublishImpl struct {
	shell *api.Shell
}

func NewPublishImpl(url string) *PublishImpl {
	return &PublishImpl{
		shell: api.NewShell(url),
	}
}

func (publish *PublishImpl) AddFile(filename string) (string, error) {
	contents,err := utils.ReadFile(filename)
	if err != nil {
		return "",err
	}
	reader := strings.NewReader(string(contents))
	hash,err := publish.shell.Add(reader)
	return hash,err
}

func (publish *PublishImpl) AddDir(dir string) (string, error) {
	hash,err := publish.shell.AddDir(dir)
	return hash,err
}
