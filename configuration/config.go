package config

import (
	"github.com/spf13/viper"
	"fmt"
	"sync"
	"github.com/fsnotify/fsnotify"
	"time"
	"os"
)

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

func (cf *Config) Get(key string) interface{} {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.Get(key)
}

func (cf *Config) GetString(key string) string {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetString(key)
}

func (cf *Config) GetInt(key string) int {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetInt(key)
}

func (cf *Config) GetInt64(key string) int64 {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetInt64(key)
}

func (cf *Config) GetFloat64(key string) float64 {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetFloat64(key)
}

func (cf *Config) GetBool(key string) bool {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetBool(key)
}

func (cf *Config) GetDuration(key string) time.Duration {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetDuration(key)
}

func (cf *Config) GetStringMap(key string) map[string]interface{} {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetStringMap(key)
}

func (cf *Config) GetStringSlice(key string) []string {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetStringSlice(key)
}

func (cf *Config) GetBytes(key string) uint {
	cf.lock.RLock()
	defer cf.lock.RUnlock()
	return cf.conf.GetSizeInBytes(key)
}

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

func (cf *Config) Print() {
	keys := cf.conf.AllKeys()
	for _, key := range keys {
		fmt.Printf("key: %s, value: %v\n", key, cf.Get(key))
	}
}

func (cf *Config) equals(anotherConfig *Config) bool {
	if len(cf.conf.AllKeys()) != len(anotherConfig.conf.AllKeys()) {
		return false
	}
	for _, key := range anotherConfig.conf.AllKeys() {
		if !cf.ContainsKey(key) {
			fmt.Printf("No value for key %s found \n", key)
			return false
		}
	}
	return true
}
