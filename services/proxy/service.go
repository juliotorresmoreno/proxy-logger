package proxy

import (
	"net"
	"strings"
)

// Proxy Componente encargado de servir de proxy
type Proxy struct {
	listener    net.Listener
	adminServer *AdminServer
}

// NewProxy contructor del proxy
func NewProxy() *Proxy {
	return new(Proxy)
}

// Listen Habilita un socket para escuchar peticiones
func (p *Proxy) Listen() error {
	var err error
	if p.listener != nil {
		if err = p.listener.Close(); err != nil {
			return err
		}
	}
	p.listener, err = net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}
	for {
		conn, err := p.listener.Accept()
		if err != nil {
			continue
		}
		go p.dispatch(conn)
	}
}

// ListenWithAdmin Habilita un socket para escuchar peticiones
func (p *Proxy) ListenWithAdmin() error {
	if p.adminServer == nil {
		p.adminServer = newAdminServer()
	}
	go p.adminServer.Listen()
	return p.Listen()
}

func (p *Proxy) dispatch(conn net.Conn) {
	handlers := []func(net.Conn, []byte) bool{
		p.handleHTTP,
		p.handleCONNECT,
		p.handleTCP,
	}
	for {
		tmp := make([]byte, 1024*64)
		_, err := conn.Read(tmp)
		if err != nil {
			continue
		}
		buffer := []byte(strings.Trim(string(tmp), string([]byte{0})))
		for _, handler := range handlers {
			if handler(conn, buffer) {
				break
			}
		}
	}
}
