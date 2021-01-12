package proxy

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func isCONNECT(content string) bool {
	if content[:7] == "CONNECT" {
		return true
	}
	return false
}

func (p *Proxy) handleCONNECT(conn net.Conn, upgrade []byte) bool {
	content := string(upgrade)
	if !isCONNECT(content) {
		return false
	}
	proxyConn, err := net.DialTimeout("tcp", "192.168.43.1:8080", time.Second)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(conn, "HTTP/1.1 500 Internal server error\r\n")
		fmt.Fprint(conn, "Content-Type: text/plain\r\n")
		fmt.Fprint(conn, "Connection: close\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, err.Error())
		fmt.Fprint(conn, "\r\n")
		conn.Close()
		return true
	}
	vProxyConn := &validConn{proxyConn, true}
	vConn := &validConn{conn, true}
	go pipe(vConn, vProxyConn)
	go pipe(vProxyConn, vConn)
	proxyConn.Write(upgrade)

	os.Stdout.Write(upgrade)
	return true
}
