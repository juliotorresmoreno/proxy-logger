package main

import (
	"fmt"
	"log"
	"os"

	"github.com/juliotorresmoreno/proxy-logger/config"
	"github.com/juliotorresmoreno/proxy-logger/services/adminservice"
	"github.com/juliotorresmoreno/proxy-logger/services/loggerservice"
	"github.com/juliotorresmoreno/proxy-logger/services/proxyservice"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	loggerservice.SetLoggerWriter(os.Stdout)
	go adminservice.NewServer().Listen()
	go proxyservice.NewServer().Listen()
	fmt.Println("Listening on", config.Addr)
	<-make(chan byte)
}
