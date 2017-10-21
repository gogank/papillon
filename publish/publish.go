package publish

import (
	api"github.com/ipfs/go-ipfs-api"
	"github.com/gogank/papillon/utils"
	"strings"
	"os/exec"
)

type Publish interface {
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

func (publish *PublishImpl) AddDirCmd(dir string) (string,error) {
	res,err := exec.Command("ipfs", "add","-r",dir).Output()
	if err!= nil {
		return "",err
	}
	str := string(res)
	strs := strings.Split(str," ")

	return strs[len(strs)-2],nil
}
