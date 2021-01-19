package adminroutes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (el justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := el.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return []os.FileInfo{}, nil
}

// NewRouter .
func NewRouter() http.Handler {
	router := mux.NewRouter()

	staticFiles := justFilesFilesystem{http.Dir("./bower_components")}
	router.PathPrefix("/static").
		Handler(http.StripPrefix("/static", http.FileServer(staticFiles))).
		Methods("GET")
	jsFiles := justFilesFilesystem{http.Dir("./public/js")}
	router.PathPrefix("/js").
		Handler(http.StripPrefix("/js", http.FileServer(jsFiles))).
		Methods("GET")
	cssFiles := justFilesFilesystem{http.Dir("./public/css")}
	router.PathPrefix("/css").
		Handler(http.StripPrefix("/css", http.FileServer(cssFiles))).
		Methods("GET")
	router.PathPrefix("/").
		HandlerFunc(indexFile).
		Methods("GET")
	return router
}
