package config

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"github.com/rs/zerolog"
	"os"
)

type Spotify struct {
	ResponseType string `yaml:"response_type"`
	RedirectUri string `yaml:"redirect_uri"`
	State string `yaml:"state"`
	Scope string `yaml:"scope"`
	ShowDialog string `yaml:"show_dialog"`
}

type Config struct {
	Spotify Spotify
}

func LoadConfig(ctx context.Context, path string) {
	file, err := os.Open(path)
	if err != nil {

	}
}
