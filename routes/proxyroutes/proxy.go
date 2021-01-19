package proxyroutes

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/juliotorresmoreno/proxy-logger/services/authservice"
	"github.com/juliotorresmoreno/proxy-logger/services/loggerservice"
)

// BasicAuth .
func BasicAuth(credentials string) error {
	decoded, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		return errors.New("Unauthorized")
	}
	splitData := strings.Split(string(decoded), ":")
	username := splitData[0]
	password := splitData[1]
	u := &authservice.User{
		Username: username,
		Password: password,
	}
	if ok, _ := authservice.SignIn(u); !ok {
		return errors.New("Unauthorized")
	}
	return nil
}

// NewRouter .
func NewRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		credentials := r.Header.Get("Authorization")
		if len(credentials) < 6 {
			credentials = r.Header.Get("Proxy-Authorization")
		}
		if len(credentials) > 6 && credentials[:5] == "Basic" {
			credentials = credentials[6:]
			if BasicAuth(credentials) != nil {
				w.WriteHeader(401)
				w.Write([]byte("Unauthorized"))
				return
			}
		} else {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized"))
			return
		}
		httpWriter := loggerservice.NewHTTPWriter(w, r)
		if r.Method == http.MethodConnect {
			handleTunneling(w, r)
			httpWriter.Protocol = "https"
			httpWriter.Register()
		} else {
			handleHTTP(httpWriter, r)
			httpWriter.Protocol = "http"
			httpWriter.Register()
		}
	}
}
