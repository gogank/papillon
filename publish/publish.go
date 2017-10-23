package publish

import (
	"os/exec"
	"strings"

	"github.com/gogank/papillon/configuration"
	"github.com/gogank/papillon/utils"
	api "github.com/mikesun/go-ipfs-api"
	"github.com/pkg/errors"
)

//Publish interface supply the core ipfs upload functions
type Publish interface {
	AddFile(filename string) (string, error)
	AddDir(dir string) (string, error)
}

//Impl implements the Publish interface
type Impl struct {
	shell *api.Shell
	cnf   *config.Config
}

//NewImpl return a publish instance
func NewImpl() *Impl {
	con := config.NewConfig("./config.toml")
	return &Impl{
		shell: api.NewShell(con.GetString(utils.CommonURL)),
		cnf:   con,
	}
}

//AddFile add a file into ipfs network
func (publish *Impl) AddFile(filename string) (string, error) {
	contents, err := utils.ReadFile(filename)
	if err != nil {
		return "", err
	}
	reader := strings.NewReader(string(contents))
	hash, err := publish.shell.Add(reader)
	return hash, err
}

//AddDir add a Dir into ipfs network
func (publish *Impl) AddDir(dir string) (string, error) {
	hash, err := publish.shell.AddDir(dir)
	return hash, err
}

//AddDirCmd add a Dir into ipfs network by native shell command
func (publish *Impl) AddDirCmd(dir string) (string, error) {
	res, err := exec.Command("ipfs", "add", "-r", dir).Output()
	if err != nil {
		return "", err
	}
	str := string(res)
	strs := strings.Split(str, " ")

	return strs[len(strs)-2], nil
}

//NamePublish same as `ipfs name publish <hash>`
func (publish *Impl) NamePublish(name, hash string) error {
	return publish.shell.Publish(name, hash)
}

//LocalID get local peerID
func (publish *Impl) LocalID() (string, error) {
	id, err := publish.shell.ID()
	if err != nil {
		return "", err
	}
	return id.ID, nil
}

//PublishCmd use native command shell to `ipfs name publish`
func (publish *Impl) PublishCmd() (string, error) {
	dir := publish.cnf.GetString(utils.DirPublic) + "/index.html"
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
		return "", errors.New("publish Failed,please check the ipfs server")
	}
	peer := strs[2]
	length := len(peer)
	peer = peer[:length-1]
	return peer, nil
}
