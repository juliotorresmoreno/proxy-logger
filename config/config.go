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
	Enabled string `yaml:"enabled"`
	Proto   string `yaml:"proto"`
	Addr    string `yaml:"addr"`
	Secret  string `yaml:"secret"`
	PemPath string `yaml:"pem_path"`
	KeyPath string `yaml:"key_path"`
}

// ACL .
type ACL struct {
	Default string   `yaml:"default"`
	Permit  []string `yaml:"permit"`
	Block   []string `yaml:"block"`
}

// Reverse .
type Reverse []struct {
	Host    string `yaml:"host"`
	Forward string `yaml:"forward"`
}

// Config .
type Config struct {
	Addr        string        `yaml:"addr"`
	RedisURL    string        `yaml:"redis_url"`
	PemPath     string        `yaml:"pem_path"`
	KeyPath     string        `yaml:"key_path"`
	Proto       string        `yaml:"proto"`
	Credentials []Credentials `yaml:"credentials"`
	Admin       Admin         `yaml:"admin"`
	ACL         ACL           `yaml:"ACL"`
	Reverse     Reverse       `yaml:"reverse"`
	credentials map[string]Credentials
	acl         struct {
		permit map[string]bool
		block  map[string]bool
	}
	reverse map[string]string
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

// GetACLPermit .
func (c Config) GetACLPermit() map[string]bool {
	return c.acl.permit
}

// GetACLBlock .
func (c Config) GetACLBlock() map[string]bool {
	return c.acl.block
}

// MapACL .
func (c *Config) MapACL() {
	if c.acl.block == nil {
		c.acl.block = map[string]bool{}
	}
	if c.acl.permit == nil {
		c.acl.permit = map[string]bool{}
	}
	for _, host := range c.ACL.Permit {
		c.acl.permit[host] = true
	}
	for _, host := range c.ACL.Block {
		c.acl.block[host] = true
	}
}

// GetReverse .
func (c Config) GetReverse() map[string]string {
	return c.reverse
}

// MapReverse .
func (c *Config) MapReverse() {
	if c.reverse == nil {
		c.reverse = map[string]string{}
	}
	for _, reverse := range c.Reverse {
		c.reverse[reverse.Host] = reverse.Forward
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
		result.MapACL()
		result.MapReverse()
		config = result
	}
	return config.(Config), nil
}
