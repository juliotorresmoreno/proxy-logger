package main

import (
	"github.com/juliotorresmoreno/proxy-logger/services/proxy"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	p := proxy.NewProxy()
	go p.ListenWithAdmin()
	go e.Start(":8080")
	<-make(chan bool)
}
