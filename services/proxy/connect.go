package proxy

import (
	"fmt"
	"log"
	"net"
	"time"
)

func isCONNECT(content string) bool {
	if content[:7] == "CONNECT" {
		return true
	}
	return false
}

func (p *Proxy) handleCONNECT(vConn *validConn, upgrade []byte) bool {
	var err error
	content := string(upgrade)
	if !isCONNECT(content) {
		return false
	}
	proxyConn, err := net.DialTimeout("tcp", getProxyHost(), 5*time.Second)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(vConn, "HTTP/1.1 500 Internal server error"+sep)
		fmt.Fprint(vConn, "Content-Type: text/plain"+sep)
		fmt.Fprint(vConn, "Connection: close"+sep+sep)
		fmt.Fprint(vConn, err.Error())
		fmt.Fprint(vConn, sep)
		vConn.Close()
		return true
	}
	vProxyConn := &validConn{"proxy", proxyConn, true}

	go pipe(vProxyConn, vConn)
	go pipe(vConn, vProxyConn)
	_, err = vProxyConn.Write(upgrade)
	if err != nil {
		vConn.Close()
		vProxyConn.Close()
	}
	return true
}
