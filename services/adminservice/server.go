package adminservice

import (
	"log"
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

func (s *Server) Listen() {
	config, err := config.GetConfig()
	if err != nil {
		log.Println(err)
		return
	}
	if config.Admin.Enabled != "true" {
		return
	}
	if config.Admin.Proto == "https" {
		err = s.Server.ListenAndServeTLS(config.PemPath, config.KeyPath)
	} else {
		err = s.Server.ListenAndServe()
	}
	if err != nil {
		log.Println("Admin error:", err)
	}
}
