package authservice

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/juliotorresmoreno/proxy-logger/config"
)

type User struct {
	Username string
	Password string
}

func password(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func SignIn(u *User) (bool, error) {
	config, _ := config.GetConfig()
	credential := config.GetCredencial(u.Username)
	if credential.Username != u.Username {
		return false, errors.New("Unauthorized")
	}
	if credential.Password != password(u.Password) {
		return false, errors.New("Unauthorized")
	}
	return true, nil
}
