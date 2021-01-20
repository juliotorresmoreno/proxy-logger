package proxyroutes

import (
	"strings"

	"github.com/juliotorresmoreno/proxy-logger/config"
)

func isSecure(host string) bool {
	config, err := config.GetConfig()
	if err != nil {
		return false
	}
	if config.ACL.Default == "block" {
		return isSecureFromPermit(host)
	}
	return isSecureFromBlock(host)
}

func isSecureFromBlock(host string) bool {
	config, _ := config.GetConfig()
	acl := config.GetACLBlock()
	if _, ok := acl[host]; ok {
		return false
	}
	host = strings.Split(host, ":")[0]
	if _, ok := acl[host]; ok {
		return false
	}
	return true
}

func isSecureFromPermit(host string) bool {
	config, _ := config.GetConfig()
	acl := config.GetACLPermit()
	if _, ok := acl[host]; ok {
		return true
	}
	host = strings.Split(host, ":")[0]
	if _, ok := acl[host]; ok {
		return true
	}
	return false
}
