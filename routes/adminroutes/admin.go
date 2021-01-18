package adminroutes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter .
func NewRouter() http.Handler {
	router := mux.NewRouter()

	staticFiles := http.FileServer(http.Dir("./bower_components"))
	router.PathPrefix("/static").
		Handler(http.StripPrefix("/static", staticFiles)).
		Methods("GET")
	jsFiles := http.FileServer(http.Dir("./public/js"))
	router.PathPrefix("/js").
		Handler(http.StripPrefix("/js", jsFiles)).
		Methods("GET")
	cssFiles := http.FileServer(http.Dir("./public/css"))
	router.PathPrefix("/css").
		Handler(http.StripPrefix("/css", cssFiles)).
		Methods("GET")
	router.PathPrefix("/").
		HandlerFunc(indexFile).
		Methods("GET")
	return router
}
