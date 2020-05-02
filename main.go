package main

import (
	"fmt"
	"log"

	"github.com/juliotorresmoreno/proxy-logger/config"
	"github.com/juliotorresmoreno/proxy-logger/server"
	"github.com/juliotorresmoreno/proxy-logger/status"
)

func main() {
	loop := make(chan bool)

	log.SetFlags(log.Lshortfile | log.Ltime | log.LstdFlags)
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	if conf.Status {
		statusSvr := status.NewServer()
		go statusSvr.Start(":4040")
	}

	svr := server.NewServer(conf)
	go svr.Listen()
	fmt.Printf("Listening server on %v\n", conf.Addr)

	<-loop
}
