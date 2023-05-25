package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConfigReaderInterface interface {
	ReadConfig(path string) (*Config, error)
}

type ConfigReader struct{}

func (d *ConfigReader) ReadConfig(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return &Config{}, err
	}

	return &cfg, nil
}
