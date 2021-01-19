package adminservice

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/proxy-logger/config"
	"github.com/juliotorresmoreno/proxy-logger/middlewares/loggermiddleware"
	"github.com/juliotorresmoreno/proxy-logger/routes/adminroutes"
)

type Server struct {
	*http.Server
}

func NewServer() *Server {
	s := &Server{}
	config, _ := config.GetConfig()
	router := adminroutes.NewRouter().(*mux.Router)
	router.Use(loggermiddleware.LogRequest)
	router.Use(gziphandler.GzipHandler)
	s.Server = &http.Server{
		Addr:    config.Admin.Addr,
		Handler: router,
	}
	return s
}
