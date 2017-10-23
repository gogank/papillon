package publish

import (
	api"github.com/mikesun/go-ipfs-api"
	"github.com/gogank/papillon/utils"
	"strings"
	"os/exec"
	"github.com/pkg/errors"
	"github.com/gogank/papillon/configuration"
)

type Publish interface {
	AddFile(filename string) (string, error)
	AddDir(dir string) (string, error)
}

type PublishImpl struct {
	shell *api.Shell
	cnf   *config.Config
}

func NewPublishImpl() *PublishImpl {
	con := config.NewConfig("./config.toml")
	return &PublishImpl{
		shell: api.NewShell(con.GetString(utils.COMMON_URL)),
		cnf:   con,
	}
}

func (publish *PublishImpl) AddFile(filename string) (string, error) {
	contents, err := utils.ReadFile(filename)
	if err != nil {
		return "", err
	}
	reader := strings.NewReader(string(contents))
	hash, err := publish.shell.Add(reader)
	return hash, err
}

func (publish *PublishImpl) AddDir(dir string) (string, error) {
	hash, err := publish.shell.AddDir(dir)
	return hash, err
}

func (publish *PublishImpl) AddDirCmd(dir string) (string, error) {
	res, err := exec.Command("ipfs", "add", "-r", dir).Output()
	if err != nil {
		return "", err
	}
	str := string(res)
	strs := strings.Split(str, " ")

	return strs[len(strs)-2], nil
}

func (publish *PublishImpl) NamePublish(name, hash string) (error) {
	return publish.shell.Publish(name, hash)
}

func (publish *PublishImpl) LocalID() (string, error) {
	id, err := publish.shell.ID()
	if err != nil {
		return "", err
	}
	return id.ID, nil
}

func (publish *PublishImpl) PublishCmd() (string, error) {
	dir := publish.cnf.GetString(utils.DIR_PUBLIC) + "/index.html"
	hash, err := publish.AddDirCmd(dir)
	if err != nil {
		return "", err
	}
	res, err := exec.Command("ipfs", "name", "publish", hash).Output()
	if err != nil {
		return "", err
	}
	str := string(res)
	strs := strings.Split(str, " ")
	if len(strs) != 4 {
		return "", errors.New("Publish Failed,please check the ipfs server.")
	}
	peer := strs[2]
	length := len(peer)
	peer = peer[:length-1]
	return peer, nil
}
