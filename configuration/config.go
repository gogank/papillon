package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Config manager
type Config struct {
	conf *viper.Viper
	lock *sync.RWMutex
}

// NewConfig returns a new instance of Config by configPath.
func NewConfig(configPath string) *Config {
	vp := viper.New()
	vp.SetConfigFile(configPath)
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Reading config file: %s error, %s !", configPath, err.Error()))
	}
	return &Config{
		conf: vp,
		lock: &sync.RWMutex{},
	}
}

// NewRawConfig new config without underlying config file.
func NewRawConfig() *Config {
	return &Config{
		conf: viper.New(),
		lock: &sync.RWMutex{},
	}
}

//Get a key return interface
func (cf *Config) Get(key string) interface{} {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.Get(key)
}

//GetString a key return string
func (cf *Config) GetString(key string) string {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetString(key)
}

//GetInt a key return int
func (cf *Config) GetInt(key string) int {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetInt(key)
}

//GetInt64 a key return int64
func (cf *Config) GetInt64(key string) int64 {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetInt64(key)
}

//GetFloat64 a key return Float64
func (cf *Config) GetFloat64(key string) float64 {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetFloat64(key)
}

//GetBool a key return int64
func (cf *Config) GetBool(key string) bool {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetBool(key)
}

//Set key:value
func (cf *Config) Set(key string, value interface{}) {
	cf.lock.Lock()
	defer cf.lock.Unlock()
	cf.conf.Set(key, value)
}

// ContainsKey judge whether the key is set in the config.
func (cf *Config) ContainsKey(key string) bool {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.IsSet(key)
}

// MergeConfig merge config by the config file path, the file try to merge should have same format.
func (cf *Config) MergeConfig(configPath string) (*Config, error) {
	cf.lock.Lock()
	defer cf.lock.Unlock()
	f, err := os.Open(configPath)
	if err != nil {
		return cf, err
	}
	err = cf.conf.MergeConfig(f)
	return cf, err
}

// OnConfigChange register function to invoke when config file change.
func (cf *Config) OnConfigChange(run func(in fsnotify.Event)) {
	cf.conf.OnConfigChange(run)
}
