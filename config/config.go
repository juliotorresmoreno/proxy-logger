package config

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Algorithm string
	Addr      string
	Hosts     []string
	ProxyHTTP string `yaml:"proxy_http"`
}

func GetConfig() (Config, error) {
	c := Config{}
	f, err := os.Open("./config.yaml")
	if err != nil {
		return c, err
	}
	err = yaml.NewDecoder(f).Decode(&c)
	return c, err
}
