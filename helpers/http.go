package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/proxy-logger/config"
)

func HandleStripPrefix(mux *mux.Router, prefix string, handler http.Handler) *mux.Route {
	return mux.PathPrefix(prefix).Handler(http.StripPrefix(prefix, handler))
}

func ReadData(reader io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func HandleHTTPError(w http.ResponseWriter, r *http.Request, err error, status int) bool {
	if err == nil {
		return false
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"message": err.Error(),
	})
	return true
}

// CreateJwtToken .
func CreateJwtToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Unix(),
	})
	config, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	hmacSampleSecret := config.Admin.Secret

	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}
