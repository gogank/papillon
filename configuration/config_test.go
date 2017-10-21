package config

import (
	"testing"
	"github.com/gogank/papillon/utils"
)

func TestNewConfig(t *testing.T) {
	path := "./config/config.toml"

	config := NewConfig(path)

	t.Logf(config.GetString(utils.DIR_PUBLIC))
}