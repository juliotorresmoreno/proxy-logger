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
	proxyConn, err := net.DialTimeout("tcp", "127.0.0.1:6000", time.Second)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(vConn, "HTTP/1.1 500 Internal server error\r\n")
		fmt.Fprint(vConn, "Content-Type: text/plain\r\n")
		fmt.Fprint(vConn, "Connection: close\r\n")
		fmt.Fprint(vConn, "\r\n")
		fmt.Fprint(vConn, err.Error())
		fmt.Fprint(vConn, "\r\n")
		vConn.Close()
		return true
	}
	vProxyConn := &validConn{"proxy", proxyConn, true}

	/*fmt.Fprint(vConn, "HTTP/1.1 200 Connection established\r\n",
		"Connection: keep-alive\r\n",
		"Via: 1.1 localhost\r\n\r\n")
	go func() {
		body := make([]byte, 1024*64)
		n, _ := vConn.Read(body)
		body = body[:n]
		os.Stdout.Write(body)
	}()*/

	go pipe(vProxyConn, vConn)
	go pipe(vConn, vProxyConn)
	_, err = proxyConn.Write(upgrade)
	if err != nil {
		vConn.Close()
		vProxyConn.Close()
		fmt.Println(err.Error())
	}
	return true
}
