package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

var config Config
var mutex = &sync.Mutex{}

//InitConfig creates a config object from the give filename
func InitConfig(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		return err
	}
	mutex.Lock()
	config = *cfg
	mutex.Unlock()
	return nil
}

//GetConfig returns the static config object
func GetConfig() Config {
	return config
}
