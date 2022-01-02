package utils

import (
	"context"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const UtilsConfigPath = "configs/utils.yaml"

type Logger struct {
	Output string `yaml:"output"`
	FileOutput string `yaml:"file_output"`
	Level string `yaml:"level"`
}

type Config struct {
	Logger Logger `yaml:"logger"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	fileData, err := ioutil.ReadFile(UtilsConfigPath)
	if err != nil {
		return nil, err
	}

	resConfig := Config{}
	if err := yaml.Unmarshal(fileData, &resConfig); err != nil {
		return nil, err
	}

	return &resConfig, nil
}

