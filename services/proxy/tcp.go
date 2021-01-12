package proxy

import (
	"fmt"
	"net"
)

func isTCP(content string) bool {
	return !isHTTP(content) && !isCONNECT(content)
}

func (p *Proxy) handleTCP(conn net.Conn, buffer []byte) bool {
	content := string(buffer)
	if !isHTTP(content) {
		return false
	}
	if p.adminServer != nil {
		connection := newConnection()
		connection.buffer.Write(buffer)
		p.adminServer.connection <- connection
		defer func() {
			p.adminServer.closeConnection <- connection
		}()
	}
	fmt.Println(content)
	return true
}
