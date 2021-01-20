package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Credentials .
type Credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Admin .
type Admin struct {
	Addr   string `yaml:"addr"`
	Secret string `yaml:"secret"`
}

// Config .
type Config struct {
	Addr        string        `yaml:"addr"`
	RedisURL    string        `yaml:"redis_url"`
	Credentials []Credentials `yaml:"credentials"`
	Admin       Admin         `yaml:"admin"`
	credentials map[string]Credentials
}

// GetCredencial .
func (c Config) GetCredencial(username string) Credentials {
	if c.credentials == nil {
		return Credentials{}
	}
	return c.credentials[username]
}

// MapCredencials .
func (c *Config) MapCredencials() {
	if c.credentials == nil {
		c.credentials = map[string]Credentials{}
	}
	for _, credential := range c.Credentials {
		c.credentials[credential.Username] = credential
	}
}

var config interface{}

// GetConfig .
func GetConfig() (Config, error) {
	if config == nil {
		result := Config{}
		f, err := os.Open("config.yml")
		if err != nil {
			return result, err
		}
		buff, err := ioutil.ReadAll(f)
		if err != nil {
			return result, err
		}
		err = yaml.Unmarshal(buff, &result)
		if err != nil {
			return result, err
		}
		result.MapCredencials()
		config = result
	}
	return config.(Config), nil
}
