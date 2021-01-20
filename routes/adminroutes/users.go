package adminroutes

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/proxy-logger/helpers"
	"github.com/juliotorresmoreno/proxy-logger/services/authservice"
)

func newUsersRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/sign-in", usersSignIn).Methods("POST")
	router.HandleFunc("/sign-up", usersSignUp).Methods("POST")

	return router
}

type usersSignInBody struct {
	Username string
	Password string
}

func usersSignIn(w http.ResponseWriter, r *http.Request) {
	var err error
	u := &usersSignInBody{}
	err = helpers.ReadData(r.Body, u)
	if helpers.HandleHTTPError(w, r, err, 400) {
		return
	}
	_, err = authservice.SignIn(&authservice.User{
		Username: u.Username,
		Password: u.Password,
	})
	if helpers.HandleHTTPError(w, r, err, 401) {
		return
	}
	jwtToken, err := createJwtToken(u.Username)
	if helpers.HandleHTTPError(w, r, err, 500) {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(map[string]string{
		"jwtToken": jwtToken,
	})
	if helpers.HandleHTTPError(w, r, err, 500) {
		return
	}
}

type usersSignUpBody struct {
	Username string
	Password string
}

func usersSignUp(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Not implemnted")
	helpers.HandleHTTPError(w, r, err, 400)
}

func createJwtToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Unix(),
	})
	hmacSampleSecret := ""

	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}
