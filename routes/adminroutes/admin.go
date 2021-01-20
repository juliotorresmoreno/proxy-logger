package adminroutes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/proxy-logger/helpers"
)

type onlyFiles struct {
	fs http.FileSystem
}

func (el onlyFiles) Open(name string) (http.File, error) {
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

	attachAPI(router)
	attachFiles(router)

	return router
}

func attachAPI(router *mux.Router) {
	api := mux.NewRouter()
	helpers.HandleStripPrefix(router, "/api", api)

	helpers.HandleStripPrefix(api, "/users", newUsersRouter())
}

func attachFiles(router *mux.Router) {
	staticFiles := http.FileServer(onlyFiles{http.Dir("./bower_components")})
	helpers.HandleStripPrefix(router, "/static", staticFiles).Methods("GET")

	jsFiles := http.FileServer(onlyFiles{http.Dir("./public/js")})
	helpers.HandleStripPrefix(router, "/js", jsFiles).Methods("GET")

	cssFiles := http.FileServer(onlyFiles{http.Dir("./public/css")})
	helpers.HandleStripPrefix(router, "/css", cssFiles).Methods("GET")

	router.PathPrefix("/").HandlerFunc(indexFile).Methods("GET")
}
