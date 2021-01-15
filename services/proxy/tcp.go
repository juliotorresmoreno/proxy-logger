package proxy

import (
	"fmt"
	"log"
	"net"
	"time"
)

func isTCP(content string) bool {
	return !isHTTP(content) && !isCONNECT(content)
}

func (p *Proxy) handleTCP(vConn *validConn, upgrade []byte) bool {
	var err error
	content := string(upgrade)
	if !isTCP(content) {
		return false
	}
	proxyConn, err := net.DialTimeout("tcp", getProxyHost(), 5*time.Second)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(vConn, err.Error())
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
