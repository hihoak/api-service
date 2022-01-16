package server

import (
"context"
"gopkg.in/yaml.v3"
"io/ioutil"
)

const serverConfigPath = "configs/server.yaml"

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Server Server `yaml:"server"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	fileData, err := ioutil.ReadFile(serverConfigPath)
	if err != nil {
		return nil, err
	}

	resConfig := Config{}
	if err := yaml.Unmarshal(fileData, &resConfig); err != nil {
		return nil, err
	}

	return &resConfig, nil
}


