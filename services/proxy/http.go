package proxy

import (
	"fmt"
	"net"
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

type httpData struct {
	method  string
	headers map[string]string
	body    []byte
}

func parseHTTPData(b []byte) *httpData {
	return &httpData{}
}

func (p *Proxy) handleHTTP(conn net.Conn, b []byte) bool {
	content := string(b)
	if !isHTTP(content) {
		return false
	}

	fmt.Println("handleHTTP")
	return true
}
