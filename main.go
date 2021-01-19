package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/juliotorresmoreno/proxy-logger/config"
	"github.com/juliotorresmoreno/proxy-logger/routes/proxyroutes"
	"github.com/juliotorresmoreno/proxy-logger/services/adminservice"
	"github.com/juliotorresmoreno/proxy-logger/services/loggerservice"
)

func main() {
	var pemPath string
	flag.StringVar(&pemPath, "pem", "server.pem", "path to pem file")
	var keyPath string
	flag.StringVar(&keyPath, "key", "server.key", "path to key file")
	var proto string
	flag.StringVar(&proto, "proto", "http", "Proxy protocol (http or https)")
	flag.Parse()
	if proto != "http" && proto != "https" {
		log.Fatal("Protocol must be either http or https")
	}
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	loggerservice.SetLoggerWriter(os.Stdout)
	server := &http.Server{
		Addr:    config.Addr,
		Handler: proxyroutes.NewRouter(),
	}
	adminServer := adminservice.NewServer()
	go adminServer.ListenAndServe()
	if proto == "http" {
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
	}
}
