package proxy

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func isHTTP(content string) bool {
	if content[:3] == "GET" {
		return true
	}
	if content[:4] == "POST" {
		return true
	}
	if content[:3] == "PUT" {
		return true
	}
	if content[:4] == "PATH" {
		return true
	}
	if content[:6] == "DELETE" {
		return true
	}
	if content[:7] == "OPTIONS" {
		return true
	}
	return false
}

func readHTTPALLData(conn net.Conn, b []byte) []byte {
	return make([]byte, 0)
}

func getHost(upgrade string) string {
	lines := strings.Split(upgrade, "\r\n")
	for _, line := range lines {
		lower := strings.ToLower(line)
		if lower[:5] == "host:" {
			host := strings.Split(line, ": ")[1]
			if !strings.Contains(host, ":") {
				return host + ":80"
			}
			return host
		}
	}
	return ""
}

func (p *Proxy) handleHTTP(conn *validConn, upgrade []byte) bool {
	content := string(upgrade)
	if !isHTTP(content) {
		return false
	}

	host := getHost(content)
	serverConn, err := net.DialTimeout("tcp", host, 5*time.Second)
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
	vServerConn := &validConn{"", serverConn, true}
	vConn := &validConn{"", conn, true}
	go pipe(vConn, vServerConn)
	go pipe(vServerConn, vConn)
	serverConn.Write(upgrade)

	os.Stdout.Write(upgrade)
	return true
}
