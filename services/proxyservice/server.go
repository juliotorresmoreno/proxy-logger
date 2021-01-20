package proxyservice

import (
	"log"
	"net/http"

	"github.com/juliotorresmoreno/proxy-logger/config"
	"github.com/juliotorresmoreno/proxy-logger/routes/proxyroutes"
)

type Server struct {
	*http.Server
}

func NewServer() *Server {
	s := &Server{}
	config, _ := config.GetConfig()
	s.Server = &http.Server{
		Addr:    config.Addr,
		Handler: proxyroutes.NewRouter(),
	}
	return s
}
func (s *Server) Listen() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	if config.Proto == "https" {
		err = s.ListenAndServeTLS(config.PemPath, config.KeyPath)
	} else {
		err = s.ListenAndServe()
	}
	if err != nil {
		log.Fatal("Admin error:", err)
	}
}
