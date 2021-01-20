package adminroutes

import (
	"encoding/json"
	"errors"
	"net/http"

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
	jwtToken, err := helpers.CreateJwtToken(u.Username)
	if helpers.HandleHTTPError(w, r, err, 500) {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(map[string]string{
		"token":    jwtToken,
		"username": u.Username,
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
