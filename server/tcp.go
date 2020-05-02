package server

import (
	"errors"
	"io"
	"net"
	"strings"
	"time"

	"github.com/juliotorresmoreno/proxy-logger/status"
)

func isTCP(h []byte) bool {
	request := strings.ToLower(string(h))
	if !strings.Contains(request, "sec-websocket-key:") {
		return false
	}
	if !strings.Contains(request, "upgrade: websocket") {
		return false
	}
	if strings.Contains(request, "HTTP/1.1") {
		return false
	}
	return true
}

func (svr Server) handleTCP(conn net.Conn, upgrade []byte, connect bool, log *status.Request) error {
	host := svr.getHost(connect)
	if host == "" {
		conn.Close()
		return errors.New("El host no es valido")
	}
	host = host[strings.Index(host, "//")+2:]
	remote, err := net.Dial("tcp", host)
	if err != nil {
		conn.Close()
		return err
	}
	vremote := &validConn{remote, true, log}
	vconn := &validConn{conn, true, log}
	go pipe(vremote, vconn)
	go pipe(vconn, vremote)
	_, err = remote.Write(upgrade)
	if log != nil {
		log.Raw = log.Raw + string(upgrade)
	}
	if err != nil {
		conn.Close()
		remote.Close()
	}
	return err
}

type validConn struct {
	net.Conn
	status bool
	log    *status.Request
}

func pipe(src, dst *validConn) {
	defer func() {
		recover()
		src.Close()
	}()
	b := make([]byte, 1024*16)
	for {
		if !src.status {
			break
		}
		src.SetReadDeadline(time.Now().Add(time.Second * 30))
		n, err := src.Read(b)
		if err == io.EOF {
			dst.status = false
			dst.Close()
			break
		}
		if n > 0 {
			dst.Write(b[:n])
			if dst.log != nil {
				dst.log.Raw = dst.log.Raw + string(b[:n])
			}
		}
	}
}
