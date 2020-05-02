package server

import (
	"log"
	"math/rand"
	"net"

	"github.com/juliotorresmoreno/proxy-logger/config"
	"github.com/labstack/echo"
)

type Server struct {
	Addr      string
	Hosts     []string
	Algorithm string
	ProxyHTTP string
}

func NewServer(conf config.Config) *Server {
	return &Server{
		Addr:      conf.Addr,
		Hosts:     conf.Hosts,
		Algorithm: conf.Algorithm,
		ProxyHTTP: conf.ProxyHTTP,
	}
}

func (svr Server) getHost(isCONNECT bool) string {
	if isCONNECT {
		return svr.ProxyHTTP
	}
	n := rand.Intn(1)
	host := svr.Hosts[n]
	return host
}

func (svr Server) Listen() {
	s, err := net.Listen("tcp", svr.Addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Println(err)
		}
		go svr.listenClient(conn)
	}
}

func (svr Server) listenClient(conn net.Conn) error {
	err := svr.handle(conn)
	if err != nil {
		log.Println("error", err.Error())
		return echo.NewHTTPError(500, err)
	}
	return nil
}
