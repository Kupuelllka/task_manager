package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	General struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"general"`
	Log struct {
		Target string `yaml:"target"`
		Level  string `yaml:"level"`
		File   string `yaml:"file"`
	} `yaml:"logger"`
}

func GetConfig(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
