package proxyroutes

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/juliotorresmoreno/proxy-logger/services/authservice"
	"github.com/juliotorresmoreno/proxy-logger/services/loggerservice"
)

const authenticationRequiredHTML = `
<!DOCTYPE HTML "-//IETF//DTD HTML 2.0//EN">
<html><head>
<title>407 Proxy Authentication Required</title>
</head><body>
<h1>Proxy Authentication Required</h1>
<p>This server could not verify that you
are authorized to access the document
requested.  Either you supplied the wrong
credentials (e.g., bad password), or your
browser doesn't understand how to supply
the credentials required.</p>
</body></html>
`

func AuthRequired(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Proxy-Authenticate", "Basic realm=\"Proxy Logger\"")
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(407)
	fmt.Fprintf(w, authenticationRequiredHTML)
}

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
				AuthRequired(w, r)
				return
			}
		} else {
			AuthRequired(w, r)
			return
		}
		httpWriter := loggerservice.NewHTTPWriter(w, r)
		if r.Method == http.MethodConnect {
			handleTunneling(w, r)
			httpWriter.Protocol = "https"
			httpWriter.Register()
		} else {
			handleHTTP(httpWriter, r)
			httpWriter.Protocol = strings.Split(r.RequestURI, ":")[0]
			httpWriter.Register()
		}
	}
}
