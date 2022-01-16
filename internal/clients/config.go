package clients

import (
	"context"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const clientsConfigPath = "configs/clients.yaml"

type Spotify struct {
	ResponseType string `yaml:"response_type"`
	RedirectUri string `yaml:"redirect_uri"`
	State string `yaml:"state"`
	Scope string `yaml:"scope"`
	ShowDialog string `yaml:"show_dialog"`
	ClientId string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type Config struct {
	Spotify Spotify `yaml:"spotify"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	fileData, err := ioutil.ReadFile(clientsConfigPath)
	if err != nil {
		return nil, err
	}

	resConfig := Config{}
	if err := yaml.Unmarshal(fileData, &resConfig); err != nil {
		return nil, err
	}

	return &resConfig, nil
}
