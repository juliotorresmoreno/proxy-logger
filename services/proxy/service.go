package proxy

import (
	"net"
	"strconv"
	"strings"
)

const sep = "\r\n"

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

func getBodyLen(b []byte) int {
	head := strings.ToLower(string(b))
	pos := strings.Index(head, "content-length")
	defaultValue := 256 * 1024
	if pos == -1 {
		return defaultValue
	}
	tmp := head[pos+15:]
	pos = strings.Index(tmp, "\n")
	tmp = strings.TrimSpace(tmp[:pos])
	if tmp != "" {
		length, err := strconv.Atoi(tmp)
		if err != nil {
			return defaultValue
		}
		return length
	}
	return defaultValue
}

func getUpgrade(conn net.Conn) ([]byte, error) {
	var b []byte
	upgrade := make([]byte, 1024*8)
	n, err := conn.Read(upgrade)
	if err != nil {
		return b, err
	}
	return upgrade[:n], nil
}

func (p *Proxy) dispatch(conn net.Conn) {
	handlers := []func(*validConn, []byte) bool{
		p.handleHTTP,
		p.handleCONNECT,
		p.handleTCP,
	}
	vConn := &validConn{"connection", conn, true}
	upgrade, err := getUpgrade(conn)
	if err != nil {
		return
	}

	for _, handler := range handlers {
		if handler(vConn, upgrade) {
			break
		}
	}
}
