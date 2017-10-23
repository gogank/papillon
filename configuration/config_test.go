package config

import (
	"github.com/gogank/papillon/utils"
	"testing"
)

func TestNewConfig(t *testing.T) {
	path := "./config/config.toml"

	config := NewConfig(path)

	t.Logf(config.GetString(utils.DIR_PUBLIC))
}
